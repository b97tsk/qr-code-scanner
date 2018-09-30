package main

import "golang.org/x/sys/windows"

func init() {
	user32 := windows.NewLazySystemDLL("user32.dll")

	defer func() {
		_ = recover()

		windows.FreeLibrary(windows.Handle(user32.Handle()))
	}()

	user32.NewProc("SetProcessDPIAware").Call()
}
