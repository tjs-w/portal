//go:build !windows && !plan9 && !solaris

package sys

import (
	"log"
	"os"

	"golang.org/x/sys/unix"
)

// winSize makes the unix system call IOCTL to get the dimensions of the TTY window
func winSize() (*unix.Winsize, error) {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil, os.NewSyscallError("IoctlGetWinsize", err)
	}
	return ws, nil
}

// TermHeight returns the height of terminal, in # of lines
func TermHeight() int {
	w, err := winSize()
	if err != nil {
		log.Fatalln(err)
	}
	return int(w.Row)
}

// TermWidth returns the width of the terminal, in # of runes
func TermWidth() int {
	w, err := winSize()
	if err != nil {
		log.Fatalln(err)
	}
	return int(w.Col)
}

// SplitAtNewLine takes an input line and breaks it at '\n' to form the output.
func SplitAtNewLine(in <-chan string) <-chan string {
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
