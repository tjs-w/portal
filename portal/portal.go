package portal

import (
	"container/ring"
	"fmt"
	"log"
	"os"

	"github.com/mattn/go-isatty"

	"github.com/tjs-w/portal/sys"
)

// Options configure the working of Portal.
type Options struct {
	Height  int    // Height of the Portal
	Width   int    // Width of the Portal
	OutFile string // OutFile is the output-file to log the output to
}

// Portal keeps the state of the portal from creation to end.
type Portal struct {
	opt     *Options
	cnt     int
	ringBuf *ring.Ring
	ch      chan string
	isTTY   bool
	oFile   *os.File
}

// New creates a new instance of Portal.
func New(opt *Options) *Portal {
	if !(isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())) {
		return &Portal{
			ch:    make(chan string),
			isTTY: false,
		}
	}

	// Evaluate options and set defaults if needed
	if opt.Height <= 0 || opt.Height > sys.TermHeight() {
		opt.Height = sys.TermHeight() - 1
	}

	var of *os.File
	var err error
	if opt.OutFile != "" {
		if of, err = os.OpenFile(opt.OutFile, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644); err != nil {
			log.Fatalln(err)
		}
	}

	return &Portal{
		opt:     opt,
		ringBuf: ring.New(opt.Height),
		ch:      make(chan string),
		isTTY:   true,
		oFile:   of,
	}
}

func (p *Portal) fileWriter(in <-chan string) <-chan string {
	if p.oFile == nil {
		return in
	}

	ch := make(chan string)
	go func() {
		for line := range in {
			ch <- line
			if _, err := p.oFile.WriteString(line + "\n"); err != nil {
				log.Fatalln(err)
			}
		}
	}()
	return ch
}

// Open starts the portal to process the incoming text. It returns a channel to input the text.
func (p *Portal) Open() chan<- string {
	if !p.isTTY {
		go func() {
			for line := range p.ch {
				fmt.Print(line)
			}
		}()
		return p.ch
	}

	out := p.fileWriter(p.ch)
	out1 := splitAtNewLine(out)
	out2 := p.foldLine(out1)
	go func(in <-chan string) {
		fmt.Println()
		for line := range in {
			p.ringBuf.Value = line
			p.ringBuf = p.ringBuf.Next()
			p.cnt++

			if p.cnt <= p.opt.Height {
				fmt.Print(line)
				continue
			}

			p.reset()
			p.ringBuf.Do(func(line interface{}) {
				fmt.Print(line.(string))
			})
		}
	}(out2)

	return p.ch
}

// Close releases the resources being used by the portal and closes it.
func (p *Portal) Close() {
	close(p.ch)
	if p.oFile == nil {
		return
	}
	fmt.Println("Output written to ", p.opt.OutFile)
	if err := p.oFile.Close(); err != nil {
		log.Fatalln(err)
	}
}

// reset clears the lines above and moves the cursor to initial position.
func (p *Portal) reset() {
	for i := 0; i < p.opt.Height; i++ {
		moveUp(1)
		clearLine()
	}
}

// clearLine clears the current line and moves the cursor to its start.
func clearLine() {
	fmt.Print("\x1b[2K")
}

// moveUp moves the cursor n line(s) up
func moveUp(n int) {
	fmt.Printf("\x1b[%dA", n)
}

// splitAtNewLine takes an input line and breaks it at '\n' to form the output.
func splitAtNewLine(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for line := range in {
			st := 0
			for i, r := range line {
				if r == '\n' {
					out <- line[st:i]
					st = i + 1
				}
			}
			if st < len(line) {
				out <- line[st:]
			}
		}
		close(out)
	}()
	return out
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// currWidth is used to determine current width to dynamically change the value if the terminal width is changed
// on the fly.
func (p *Portal) currWidth() int {
	if p.opt.Width <= 0 || p.opt.Width > sys.TermWidth() {
		return sys.TermWidth()
	}
	return p.opt.Width
}

// foldLine takes a line, determines the size of current terminal or configured width and then breaks it accordingly.
func (p *Portal) foldLine(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for line := range in {
			l := len(line)
			w := p.currWidth()
			for s, e := 0, min(w, l); ; s, e = e, min(e+w, l) {
				out <- fmt.Sprintln(line[s:e])
				if e == l {
					break
				}
			}
		}
		close(out)
	}()
	return out
}
