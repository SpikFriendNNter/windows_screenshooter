package print

import (
	"image"
	"syscall"
	"log"
	"errors"
	"unsafe"

	"windows_screenshooter/screen"
	"windows_screenshooter/winapi"
)

func Take_Screenshot() (*image.RGBA, error) {
	// Getting window HANDLE
	w_handle := screen.GetWindowHwnd()
	m_handle, err := screen.GetMonitorWindowHwnd(w_handle)
	if err != nil {
		return nil, err
	}

	// Getting window specifications
	monitor_info, _ := screen.GetMonitorInfo(m_handle)

	// Getting width and height
	width  := monitor_info.RcMonitor.Right  - monitor_info.RcMonitor.Left
	height := monitor_info.RcMonitor.Bottom - monitor_info.RcMonitor.Top

	// Creating image
	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))

	// Getting window DC and setting when to free it
	dc := getDC(w_handle)
	defer releaseDC(w_handle, dc)

	// Creating a new DC (device context)
	dc_dest, _ := createCompatibleDC(dc)

	bitmap_h, err := createBitmapObject(dc, int32(width), int32(height))
	if err != nil {
		return nil, err
	}

	// Setting basic BITMAP
	// informations
	var header BITMAPINFOHEADER
	header.BiSize = winapi.DWORD(unsafe.Sizeof(header))
	header.BiPlanes = 1
	header.BiBitCount = 32
	header.BiWidth = winapi.LONG(width)
	header.BiHeight = winapi.LONG(-height)
	header.BiCompression = winapi.BI_RGB
	header.BiSizeImage = 0

	bitmapSize := uintptr(((int64(width)*int64(header.BiBitCount) + 31) / 32) * 4 * int64(height))

	// Allocating a buffer memory to store
	// bitmap info
	alloc_h, _ := globalAlloc(bitmapSize)
	lock_p,  _ := globalLock(alloc_h) // Getting pointer from alloc_h above
	defer globalUnlock(lock_p)
	defer globalFree(alloc_h)

	// Selecting Bitmap Object
	hOld, err := selectBitmap(dc_dest, bitmap_h)
	if err != nil {
		return nil, err
	}

	// Copying color-bits informations to destination DC
	_, err = transferColorsToBitmap_O(dc_dest, dc, int32(width), int32(height))
	selectBitmap(dc_dest, hOld)

	// Getting information from the new bitmap object
	// and storing it inside of our global allocated space of memory
	// created above.
	err = sendingBitmapInfoToBuffer(dc, bitmap_h, uint32(height), lock_p, (*BITMAPINFO)(unsafe.Pointer(&header)))
	if err != nil {
		return nil, err
	}

	i := 0
	src := lock_p
	for x := 0; x < int(width); x++ {
		for y := 0; y < int(height); y++ {
			r := *(*uint8)(unsafe.Pointer(src +2))
			g := *(*uint8)(unsafe.Pointer(src +1))
			b := *(*uint8)(unsafe.Pointer(src))

			// Adding rgba data inside of image
			img.Pix[i], img.Pix[i+1], img.Pix[i+2], img.Pix[i+3] = r, g, b, 255

			i += 4
			src += 4
		}
	}

	return img, nil
}

func loadDll(dll_name string) *syscall.DLL {
	dll, err := syscall.LoadDLL(dll_name)
	if err != nil {
		log.Fatal("ERROR LOADING DLL: ", err)
	}

	return dll
}

func getDC(hwnd uintptr) uintptr {
	user32 := loadDll("user32.dll")
	getDC := user32.MustFindProc("GetDC")
	dc, _, _ := getDC.Call(hwnd)
	return dc
}

func releaseDC(hwnd, DC uintptr) (int8, error) {
	user32 := loadDll("user32.dll")
	releaseDC, err := user32.FindProc("ReleaseDC")
	if err != nil {
		return 0, err
	}

	free, _, _ := releaseDC.Call(hwnd, DC)

	return int8(free), nil
}

func createBitmapObject(DC uintptr, width, height int32) (uintptr, error) {
	// loading dll and check if ocurred some error
	gdi32 := loadDll("gdi32.dll")

	// Loading CreateCompatibleBitmap procedure from
	// gdi32.dll
	createCompatibleBitmap, err := gdi32.FindProc("CreateCompatibleBitmap")
	if err != nil {
		return uintptr(0), err
	}

	bitmap_handle, _, _ := createCompatibleBitmap.Call(DC,
	uintptr(width),
	uintptr(height))

	// Checking if procedure result on error
	if bitmap_handle == 0 {
		return uintptr(0), nil
	}

	return bitmap_handle, nil
}

func createCompatibleDC(DC uintptr) (uintptr, error) {
	gdi32 := loadDll("gdi32.dll")
	createCompatibleDC, err := gdi32.FindProc("CreateCompatibleDC")
	if err != nil {
		return uintptr(0), err
	}

	DC_handle, _, _ := createCompatibleDC.Call(DC)
	if DC_handle == 0 {
		return uintptr(0), nil
	}

	return DC_handle, nil
}

func selectBitmap(dc, hBitmap uintptr) (uintptr, error) {
	gdi32 := loadDll("gdi32.dll")
	selectObject, err := gdi32.FindProc("SelectObject")
	if err != nil {
		return uintptr(0), err
	}

	hOld, _, _ := selectObject.Call(dc, hBitmap)
	return hOld, nil
}

func transferColorsToBitmap_O(dc_d, dc_s uintptr, width, height int32) (int8, error) {
	gdi32 := loadDll("gdi32.dll")
	bit_blt, err := gdi32.FindProc("BitBlt")
	if err != nil {
		return int8(0), err
	}

	result, _, _ := bit_blt.Call(
		dc_d,
		0,
		0,
		uintptr(width),
		uintptr(height),
		dc_s,
		0,
		0,
		uintptr(winapi.SRCCOPY))

	return int8(result), nil
}

func globalAlloc(bitmap_S uintptr) (uintptr, error) {
	kernel32 := loadDll("kernel32.dll")
	globalAllocation, err := kernel32.FindProc("GlobalAlloc")

	if err != nil {
		return uintptr(0), err
	}

	memory_h, _, _ := globalAllocation.Call(
		uintptr(winapi.GMEM_MOVEABLE),
		uintptr(bitmap_S))

	if memory_h == 0 {
		return uintptr(0), nil
	}

	return memory_h, nil
}

func globalFree(global_mem_h uintptr) (error) {
	kernel32 := loadDll("kernel32.dll")
	globalFreeMem, err := kernel32.FindProc("GlobalFree")

	if err != nil {
		return nil
	} else if r, _, _ := globalFreeMem.Call(global_mem_h); r != 0 {
		return errors.New("Failed to free global memory allocation.")
	}

	return nil
}

func globalLock(globalMem_h uintptr) (uintptr, error) {
	kernel32 := loadDll("kernel32.dll")
	_globalLock, err := kernel32.FindProc("GlobalLock")
	if err != nil {
		return uintptr(0), nil
	}

	// Trying to get pointer of GlobalAlloc
	result, _, _ := _globalLock.Call(
		globalMem_h)

	if result == 0 {
		return uintptr(0), nil
	}

	return result, nil
}

func globalUnlock(pointer uintptr) (int, error) {
	kernel32 := loadDll("kernel32.dll")
	_globalUnlock, err := kernel32.FindProc("GlobalUnlock")
	if err != nil {
		return -1, err
	}

	lock, _, _ := _globalUnlock.Call(pointer)
	if lock != 0 {
		return int(lock), errors.New("Failed trying unlock.")
	}

	return int(lock), nil
 }

func sendingBitmapInfoToBuffer(dc_source, bitmap_h uintptr, height uint32,
	gMemPtr uintptr,
	lpbmi *BITMAPINFO) (error) {
		gdi32 := loadDll("gdi32.dll")
		getDIBits, err := gdi32.FindProc("GetDIBits")

		if err != nil {
			return err
		}

		result, _, _ := getDIBits.Call(
			dc_source,
			bitmap_h,
			0,
			uintptr(height),
			gMemPtr,
			uintptr(unsafe.Pointer(lpbmi)),
			0)

		if result == 0 {
			return errors.New("GetDIBits function failed.")
		}

		return nil
}
