package zebrazxp13

import (
    "syscall"
    "unsafe"
)

var (
    zBRGraphics                    = syscall.NewLazyDLL("ZBRGraphics.dll")
    zBRGDICancelJob                = zBRGraphics.NewProc("ZBRGDICancelJob")
    zBRGDIClearGraphics            = zBRGraphics.NewProc("ZBRGDIClearGraphics")
    zBRGDICloseGraphics            = zBRGraphics.NewProc("ZBRGDICloseGraphics")
    zBRGDIDrawBarCode              = zBRGraphics.NewProc("ZBRGDIDrawBarCode")
    zBRGDIDrawEllipse              = zBRGraphics.NewProc("ZBRGDIDrawEllipse")
    zBRGDIDrawImage                = zBRGraphics.NewProc("ZBRGDIDrawImage")
    zBRGDIDrawImageEx              = zBRGraphics.NewProc("ZBRGDIDrawImageEx")
    zBRGDIDrawImagePos             = zBRGraphics.NewProc("ZBRGDIDrawImagePos")
    zBRGDIDrawImagePosEx           = zBRGraphics.NewProc("ZBRGDIDrawImagePosEx")
    zBRGDIDrawImageRect            = zBRGraphics.NewProc("ZBRGDIDrawImageRect")
    zBRGDIDrawImageRectEx          = zBRGraphics.NewProc("ZBRGDIDrawImageRectEx")
    zBRGDIDrawLine                 = zBRGraphics.NewProc("ZBRGDIDrawLine")
    zBRGDIDrawRectangle            = zBRGraphics.NewProc("ZBRGDIDrawRectangle")
    zBRGDIDrawText                 = zBRGraphics.NewProc("ZBRGDIDrawText")
    zBRGDIDrawTextEx               = zBRGraphics.NewProc("ZBRGDIDrawTextEx")
    zBRGDIDrawTextRect             = zBRGraphics.NewProc("ZBRGDIDrawTextRect")
    zBRGDIDrawTextRectEx           = zBRGraphics.NewProc("ZBRGDIDrawTextRectEx")
    zBRGDIDrawTextUnicode          = zBRGraphics.NewProc("ZBRGDIDrawTextUnicode")
    zBRGDIDrawTextUnicodeEx        = zBRGraphics.NewProc("ZBRGDIDrawTextUnicodeEx")
    zBRGDIEndPage                  = zBRGraphics.NewProc("ZBRGDIEndPage")
    zBRGDIGetSDKProductVer         = zBRGraphics.NewProc("ZBRGDIGetSDKProductVer")
    zBRGDIGetSDKVer                = zBRGraphics.NewProc("ZBRGDIGetSDKVer")
    zBRGDIGetSDKVsn                = zBRGraphics.NewProc("ZBRGDIGetSDKVsn")
    zBRGDIInitGraphics             = zBRGraphics.NewProc("ZBRGDIInitGraphics")
    zBRGDIInitGraphicsEx           = zBRGraphics.NewProc("ZBRGDIInitGraphicsEx")
    zBRGDIInitGraphicsFromPrintDlg = zBRGraphics.NewProc("ZBRGDIInitGraphicsFromPrintDlg")
    zBRGDIIsPrinterReady           = zBRGraphics.NewProc("ZBRGDIIsPrinterReady")
    zBRGDIPreviewGraphics          = zBRGraphics.NewProc("ZBRGDIPreviewGraphics")
    zBRGDIPrintFilePos             = zBRGraphics.NewProc("ZBRGDIPrintFilePos")
    zBRGDIPrintFileRect            = zBRGraphics.NewProc("ZBRGDIPrintFileRect")
    zBRGDIPrintGraphics            = zBRGraphics.NewProc("ZBRGDIPrintGraphics")
    zBRGDIPrintGraphicsEx          = zBRGraphics.NewProc("ZBRGDIPrintGraphicsEx")
    zBRGDIPrintImagePos            = zBRGraphics.NewProc("ZBRGDIPrintImagePos")
    zBRGDIPrintImageRect           = zBRGraphics.NewProc("ZBRGDIPrintImageRect")
    zBRGDIStartPage                = zBRGraphics.NewProc("ZBRGDIStartPage")
)

// extern void ZBRGDIGetSDKVer(out int major, out int minor, out int englevel);
func ZBRGDIGetSDKVer() (major, minor, engLevel int) {
    _, _, _ =  zBRGDIGetSDKVer.Call(
        uintptr(unsafe.Pointer(&major)),
        uintptr(unsafe.Pointer(&minor)),
        uintptr(unsafe.Pointer(&engLevel)),
    )
    return
}

// extern int ZBRGDIIsPrinterReady(byte[] printername, out int err);
func ZBRGDIIsPrinterReady(strPrinterName string) (ret uintptr,err uint) {
        ret, _, _ = zBRGDIIsPrinterReady.Call(
                uintptr(unsafe.Pointer(&([]byte(strPrinterName))[0])),
                uintptr(unsafe.Pointer(&err)),
        )
        return ret, err
}

// extern void ZBRGDIGetSDKVsn(out int major, out int minor, out int englevel);
func ZBRGDIGetSDKVsn() (ret uintptr) {
    panic("ZBRGDIGetSDKVsn not implemented")
    ret, _, _ =  zBRGDIGetSDKVsn.Call(
        )
    return ret
}

// extern int ZBRGDIGetSDKProductVer(byte[] productversion, out int buffsize, out int err);
func ZBRGDIGetSDKProductVer() (ret uintptr) {
    panic("ZBRGDIGetSDKProductVer not implemented")
    ret, _, _ =  zBRGDIGetSDKProductVer.Call(
        )
    return ret
}

// extern int ZBRGDIInitGraphics(byte[] printername, out IntPtr hdc, out int err);
func ZBRGDIInitGraphics(strPrinterName string) (ret uintptr, handle syscall.Handle, err uint) {
        ret, _, _ = zBRGDIInitGraphics.Call(
                uintptr(unsafe.Pointer(&([]byte(strPrinterName))[0])),
                uintptr(unsafe.Pointer(&handle)),
                uintptr(unsafe.Pointer(&err)),
        )
        return ret, handle, err
}

// extern int ZBRGDIInitGraphicsEx(byte[] printername, out IntPtr hdc, byte[] jobname, out int jobid, out int err);
func ZBRGDIInitGraphicsEx() (ret uintptr) {
    panic("ZBRGDIInitGraphicsEx not implemented")
    ret, _, _ =  zBRGDIInitGraphicsEx.Call(
        )
    return ret
}

// extern int ZBRGDIInitGraphicsFromPrintDlg(out IntPtr hdc, out int err);
func ZBRGDIInitGraphicsFromPrintDlg() (ret uintptr) {
    panic("ZBRGDIInitGraphicsFromPrintDlg not implemented")
    ret, _, _ =  zBRGDIInitGraphicsFromPrintDlg.Call(
        )
    return ret
}

// extern int ZBRGDICloseGraphics(IntPtr hdc, out int err);
func ZBRGDICloseGraphics(handle syscall.Handle) (ret uintptr, err uint){
    ret, _, _ = zBRGDICloseGraphics.Call(
        uintptr(handle),
        uintptr(unsafe.Pointer(&err)),
    )
    return ret, err
}

// extern int ZBRGDIClearGraphics(out int err);
func ZBRGDIClearGraphics() (ret uintptr, err uint) {
    ret, _, _ = zBRGDIClearGraphics.Call(
        uintptr(unsafe.Pointer(&err)),
    )
    return ret, err
}

// extern int ZBRGDIStartPage(IntPtr hdc, out int err);
func ZBRGDIStartPage() (ret uintptr) {
    panic("ZBRGDIStartPage not implemented")
    ret, _, _ =  zBRGDIStartPage.Call(
        )
    return ret
}

// extern int ZBRGDIEndPage(IntPtr hdc, out int err);
func ZBRGDIEndPage() (ret uintptr) {
    panic("ZBRGDIEndPage not implemented")
    ret, _, _ =  zBRGDIEndPage.Call(
        )
    return ret
}

// extern int ZBRGDIPreviewGraphics(IntPtr hwnd, out int err);
func ZBRGDIPreviewGraphics() (ret uintptr) {
    panic("ZBRGDIPreviewGraphics not implemented")
    ret, _, _ =  zBRGDIPreviewGraphics.Call(
        )
    return ret
}

// extern int ZBRGDIPrintGraphics(IntPtr hdc, out int err);
func ZBRGDIPrintGraphics(handle syscall.Handle) (ret uintptr, err uint) {
        ret, _, _ = zBRGDIPrintGraphics.Call(
                uintptr(handle),
                uintptr(unsafe.Pointer(&err)),
        )
        return ret, err
}

// extern int ZBRGDIPrintGraphicsEx(IntPtr hdc, out int err);
func ZBRGDIPrintGraphicsEx() (ret uintptr) {
    panic("ZBRGDIPrintGraphicsEx not implemented")
    ret, _, _ =  zBRGDIPrintGraphicsEx.Call(
        )
    return ret
}

// extern int ZBRGDIPrintFilePos(IntPtr hdc, byte[] filename, int position, out int err);
func ZBRGDIPrintFilePos() (ret uintptr) {
    panic("ZBRGDIPrintFilePos not implemented")
    ret, _, _ =  zBRGDIPrintFilePos.Call(
        )
    return ret
}

// extern int ZBRGDIPrintImagePos(IntPtr hdc, byte[] imagedata, int imagesize, int position, out int err);
func ZBRGDIPrintImagePos() (ret uintptr) {
    panic("ZBRGDIPrintImagePos not implemented")
    ret, _, _ =  zBRGDIPrintImagePos.Call(
        )
    return ret
}

// extern int ZBRGDIPrintImageRect(IntPtr hdc, byte[] imagedata, int imagesize, int x, int y, int width, int height, out int err);
func ZBRGDIPrintImageRect() (ret uintptr) {
    panic("ZBRGDIPrintImageRect not implemented")
    ret, _, _ =  zBRGDIPrintImageRect.Call(
        )
    return ret
}

// extern int ZBRGDICancelJob(byte[] printername, int jobid, out int err);
func ZBRGDICancelJob() (ret uintptr) {
    panic("ZBRGDICancelJob not implemented")
    ret, _, _ =  zBRGDICancelJob.Call(
        )
    return ret
}

// extern int ZBRGDIDrawText(int x, int y, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawText(x uint, y uint, text string, font string, fontsize uint, fontstyle TextFontStyle, color uint32) (ret uintptr, err uint) {
    ret, _, _ =  zBRGDIDrawText.Call(
        uintptr(x),
        uintptr(y),
        uintptr(unsafe.Pointer(&([]byte(text))[0])),
        uintptr(unsafe.Pointer(&([]byte(font))[0])),
        uintptr(fontsize),
        uintptr(fontstyle),
        uintptr(color),
        uintptr(unsafe.Pointer(&err)),
    )
    return ret, err
}

// extern int ZBRGDIDrawTextEx(int x, int y, int angle, int alignment, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawTextEx() (ret uintptr) {
    panic("ZBRGDIDrawTextEx not implemented")
    ret, _, _ =  zBRGDIDrawTextEx.Call(
        )
    return ret
}

// extern int ZBRGDIDrawTextUnicode(int x, int y, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawTextUnicode(x uint, y uint, text string, font string, fontsize uint, fontstyle TextFontStyle, color uint32) (ret uintptr, err uint) {
    ret, _, _ =  zBRGDIDrawTextUnicode.Call(
        uintptr(x),
        uintptr(y),
        uintptr(unsafe.Pointer(&([]byte(text))[0])),
        uintptr(unsafe.Pointer(&([]byte(font))[0])),
        uintptr(fontsize),
        uintptr(fontstyle),
        uintptr(color),
        uintptr(unsafe.Pointer(&err)),
    )
    return ret, err
}

// extern int ZBRGDIDrawTextUnicodeEx(int x, int y, int angle, int alignment, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawTextUnicodeEx() (ret uintptr) {
    panic("ZBRGDIDrawTextUnicodeEx not implemented")
    ret, _, _ =  zBRGDIDrawTextUnicodeEx.Call(
        )
    return ret
}

// extern int ZBRGDIDrawTextRect(int x, int y, int width, int height, int alignment, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawTextRect(x uint, y uint, width uint, height uint, alignment TextAlignment, text string, font string, fontsize uint, fontstyle TextFontStyle, color uint32) (ret uintptr, err uint) {
    ret, _, _ =  zBRGDIDrawTextRect.Call(
        uintptr(x),
        uintptr(y),
        uintptr(width),
        uintptr(height),
        uintptr(alignment),
        uintptr(unsafe.Pointer(&([]byte(text))[0])),
        uintptr(unsafe.Pointer(&([]byte(font))[0])),
        uintptr(fontsize),
        uintptr(fontstyle),
        uintptr(color),
        uintptr(unsafe.Pointer(&err)),
    )
    return ret, err
}

// extern int ZBRGDIDrawTextRectEx(int x, int y, int width, int height, int angle, int alignment, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawTextRectEx() (ret uintptr) {
    panic("ZBRGDIDrawTextRectEx not implemented")
    ret, _, _ =  zBRGDIDrawTextRectEx.Call(
        )
    return ret
}

// extern int ZBRGDIDrawLine(int x1, int y1, int x2, int y2, int color, float thickness, out int err);
func ZBRGDIDrawLine(x1 uint, y1 uint, x2 uint, y2 uint, color uint32, thickness float32) (ret uintptr, err uint) {
    ret, _, _ =  zBRGDIDrawLine.Call(
        uintptr(x1),
        uintptr(y1),
        uintptr(x2),
        uintptr(y2),
        uintptr(color),
        uintptr(thickness),
        uintptr(unsafe.Pointer(&err)),
    )
    return ret, err
}

// extern int ZBRGDIDrawImage(byte[] filename, int x, int y, out int err);
func ZBRGDIDrawImage() (ret uintptr) {
    panic("ZBRGDIDrawImage not implemented")
    ret, _, _ =  zBRGDIDrawImage.Call(
        )
    return ret
}

// extern int ZBRGDIDrawImageEx(byte[] filename, int imagesize, int x, int y, out int err);
func ZBRGDIDrawImageEx() (ret uintptr) {
    panic("ZBRGDIDrawImageEx not implemented")
    ret, _, _ =  zBRGDIDrawImageEx.Call(
        )
    return ret
}

// extern int ZBRGDIDrawImagePos(byte[] filename, int position, out int err);
func ZBRGDIDrawImagePos() (ret uintptr) {
    panic("ZBRGDIDrawImagePos not implemented")
    ret, _, _ =  zBRGDIDrawImagePos.Call(
        )
    return ret
}

// extern int ZBRGDIDrawImagePosEx(byte[] filename, int imagesize, int position, out int err);
func ZBRGDIDrawImagePosEx() (ret uintptr) {
    panic("ZBRGDIDrawImagePosEx not implemented")
    ret, _, _ =  zBRGDIDrawImagePosEx.Call(
        )
    return ret
}

// extern int ZBRGDIDrawImageRect(byte[] filename, int x, int y, int width, int height, out int err);
func ZBRGDIDrawImageRect() (ret uintptr) {
    panic("ZBRGDIDrawImageRect not implemented")
    ret, _, _ =  zBRGDIDrawImageRect.Call(
        )
    return ret
}

// extern int ZBRGDIDrawImageRectEx(byte[] filename, int imagesize, int x, int y, int width, int height, out int err);
func ZBRGDIDrawImageRectEx() (ret uintptr) {
    panic("ZBRGDIDrawImageRectEx not implemented")
    ret, _, _ =  zBRGDIDrawImageRectEx.Call(
        )
    return ret
}

// extern int ZBRGDIDrawRectangle(int x, int y, int width, int height, float thickness, int color, out int err);
func ZBRGDIDrawRectangle(x uint, y uint, width uint, height uint, thickness float32, color uint32) (ret uintptr, err uint) {
    ret, _, _ =  zBRGDIDrawRectangle.Call(
        uintptr(x),
        uintptr(y),
        uintptr(width),
        uintptr(height),
        uintptr(thickness),
        uintptr(color),
        uintptr(unsafe.Pointer(&err)),
    )
    return ret, err
}

// extern int ZBRGDIDrawEllipse(int x, int y, int width, int height, float thickness, int color, out int err);
func ZBRGDIDrawEllipse(x uint, y uint, width uint, height uint, thickness float32, color uint32) (ret uintptr, err uint) {
    ret, _, _ =  zBRGDIDrawEllipse.Call(
        uintptr(x),
        uintptr(y),
        uintptr(width),
        uintptr(height),
        uintptr(thickness),
        uintptr(color),
        uintptr(unsafe.Pointer(&err)),
    )
    return ret, err
}

// extern int ZBRGDIDrawBarCode(int x, int y, int rotation, int barcodetype, int barwidthratio, int barcodemultiplier, int barcodeheight, int textunder, byte[] barcodedata, out int err);
func ZBRGDIDrawBarCode(x uint, y uint, rotation BarCodeRotation, barcodetype BarCodeType, barwidthratio BarCodeWidth, barcodemultiplier uint, barcodeheight uint, textunder BarCodeTextUnder, barcodedata string) (ret uintptr,err uint) {
    ret, _, _ = zBRGDIDrawBarCode.Call(
            uintptr(x),
            uintptr(y),
            uintptr(rotation),
            uintptr(barcodetype),
            uintptr(barwidthratio),
            uintptr(barcodemultiplier),
            uintptr(barcodeheight),
            uintptr(textunder),
            uintptr(unsafe.Pointer(&([]byte(barcodedata))[0])),
            uintptr(unsafe.Pointer(&err)),
    )
    return ret, err
}

