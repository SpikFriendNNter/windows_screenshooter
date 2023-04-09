# <strong>WARNING: This code is only for educational purposes. Please, do not use as part of any malicious code. Spreading malware is crime.</strong>

# <strong>How is the screenshot process done using the Windows API?</strong>

<img src="https://assets.labs.ine.com/web/badges/low/winapi.png" width=200>

## <strong>Process</strong>

### The image capture process using the windows api is given by the following structure:

- Get the DC(Device Context) of the screen.
    
- Create a Bitmap object compatible with the source DC..
    
- Duplicate that DC, creating a new DC compatible with the original.
    
- You need to define some default values ​​that a bitmap header loads.
    
- After creating the bitmap object, it is necessary to allocate enough space to contain the size of the bitmap in the memory of the operating system to serve as a buffer, and then we can access the bitmap data, which later will contain the color data obtained by the DC source.

- We use Windows API functions to transfer/copy color data from the source DC to the destination DC. The Bitmap object inside the destination DC will contain the color data obtained from the source DC. Usually the function that performs such a task is BitBlt from the gdi32 dll.

- After that, we use the GetDIBits function that will transfer the data from the Bitmap object in the destination DC to the buffer we allocated in memory.

- With that done, we can get the RGB color data from the buffer as a roll of bits (like an array), which has the following data order: BGR => Blue, Green, Red.

- Finally, remember to free the buffer allocated in memory, free the DCs and window handlers, and close the bitmap objects.

## <strong>O que é DC?</strong>
- According to the windows documentation (https://learn.microsoft.com/en-us/cpp/mfc/device-contexts?view=msvc-170), a DC or Device Context is a Windows data structure containing information about drawing attributes of a device such as a monitor or printer. In short, it is a data structure for representing a graphics device. For more details, refer to the documentation indicated in the link above.

## <strong>How to get screen specs?</strong>

- <span style="color:yellow;">GetMonitorInfoW:</span> function that gets the specifications of the specified monitor. Remembering that, for this function to work, it is necessary to provide a parameter, which is a <i style="color:#9dcec7;">HANDLE</i> for the specified window monitor. <strong style="color:red;">Atention:</strong> Passing the HANDLE of the specified window as a parameter will cause an error. The correct way is to get the <i style="color:#9dcec7;">HANDLE</i> from the monitor using the function MonitorFromWindow.

- <span style="color:yellow;">MonitorFromWindow</span>: function that takes a <i style="color:#9dcec7;">HANDLE</i> as a parameter for a specified window and returns a <i style="color:#9dcec7;">HANDLE</i> of that window's monitor.

## <strong>Windows Api functions used in this project:</strong>
- <span style="color:yellow;">GetDesktopWindow</span>;
- <span style="color:yellow;">MonitorFromWindow</span>;
- <span style="color:yellow;">GetMonitorInfoW</span>;
- <span style="color:yellow;">GetDC</span>;
- <span style="color:yellow;">ReleaseDC</span>;
- <span style="color:yellow;">CreateCompatibleDC</span>;
- <span style="color:yellow;">CreateCompatibleBitmap</span>;
- <span style="color:yellow;">SelectObject</span>;
- <span style="color:yellow;">BitBlt</span>;
- <span style="color:yellow;">GetDIBits</span>;
- <span style="color:yellow;">GlobalAlloc</span>;
- <span style="color:yellow;">GlobalFree</span>;
- <span style="color:yellow;">GlobalLock</span>;

## Project used as base
- https://github.com/kbinani/screenshot
