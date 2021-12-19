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
