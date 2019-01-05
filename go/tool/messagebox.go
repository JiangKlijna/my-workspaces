package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"unsafe"
)

func messageBox(title, text string) bool {
	switch runtime.GOOS {
	case "windows":
		user32 := syscall.NewLazyDLL("user32.dll")
		messageBoxW := user32.NewProc("MessageBoxW")
		mbYesNo := 0x00000004
		mbIconQuestion := 0x00000020
		idYes := 6
		ret, _, _ := messageBoxW.Call(0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
			uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), uintptr(uint(mbYesNo|mbIconQuestion)))
		return int(ret) == idYes
	case "linux":
		err := exec.Command("zenity", "--question", "--title", title, "--text", text).Run()
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				return exitError.Sys().(syscall.WaitStatus).ExitStatus() == 0
			}
		}
		return true
	case "darwin":
		script := `set T to button returned of ` +
			`(display dialog "%s" with title "%s" buttons {"No", "Yes"} default button Yes")`
		out, err := exec.Command("osascript", "-e", fmt.Sprintf(script, text, title)).Output()
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				return exitError.Sys().(syscall.WaitStatus).ExitStatus() == 0
			}
		}
		return strings.TrimSpace(string(out)) == "Yes"
	default:
		return false
	}
}

func main() {
	messageBox("title", "text")
}
