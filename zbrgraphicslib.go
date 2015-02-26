package zbrgraphics

func abort(funcname string, err error) {
        panic(fmt.Sprintf("%s failed: %v", funcname, err))
}

var (
    zBRGraphics                    , _ = syscall.NewLazyDLL("ZBRGraphics.dll")
    zBRGDICancelJob                , _ = zBRGraphics.NewProc("ZBRGDICancelJob")
    zBRGDIClearGraphics            , _ = zBRGraphics.NewProc("ZBRGDIClearGraphics")
    zBRGDICloseGraphics            , _ = zBRGraphics.NewProc("ZBRGDICloseGraphics")
    zBRGDIDrawBarCode              , _ = zBRGraphics.NewProc("ZBRGDIDrawBarCode")
    zBRGDIDrawEllipse              , _ = zBRGraphics.NewProc("ZBRGDIDrawEllipse")
    zBRGDIDrawImage                , _ = zBRGraphics.NewProc("ZBRGDIDrawImage")
    zBRGDIDrawImageEx              , _ = zBRGraphics.NewProc("ZBRGDIDrawImageEx")
    zBRGDIDrawImagePos             , _ = zBRGraphics.NewProc("ZBRGDIDrawImagePos")
    zBRGDIDrawImagePosEx           , _ = zBRGraphics.NewProc("ZBRGDIDrawImagePosEx")
    zBRGDIDrawImageRect            , _ = zBRGraphics.NewProc("ZBRGDIDrawImageRect")
    zBRGDIDrawImageRectEx          , _ = zBRGraphics.NewProc("ZBRGDIDrawImageRectEx")
    zBRGDIDrawLine                 , _ = zBRGraphics.NewProc("ZBRGDIDrawLine")
    zBRGDIDrawRectangle            , _ = zBRGraphics.NewProc("ZBRGDIDrawRectangle")
    zBRGDIDrawText                 , _ = zBRGraphics.NewProc("ZBRGDIDrawText")
    zBRGDIDrawTextEx               , _ = zBRGraphics.NewProc("ZBRGDIDrawTextEx")
    zBRGDIDrawTextRect             , _ = zBRGraphics.NewProc("ZBRGDIDrawTextRect")
    zBRGDIDrawTextRectEx           , _ = zBRGraphics.NewProc("ZBRGDIDrawTextRectEx")
    zBRGDIDrawTextUnicode          , _ = zBRGraphics.NewProc("ZBRGDIDrawTextUnicode")
    zBRGDIDrawTextUnicodeEx        , _ = zBRGraphics.NewProc("ZBRGDIDrawTextUnicodeEx")
    zBRGDIEndPage                  , _ = zBRGraphics.NewProc("ZBRGDIEndPage")
    zBRGDIGetSDKProductVer         , _ = zBRGraphics.NewProc("ZBRGDIGetSDKProductVer")
    zBRGDIGetSDKVer                , _ = zBRGraphics.NewProc("ZBRGDIGetSDKVer")
    zBRGDIGetSDKVsn                , _ = zBRGraphics.NewProc("ZBRGDIGetSDKVsn")
    zBRGDIInitGraphics             , _ = zBRGraphics.NewProc("ZBRGDIInitGraphics")
    zBRGDIInitGraphicsEx           , _ = zBRGraphics.NewProc("ZBRGDIInitGraphicsEx")
    zBRGDIInitGraphicsFromPrintDlg , _ = zBRGraphics.NewProc("ZBRGDIInitGraphicsFromPrintDlg")
    zBRGDIIsPrinterReady           , _ = zBRGraphics.NewProc("ZBRGDIIsPrinterReady")
    zBRGDIPreviewGraphics          , _ = zBRGraphics.NewProc("ZBRGDIPreviewGraphics")
    zBRGDIPrintFilePos             , _ = zBRGraphics.NewProc("ZBRGDIPrintFilePos")
    zBRGDIPrintFileRect            , _ = zBRGraphics.NewProc("ZBRGDIPrintFileRect")
    zBRGDIPrintGraphics            , _ = zBRGraphics.NewProc("ZBRGDIPrintGraphics")
    zBRGDIPrintGraphicsEx          , _ = zBRGraphics.NewProc("ZBRGDIPrintGraphicsEx")
    zBRGDIPrintImagePos            , _ = zBRGraphics.NewProc("ZBRGDIPrintImagePos")
    zBRGDIPrintImageRect           , _ = zBRGraphics.NewProc("ZBRGDIPrintImageRect")
    zBRGDIStartPage                , _ = zBRGraphics.NewProc("ZBRGDIStartPage")
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
func ZBRGDIGetSDKVer() (ret uintptr) {
    panic("ZBRGDIGetSDKVer not implemented")
    ret, _, _ =  zBRGDIGetSDKVer.Call(
        )
    return ret
}
// extern int ZBRGDIIsPrinterReady(byte[] printername, out int err);
func ZBRGDIIsPrinterReady() (ret uintptr) {
    panic("ZBRGDIIsPrinterReady not implemented")
    ret, _, _ =  zBRGDIIsPrinterReady.Call(
        )
    return ret
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
func ZBRGDIInitGraphics() (ret uintptr) {
    panic("ZBRGDIInitGraphics not implemented")
    ret, _, _ =  zBRGDIInitGraphics.Call(
        )
    return ret
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
func ZBRGDICloseGraphics() (ret uintptr) {
    panic("ZBRGDICloseGraphics not implemented")
    ret, _, _ =  zBRGDICloseGraphics.Call(
        )
    return ret
}
// extern int ZBRGDIClearGraphics(out int err);
func ZBRGDIClearGraphics() (ret uintptr) {
    panic("ZBRGDIClearGraphics not implemented")
    ret, _, _ =  zBRGDIClearGraphics.Call(
        )
    return ret
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
func ZBRGDIPrintGraphics() (ret uintptr) {
    panic("ZBRGDIPrintGraphics not implemented")
    ret, _, _ =  zBRGDIPrintGraphics.Call(
        )
    return ret
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
func ZBRGDIDrawText() (ret uintptr) {
    panic("ZBRGDIDrawText not implemented")
    ret, _, _ =  zBRGDIDrawText.Call(
        )
    return ret
}
// extern int ZBRGDIDrawTextEx(int x, int y, int angle, int alignment, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawTextEx() (ret uintptr) {
    panic("ZBRGDIDrawTextEx not implemented")
    ret, _, _ =  zBRGDIDrawTextEx.Call(
        )
    return ret
}
// extern int ZBRGDIDrawTextUnicode(int x, int y, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawTextUnicode() (ret uintptr) {
    panic("ZBRGDIDrawTextUnicode not implemented")
    ret, _, _ =  zBRGDIDrawTextUnicode.Call(
        )
    return ret
}
// extern int ZBRGDIDrawTextUnicodeEx(int x, int y, int angle, int alignment, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawTextUnicodeEx() (ret uintptr) {
    panic("ZBRGDIDrawTextUnicodeEx not implemented")
    ret, _, _ =  zBRGDIDrawTextUnicodeEx.Call(
        )
    return ret
}
// extern int ZBRGDIDrawTextRect(int x, int y, int width, int height, int alignment, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawTextRect() (ret uintptr) {
    panic("ZBRGDIDrawTextRect not implemented")
    ret, _, _ =  zBRGDIDrawTextRect.Call(
        )
    return ret
}
// extern int ZBRGDIDrawTextRectEx(int x, int y, int width, int height, int angle, int alignment, byte[] text, byte[] font, int fontsize, int fontstyle, int color, out int err);
func ZBRGDIDrawTextRectEx() (ret uintptr) {
    panic("ZBRGDIDrawTextRectEx not implemented")
    ret, _, _ =  zBRGDIDrawTextRectEx.Call(
        )
    return ret
}
// extern int ZBRGDIDrawLine(int x1, int y1, int x2, int y2, int color, float thickness, out int err);
func ZBRGDIDrawLine() (ret uintptr) {
    panic("ZBRGDIDrawLine not implemented")
    ret, _, _ =  zBRGDIDrawLine.Call(
        )
    return ret
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
func ZBRGDIDrawRectangle() (ret uintptr) {
    panic("ZBRGDIDrawRectangle not implemented")
    ret, _, _ =  zBRGDIDrawRectangle.Call(
        )
    return ret
}
// extern int ZBRGDIDrawEllipse(int x, int y, int width, int height, float thickness, int color, out int err);
func ZBRGDIDrawEllipse() (ret uintptr) {
    panic("ZBRGDIDrawEllipse not implemented")
    ret, _, _ =  zBRGDIDrawEllipse.Call(
        )
    return ret
}
// extern int ZBRGDIDrawBarcode(int x, int y, int rotation, int barcodetype, int barwidthratio, int barcodemultiplier, int barcodeheight, int textunder, byte[] barcodedata, out int err);
func ZBRGDIDrawBarcode() (ret uintptr) {
    panic("ZBRGDIDrawBarcode not implemented")
    ret, _, _ =  zBRGDIDrawBarcode.Call(
        )
    return ret
}
