package screen

import (
	"syscall"
	"unsafe"
	"log"

	"github.com/spikfriendnnter/windows_screenshooter/winapi"
)

func loadDll(dll_name string) *syscall.DLL {
	dll, err := syscall.LoadDLL(dll_name)
	if err != nil {
		log.Fatal("ERROR LOADING DLL: ", err)
	}

	return dll
}

func GetWindowHwnd() (hwnd uintptr) {
	user32 := loadDll("user32.dll")
	getDesktopWindow := user32.MustFindProc("GetDesktopWindow")

	hwnd, _, _ = getDesktopWindow.Call()
	return
}

func GetMonitorWindowHwnd(windowHWND uintptr) (uintptr, error) {
	user32 := loadDll("user32.dll")
	monitorFromWindow, err := user32.FindProc("MonitorFromWindow")
	if err  != nil {
		return uintptr(0), err
	}

	monitor_h, _, _ := monitorFromWindow.Call(windowHWND, uintptr(winapi.MONITOR_DEFAULTTOPRIMARY))
	return monitor_h, nil
}

func GetMonitorInfo(windowHwnd uintptr) (*MONITORINFO, error) {
	user32 := loadDll("user32.dll")

	getMonitorInfoW := user32.MustFindProc("GetMonitorInfoW")
	lpmi := MONITORINFO{}
	lpmi.CbSize = winapi.DWORD(unsafe.Sizeof(lpmi))

	result, _, _ := getMonitorInfoW.Call(windowHwnd, uintptr(unsafe.Pointer(&lpmi)))

	if result == 0 {
		return nil, nil 
	}

	return &lpmi, nil
}
