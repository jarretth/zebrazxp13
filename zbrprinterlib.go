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
    zBRPrinter                     , _ = syscall.LoadLibrary("ZBRPrinter.dll")
    zBRCloseHandle                 , _ = syscall.GetProcAddress(zBRPrinter, "ZBRCloseHandle")
    zBRGetHandle                   , _ = syscall.GetProcAddress(zBRPrinter, "ZBRGetHandle")
    zBRPRNCheckForErrors           , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNCheckForErrors")
    zBRPRNChkDueForCleaning        , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNChkDueForCleaning")
    zBRPRNClrErrStatusLn           , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNClrErrStatusLn")
    zBRPRNClrMediaPath             , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNClrMediaPath")
    zBRPRNGetChecksum              , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetChecksum")
    zBRPRNGetCleaningParam         , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetCleaningParam")
    zBRPRNGetMsgSuppressionState   , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetMsgSuppressionState")
    zBRPRNGetOpParam               , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetOpParam")
    zBRPRNGetPanelsPrinted         , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetPanelsPrinted")
    zBRPRNGetPanelsRemaining       , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetPanelsRemaining")
    zBRPRNGetPrintCount            , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetPrintCount")
    zBRPRNGetPrinterOptions        , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetPrinterOptions")
    zBRPRNGetPrinterSerialNumb     , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetPrinterSerialNumb")
    zBRPRNGetPrinterSerialNumber   , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetPrinterSerialNumber")
    zBRPRNGetPrinterStatus         , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetPrinterStatus")
    zBRPRNGetPrintHeadSerialNumb   , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetPrintHeadSerialNumb")
    zBRPRNGetPrintHeadSerialNumber , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetPrintHeadSerialNumber")
    zBRPRNGetPrintStatus           , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetPrintStatus")
    zBRPRNGetSDKProductVer         , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetSDKProductVer")
    zBRPRNGetSDKVer                , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetSDKVer")
    zBRPRNGetSDKVsn                , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetSDKVsn")
    zBRPRNGetSensorStatus          , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNGetSensorStatus")
    zBRPRNImmediateParamSave       , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNImmediateParamSave")
    zBRPRNIsPrinterReady           , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNIsPrinterReady")
    zBRPRNPrintPrnFile             , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNPrintPrnFile")
    zBRPRNResetPrinter             , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNResetPrinter")
    zBRPRNSelfAdj                  , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSelfAdj")
    zBRPRNSetCardFeedingMode       , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetCardFeedingMode")
    zBRPRNSetCleaningParam         , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetCleaningParam")
    zBRPRNSetOverlayMode           , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetOverlayMode")
    zBRPRNSetPrintHeadResistance   , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetPrintHeadResistance")
    zBRPRNSetRelativeXOffset       , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetRelativeXOffset")
    zBRPRNSetRelativeYOffset       , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSetRelativeYOffset")
    zBRPRNStartCleaningCardSeq     , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNStartCleaningCardSeq")
    zBRPRNStartCleaningSeq         , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNStartCleaningSeq")
    zBRPRNSuppressStatusMsgs       , _ = syscall.GetProcAddress(zBRPrinter, "ZBRPRNSuppressStatusMsgs")
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

// void ZBRPRNGetSDKVer(out int major, out int minor, out int engLevel)
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

func ZBRPRNGetSDKVsn() () {
    if zBRPRNGetSDKVsn == 0 {
        panic("ZBRPRNGetSDKVsn not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNGetSDKVsn not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetSDKVsn),
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

func ZBRPRNGetSDKProductVer() () {
    if zBRPRNGetSDKProductVer == 0 {
        panic("ZBRPRNGetSDKProductVer not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNGetSDKProductVer not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetSDKProductVer),
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

func ZBRCloseHandle() () {
    if zBRCloseHandle == 0 {
        panic("ZBRCloseHandle not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRCloseHandle not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRCloseHandle),
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

func ZBRPRNGetMsgSuppressionState() () {
    if zBRPRNGetMsgSuppressionState == 0 {
        panic("ZBRPRNGetMsgSuppressionState not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNGetMsgSuppressionState not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetMsgSuppressionState),
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

func ZBRPRNSuppressStatusMsgs() () {
    if zBRPRNSuppressStatusMsgs == 0 {
        panic("ZBRPRNSuppressStatusMsgs not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNSuppressStatusMsgs not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSuppressStatusMsgs),
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

func ZBRPRNSetOverlayMode() () {
    if zBRPRNSetOverlayMode == 0 {
        panic("ZBRPRNSetOverlayMode not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNSetOverlayMode not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetOverlayMode),
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

func ZBRPRNPrintPrnFile() () {
    if zBRPRNPrintPrnFile == 0 {
        panic("ZBRPRNPrintPrnFile not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNPrintPrnFile not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNPrintPrnFile),
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

func ZBRPRNCheckForErrors() () {
    if zBRPRNCheckForErrors == 0 {
        panic("ZBRPRNCheckForErrors not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNCheckForErrors not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNCheckForErrors),
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

func ZBRPRNClrErrStatusLn() () {
    if zBRPRNClrErrStatusLn == 0 {
        panic("ZBRPRNClrErrStatusLn not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNClrErrStatusLn not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNClrErrStatusLn),
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

func ZBRPRNGetPrintCount() () {
    if zBRPRNGetPrintCount == 0 {
        panic("ZBRPRNGetPrintCount not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNGetPrintCount not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetPrintCount),
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

func ZBRPRNGetPrintStatus() () {
    if zBRPRNGetPrintStatus == 0 {
        panic("ZBRPRNGetPrintStatus not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNGetPrintStatus not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetPrintStatus),
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

func ZBRPRNGetPanelsRemaining() () {
    if zBRPRNGetPanelsRemaining == 0 {
        panic("ZBRPRNGetPanelsRemaining not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNGetPanelsRemaining not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetPanelsRemaining),
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

func ZBRPRNGetPanelsPrinted() () {
    if zBRPRNGetPanelsPrinted == 0 {
        panic("ZBRPRNGetPanelsPrinted not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNGetPanelsPrinted not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetPanelsPrinted),
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

func ZBRPRNGetPrinterSerialNumber() () {
    if zBRPRNGetPrinterSerialNumber == 0 {
        panic("ZBRPRNGetPrinterSerialNumber not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNGetPrinterSerialNumber not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetPrinterSerialNumber),
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

func ZBRPRNGetPrinterSerialNumb() () {
    if zBRPRNGetPrinterSerialNumb == 0 {
        panic("ZBRPRNGetPrinterSerialNumb not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNGetPrinterSerialNumb not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetPrinterSerialNumb),
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

func ZBRPRNGetPrinterOptions() () {
    if zBRPRNGetPrinterOptions == 0 {
        panic("ZBRPRNGetPrinterOptions not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNGetPrinterOptions not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetPrinterOptions),
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

func ZBRPRNGetPrintHeadSerialNumber() () {
    if zBRPRNGetPrintHeadSerialNumber == 0 {
        panic("ZBRPRNGetPrintHeadSerialNumber not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNGetPrintHeadSerialNumber not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetPrintHeadSerialNumber),
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

func ZBRPRNGetPrintHeadSerialNumb() () {
    if zBRPRNGetPrintHeadSerialNumb == 0 {
        panic("ZBRPRNGetPrintHeadSerialNumb not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNGetPrintHeadSerialNumb not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetPrintHeadSerialNumb),
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

func ZBRPRNGetOpParam() () {
    if zBRPRNGetOpParam == 0 {
        panic("ZBRPRNGetOpParam not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNGetOpParam not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetOpParam),
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

func ZBRPRNGetPrinterStatus() () {
    if zBRPRNGetPrinterStatus == 0 {
        panic("ZBRPRNGetPrinterStatus not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNGetPrinterStatus not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetPrinterStatus),
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

func ZBRPRNGetSensorStatus() () {
    if zBRPRNGetSensorStatus == 0 {
        panic("ZBRPRNGetSensorStatus not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNGetSensorStatus not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetSensorStatus),
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

func ZBRPRNIsPrinterReady() () {
    if zBRPRNIsPrinterReady == 0 {
        panic("ZBRPRNIsPrinterReady not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNIsPrinterReady not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNIsPrinterReady),
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

func ZBRPRNChkDueForCleaning() () {
    if zBRPRNChkDueForCleaning == 0 {
        panic("ZBRPRNChkDueForCleaning not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNChkDueForCleaning not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNChkDueForCleaning),
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

func ZBRPRNStartCleaningSeq() () {
    if zBRPRNStartCleaningSeq == 0 {
        panic("ZBRPRNStartCleaningSeq not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNStartCleaningSeq not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNStartCleaningSeq),
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

func ZBRPRNStartCleaningCardSeq() () {
    if zBRPRNStartCleaningCardSeq == 0 {
        panic("ZBRPRNStartCleaningCardSeq not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNStartCleaningCardSeq not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNStartCleaningCardSeq),
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

func ZBRPRNGetCleaningParam() () {
    if zBRPRNGetCleaningParam == 0 {
        panic("ZBRPRNGetCleaningParam not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNGetCleaningParam not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetCleaningParam),
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

func ZBRPRNSetCleaningParam() () {
    if zBRPRNSetCleaningParam == 0 {
        panic("ZBRPRNSetCleaningParam not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNSetCleaningParam not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetCleaningParam),
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

func ZBRPRNResetPrinter() () {
    if zBRPRNResetPrinter == 0 {
        panic("ZBRPRNResetPrinter not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNResetPrinter not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNResetPrinter),
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

func ZBRPRNSelfAdj() () {
    if zBRPRNSelfAdj == 0 {
        panic("ZBRPRNSelfAdj not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNSelfAdj not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSelfAdj),
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

func ZBRPRNGetChecksum() () {
    if zBRPRNGetChecksum == 0 {
        panic("ZBRPRNGetChecksum not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNGetChecksum not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNGetChecksum),
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

func ZBRPRNSetCardFeedingMode() () {
    if zBRPRNSetCardFeedingMode == 0 {
        panic("ZBRPRNSetCardFeedingMode not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNSetCardFeedingMode not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetCardFeedingMode),
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

func ZBRPRNSetPrintHeadResistance() () {
    if zBRPRNSetPrintHeadResistance == 0 {
        panic("ZBRPRNSetPrintHeadResistance not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNSetPrintHeadResistance not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetPrintHeadResistance),
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

func ZBRPRNClrMediaPath() () {
    if zBRPRNClrMediaPath == 0 {
        panic("ZBRPRNClrMediaPath not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNClrMediaPath not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNClrMediaPath),
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

func ZBRPRNImmediateParamSave() () {
    if zBRPRNImmediateParamSave == 0 {
        panic("ZBRPRNImmediateParamSave not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNImmediateParamSave not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNImmediateParamSave),
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

func ZBRPRNSetRelativeXOffset() () {
    if zBRPRNSetRelativeXOffset == 0 {
        panic("ZBRPRNSetRelativeXOffset not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNSetRelativeXOffset not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetRelativeXOffset),
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

func ZBRPRNSetRelativeYOffset() () {
    if zBRPRNSetRelativeYOffset == 0 {
        panic("ZBRPRNSetRelativeYOffset not defined. Check library")
    }
    var nargs uintptr = 0
    abort("Call ZBRPRNSetRelativeYOffset not implemented")
    _, _, callErr := syscall.Syscall9(uintptr(zBRPRNSetRelativeYOffset),
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
