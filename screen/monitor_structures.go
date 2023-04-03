package screen

import "windows_screenshooter/winapi"

type MONITORINFO struct {
	CbSize		winapi.DWORD
	RcMonitor	RECT
	RcWork		RECT
	dwFlags		winapi.DWORD
}

type RECT struct {
	Left	winapi.LONG
	Top		winapi.LONG
	Right	winapi.LONG
	Bottom	winapi.LONG
}