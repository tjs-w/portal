//go:build windows

package sys

import (
	"log"
	"os"

	"golang.org/x/sys/windows"
)

func winSize() (*windows.ConsoleScreenBufferInfo, error) {
	fd := os.Stdout.Fd()
	var info windows.ConsoleScreenBufferInfo
	if err := windows.GetConsoleScreenBufferInfo(windows.Handle(fd), &info); err != nil {
		return nil, err
	}
	return info, nil
}

func TermHeight() int {
	info, err := winSize()
	if err != nil {
		log.Fatalln(err)
	}
	return info.Window.Bottom - info.Window.Top + 1
}

func TermWidth() int {
	info, err := winSize()
	if err != nil {
		log.Fatalln(err)
	}
	return info.Window.Right - info.Window.Left + 1
}

// SplitAtNewLine takes an input line and breaks it at '\n' to form the output.
func SplitAtNewLine(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		const (
			ANY_RUNE        = iota
			CARRIAGE_RETURN // "\r"
			LINE_FEED       // "\n"
		)

		for line := range in {
			start := 0
			prev := ANY_RUNE
			cr_idx := 0
			for i, r := range line {
				if r == '\r' {
					prev = CARRIAGE_RETURN
					continue
				}
				if r == '\n' && prev == CARRIAGE_RETURN {
					out <- line[start : i-1]
					start = i + 1
					prev = LINE_FEED
					continue
				}
				prev = ANY_RUNE
			}
			if start < len(line) {
				out <- line[start:]
			}
		}
		close(out)
	}()
	return out
}
