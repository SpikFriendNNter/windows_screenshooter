package print

import (
	"windows_screenshooter/winapi"
)

type BITMAPINFOHEADER struct {
  BiSize 				      winapi.DWORD 
  BiWidth 				    winapi.LONG
  BiHeight 				    winapi.LONG  
  BiPlanes 				    winapi.WORD  
  BiBitCount 			    winapi.WORD  
  BiCompression 	    winapi.DWORD 
  BiSizeImage 		    winapi.DWORD 
  BiXPelsPerMeter     winapi.LONG  
  BiYPelsPerMeter     winapi.LONG  
  BiClrUsed 			    winapi.DWORD 
  BiClrImportant 	    winapi.DWORD 
}

type RGBQUAD struct {
  rgbBlue	 	     uint8
  rgbGreen	 	   uint8
  rgbRed	 	     uint8
  rgbReserved    uint8
}

type BITMAPINFO struct {
	BmiHeader	BITMAPINFOHEADER
	BmiColors	[1]RGBQUAD
}