package zbrgraphics


func abort(funcname string, err error) {
        panic(fmt.Sprintf("%s failed: %v", funcname, err))
}

var (
    zBRGraphics                    , _ = syscall.LoadLibrary("ZBRGraphics.dll")
    zBRGDICancelJob                , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDICancelJob")
    zBRGDIClearGraphics            , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIClearGraphics")
    zBRGDICloseGraphics            , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDICloseGraphics")
    zBRGDIDrawBarCode              , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIDrawBarCode")
    zBRGDIDrawEllipse              , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIDrawEllipse")
    zBRGDIDrawImage                , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIDrawImage")
    zBRGDIDrawImageEx              , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIDrawImageEx")
    zBRGDIDrawImagePos             , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIDrawImagePos")
    zBRGDIDrawImagePosEx           , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIDrawImagePosEx")
    zBRGDIDrawImageRect            , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIDrawImageRect")
    zBRGDIDrawImageRectEx          , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIDrawImageRectEx")
    zBRGDIDrawLine                 , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIDrawLine")
    zBRGDIDrawRectangle            , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIDrawRectangle")
    zBRGDIDrawText                 , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIDrawText")
    zBRGDIDrawTextEx               , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIDrawTextEx")
    zBRGDIDrawTextRect             , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIDrawTextRect")
    zBRGDIDrawTextRectEx           , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIDrawTextRectEx")
    zBRGDIDrawTextUnicode          , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIDrawTextUnicode")
    zBRGDIDrawTextUnicodeEx        , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIDrawTextUnicodeEx")
    zBRGDIEndPage                  , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIEndPage")
    zBRGDIGetSDKProductVer         , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIGetSDKProductVer")
    zBRGDIGetSDKVer                , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIGetSDKVer")
    zBRGDIGetSDKVsn                , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIGetSDKVsn")
    zBRGDIInitGraphics             , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIInitGraphics")
    zBRGDIInitGraphicsEx           , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIInitGraphicsEx")
    zBRGDIInitGraphicsFromPrintDlg , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIInitGraphicsFromPrintDlg")
    zBRGDIIsPrinterReady           , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIIsPrinterReady")
    zBRGDIPreviewGraphics          , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIPreviewGraphics")
    zBRGDIPrintFilePos             , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIPrintFilePos")
    zBRGDIPrintFileRect            , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIPrintFileRect")
    zBRGDIPrintGraphics            , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIPrintGraphics")
    zBRGDIPrintGraphicsEx          , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIPrintGraphicsEx")
    zBRGDIPrintImagePos            , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIPrintImagePos")
    zBRGDIPrintImageRect           , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIPrintImageRect")
    zBRGDIStartPage                , _ = syscall.GetProcAddress(zBRGraphics, "ZBRGDIStartPage")
)

const (
    ZBR_GDI_ERROR_GENERIC_ERROR                = 8001
    ZBR_GDI_ERROR_INVALID_PARAMETER            = 8002
    ZBR_GDI_ERROR_OUT_OF_MEMORY                = 8003
    ZBR_GDI_ERROR_OBJECT_BUSY                  = 8004
    ZBR_GDI_ERROR_INSUFFICIENT_BUFFER          = 8005
    ZBR_GDI_ERROR_NOT_IMPLEMENTED              = 8006
    ZBR_GDI_ERROR_WIN32_ERROR                  = 8007
    ZBR_GDI_ERROR_WRONG_STATE                  = 8008
    ZBR_GDI_ERROR_ABORTED                      = 8009
    ZBR_GDI_ERROR_FILE_NOT_FOUND               = 8010
    ZBR_GDI_ERROR_VALUE_OVERFLOW               = 8011
    ZBR_GDI_ERROR_ACCESS_DENIED                = 8012
    ZBR_GDI_ERROR_UNKNOWN_IMAGE_FORMAT         = 8013
    ZBR_GDI_ERROR_FONT_FAMILY_NOT_FOUND        = 8014
    ZBR_GDI_ERROR_FONT_STYLE_NOT_FOUND         = 8015
    ZBR_GDI_ERROR_NOT_TRUE_TYPE_FONT           = 8016
    ZBR_GDI_ERROR_UNSUPPORTED_GDIPLUS_         = 8017
    ZBR_GDI_ERROR_GDIPLUS_NOT_INITIALIZED      = 8018
    ZBR_GDI_ERROR_PROPERTY_NOT_FOUND           = 8019
    ZBR_GDI_ERROR_PROPERTY_NOT_SUPPORTED       = 8020
    ZBR_GDI_ERROR_GRAPHICS_ALREADY_INITIALIZED = 8021
    ZBR_GDI_ERROR_NO_GRAPHIC_DATA              = 8022
    ZBR_GDI_ERROR_GRAPHICS_NOT_INITIALIZED     = 8023
    ZBR_GDI_ERROR_GETTING_DEVICE_CONTEXT       = 8024
    ZBR_DLG_ERROR_DLG_CANCELED                 = 8025
    ZBR_DLG_ERROR_SETUP_FAILURE                = 8026
    ZBR_DLG_ERROR_PARSE_FAILURE                = 8027
    ZBR_DLG_ERROR_RET_DEFAULT_FAILURE          = 8028
    ZBR_DLG_ERROR_LOAD_DRV_FAILURE             = 8029
    ZBR_DLG_ERROR_GET_DEVMODE_FAIL             = 8030
    ZBR_DLG_ERROR_INIT_FAILURE                 = 8031
    ZBR_DLG_ERROR_NO_DEVICES                   = 8032
    ZBR_DLG_ERROR_NO_DEFAULT_PRINTER           = 8033
    ZBR_DLG_ERROR_DN_DM_MISMATCH               = 8034
    ZBR_DLG_ERROR_CREATE_IC_FAILURE            = 8035
    ZBR_DLG_ERROR_PRINTER_NOT_FOUND            = 8036
    ZBR_DLG_ERROR_DEFAULT_DIFFERENT            = 8037
)

// extern void ZBRGDIGetSDKVer(out int major, out int minor, out int englevel);
func ZBRGDIGetSDKVer() (major uint, minor uint, engLevel uint) {
    if zBRGDIGetSDKVer == 0 {
        panic("ZBRGDIGetSDKVer not defined. Check library")
    }
    var nargs uintptr = 3
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIGetSDKVer),
            nargs,
            uintptr(unsafe.Pointer(&major)),
            uintptr(unsafe.Pointer(&minor)),
            uintptr(unsafe.Pointer(&engLevel)),
            0,
            0,
            0,
            0,
            0,
            0)
    if callErr != 0 {
            abort("Call ZBRGDIGetSDKVer", callErr)
    }
    return major,minor,engLevel
}

// extern int ZBRGDIIsPrinterReady(byte[] printername, out int err);
func ZBRGDIIsPrinterReady(strPrinterName []byte) (ret uintptr, err uint) {
    if zBRGDIIsPrinterReady == 0 {
        panic("ZBRGDIIsPrinterReady not defined. Check library")
    }
    var nargs uintptr = 2
    ret, _, callErr := syscall.Syscall9(uintptr(zBRGDIIsPrinterReady),
            nargs,
            // uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(strPrinterName))),
            uintptr(unsafe.Pointer(&strPrinterName)),
            uintptr(err),
            0,
            0,
            0,
            0,
            0,
            0,
            0)
    if callErr != 0 {
            abort("Call ZBRGDIIsPrinterReady", callErr)
    }
    return ret, err
}

// extern void ZBRGDIGetSDKVsn(out int major, out int minor, out int englevel);
func ZBRGDIGetSDKVsn() () {
    if zBRGDIGetSDKVsn == 0 {
        panic("ZBRGDIGetSDKVsn not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIGetSDKVsn is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIGetSDKVsn),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIGetSDKVsn",callErr)
    }
    return
}

// extern int ZBRGDIGetSDKProductVer(byte[] productversion, out int buffsize, out int err);
func ZBRGDIGetSDKProductVer() () {
    if zBRGDIGetSDKProductVer == 0 {
        panic("ZBRGDIGetSDKProductVer not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIGetSDKProductVer is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIGetSDKProductVer),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIGetSDKProductVer",callErr)
    }
    return
}

// extern int ZBRGDIInitGraphics(byte[] printername, out IntPtr hdc, out int err);
func ZBRGDIInitGraphics() () {
    if zBRGDIInitGraphics == 0 {
        panic("ZBRGDIInitGraphics not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIInitGraphics is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIInitGraphics),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIInitGraphics",callErr)
    }
    return
}

// extern int ZBRGDIInitGraphicsEx(byte[] printername, out IntPtr hdc, byte[] jobname, out int jobid, out int err);
func ZBRGDIInitGraphicsEx() () {
    if zBRGDIInitGraphicsEx == 0 {
        panic("ZBRGDIInitGraphicsEx not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIInitGraphicsEx is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIInitGraphicsEx),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIInitGraphicsEx",callErr)
    }
    return
}

// extern int ZBRGDIInitGraphicsFromPrintDlg(out IntPtr hdc, out int err);
func ZBRGDIInitGraphicsFromPrintDlg() () {
    if zBRGDIInitGraphicsFromPrintDlg == 0 {
        panic("ZBRGDIInitGraphicsFromPrintDlg not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIInitGraphicsFromPrintDlg is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIInitGraphicsFromPrintDlg),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIInitGraphicsFromPrintDlg",callErr)
    }
    return
}

// extern int ZBRGDICloseGraphics(IntPtr hdc, out int err);
func ZBRGDICloseGraphics() () {
    if zBRGDICloseGraphics == 0 {
        panic("ZBRGDICloseGraphics not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDICloseGraphics is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDICloseGraphics),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDICloseGraphics",callErr)
    }
    return
}

// extern int ZBRGDIClearGraphics(out int err);
func ZBRGDIClearGraphics() () {
    if zBRGDIClearGraphics == 0 {
        panic("ZBRGDIClearGraphics not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIClearGraphics is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIClearGraphics),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIClearGraphics",callErr)
    }
    return
}

// extern int ZBRGDIStartPage(IntPtr hdc, out int err);
func ZBRGDIStartPage() () {
    if zBRGDIStartPage == 0 {
        panic("ZBRGDIStartPage not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIStartPage is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIStartPage),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIStartPage",callErr)
    }
    return
}

// extern int ZBRGDIEndPage(IntPtr hdc, out int err);
func ZBRGDIEndPage() () {
    if zBRGDIEndPage == 0 {
        panic("ZBRGDIEndPage not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIEndPage is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIEndPage),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIEndPage",callErr)
    }
    return
}

// extern int ZBRGDIPreviewGraphics(IntPtr hwnd, out int err);
func ZBRGDIPreviewGraphics() () {
    if zBRGDIPreviewGraphics == 0 {
        panic("ZBRGDIPreviewGraphics not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIPreviewGraphics is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIPreviewGraphics),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIPreviewGraphics",callErr)
    }
    return
}

// extern int ZBRGDIPrintGraphics(IntPtr hdc, out int err);
func ZBRGDIPrintGraphics() () {
    if zBRGDIPrintGraphics == 0 {
        panic("ZBRGDIPrintGraphics not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIPrintGraphics is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIPrintGraphics),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIPrintGraphics",callErr)
    }
    return
}

// extern int ZBRGDIPrintGraphicsEx(IntPtr hdc, out int err);
func ZBRGDIPrintGraphicsEx() () {
    if zBRGDIPrintGraphicsEx == 0 {
        panic("ZBRGDIPrintGraphicsEx not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIPrintGraphicsEx is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIPrintGraphicsEx),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIPrintGraphicsEx",callErr)
    }
    return
}

// extern int ZBRGDIPrintFilePos(IntPtr hdc, byte[] filename, int position, out int err);
func ZBRGDIPrintFilePos() () {
    if zBRGDIPrintFilePos == 0 {
        panic("ZBRGDIPrintFilePos not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIPrintFilePos is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIPrintFilePos),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIPrintFilePos",callErr)
    }
    return
}

// extern int ZBRGDIPrintFilePos(IntPtr hdc, byte[] filename, int position, out int err);
func ZBRGDIPrintFileRect() () {
    if zBRGDIPrintFileRect == 0 {
        panic("ZBRGDIPrintFileRect not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIPrintFileRect is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIPrintFileRect),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIPrintFileRect",callErr)
    }
    return
}

// extern int ZBRGDIPrintImagePos(IntPtr hdc, byte[] imagedata, int imagesize, int position, out int err);
func ZBRGDIPrintImagePos() () {
    if zBRGDIPrintImagePos == 0 {
        panic("ZBRGDIPrintImagePos not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIPrintImagePos is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIPrintImagePos),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIPrintImagePos",callErr)
    }
    return
}

// extern int ZBRGDIPrintImageRect(IntPtr hdc, byte[] imagedata, int imagesize, int x, int y, int width, int height, out int err);
func ZBRGDIPrintImageRect() () {
    if zBRGDIPrintImageRect == 0 {
        panic("ZBRGDIPrintImageRect not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIPrintImageRect is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIPrintImageRect),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIPrintImageRect",callErr)
    }
    return
}

// extern int ZBRGDICancelJob(byte[] printername, int jobid, out int err);
func ZBRGDICancelJob() () {
    if zBRGDICancelJob == 0 {
        panic("ZBRGDICancelJob not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDICancelJob is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDICancelJob),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDICancelJob",callErr)
    }
    return
}

// extern int ZBRGDIDrawText(int x, int y, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawText() () {
    if zBRGDIDrawText == 0 {
        panic("ZBRGDIDrawText not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIDrawText is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIDrawText),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIDrawText",callErr)
    }
    return
}

// extern int ZBRGDIDrawTextEx(int x, int y, int angle, int alignment, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawTextEx() () {
    if zBRGDIDrawTextEx == 0 {
        panic("ZBRGDIDrawTextEx not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIDrawTextEx is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIDrawTextEx),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIDrawTextEx",callErr)
    }
    return
}

// extern int ZBRGDIDrawTextUnicode(int x, int y, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawTextUnicode() () {
    if zBRGDIDrawTextUnicode == 0 {
        panic("ZBRGDIDrawTextUnicode not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIDrawTextUnicode is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIDrawTextUnicode),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIDrawTextUnicode",callErr)
    }
    return
}

// extern int ZBRGDIDrawTextUnicodeEx(int x, int y, int angle, int alignment, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawTextUnicodeEx() () {
    if zBRGDIDrawTextUnicodeEx == 0 {
        panic("ZBRGDIDrawTextUnicodeEx not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIDrawTextUnicodeEx is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIDrawTextUnicodeEx),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIDrawTextUnicodeEx",callErr)
    }
    return
}

// extern int ZBRGDIDrawTextRect(int x, int y, int width, int height, int alignment, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawTextRect() () {
    if zBRGDIDrawTextRect == 0 {
        panic("ZBRGDIDrawTextRect not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIDrawTextRect is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIDrawTextRect),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIDrawTextRect",callErr)
    }
    return
}

// extern int ZBRGDIDrawTextRectEx(int x, int y, int width, int height, int angle, int alignment, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawTextRectEx() () {
    if zBRGDIDrawTextRectEx == 0 {
        panic("ZBRGDIDrawTextRectEx not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIDrawTextRectEx is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIDrawTextRectEx),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIDrawTextRectEx",callErr)
    }
    return
}

// extern int ZBRGDIDrawLine(int x1, int y1, int x2, int y2, int color, float thickness, out int err);
func ZBRGDIDrawLine() () {
    if zBRGDIDrawLine == 0 {
        panic("ZBRGDIDrawLine not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIDrawLine is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIDrawLine),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIDrawLine",callErr)
    }
    return
}

// extern int ZBRGDIDrawImage(byte[] filename, int x, int y, out int err);
func ZBRGDIDrawImage() () {
    if zBRGDIDrawImage == 0 {
        panic("ZBRGDIDrawImage not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIDrawImage is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIDrawImage),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIDrawImage",callErr)
    }
    return
}

// extern int ZBRGDIDrawImageEx(byte[] filename, int imagesize, int x, int y, out int err);
func ZBRGDIDrawImageEx() () {
    if zBRGDIDrawImageEx == 0 {
        panic("ZBRGDIDrawImageEx not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIDrawImageEx is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIDrawImageEx),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIDrawImageEx",callErr)
    }
    return
}

// extern int ZBRGDIDrawImagePos(byte[] filename, int position, out int err);
func ZBRGDIDrawImagePos() () {
    if zBRGDIDrawImagePos == 0 {
        panic("ZBRGDIDrawImagePos not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIDrawImagePos is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIDrawImagePos),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIDrawImagePos",callErr)
    }
    return
}

// extern int ZBRGDIDrawImagePosEx(byte[] filename, int imagesize, int position, out int err);
func ZBRGDIDrawImagePosEx() () {
    if zBRGDIDrawImagePosEx == 0 {
        panic("ZBRGDIDrawImagePosEx not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIDrawImagePosEx is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIDrawImagePosEx),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIDrawImagePosEx",callErr)
    }
    return
}

// extern int ZBRGDIDrawImageRect(byte[] filename, int x, int y, int width, int height, out int err);
func ZBRGDIDrawImageRect() () {
    if zBRGDIDrawImageRect == 0 {
        panic("ZBRGDIDrawImageRect not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIDrawImageRect is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIDrawImageRect),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIDrawImageRect",callErr)
    }
    return
}

// extern int ZBRGDIDrawImageRectEx(byte[] filename, int imagesize, int x, int y, int width, int height, out int err);
func ZBRGDIDrawImageRectEx() () {
    if zBRGDIDrawImageRectEx == 0 {
        panic("ZBRGDIDrawImageRectEx not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIDrawImageRectEx is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIDrawImageRectEx),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIDrawImageRectEx",callErr)
    }
    return
}

// extern int ZBRGDIDrawRectangle(int x, int y, int width, int height, float thickness, int color, out int err);
func ZBRGDIDrawRectangle() () {
    if zBRGDIDrawRectangle == 0 {
        panic("ZBRGDIDrawRectangle not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIDrawRectangle is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIDrawRectangle),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIDrawRectangle",callErr)
    }
    return
}

// extern int ZBRGDIDrawEllipse(int x, int y, int width, int height, float thickness, int color, out int err);
func ZBRGDIDrawEllipse() () {
    if zBRGDIDrawEllipse == 0 {
        panic("ZBRGDIDrawEllipse not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIDrawEllipse is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIDrawEllipse),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIDrawEllipse",callErr)
    }
    return
}

// extern int ZBRGDIDrawBarcode(int x, int y, int rotation, int barcodetype, int barwidthratio, int barcodemultiplier, int barcodeheight, int textunder, byte[] barcodedata, out int err);
func ZBRGDIDrawBarCode() () {
    if zBRGDIDrawBarCode == 0 {
        panic("ZBRGDIDrawBarCode not defined. Check library")
    }
    var nargs uintptr = 0
    panic("Call ZBRGDIDrawBarCode is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRGDIDrawBarCode),
        nargs,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0)
    if callErr != 0 {
        abort("ZBRGDIDrawBarCode",callErr)
    }
    return
}
