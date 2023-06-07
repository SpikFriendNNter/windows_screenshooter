package winapi

type HWND		uintptr
type LONG		int32
type LONGLONG		uint64
type WORD		uint16
type DWORD		uint32
type DWORDLONG		LONGLONG


const (
	MONITOR_DEFAULTTOPRIMARY 	 = 0x00000001
	SRCCOPY 			 = int32(0x00CC0020)
	BI_RGB				 = 0
	BI_BITFIELDS			 = 3
	GMEM_MOVEABLE			 = 0x0002
	DIB_RGB_COLORS 			 = 0x00
	DIB_PAL_COLORS 		 	 = 0x01
)
