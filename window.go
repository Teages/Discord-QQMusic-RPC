package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	user32        = syscall.MustLoadDLL("User32.dll")
	findWindow    = user32.MustFindProc("FindWindowW")
	getWindowText = user32.MustFindProc("GetWindowTextW")

	getWindowTextLength = user32.MustFindProc("GetWindowTextLengthW")
)

func GetDesktopWindowName(appName string) string {
	match, _ := syscall.UTF16FromString(appName)
	handel, _, err := findWindow.Call(uintptr(unsafe.Pointer(&match[0])), uintptr(unsafe.Pointer(nil)))
	if err != nil {
		fmt.Println(err)
	}
	if handel == 0 {
		return ""
	}

	len, _, err := getWindowTextLength.Call(handel)
	if err != nil {
		fmt.Println(err)
	}

	titleW := make([]uint16, len)
	getWindowText.Call(handel, uintptr(unsafe.Pointer(&titleW[0])), len)
	title := syscall.UTF16ToString(titleW)
	fmt.Println(title)
	return title
}
