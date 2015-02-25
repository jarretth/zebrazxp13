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

//int
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

func ZBRGDIGetSDKVsn() () {
    if zBRGDIGetSDKVsn == 0 {
        panic("ZBRGDIGetSDKVsn not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIGetSDKVsn not implemented")
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
    return
}
func ZBRGDIGetSDKProductVer() () {
    if zBRGDIGetSDKProductVer == 0 {
        panic("ZBRGDIGetSDKProductVer not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIGetSDKProductVer not implemented")
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
    return
}
func ZBRGDIInitGraphics() () {
    if zBRGDIInitGraphics == 0 {
        panic("ZBRGDIInitGraphics not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIInitGraphics not implemented")
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
    return
}
func ZBRGDIInitGraphicsEx() () {
    if zBRGDIInitGraphicsEx == 0 {
        panic("ZBRGDIInitGraphicsEx not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIInitGraphicsEx not implemented")
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
    return
}
func ZBRGDIInitGraphicsFromPrintDlg() () {
    if zBRGDIInitGraphicsFromPrintDlg == 0 {
        panic("ZBRGDIInitGraphicsFromPrintDlg not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIInitGraphicsFromPrintDlg not implemented")
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
    return
}
func ZBRGDICloseGraphics() () {
    if zBRGDICloseGraphics == 0 {
        panic("ZBRGDICloseGraphics not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDICloseGraphics not implemented")
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
    return
}
func ZBRGDIClearGraphics() () {
    if zBRGDIClearGraphics == 0 {
        panic("ZBRGDIClearGraphics not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIClearGraphics not implemented")
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
    return
}
func ZBRGDIStartPage() () {
    if zBRGDIStartPage == 0 {
        panic("ZBRGDIStartPage not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIStartPage not implemented")
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
    return
}
func ZBRGDIEndPage() () {
    if zBRGDIEndPage == 0 {
        panic("ZBRGDIEndPage not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIEndPage not implemented")
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
    return
}
func ZBRGDIPreviewGraphics() () {
    if zBRGDIPreviewGraphics == 0 {
        panic("ZBRGDIPreviewGraphics not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIPreviewGraphics not implemented")
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
    return
}
func ZBRGDIPrintGraphics() () {
    if zBRGDIPrintGraphics == 0 {
        panic("ZBRGDIPrintGraphics not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIPrintGraphics not implemented")
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
    return
}
func ZBRGDIPrintGraphicsEx() () {
    if zBRGDIPrintGraphicsEx == 0 {
        panic("ZBRGDIPrintGraphicsEx not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIPrintGraphicsEx not implemented")
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
    return
}
func ZBRGDIPrintFilePos() () {
    if zBRGDIPrintFilePos == 0 {
        panic("ZBRGDIPrintFilePos not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIPrintFilePos not implemented")
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
    return
}
func ZBRGDIPrintFileRect() () {
    if zBRGDIPrintFileRect == 0 {
        panic("ZBRGDIPrintFileRect not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIPrintFileRect not implemented")
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
    return
}
func ZBRGDIPrintImagePos() () {
    if zBRGDIPrintImagePos == 0 {
        panic("ZBRGDIPrintImagePos not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIPrintImagePos not implemented")
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
    return
}
func ZBRGDIPrintImageRect() () {
    if zBRGDIPrintImageRect == 0 {
        panic("ZBRGDIPrintImageRect not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIPrintImageRect not implemented")
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
    return
}
func ZBRGDICancelJob() () {
    if zBRGDICancelJob == 0 {
        panic("ZBRGDICancelJob not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDICancelJob not implemented")
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
    return
}
func ZBRGDIDrawText() () {
    if zBRGDIDrawText == 0 {
        panic("ZBRGDIDrawText not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIDrawText not implemented")
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
    return
}
func ZBRGDIDrawTextEx() () {
    if zBRGDIDrawTextEx == 0 {
        panic("ZBRGDIDrawTextEx not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIDrawTextEx not implemented")
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
    return
}
func ZBRGDIDrawTextUnicode() () {
    if zBRGDIDrawTextUnicode == 0 {
        panic("ZBRGDIDrawTextUnicode not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIDrawTextUnicode not implemented")
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
    return
}
func ZBRGDIDrawTextUnicodeEx() () {
    if zBRGDIDrawTextUnicodeEx == 0 {
        panic("ZBRGDIDrawTextUnicodeEx not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIDrawTextUnicodeEx not implemented")
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
    return
}
func ZBRGDIDrawTextRect() () {
    if zBRGDIDrawTextRect == 0 {
        panic("ZBRGDIDrawTextRect not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIDrawTextRect not implemented")
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
    return
}
func ZBRGDIDrawTextRectEx() () {
    if zBRGDIDrawTextRectEx == 0 {
        panic("ZBRGDIDrawTextRectEx not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIDrawTextRectEx not implemented")
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
    return
}
func ZBRGDIDrawLine() () {
    if zBRGDIDrawLine == 0 {
        panic("ZBRGDIDrawLine not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIDrawLine not implemented")
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
    return
}
func ZBRGDIDrawImage() () {
    if zBRGDIDrawImage == 0 {
        panic("ZBRGDIDrawImage not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIDrawImage not implemented")
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
    return
}
func ZBRGDIDrawImageEx() () {
    if zBRGDIDrawImageEx == 0 {
        panic("ZBRGDIDrawImageEx not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIDrawImageEx not implemented")
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
    return
}
func ZBRGDIDrawImagePos() () {
    if zBRGDIDrawImagePos == 0 {
        panic("ZBRGDIDrawImagePos not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIDrawImagePos not implemented")
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
    return
}
func ZBRGDIDrawImagePosEx() () {
    if zBRGDIDrawImagePosEx == 0 {
        panic("ZBRGDIDrawImagePosEx not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIDrawImagePosEx not implemented")
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
    return
}
func ZBRGDIDrawImageRect() () {
    if zBRGDIDrawImageRect == 0 {
        panic("ZBRGDIDrawImageRect not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIDrawImageRect not implemented")
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
    return
}
func ZBRGDIDrawImageRectEx() () {
    if zBRGDIDrawImageRectEx == 0 {
        panic("ZBRGDIDrawImageRectEx not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIDrawImageRectEx not implemented")
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
    return
}
func ZBRGDIDrawRectangle() () {
    if zBRGDIDrawRectangle == 0 {
        panic("ZBRGDIDrawRectangle not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIDrawRectangle not implemented")
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
    return
}
func ZBRGDIDrawEllipse() () {
    if zBRGDIDrawEllipse == 0 {
        panic("ZBRGDIDrawEllipse not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIDrawEllipse not implemented")
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
    return
}
func ZBRGDIDrawBarCode() () {
    if zBRGDIDrawBarCode == 0 {
        panic("ZBRGDIDrawBarCode not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRGDIDrawBarCode not implemented")
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
    return
}
