//go:build !windows && !plan9 && !solaris

package sys

import (
	"log"
	"os"

	"golang.org/x/sys/unix"
)

func winSize() (*unix.Winsize, error) {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil, os.NewSyscallError("IoctlGetWinsize", err)
	}
	return ws, nil
}

func TermHeight() int {
	w, err := winSize()
	if err != nil {
		log.Fatalln(err)
	}
	return int(w.Row)
}

func TermWidth() int {
	w, err := winSize()
	if err != nil {
		log.Fatalln(err)
	}
	return int(w.Col)
}
