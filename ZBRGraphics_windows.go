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
func ZBRGDIIsPrinterReady(strPrinterName string) (success SuccessReturn,err uint) {
        ret, _, _ := zBRGDIIsPrinterReady.Call(
                uintptr(unsafe.Pointer(&([]byte(strPrinterName))[0])),
                uintptr(unsafe.Pointer(&err)),
        )
        return SuccessReturn(ret), err
}

// extern void ZBRGDIGetSDKVsn(out int major, out int minor, out int englevel);
func ZBRGDIGetSDKVsn() (success SuccessReturn) {
    panic("ZBRGDIGetSDKVsn not implemented")
    ret, _, _ :=  zBRGDIGetSDKVsn.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDIGetSDKProductVer(byte[] productversion, out int buffsize, out int err);
func ZBRGDIGetSDKProductVer() (success SuccessReturn) {
    panic("ZBRGDIGetSDKProductVer not implemented")
    ret, _, _ :=  zBRGDIGetSDKProductVer.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDIInitGraphics(byte[] printername, out IntPtr hdc, out int err);
func ZBRGDIInitGraphics(strPrinterName string) (success SuccessReturn, handle syscall.Handle, err uint) {
        ret, _, _ := zBRGDIInitGraphics.Call(
                uintptr(unsafe.Pointer(&([]byte(strPrinterName))[0])),
                uintptr(unsafe.Pointer(&handle)),
                uintptr(unsafe.Pointer(&err)),
        )
        return SuccessReturn(ret), handle, err
}

// extern int ZBRGDIInitGraphicsEx(byte[] printername, out IntPtr hdc, byte[] jobname, out int jobid, out int err);
func ZBRGDIInitGraphicsEx() (success SuccessReturn) {
    panic("ZBRGDIInitGraphicsEx not implemented")
    ret, _, _ :=  zBRGDIInitGraphicsEx.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDIInitGraphicsFromPrintDlg(out IntPtr hdc, out int err);
func ZBRGDIInitGraphicsFromPrintDlg() (success SuccessReturn) {
    panic("ZBRGDIInitGraphicsFromPrintDlg not implemented")
    ret, _, _ :=  zBRGDIInitGraphicsFromPrintDlg.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDICloseGraphics(IntPtr hdc, out int err);
func ZBRGDICloseGraphics(handle syscall.Handle) (success SuccessReturn, err uint){
    ret, _, _ := zBRGDICloseGraphics.Call(
        uintptr(handle),
        uintptr(unsafe.Pointer(&err)),
    )
    return SuccessReturn(ret), err
}

// extern int ZBRGDIClearGraphics(out int err);
func ZBRGDIClearGraphics() (success SuccessReturn, err uint) {
    ret, _, _ := zBRGDIClearGraphics.Call(
        uintptr(unsafe.Pointer(&err)),
    )
    return SuccessReturn(ret), err
}

// extern int ZBRGDIStartPage(IntPtr hdc, out int err);
func ZBRGDIStartPage() (success SuccessReturn) {
    panic("ZBRGDIStartPage not implemented")
    ret, _, _ :=  zBRGDIStartPage.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDIEndPage(IntPtr hdc, out int err);
func ZBRGDIEndPage() (success SuccessReturn) {
    panic("ZBRGDIEndPage not implemented")
    ret, _, _ :=  zBRGDIEndPage.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDIPreviewGraphics(IntPtr hwnd, out int err);
func ZBRGDIPreviewGraphics() (success SuccessReturn) {
    panic("ZBRGDIPreviewGraphics not implemented")
    ret, _, _ :=  zBRGDIPreviewGraphics.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDIPrintGraphics(IntPtr hdc, out int err);
func ZBRGDIPrintGraphics(handle syscall.Handle) (success SuccessReturn, err uint) {
        ret, _, _ := zBRGDIPrintGraphics.Call(
                uintptr(handle),
                uintptr(unsafe.Pointer(&err)),
        )
        return SuccessReturn(ret), err
}

// extern int ZBRGDIPrintGraphicsEx(IntPtr hdc, out int err);
func ZBRGDIPrintGraphicsEx() (success SuccessReturn) {
    panic("ZBRGDIPrintGraphicsEx not implemented")
    ret, _, _ :=  zBRGDIPrintGraphicsEx.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDIPrintFilePos(IntPtr hdc, byte[] filename, int position, out int err);
func ZBRGDIPrintFilePos() (success SuccessReturn) {
    panic("ZBRGDIPrintFilePos not implemented")
    ret, _, _ :=  zBRGDIPrintFilePos.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDIPrintImagePos(IntPtr hdc, byte[] imagedata, int imagesize, int position, out int err);
func ZBRGDIPrintImagePos() (success SuccessReturn) {
    panic("ZBRGDIPrintImagePos not implemented")
    ret, _, _ :=  zBRGDIPrintImagePos.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDIPrintImageRect(IntPtr hdc, byte[] imagedata, int imagesize, int x, int y, int width, int height, out int err);
func ZBRGDIPrintImageRect() (success SuccessReturn) {
    panic("ZBRGDIPrintImageRect not implemented")
    ret, _, _ :=  zBRGDIPrintImageRect.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDICancelJob(byte[] printername, int jobid, out int err);
func ZBRGDICancelJob() (success SuccessReturn) {
    panic("ZBRGDICancelJob not implemented")
    ret, _, _ :=  zBRGDICancelJob.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDIDrawText(int x, int y, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawText(x uint, y uint, text string, font string, fontsize uint, fontstyle TextFontStyle, color uint32) (success SuccessReturn, err uint) {
    ret, _, _ :=  zBRGDIDrawText.Call(
        uintptr(x),
        uintptr(y),
        uintptr(unsafe.Pointer(&([]byte(text))[0])),
        uintptr(unsafe.Pointer(&([]byte(font))[0])),
        uintptr(fontsize),
        uintptr(fontstyle),
        uintptr(color),
        uintptr(unsafe.Pointer(&err)),
    )
    return SuccessReturn(ret), err
}

// extern int ZBRGDIDrawTextEx(int x, int y, int angle, int alignment, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawTextEx() (success SuccessReturn) {
    panic("ZBRGDIDrawTextEx not implemented")
    ret, _, _ :=  zBRGDIDrawTextEx.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDIDrawTextUnicode(int x, int y, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawTextUnicode(x uint, y uint, text string, font string, fontsize uint, fontstyle TextFontStyle, color uint32) (success SuccessReturn, err uint) {
    ret, _, _ :=  zBRGDIDrawTextUnicode.Call(
        uintptr(x),
        uintptr(y),
        uintptr(unsafe.Pointer(&([]byte(text))[0])),
        uintptr(unsafe.Pointer(&([]byte(font))[0])),
        uintptr(fontsize),
        uintptr(fontstyle),
        uintptr(color),
        uintptr(unsafe.Pointer(&err)),
    )
    return SuccessReturn(ret), err
}

// extern int ZBRGDIDrawTextUnicodeEx(int x, int y, int angle, int alignment, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawTextUnicodeEx() (success SuccessReturn) {
    panic("ZBRGDIDrawTextUnicodeEx not implemented")
    ret, _, _ :=  zBRGDIDrawTextUnicodeEx.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDIDrawTextRect(int x, int y, int width, int height, int alignment, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawTextRect(x uint, y uint, width uint, height uint, alignment TextAlignment, text string, font string, fontsize uint, fontstyle TextFontStyle, color uint32) (success SuccessReturn, err uint) {
    ret, _, _ :=  zBRGDIDrawTextRect.Call(
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
    return SuccessReturn(ret), err
}

// extern int ZBRGDIDrawTextRectEx(int x, int y, int width, int height, int angle, int alignment, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawTextRectEx() (success SuccessReturn) {
    panic("ZBRGDIDrawTextRectEx not implemented")
    ret, _, _ :=  zBRGDIDrawTextRectEx.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDIDrawLine(int x1, int y1, int x2, int y2, int color, float thickness, out int err);
func ZBRGDIDrawLine(x1 uint, y1 uint, x2 uint, y2 uint, color uint32, thickness float32) (success SuccessReturn, err uint) {
    ret, _, _ :=  zBRGDIDrawLine.Call(
        uintptr(x1),
        uintptr(y1),
        uintptr(x2),
        uintptr(y2),
        uintptr(color),
        uintptr(thickness),
        uintptr(unsafe.Pointer(&err)),
    )
    return SuccessReturn(ret), err
}

// extern int ZBRGDIDrawImage(byte[] filename, int x, int y, out int err);
func ZBRGDIDrawImage() (success SuccessReturn) {
    panic("ZBRGDIDrawImage not implemented")
    ret, _, _ :=  zBRGDIDrawImage.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDIDrawImageEx(byte[] filename, int imagesize, int x, int y, out int err);
func ZBRGDIDrawImageEx() (success SuccessReturn) {
    panic("ZBRGDIDrawImageEx not implemented")
    ret, _, _ :=  zBRGDIDrawImageEx.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDIDrawImagePos(byte[] filename, int position, out int err);
func ZBRGDIDrawImagePos() (success SuccessReturn) {
    panic("ZBRGDIDrawImagePos not implemented")
    ret, _, _ :=  zBRGDIDrawImagePos.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDIDrawImagePosEx(byte[] filename, int imagesize, int position, out int err);
func ZBRGDIDrawImagePosEx() (success SuccessReturn) {
    panic("ZBRGDIDrawImagePosEx not implemented")
    ret, _, _ :=  zBRGDIDrawImagePosEx.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDIDrawImageRect(byte[] filename, int x, int y, int width, int height, out int err);
func ZBRGDIDrawImageRect() (success SuccessReturn) {
    panic("ZBRGDIDrawImageRect not implemented")
    ret, _, _ :=  zBRGDIDrawImageRect.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDIDrawImageRectEx(byte[] filename, int imagesize, int x, int y, int width, int height, out int err);
func ZBRGDIDrawImageRectEx() (success SuccessReturn) {
    panic("ZBRGDIDrawImageRectEx not implemented")
    ret, _, _ :=  zBRGDIDrawImageRectEx.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRGDIDrawRectangle(int x, int y, int width, int height, float thickness, int color, out int err);
func ZBRGDIDrawRectangle(x uint, y uint, width uint, height uint, thickness float32, color uint32) (success SuccessReturn, err uint) {
    ret, _, _ :=  zBRGDIDrawRectangle.Call(
        uintptr(x),
        uintptr(y),
        uintptr(width),
        uintptr(height),
        uintptr(thickness),
        uintptr(color),
        uintptr(unsafe.Pointer(&err)),
    )
    return SuccessReturn(ret), err
}

// extern int ZBRGDIDrawEllipse(int x, int y, int width, int height, float thickness, int color, out int err);
func ZBRGDIDrawEllipse(x uint, y uint, width uint, height uint, thickness float32, color uint32) (success SuccessReturn, err uint) {
    ret, _, _ :=  zBRGDIDrawEllipse.Call(
        uintptr(x),
        uintptr(y),
        uintptr(width),
        uintptr(height),
        uintptr(thickness),
        uintptr(color),
        uintptr(unsafe.Pointer(&err)),
    )
    return SuccessReturn(ret), err
}

// extern int ZBRGDIDrawBarCode(int x, int y, int rotation, int barcodetype, int barwidthratio, int barcodemultiplier, int barcodeheight, int textunder, byte[] barcodedata, out int err);
func ZBRGDIDrawBarCode(x uint, y uint, rotation BarCodeRotation, barcodetype BarCodeType, barwidthratio BarCodeWidth, barcodemultiplier uint, barcodeheight uint, textunder BarCodeTextUnder, barcodedata string) (success SuccessReturn,err uint) {
    ret, _, _ := zBRGDIDrawBarCode.Call(
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
    return SuccessReturn(ret), err
}

