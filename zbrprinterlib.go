package zbrprinter

import (
    "fmt"
    "syscall"
    "unsafe"
)

func abort(funcname string, err error) {
    panic(fmt.Sprintf("%s failed: %v", funcname, err))
}

var (
    zBRPrinter                        , _ = syscall.LoadLibrary("ZBRPrinter.dll")
    zBRCloseHandle                    , _ = syscall.GetProcAddress(zBRPrinter, "ZBRCloseHandle")
    zBRGetHandle                      , _ = syscall.GetProcAddress(zBRPrinter, "ZBRGetHandle")
    zBRPRNCheckForErrors              , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNCheckForErrors")
    zBRPRNChkDueForCleaning           , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNChkDueForCleaning")
    zBRPRNClrColorImgBuf              , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNClrColorImgBuf")
    zBRPRNClrColorImgBufs             , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNClrColorImgBufs")
    zBRPRNClrErrStatusLn              , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNClrErrStatusLn")
    zBRPRNClrMediaPath                , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNClrMediaPath")
    zBRPRNClrMonoImgBuf               , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNClrMonoImgBuf")
    zBRPRNClrMonoImgBufs              , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNClrMonoImgBufs")
    zBRPRNClrSpecifiedBmp             , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNClrSpecifiedBmp")
    zBRPRNEjectCard                   , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNEjectCard")
    zBRPRNEnableMagReadDataEncryption , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNEnableMagReadDataEncryption")
    zBRPRNEndSmartCard                , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNEndSmartCard")
    zBRPRNFlipCard                    , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNFlipCard")
    zBRPRNGetChecksum                 , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetChecksum")
    zBRPRNGetCleaningParam            , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetCleaningParam")
    zBRPRNGetIPAddress                , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetIPAddress")
    zBRPRNGetMsgSuppressionState      , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetMsgSuppressionState")
    zBRPRNGetOpParam                  , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetOpParam")
    zBRPRNGetPanelsPrinted            , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetPanelsPrinted")
    zBRPRNGetPanelsRemaining          , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetPanelsRemaining")
    zBRPRNGetPrintCount               , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetPrintCount")
    zBRPRNGetPrinterOptions           , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetPrinterOptions")
    zBRPRNGetPrinterSerialNumb        , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetPrinterSerialNumb")
    zBRPRNGetPrinterSerialNumber      , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetPrinterSerialNumber")
    zBRPRNGetPrinterStatus            , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetPrinterStatus")
    zBRPRNGetPrintHeadSerialNumb      , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetPrintHeadSerialNumb")
    zBRPRNGetPrintHeadSerialNumber    , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetPrintHeadSerialNumber")
    zBRPRNGetPrintStatus              , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetPrintStatus")
    zBRPRNGetSDKProductVer            , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetSDKProductVer")
    zBRPRNGetSDKVer                   , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetSDKVer")
    zBRPRNGetSDKVsn                   , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetSDKVsn")
    zBRPRNGetSensorStatus             , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetSensorStatus")
    zBRPRNImmediateParamSave          , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNImmediateParamSave")
    zBRPRNIsPrinterReady              , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNIsPrinterReady")
    zBRPRNMoveCard                    , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNMoveCard")
    zBRPRNMoveCardBkwd                , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNMoveCardBkwd")
    zBRPRNMoveCardFwd                 , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNMoveCardFwd")
    zBRPRNMovePrintReady              , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNMovePrintReady")
    zBRPRNPrintCardPanel              , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNPrintCardPanel")
    zBRPRNPrintClearVarnish           , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNPrintClearVarnish")
    zBRPRNPrintColorImgBuf            , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNPrintColorImgBuf")
    zBRPRNPrintHologramOverlay        , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNPrintHologramOverlay")
    zBRPRNPrintMonoImgBuf             , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNPrintMonoImgBuf")
    zBRPRNPrintMonoImgBufEx           , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNPrintMonoImgBufEx")
    zBRPRNPrintMonoPanel              , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNPrintMonoPanel")
    zBRPRNPrintPrnFile                , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNPrintPrnFile")
    zBRPRNPrintTestCard               , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNPrintTestCard")
    zBRPRNPrintVarnish                , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNPrintVarnish")
    zBRPRNPrintVarnishEx              , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNPrintVarnishEx")
    zBRPRNReadMag                     , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNReadMag")
    zBRPRNReadMagByTrk                , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNReadMagByTrk")
    zBRPRNResetMagEncoder             , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNResetMagEncoder")
    zBRPRNResetPrinter                , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNResetPrinter")
    zBRPRNResync                      , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNResync")
    zBRPRNReversePrintReady           , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNReversePrintReady")
    zBRPRNSelfAdj                     , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSelfAdj")
    zBRPRNSetCardFeedingMode          , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetCardFeedingMode")
    zBRPRNSetCleaningParam            , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetCleaningParam")
    zBRPRNSetColorContrast            , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetColorContrast")
    zBRPRNSetContrastIntensityLvl     , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetContrastIntensityLvl")
    zBRPRNSetEncoderCoercivity        , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetEncoderCoercivity")
    zBRPRNSetEncodingDir              , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetEncodingDir")
    zBRPRNSetEndOfPrint               , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetEndOfPrint")
    zBRPRNSetHologramIntensity        , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetHologramIntensity")
    zBRPRNSetMagEncodingStd           , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetMagEncodingStd")
    zBRPRNSetMonoContrast             , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetMonoContrast")
    zBRPRNSetMonoIntensity            , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetMonoIntensity")
    zBRPRNSetOverlayMode              , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetOverlayMode")
    zBRPRNSetPrintHeadResistance      , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetPrintHeadResistance")
    zBRPRNSetRelativeXOffset          , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetRelativeXOffset")
    zBRPRNSetRelativeYOffset          , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetRelativeYOffset")
    zBRPRNSetStartPrintSideBOffset    , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetStartPrintSideBOffset")
    zBRPRNSetStartPrintSideBXOffset   , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetStartPrintSideBXOffset")
    zBRPRNSetStartPrintSideBYOffset   , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetStartPrintSideBYOffset")
    zBRPRNSetStartPrintXOffset        , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetStartPrintXOffset")
    zBRPRNSetStartPrintYOffset        , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetStartPrintYOffset")
    zBRPRNSetTrkDensity               , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetTrkDensity")
    zBRPRNStartCleaningCardSeq        , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNStartCleaningCardSeq")
    zBRPRNStartCleaningSeq            , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNStartCleaningSeq")
    zBRPRNStartSmartCard              , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNStartSmartCard")
    zBRPRNSuppressStatusMsgs          , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSuppressStatusMsgs")
    zBRPRNUpgradeFirmware             , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNUpgradeFirmware")
    zBRPRNWriteBarCode                , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNWriteBarCode")
    zBRPRNWriteBox                    , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNWriteBox")
    zBRPRNWriteBoxEx                  , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNWriteBoxEx")
    zBRPRNWriteMag                    , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNWriteMag")
    zBRPRNWriteMagByTrk               , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNWriteMagByTrk")
    zBRPRNWriteMagPassThru            , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNWriteMagPassThru")
    zBRPRNWriteText                   , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNWriteText")
    zBRPRNWriteTextEx                 , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNWriteTextEx")
)

const (
    ZBR_ERROR_PRINTER_MECHANICAL_ERROR              = -1
    ZBR_ERROR_BROKEN_RIBBON                         = 1
    ZBR_ERROR_TEMPERATURE                           = 2
    ZBR_ERROR_MECHANICAL_ERROR                      = 3
    ZBR_ERROR_OUT_OF_CARD                           = 4
    ZBR_ERROR_CARD_IN_ENCODER                       = 5
    ZBR_ERROR_CARD_NOT_IN_ENCODER                   = 6
    ZBR_ERROR_PRINT_HEAD_OPEN                       = 7
    ZBR_ERROR_OUT_OF_RIBBON                         = 8
    ZBR_ERROR_REMOVE_RIBBON                         = 9
    ZBR_ERROR_PARAMETERS_ERROR                      = 10
    ZBR_ERROR_INVALID_COORDINATES                   = 11
    ZBR_ERROR_UNKNOWN_BARCODE                       = 12
    ZBR_ERROR_UNKNOWN_TEXT                          = 13
    ZBR_ERROR_COMMAND_ERROR                         = 14
    ZBR_ERROR_BARCODE_DATA_SYNTAX                   = 20
    ZBR_ERROR_TEXT_DATA_SYNTAX                      = 21
    ZBR_ERROR_GRAPHIC_DATA_SYNTAX                   = 22
    ZBR_ERROR_GRAPHIC_IMAGE_INITIALIZATION          = 30
    ZBR_ERROR_GRAPHIC_IMAGE_MAXIMUM_WIDTH_EXCEEDED  = 31
    ZBR_ERROR_GRAPHIC_IMAGE_MAXIMUM_HEIGHT_EXCEEDED = 32
    ZBR_ERROR_GRAPHIC_IMAGE_DATA_CHECKSUM_ERROR     = 33
    ZBR_ERROR_DATA_TRANSFER_TIME_OUT                = 34
    ZBR_ERROR_CHECK_RIBBON                          = 35
    ZBR_ERROR_INVALID_MAGNETIC_DATA                 = 40
    ZBR_ERROR_MAG_ENCODER_WRITE                     = 41
    ZBR_ERROR_READING_ERROR                         = 42
    ZBR_ERROR_MAG_ENCODER_MECHANICAL                = 43
    ZBR_ERROR_MAG_ENCODER_NOT_RESPONDING            = 44
    ZBR_ERROR_MAG_ENCODER_MISSING_OR_CARD_          = 45
    ZBR_ERROR_ROTATION_ERROR                        = 47
    ZBR_ERROR_COVER_OPEN                            = 48
    ZBR_ERROR_ENCODING_ERROR                        = 49
    ZBR_ERROR_MAGNETIC_ERROR                        = 50
    ZBR_ERROR_BLANK_TRACK                           = 51
    ZBR_ERROR_FLASH_ERROR                           = 52
    ZBR_ERROR_NO_ACCESS                             = 53
    ZBR_ERROR_SEQUENCE_ERROR                        = 54
    ZBR_ERROR_PROX_ERROR                            = 55
    ZBR_ERROR_CONTACT_DATA_ERROR                    = 56
    ZBR_ERROR_PROX_DATA_ERROR                       = 57
    ZBR_SDK_ERROR_PRINTER_NOT_SUPPORTED             = 60
    ZBR_SDK_ERROR_CANNOT_GET_PRINTER_HANDLE         = 61
    ZBR_SDK_ERROR_CANNOT_GET_PRINTER_DRIVER         = 62
    ZBR_SDK_ERROR_INVALID_PARAMETER                 = 63
    ZBR_SDK_ERROR_PRINTER_BUSY                      = 64
    ZBR_SDK_ERROR_INVALID_PRINTER_HANDLE            = 65
    ZBR_SDK_ERROR_CLOSE_HANDLE_ERROR                = 66
    ZBR_SDK_ERROR_COMMUNICATION_ERROR               = 67
    ZBR_SDK_ERROR_BUFFER_OVERFLOW                   = 68
    ZBR_SDK_ERROR_READ_DATA_ERROR                   = 69
    ZBR_SDK_ERROR_WRITE_DATA_ERROR                  = 70
    ZBR_SDK_ERROR_LOAD_LIBRARY_ERROR                = 71
    ZBR_SDK_ERROR_INVALID_STRUCT_ALIGNMENT          = 72
    ZBR_SDK_ERROR_GETTING_DEVICE_CONTEXT            = 73
    ZBR_SDK_ERROR_SPOOLER_ERROR                     = 74
    ZBR_SDK_ERROR_OUT_OF_MEMORY                     = 75
    ZBR_SDK_ERROR_OUT_OF_DISK_SPACE                 = 76
    ZBR_SDK_ERROR_USER_ABORT                        = 77
    ZBR_SDK_ERROR_APPLICATION_ABORT                 = 78
    ZBR_SDK_ERROR_CREATE_FILE_ERROR                 = 79
    ZBR_SDK_ERROR_WRITE_FILE_ERROR                  = 80
    ZBR_SDK_ERROR_READ_FILE_ERROR                   = 81
    ZBR_SDK_ERROR_INVALID_MEDIA                     = 82
    ZBR_SDK_ERROR_MEMORY_ALLOCATION                 = 83
    ZBR_SDK_ERROR_UNKNOWN_ERROR                     = 255
)

// extern int ZBRCloseHandle(IntPtr handle, out int err);
func ZBRCloseHandle() () {
    if zBRCloseHandle == 0 {
        panic("ZBRCloseHandle is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRCloseHandle is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRCloseHandle),
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
        abort("ZBRCloseHandle",callErr)
    }
    return
}

// extern int ZBRGetHandle(out IntPtr handle, byte[] drvname, out int printertype, out int err);
func ZBRGetHandle(drvName string) (handle syscall.Handle, prn_type int, err int) {
    if zBRGetHandle == 0 {
        panic("ZBRGetHandle not defined. Check library")
    }
    var nargs uintptr = 4
    ret, _, callErr := syscall.Syscall9(uintptr(zBRGetHandle),
            nargs,
            uintptr(unsafe.Pointer(&handle)),
            uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(drvName))),
            uintptr(unsafe.Pointer(&prn_type)),
            uintptr(unsafe.Pointer(&err)),
            0,
            0,
            0,
            0,
            0)
    return handle, prn_type, err
}

// extern int ZBRPRNCheckForErrors(IntPtr hprinter, int printertype);
func ZBRPRNCheckForErrors() () {
    if zBRPRNCheckForErrors == 0 {
        panic("ZBRPRNCheckForErrors is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNCheckForErrors is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNCheckForErrors),
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
        abort("ZBRPRNCheckForErrors",callErr)
    }
    return
}

// extern int ZBRPRNChkDueForCleaning(IntPtr hprinter, int printertype, out int imgcounter, out int cleancounter, out int cleancardcounter, out int err);
func ZBRPRNChkDueForCleaning() () {
    if zBRPRNChkDueForCleaning == 0 {
        panic("ZBRPRNChkDueForCleaning is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNChkDueForCleaning is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNChkDueForCleaning),
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
        abort("ZBRPRNChkDueForCleaning",callErr)
    }
    return
}

// extern int ZBRPRNClrColorImgBuf(IntPtr hprinter, int printertype, int colorbufidx, out int err);
func ZBRPRNClrColorImgBuf() () {
    if zBRPRNClrColorImgBuf == 0 {
        panic("ZBRPRNClrColorImgBuf is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNClrColorImgBuf is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNClrColorImgBuf),
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
        abort("ZBRPRNClrColorImgBuf",callErr)
    }
    return
}

// extern int ZBRPRNClrColorImgBufs(HANDLE hPrinter, int printerType, int *err)
func ZBRPRNClrColorImgBufs() () {
    if zBRPRNClrColorImgBufs == 0 {
        panic("ZBRPRNClrColorImgBufs is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNClrColorImgBufs is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNClrColorImgBufs),
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
        abort("ZBRPRNClrColorImgBufs",callErr)
    }
    return
}

// extern int ZBRPRNClrErrStatusLn(IntPtr hprinter, int printertype, out int err);
func ZBRPRNClrErrStatusLn() () {
    if zBRPRNClrErrStatusLn == 0 {
        panic("ZBRPRNClrErrStatusLn is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNClrErrStatusLn is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNClrErrStatusLn),
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
        abort("ZBRPRNClrErrStatusLn",callErr)
    }
    return
}

// extern int ZBRPRNClrMediaPath(IntPtr hprinter, int printertype, out int err);
func ZBRPRNClrMediaPath() () {
    if zBRPRNClrMediaPath == 0 {
        panic("ZBRPRNClrMediaPath is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNClrMediaPath is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNClrMediaPath),
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
        abort("ZBRPRNClrMediaPath",callErr)
    }
    return
}

// extern int ZBRPRNClrMonoImgBuf(IntPtr hprinter, int printertype, int clrvarnish, out int err);
func ZBRPRNClrMonoImgBuf() () {
    if zBRPRNClrMonoImgBuf == 0 {
        panic("ZBRPRNClrMonoImgBuf is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNClrMonoImgBuf is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNClrMonoImgBuf),
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
        abort("ZBRPRNClrMonoImgBuf",callErr)
    }
    return
}

// extern int ZBRPRNClrMonoImgBufs(IntPtr hprinter, int printertype, int clrbuffer, out int err);
func ZBRPRNClrMonoImgBufs() () {
    if zBRPRNClrMonoImgBufs == 0 {
        panic("ZBRPRNClrMonoImgBufs is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNClrMonoImgBufs is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNClrMonoImgBufs),
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
        abort("ZBRPRNClrMonoImgBufs",callErr)
    }
    return
}

// extern int ZBRPRNClrSpecifiedBmp(IntPtr hprinter, int printertype, int colorbufidx, out int err);
func ZBRPRNClrSpecifiedBmp() () {
    if zBRPRNClrSpecifiedBmp == 0 {
        panic("ZBRPRNClrSpecifiedBmp is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNClrSpecifiedBmp is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNClrSpecifiedBmp),
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
        abort("ZBRPRNClrSpecifiedBmp",callErr)
    }
    return
}

// extern int ZBRPRNEjectCard(IntPtr _handle, int prn_type, out int err);
func ZBRPRNEjectCard() () {
    if zBRPRNEjectCard == 0 {
        panic("ZBRPRNEjectCard is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNEjectCard is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNEjectCard),
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
        abort("ZBRPRNEjectCard",callErr)
    }
    return
}

// extern int ZBRPRNEnableMagReadDataEncryption(IntPtr handle, int printertype, out int err);
func ZBRPRNEnableMagReadDataEncryption() () {
    if zBRPRNEnableMagReadDataEncryption == 0 {
        panic("ZBRPRNEnableMagReadDataEncryption is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNEnableMagReadDataEncryption is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNEnableMagReadDataEncryption),
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
        abort("ZBRPRNEnableMagReadDataEncryption",callErr)
    }
    return
}

// extern int ZBRPRNEndSmartCard(IntPtr _handle, int printertype, int cardtype, int movement, out int err);
func ZBRPRNEndSmartCard() () {
    if zBRPRNEndSmartCard == 0 {
        panic("ZBRPRNEndSmartCard is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNEndSmartCard is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNEndSmartCard),
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
        abort("ZBRPRNEndSmartCard",callErr)
    }
    return
}

// extern int ZBRPRNFlipCard(IntPtr hprinter, int printertype, out int err);
func ZBRPRNFlipCard() () {
    if zBRPRNFlipCard == 0 {
        panic("ZBRPRNFlipCard is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNFlipCard is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNFlipCard),
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
        abort("ZBRPRNFlipCard",callErr)
    }
    return
}

// extern int ZBRPRNGetChecksum(IntPtr hprinter, int printertype, out int checksum, out int err);
func ZBRPRNGetChecksum() () {
    if zBRPRNGetChecksum == 0 {
        panic("ZBRPRNGetChecksum is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNGetChecksum is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetChecksum),
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
        abort("ZBRPRNGetChecksum",callErr)
    }
    return
}

// extern int ZBRPRNGetCleaningParam(IntPtr hprinter, int printertype, out int imgcounter, out int cleancounter, out int cleancardcounter, out int err);
func ZBRPRNGetCleaningParam() () {
    if zBRPRNGetCleaningParam == 0 {
        panic("ZBRPRNGetCleaningParam is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNGetCleaningParam is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetCleaningParam),
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
        abort("ZBRPRNGetCleaningParam",callErr)
    }
    return
}

// extern int ZBRPRNGetIPAddress(string PrinterName, char* ipAddress);
func ZBRPRNGetIPAddress() () {
    if zBRPRNGetIPAddress == 0 {
        panic("ZBRPRNGetIPAddress is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNGetIPAddress is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetIPAddress),
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
        abort("ZBRPRNGetIPAddress",callErr)
    }
    return
}

// extern int ZBRPRNGetMsgSuppressionState(IntPtr hprinter, int printertype, out int state, out int err);
func ZBRPRNGetMsgSuppressionState() () {
    if zBRPRNGetMsgSuppressionState == 0 {
        panic("ZBRPRNGetMsgSuppressionState is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNGetMsgSuppressionState is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetMsgSuppressionState),
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
        abort("ZBRPRNGetMsgSuppressionState",callErr)
    }
    return
}

// extern int ZBRPRNGetOpParam(IntPtr hprinter, int printertype, int paramidx, byte[] opparam, out int respsize, out int err);
func ZBRPRNGetOpParam() () {
    if zBRPRNGetOpParam == 0 {
        panic("ZBRPRNGetOpParam is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNGetOpParam is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetOpParam),
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
        abort("ZBRPRNGetOpParam",callErr)
    }
    return
}

// extern int ZBRPRNGetPanelsPrinted(IntPtr hprinter, int printertype, out int panels, out int err);
func ZBRPRNGetPanelsPrinted() () {
    if zBRPRNGetPanelsPrinted == 0 {
        panic("ZBRPRNGetPanelsPrinted is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNGetPanelsPrinted is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetPanelsPrinted),
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
        abort("ZBRPRNGetPanelsPrinted",callErr)
    }
    return
}

// extern int ZBRPRNGetPanelsRemaining(IntPtr hprinter, int printertype, out int panels, out int err);
func ZBRPRNGetPanelsRemaining() () {
    if zBRPRNGetPanelsRemaining == 0 {
        panic("ZBRPRNGetPanelsRemaining is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNGetPanelsRemaining is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetPanelsRemaining),
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
        abort("ZBRPRNGetPanelsRemaining",callErr)
    }
    return
}

// extern int ZBRPRNGetPrintCount(IntPtr hprinter, int printertype, out int printcount, out int err);
func ZBRPRNGetPrintCount() () {
    if zBRPRNGetPrintCount == 0 {
        panic("ZBRPRNGetPrintCount is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNGetPrintCount is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetPrintCount),
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
        abort("ZBRPRNGetPrintCount",callErr)
    }
    return
}

// extern int ZBRPRNGetPrinterOptions(IntPtr hprinter, int printertype, byte[] options, out int respsize, out int err);
func ZBRPRNGetPrinterOptions() () {
    if zBRPRNGetPrinterOptions == 0 {
        panic("ZBRPRNGetPrinterOptions is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNGetPrinterOptions is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetPrinterOptions),
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
        abort("ZBRPRNGetPrinterOptions",callErr)
    }
    return
}

// extern int ZBRPRNGetPrinterSerialNumb(IntPtr hprinter, int printertype, byte[] serialnum, out int respsize, out int err);
func ZBRPRNGetPrinterSerialNumb() () {
    if zBRPRNGetPrinterSerialNumb == 0 {
        panic("ZBRPRNGetPrinterSerialNumb is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNGetPrinterSerialNumb is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetPrinterSerialNumb),
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
        abort("ZBRPRNGetPrinterSerialNumb",callErr)
    }
    return
}

// extern int ZBRPRNGetPrinterSerialNumber(IntPtr hprinter, int printertype, byte[] serialnum, out int respsize, out int err);
func ZBRPRNGetPrinterSerialNumber() () {
    if zBRPRNGetPrinterSerialNumber == 0 {
        panic("ZBRPRNGetPrinterSerialNumber is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNGetPrinterSerialNumber is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetPrinterSerialNumber),
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
        abort("ZBRPRNGetPrinterSerialNumber",callErr)
    }
    return
}

// extern int ZBRPRNGetPrinterStatus(out int statuscode);
func ZBRPRNGetPrinterStatus() () {
    if zBRPRNGetPrinterStatus == 0 {
        panic("ZBRPRNGetPrinterStatus is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNGetPrinterStatus is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetPrinterStatus),
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
        abort("ZBRPRNGetPrinterStatus",callErr)
    }
    return
}

// extern int ZBRPRNGetPrintHeadSerialNumb(IntPtr hprinter, int printertype, byte[] serialnum, out int respsize, out int err);
func ZBRPRNGetPrintHeadSerialNumb() () {
    if zBRPRNGetPrintHeadSerialNumb == 0 {
        panic("ZBRPRNGetPrintHeadSerialNumb is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNGetPrintHeadSerialNumb is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetPrintHeadSerialNumb),
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
        abort("ZBRPRNGetPrintHeadSerialNumb",callErr)
    }
    return
}

// extern int ZBRPRNGetPrintHeadSerialNumber(IntPtr hprinter, int printertype, byte[] serialnum, out int respsize, out int err);
func ZBRPRNGetPrintHeadSerialNumber() () {
    if zBRPRNGetPrintHeadSerialNumber == 0 {
        panic("ZBRPRNGetPrintHeadSerialNumber is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNGetPrintHeadSerialNumber is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetPrintHeadSerialNumber),
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
        abort("ZBRPRNGetPrintHeadSerialNumber",callErr)
    }
    return
}

// extern int ZBRPRNGetPrintStatus(IntPtr hprinter, int printertype);
func ZBRPRNGetPrintStatus() () {
    if zBRPRNGetPrintStatus == 0 {
        panic("ZBRPRNGetPrintStatus is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNGetPrintStatus is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetPrintStatus),
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
        abort("ZBRPRNGetPrintStatus",callErr)
    }
    return
}

// extern int ZBRPRNGetSDKProductVer(byte[] productversion, out int buffsize);
func ZBRPRNGetSDKProductVer() () {
    if zBRPRNGetSDKProductVer == 0 {
        panic("ZBRPRNGetSDKProductVer is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNGetSDKProductVer is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetSDKProductVer),
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
        abort("ZBRPRNGetSDKProductVer",callErr)
    }
    return
}

// extern void ZBRPRNGetSDKVer(out int major, out int minor, out int englevel);
func ZBRPRNGetSDKVer() (major uint,minor uint,engLevel uint) {
    if zBRPRNGetSDKVer == 0 {
        panic("ZBRPRNGetSDKVer not defined. Check library")
    }
    var nargs uintptr = 3
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetSDKVer),
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
            abort("Call ZBRPRNGetSDKVer", callErr)
    }
    return major,minor,engLevel
}

// extern void ZBRPRNGetSDKVsn(out int major, out int minor, out int englevel);
func ZBRPRNGetSDKVsn() () {
    if zBRPRNGetSDKVsn == 0 {
        panic("ZBRPRNGetSDKVsn is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNGetSDKVsn is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetSDKVsn),
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
        abort("ZBRPRNGetSDKVsn",callErr)
    }
    return
}

// extern int ZBRPRNGetSensorStatus(IntPtr hprinter, int printertype, out byte status, out int err);
func ZBRPRNGetSensorStatus() () {
    if zBRPRNGetSensorStatus == 0 {
        panic("ZBRPRNGetSensorStatus is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNGetSensorStatus is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetSensorStatus),
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
        abort("ZBRPRNGetSensorStatus",callErr)
    }
    return
}

// extern int ZBRPRNImmediateParamSave(IntPtr hprinter, int printertype, out int err);
func ZBRPRNImmediateParamSave() () {
    if zBRPRNImmediateParamSave == 0 {
        panic("ZBRPRNImmediateParamSave is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNImmediateParamSave is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNImmediateParamSave),
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
        abort("ZBRPRNImmediateParamSave",callErr)
    }
    return
}

// extern int ZBRPRNIsPrinterReady(IntPtr hprinter, int printertype, out int err);
func ZBRPRNIsPrinterReady() () {
    if zBRPRNIsPrinterReady == 0 {
        panic("ZBRPRNIsPrinterReady is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNIsPrinterReady is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNIsPrinterReady),
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
        abort("ZBRPRNIsPrinterReady",callErr)
    }
    return
}

// extern int ZBRPRNMoveCard(IntPtr hprinter, int printertype, int steps, out int err);
func ZBRPRNMoveCard() () {
    if zBRPRNMoveCard == 0 {
        panic("ZBRPRNMoveCard is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNMoveCard is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNMoveCard),
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
        abort("ZBRPRNMoveCard",callErr)
    }
    return
}

// extern int ZBRPRNMoveCardBkwd(IntPtr hprinter, int printertype, int steps, out int err);
func ZBRPRNMoveCardBkwd() () {
    if zBRPRNMoveCardBkwd == 0 {
        panic("ZBRPRNMoveCardBkwd is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNMoveCardBkwd is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNMoveCardBkwd),
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
        abort("ZBRPRNMoveCardBkwd",callErr)
    }
    return
}

// extern int ZBRPRNMoveCardFwd(IntPtr hprinter, int printertype, int steps, out int err);
func ZBRPRNMoveCardFwd() () {
    if zBRPRNMoveCardFwd == 0 {
        panic("ZBRPRNMoveCardFwd is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNMoveCardFwd is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNMoveCardFwd),
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
        abort("ZBRPRNMoveCardFwd",callErr)
    }
    return
}

// extern int ZBRPRNMovePrintReady(IntPtr handle, int printertype, out int err);
func ZBRPRNMovePrintReady() () {
    if zBRPRNMovePrintReady == 0 {
        panic("ZBRPRNMovePrintReady is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNMovePrintReady is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNMovePrintReady),
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
        abort("ZBRPRNMovePrintReady",callErr)
    }
    return
}

// extern int ZBRPRNPrintCardPanel(IntPtr hprinter, int printertype, int imgbufidx, out int err);
func ZBRPRNPrintCardPanel() () {
    if zBRPRNPrintCardPanel == 0 {
        panic("ZBRPRNPrintCardPanel is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNPrintCardPanel is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNPrintCardPanel),
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
        abort("ZBRPRNPrintCardPanel",callErr)
    }
    return
}

// extern int ZBRPRNPrintClearVarnish(IntPtr hprinter, int printertype, out int err);
func ZBRPRNPrintClearVarnish() () {
    if zBRPRNPrintClearVarnish == 0 {
        panic("ZBRPRNPrintClearVarnish is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNPrintClearVarnish is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNPrintClearVarnish),
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
        abort("ZBRPRNPrintClearVarnish",callErr)
    }
    return
}

// extern int ZBRPRNPrintColorImgBuf(IntPtr hprinter, int printertype, int imgbufidx, out int err);
func ZBRPRNPrintColorImgBuf() () {
    if zBRPRNPrintColorImgBuf == 0 {
        panic("ZBRPRNPrintColorImgBuf is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNPrintColorImgBuf is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNPrintColorImgBuf),
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
        abort("ZBRPRNPrintColorImgBuf",callErr)
    }
    return
}

// extern int ZBRPRNPrintHologramOverlay(IntPtr hprinter, int printertype, int printcardcommand, out int err);
func ZBRPRNPrintHologramOverlay() () {
    if zBRPRNPrintHologramOverlay == 0 {
        panic("ZBRPRNPrintHologramOverlay is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNPrintHologramOverlay is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNPrintHologramOverlay),
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
        abort("ZBRPRNPrintHologramOverlay",callErr)
    }
    return
}

// extern int ZBRPRNPrintMonoImgBuf(IntPtr hprinter, int printertype, out int err);
func ZBRPRNPrintMonoImgBuf() () {
    if zBRPRNPrintMonoImgBuf == 0 {
        panic("ZBRPRNPrintMonoImgBuf is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNPrintMonoImgBuf is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNPrintMonoImgBuf),
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
        abort("ZBRPRNPrintMonoImgBuf",callErr)
    }
    return
}

// extern int ZBRPRNPrintMonoImgBufEx(IntPtr hprinter, int printertype, int cardcommand, out int err);
func ZBRPRNPrintMonoImgBufEx() () {
    if zBRPRNPrintMonoImgBufEx == 0 {
        panic("ZBRPRNPrintMonoImgBufEx is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNPrintMonoImgBufEx is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNPrintMonoImgBufEx),
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
        abort("ZBRPRNPrintMonoImgBufEx",callErr)
    }
    return
}

// extern int ZBRPRNPrintMonoPanel(IntPtr hprinter, int printertype, out int err);
func ZBRPRNPrintMonoPanel() () {
    if zBRPRNPrintMonoPanel == 0 {
        panic("ZBRPRNPrintMonoPanel is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNPrintMonoPanel is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNPrintMonoPanel),
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
        abort("ZBRPRNPrintMonoPanel",callErr)
    }
    return
}

// extern int ZBRPRNPrintPRNFile(IntPtr hprinter, int printertype, byte[] filename, out int err);
func ZBRPRNPrintPrnFile() () {
    if zBRPRNPrintPrnFile == 0 {
        panic("ZBRPRNPrintPrnFile is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNPrintPrnFile is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNPrintPrnFile),
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
        abort("ZBRPRNPrintPrnFile",callErr)
    }
    return
}

// extern int ZBRPRNPrintTestCard(IntPtr hprinter, int printertype, int testcardtype, out int err);
func ZBRPRNPrintTestCard() () {
    if zBRPRNPrintTestCard == 0 {
        panic("ZBRPRNPrintTestCard is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNPrintTestCard is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNPrintTestCard),
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
        abort("ZBRPRNPrintTestCard",callErr)
    }
    return
}

// extern int ZBRPRNPrintVarnish(IntPtr hprinter, int printertype, out int err);
func ZBRPRNPrintVarnish() () {
    if zBRPRNPrintVarnish == 0 {
        panic("ZBRPRNPrintVarnish is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNPrintVarnish is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNPrintVarnish),
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
        abort("ZBRPRNPrintVarnish",callErr)
    }
    return
}

// extern int ZBRPRNPrintVarnishEx(IntPtr hprinter, int printertype, int printcardcommand, out int err);
func ZBRPRNPrintVarnishEx() () {
    if zBRPRNPrintVarnishEx == 0 {
        panic("ZBRPRNPrintVarnishEx is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNPrintVarnishEx is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNPrintVarnishEx),
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
        abort("ZBRPRNPrintVarnishEx",callErr)
    }
    return
}

// extern int ZBRPRNReadMag(IntPtr handle, int printertype, int trkstoread, byte[] trk1buf, out int trk1bytesneeded, byte[] trk2buf, out int trk2bytesneeded, byte[] trk3buf, out int trk3bytesneeded, out int err);
func ZBRPRNReadMag() () {
    if zBRPRNReadMag == 0 {
        panic("ZBRPRNReadMag is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNReadMag is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNReadMag),
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
        abort("ZBRPRNReadMag",callErr)
    }
    return
}

// extern int ZBRPRNReadMagByTrk(IntPtr hprinter, int printertype, int trknum, byte[] trkbuf, out int respsize, out int err);
func ZBRPRNReadMagByTrk() () {
    if zBRPRNReadMagByTrk == 0 {
        panic("ZBRPRNReadMagByTrk is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNReadMagByTrk is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNReadMagByTrk),
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
        abort("ZBRPRNReadMagByTrk",callErr)
    }
    return
}

// extern int ZBRPRNResetMagEncoder(IntPtr hprinter, int printertype, out int err);
func ZBRPRNResetMagEncoder() () {
    if zBRPRNResetMagEncoder == 0 {
        panic("ZBRPRNResetMagEncoder is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNResetMagEncoder is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNResetMagEncoder),
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
        abort("ZBRPRNResetMagEncoder",callErr)
    }
    return
}

// extern int ZBRPRNResetPrinter(IntPtr hprinter, int printertype, out int err);
func ZBRPRNResetPrinter() () {
    if zBRPRNResetPrinter == 0 {
        panic("ZBRPRNResetPrinter is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNResetPrinter is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNResetPrinter),
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
        abort("ZBRPRNResetPrinter",callErr)
    }
    return
}

// extern int ZBRPRNResync(IntPtr hprinter, int printertype, out int err);
func ZBRPRNResync() () {
    if zBRPRNResync == 0 {
        panic("ZBRPRNResync is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNResync is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNResync),
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
        abort("ZBRPRNResync",callErr)
    }
    return
}

// extern int ZBRPRNReversePrintReady(IntPtr hprinter, int printertype, out int err);
func ZBRPRNReversePrintReady() () {
    if zBRPRNReversePrintReady == 0 {
        panic("ZBRPRNReversePrintReady is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNReversePrintReady is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNReversePrintReady),
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
        abort("ZBRPRNReversePrintReady",callErr)
    }
    return
}

// extern int ZBRPRNSelfAdj(IntPtr hprinter, int printertype, out int err);
func ZBRPRNSelfAdj() () {
    if zBRPRNSelfAdj == 0 {
        panic("ZBRPRNSelfAdj is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSelfAdj is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSelfAdj),
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
        abort("ZBRPRNSelfAdj",callErr)
    }
    return
}

// extern int ZBRPRNSetCardFeedingMode(IntPtr hprinter, int printertype, int mode, out int err);
func ZBRPRNSetCardFeedingMode() () {
    if zBRPRNSetCardFeedingMode == 0 {
        panic("ZBRPRNSetCardFeedingMode is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetCardFeedingMode is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetCardFeedingMode),
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
        abort("ZBRPRNSetCardFeedingMode",callErr)
    }
    return
}

// extern int ZBRPRNSetCleaningParam(IntPtr hprinter, int printertype, int ribbonpanelcounter, int cleancardpass, out int err);
func ZBRPRNSetCleaningParam() () {
    if zBRPRNSetCleaningParam == 0 {
        panic("ZBRPRNSetCleaningParam is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetCleaningParam is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetCleaningParam),
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
        abort("ZBRPRNSetCleaningParam",callErr)
    }
    return
}

// extern int ZBRPRNSetColorContrast(IntPtr hprinter, int printertype, int imgbufidx, int contrast, out int err);
func ZBRPRNSetColorContrast() () {
    if zBRPRNSetColorContrast == 0 {
        panic("ZBRPRNSetColorContrast is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetColorContrast is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetColorContrast),
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
        abort("ZBRPRNSetColorContrast",callErr)
    }
    return
}

// extern int ZBRPRNSetContrastIntensityLvl(IntPtr hprinter, int printertype, int imgbufidx, int intensity, out int err);
func ZBRPRNSetContrastIntensityLvl() () {
    if zBRPRNSetContrastIntensityLvl == 0 {
        panic("ZBRPRNSetContrastIntensityLvl is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetContrastIntensityLvl is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetContrastIntensityLvl),
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
        abort("ZBRPRNSetContrastIntensityLvl",callErr)
    }
    return
}

// extern int ZBRPRNSetEncoderCoercivity(IntPtr hprinter, int printertype, int coercivity, out int err);
func ZBRPRNSetEncoderCoercivity() () {
    if zBRPRNSetEncoderCoercivity == 0 {
        panic("ZBRPRNSetEncoderCoercivity is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetEncoderCoercivity is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetEncoderCoercivity),
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
        abort("ZBRPRNSetEncoderCoercivity",callErr)
    }
    return
}

// extern int ZBRPRNSetEncodingDir(IntPtr hprinter, int printertype, int dir, out int err);
func ZBRPRNSetEncodingDir() () {
    if zBRPRNSetEncodingDir == 0 {
        panic("ZBRPRNSetEncodingDir is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetEncodingDir is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetEncodingDir),
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
        abort("ZBRPRNSetEncodingDir",callErr)
    }
    return
}

// extern int ZBRPRNSetEndOfPrint(IntPtr hprinter, int printertype, int xwidth, out int err);
func ZBRPRNSetEndOfPrint() () {
    if zBRPRNSetEndOfPrint == 0 {
        panic("ZBRPRNSetEndOfPrint is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetEndOfPrint is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetEndOfPrint),
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
        abort("ZBRPRNSetEndOfPrint",callErr)
    }
    return
}

// extern int ZBRPRNSetHologramIntensity(IntPtr hprinter, int printertype, int intensity, out int err);
func ZBRPRNSetHologramIntensity() () {
    if zBRPRNSetHologramIntensity == 0 {
        panic("ZBRPRNSetHologramIntensity is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetHologramIntensity is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetHologramIntensity),
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
        abort("ZBRPRNSetHologramIntensity",callErr)
    }
    return
}

// extern int ZBRPRNSetMagEncodingStd(IntPtr hprinter, int printertype, int standard, out int err);
func ZBRPRNSetMagEncodingStd() () {
    if zBRPRNSetMagEncodingStd == 0 {
        panic("ZBRPRNSetMagEncodingStd is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetMagEncodingStd is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetMagEncodingStd),
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
        abort("ZBRPRNSetMagEncodingStd",callErr)
    }
    return
}

// extern int ZBRPRNSetMonoContrast(IntPtr hprinter, int printertype, int contrast, out int err);
func ZBRPRNSetMonoContrast() () {
    if zBRPRNSetMonoContrast == 0 {
        panic("ZBRPRNSetMonoContrast is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetMonoContrast is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetMonoContrast),
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
        abort("ZBRPRNSetMonoContrast",callErr)
    }
    return
}

// extern int ZBRPRNSetMonoIntensity(IntPtr hprinter, int printertype, int intensity, out int err);
func ZBRPRNSetMonoIntensity() () {
    if zBRPRNSetMonoIntensity == 0 {
        panic("ZBRPRNSetMonoIntensity is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetMonoIntensity is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetMonoIntensity),
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
        abort("ZBRPRNSetMonoIntensity",callErr)
    }
    return
}

// extern int ZBRPRNSetOverlayMode(HANDLE hPrinter, int printerType, int side, OverlayType overlay, char *filename, int *err);
func ZBRPRNSetOverlayMode() () {
    if zBRPRNSetOverlayMode == 0 {
        panic("ZBRPRNSetOverlayMode is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetOverlayMode is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetOverlayMode),
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
        abort("ZBRPRNSetOverlayMode",callErr)
    }
    return
}

// extern int ZBRPRNSetPrintHeadResistance(IntPtr hprinter, int printertype, int resistance, out int err);
func ZBRPRNSetPrintHeadResistance() () {
    if zBRPRNSetPrintHeadResistance == 0 {
        panic("ZBRPRNSetPrintHeadResistance is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetPrintHeadResistance is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetPrintHeadResistance),
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
        abort("ZBRPRNSetPrintHeadResistance",callErr)
    }
    return
}

// extern int ZBRPRNSetRelativeXOffset(IntPtr hprinter, int printertype, int offset, out int err);
func ZBRPRNSetRelativeXOffset() () {
    if zBRPRNSetRelativeXOffset == 0 {
        panic("ZBRPRNSetRelativeXOffset is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetRelativeXOffset is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetRelativeXOffset),
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
        abort("ZBRPRNSetRelativeXOffset",callErr)
    }
    return
}

// extern int ZBRPRNSetRelativeYOffset(IntPtr hprinter, int printertype, int offset, out int err);
func ZBRPRNSetRelativeYOffset() () {
    if zBRPRNSetRelativeYOffset == 0 {
        panic("ZBRPRNSetRelativeYOffset is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetRelativeYOffset is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetRelativeYOffset),
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
        abort("ZBRPRNSetRelativeYOffset",callErr)
    }
    return
}

// extern int ZBRPRNSetStartPrintSideBOffset(IntPtr hprinter, int printertype, int x_y, int offset, out int err);
func ZBRPRNSetStartPrintSideBOffset() () {
    if zBRPRNSetStartPrintSideBOffset == 0 {
        panic("ZBRPRNSetStartPrintSideBOffset is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetStartPrintSideBOffset is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetStartPrintSideBOffset),
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
        abort("ZBRPRNSetStartPrintSideBOffset",callErr)
    }
    return
}

// extern int ZBRPRNSetStartPrintSideBXOffset(IntPtr hprinter, int printertype, int offset, out int err);
func ZBRPRNSetStartPrintSideBXOffset() () {
    if zBRPRNSetStartPrintSideBXOffset == 0 {
        panic("ZBRPRNSetStartPrintSideBXOffset is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetStartPrintSideBXOffset is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetStartPrintSideBXOffset),
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
        abort("ZBRPRNSetStartPrintSideBXOffset",callErr)
    }
    return
}

// extern int ZBRPRNSetStartPrintSideBYOffset(IntPtr hprinter, int printertype, int offset, out int err);
func ZBRPRNSetStartPrintSideBYOffset() () {
    if zBRPRNSetStartPrintSideBYOffset == 0 {
        panic("ZBRPRNSetStartPrintSideBYOffset is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetStartPrintSideBYOffset is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetStartPrintSideBYOffset),
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
        abort("ZBRPRNSetStartPrintSideBYOffset",callErr)
    }
    return
}

// extern int ZBRPRNSetStartPrintXOffset(IntPtr hprinter, int printertype, int offset, out int err);
func ZBRPRNSetStartPrintXOffset() () {
    if zBRPRNSetStartPrintXOffset == 0 {
        panic("ZBRPRNSetStartPrintXOffset is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetStartPrintXOffset is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetStartPrintXOffset),
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
        abort("ZBRPRNSetStartPrintXOffset",callErr)
    }
    return
}

// extern int ZBRPRNSetStartPrintYOffset(IntPtr hprinter, int printertype, int offset, out int err);
func ZBRPRNSetStartPrintYOffset() () {
    if zBRPRNSetStartPrintYOffset == 0 {
        panic("ZBRPRNSetStartPrintYOffset is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetStartPrintYOffset is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetStartPrintYOffset),
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
        abort("ZBRPRNSetStartPrintYOffset",callErr)
    }
    return
}

// extern int ZBRPRNSetTrkDensity(IntPtr hprinter, int printertype, int trknum, int density, out int err);
func ZBRPRNSetTrkDensity() () {
    if zBRPRNSetTrkDensity == 0 {
        panic("ZBRPRNSetTrkDensity is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSetTrkDensity is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetTrkDensity),
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
        abort("ZBRPRNSetTrkDensity",callErr)
    }
    return
}

// extern int ZBRPRNStartCleaningCardSeq(IntPtr hprinter, int printertype, out int err);
func ZBRPRNStartCleaningCardSeq() () {
    if zBRPRNStartCleaningCardSeq == 0 {
        panic("ZBRPRNStartCleaningCardSeq is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNStartCleaningCardSeq is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNStartCleaningCardSeq),
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
        abort("ZBRPRNStartCleaningCardSeq",callErr)
    }
    return
}

// extern int ZBRPRNStartCleaningSeq(IntPtr hprinter, int printertype, out int err);
func ZBRPRNStartCleaningSeq() () {
    if zBRPRNStartCleaningSeq == 0 {
        panic("ZBRPRNStartCleaningSeq is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNStartCleaningSeq is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNStartCleaningSeq),
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
        abort("ZBRPRNStartCleaningSeq",callErr)
    }
    return
}

// extern int ZBRPRNStartSmartCard(IntPtr _handle, int printertype, int cardtype, out int err);
func ZBRPRNStartSmartCard() () {
    if zBRPRNStartSmartCard == 0 {
        panic("ZBRPRNStartSmartCard is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNStartSmartCard is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNStartSmartCard),
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
        abort("ZBRPRNStartSmartCard",callErr)
    }
    return
}

// extern int ZBRPRNSuppressStatusMsgs(IntPtr hprinter, int printertype, int suppressmsgs, out int err);
func ZBRPRNSuppressStatusMsgs() () {
    if zBRPRNSuppressStatusMsgs == 0 {
        panic("ZBRPRNSuppressStatusMsgs is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNSuppressStatusMsgs is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSuppressStatusMsgs),
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
        abort("ZBRPRNSuppressStatusMsgs",callErr)
    }
    return
}

// extern intZBRPRNUpgradeFirmware(HANDLE hPrinter, int printerType, char *filename, int *err);
func ZBRPRNUpgradeFirmware() () {
    if zBRPRNUpgradeFirmware == 0 {
        panic("ZBRPRNUpgradeFirmware is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNUpgradeFirmware is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNUpgradeFirmware),
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
        abort("ZBRPRNUpgradeFirmware",callErr)
    }
    return
}

// extern int ZBRPRNWriteBarcode(IntPtr hprinter, int printertype, int startx, int starty, int rotation, int barcodetype, int barwidthratio, int barcodemultiplier, int barcodeheight, int textunder, byte[] barcodedata, out int err);
func ZBRPRNWriteBarCode() () {
    if zBRPRNWriteBarCode == 0 {
        panic("ZBRPRNWriteBarCode is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNWriteBarCode is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNWriteBarCode),
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
        abort("ZBRPRNWriteBarCode",callErr)
    }
    return
}

// extern int ZBRPRNWriteBox(IntPtr hprinter, int printertype, int startx, int starty, int width, int height, int thickness, out int err);
func ZBRPRNWriteBox() () {
    if zBRPRNWriteBox == 0 {
        panic("ZBRPRNWriteBox is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNWriteBox is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNWriteBox),
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
        abort("ZBRPRNWriteBox",callErr)
    }
    return
}

// extern int ZBRPRNWriteBoxEx(IntPtr hprinter, int printertype, int startx, int starty, int width, int height, int thickness, int gmode, int isvarnish, out int err);
func ZBRPRNWriteBoxEx() () {
    if zBRPRNWriteBoxEx == 0 {
        panic("ZBRPRNWriteBoxEx is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNWriteBoxEx is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNWriteBoxEx),
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
        abort("ZBRPRNWriteBoxEx",callErr)
    }
    return
}

// extern int ZBRPRNWriteMag(IntPtr handle, int printertype, int trkstowrite, byte[] trk1data, byte[] trk2data, byte[] trk3data, out int err);
func ZBRPRNWriteMag() () {
    if zBRPRNWriteMag == 0 {
        panic("ZBRPRNWriteMag is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNWriteMag is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNWriteMag),
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
        abort("ZBRPRNWriteMag",callErr)
    }
    return
}

// extern int ZBRPRNWriteMagByTrk(IntPtr hprinter, int printertype, int trknum, byte[] trkdata, out int err);
func ZBRPRNWriteMagByTrk() () {
    if zBRPRNWriteMagByTrk == 0 {
        panic("ZBRPRNWriteMagByTrk is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNWriteMagByTrk is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNWriteMagByTrk),
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
        abort("ZBRPRNWriteMagByTrk",callErr)
    }
    return
}

// extern int ZBRPRNWriteMagPassThru(IntPtr hdc, int printertype, int trknum, byte[] trkdata, out int err);
func ZBRPRNWriteMagPassThru() () {
    if zBRPRNWriteMagPassThru == 0 {
        panic("ZBRPRNWriteMagPassThru is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNWriteMagPassThru is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNWriteMagPassThru),
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
        abort("ZBRPRNWriteMagPassThru",callErr)
    }
    return
}

// extern int ZBRPRNWriteText(IntPtr hprinter, int printertype, int startx, int starty, int rotation, int isbold, int height, byte[] text, out int err);
func ZBRPRNWriteText() () {
    if zBRPRNWriteText == 0 {
        panic("ZBRPRNWriteText is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNWriteText is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNWriteText),
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
        abort("ZBRPRNWriteText",callErr)
    }
    return
}

// extern int ZBRPRNWriteTextEx(IntPtr hprinter, int printertype, int startx, int starty, int rotation, int isbold, int widht, int height, int gmode, byte[] text, int isvarnish, out int err);
func ZBRPRNWriteTextEx() () {
    if zBRPRNWriteTextEx == 0 {
        panic("ZBRPRNWriteTextEx is not defined. Check library")
    }
    var nargs uintptr = 0
    panic("ZBRPRNWriteTextEx is not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNWriteTextEx),
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
        abort("ZBRPRNWriteTextEx",callErr)
    }
    return
}
