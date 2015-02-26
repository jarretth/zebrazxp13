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
    zBRPrinter                        , _ = syscall.NewLazyDLL("ZBRPrinter.dll")
    zBRCloseHandle                    , _ = zBRPrinter.NewProc("ZBRCloseHandle")
    zBRGetHandle                      , _ = zBRPrinter.NewProc("ZBRGetHandle")
    zBRPRNCheckForErrors              , _ = zBRPrinter.NewProc("ZBRPRNCheckForErrors")
    zBRPRNChkDueForCleaning           , _ = zBRPrinter.NewProc("ZBRPRNChkDueForCleaning")
    zBRPRNClrColorImgBuf              , _ = zBRPrinter.NewProc("ZBRPRNClrColorImgBuf")
    zBRPRNClrColorImgBufs             , _ = zBRPrinter.NewProc("ZBRPRNClrColorImgBufs")
    zBRPRNClrErrStatusLn              , _ = zBRPrinter.NewProc("ZBRPRNClrErrStatusLn")
    zBRPRNClrMediaPath                , _ = zBRPrinter.NewProc("ZBRPRNClrMediaPath")
    zBRPRNClrMonoImgBuf               , _ = zBRPrinter.NewProc("ZBRPRNClrMonoImgBuf")
    zBRPRNClrMonoImgBufs              , _ = zBRPrinter.NewProc("ZBRPRNClrMonoImgBufs")
    zBRPRNClrSpecifiedBmp             , _ = zBRPrinter.NewProc("ZBRPRNClrSpecifiedBmp")
    zBRPRNEjectCard                   , _ = zBRPrinter.NewProc("ZBRPRNEjectCard")
    zBRPRNEnableMagReadDataEncryption , _ = zBRPrinter.NewProc("ZBRPRNEnableMagReadDataEncryption")
    zBRPRNEndSmartCard                , _ = zBRPrinter.NewProc("ZBRPRNEndSmartCard")
    zBRPRNFlipCard                    , _ = zBRPrinter.NewProc("ZBRPRNFlipCard")
    zBRPRNGetChecksum                 , _ = zBRPrinter.NewProc("ZBRPRNGetChecksum")
    zBRPRNGetCleaningParam            , _ = zBRPrinter.NewProc("ZBRPRNGetCleaningParam")
    zBRPRNGetIPAddress                , _ = zBRPrinter.NewProc("ZBRPRNGetIPAddress")
    zBRPRNGetMsgSuppressionState      , _ = zBRPrinter.NewProc("ZBRPRNGetMsgSuppressionState")
    zBRPRNGetOpParam                  , _ = zBRPrinter.NewProc("ZBRPRNGetOpParam")
    zBRPRNGetPanelsPrinted            , _ = zBRPrinter.NewProc("ZBRPRNGetPanelsPrinted")
    zBRPRNGetPanelsRemaining          , _ = zBRPrinter.NewProc("ZBRPRNGetPanelsRemaining")
    zBRPRNGetPrintCount               , _ = zBRPrinter.NewProc("ZBRPRNGetPrintCount")
    zBRPRNGetPrinterOptions           , _ = zBRPrinter.NewProc("ZBRPRNGetPrinterOptions")
    zBRPRNGetPrinterSerialNumb        , _ = zBRPrinter.NewProc("ZBRPRNGetPrinterSerialNumb")
    zBRPRNGetPrinterSerialNumber      , _ = zBRPrinter.NewProc("ZBRPRNGetPrinterSerialNumber")
    zBRPRNGetPrinterStatus            , _ = zBRPrinter.NewProc("ZBRPRNGetPrinterStatus")
    zBRPRNGetPrintHeadSerialNumb      , _ = zBRPrinter.NewProc("ZBRPRNGetPrintHeadSerialNumb")
    zBRPRNGetPrintHeadSerialNumber    , _ = zBRPrinter.NewProc("ZBRPRNGetPrintHeadSerialNumber")
    zBRPRNGetPrintStatus              , _ = zBRPrinter.NewProc("ZBRPRNGetPrintStatus")
    zBRPRNGetSDKProductVer            , _ = zBRPrinter.NewProc("ZBRPRNGetSDKProductVer")
    zBRPRNGetSDKVer                   , _ = zBRPrinter.NewProc("ZBRPRNGetSDKVer")
    zBRPRNGetSDKVsn                   , _ = zBRPrinter.NewProc("ZBRPRNGetSDKVsn")
    zBRPRNGetSensorStatus             , _ = zBRPrinter.NewProc("ZBRPRNGetSensorStatus")
    zBRPRNImmediateParamSave          , _ = zBRPrinter.NewProc("ZBRPRNImmediateParamSave")
    zBRPRNIsPrinterReady              , _ = zBRPrinter.NewProc("ZBRPRNIsPrinterReady")
    zBRPRNMoveCard                    , _ = zBRPrinter.NewProc("ZBRPRNMoveCard")
    zBRPRNMoveCardBkwd                , _ = zBRPrinter.NewProc("ZBRPRNMoveCardBkwd")
    zBRPRNMoveCardFwd                 , _ = zBRPrinter.NewProc("ZBRPRNMoveCardFwd")
    zBRPRNMovePrintReady              , _ = zBRPrinter.NewProc("ZBRPRNMovePrintReady")
    zBRPRNPrintCardPanel              , _ = zBRPrinter.NewProc("ZBRPRNPrintCardPanel")
    zBRPRNPrintClearVarnish           , _ = zBRPrinter.NewProc("ZBRPRNPrintClearVarnish")
    zBRPRNPrintColorImgBuf            , _ = zBRPrinter.NewProc("ZBRPRNPrintColorImgBuf")
    zBRPRNPrintHologramOverlay        , _ = zBRPrinter.NewProc("ZBRPRNPrintHologramOverlay")
    zBRPRNPrintMonoImgBuf             , _ = zBRPrinter.NewProc("ZBRPRNPrintMonoImgBuf")
    zBRPRNPrintMonoImgBufEx           , _ = zBRPrinter.NewProc("ZBRPRNPrintMonoImgBufEx")
    zBRPRNPrintMonoPanel              , _ = zBRPrinter.NewProc("ZBRPRNPrintMonoPanel")
    zBRPRNPrintPrnFile                , _ = zBRPrinter.NewProc("ZBRPRNPrintPrnFile")
    zBRPRNPrintTestCard               , _ = zBRPrinter.NewProc("ZBRPRNPrintTestCard")
    zBRPRNPrintVarnish                , _ = zBRPrinter.NewProc("ZBRPRNPrintVarnish")
    zBRPRNPrintVarnishEx              , _ = zBRPrinter.NewProc("ZBRPRNPrintVarnishEx")
    zBRPRNReadMag                     , _ = zBRPrinter.NewProc("ZBRPRNReadMag")
    zBRPRNReadMagByTrk                , _ = zBRPrinter.NewProc("ZBRPRNReadMagByTrk")
    zBRPRNResetMagEncoder             , _ = zBRPrinter.NewProc("ZBRPRNResetMagEncoder")
    zBRPRNResetPrinter                , _ = zBRPrinter.NewProc("ZBRPRNResetPrinter")
    zBRPRNResync                      , _ = zBRPrinter.NewProc("ZBRPRNResync")
    zBRPRNReversePrintReady           , _ = zBRPrinter.NewProc("ZBRPRNReversePrintReady")
    zBRPRNSelfAdj                     , _ = zBRPrinter.NewProc("ZBRPRNSelfAdj")
    zBRPRNSetCardFeedingMode          , _ = zBRPrinter.NewProc("ZBRPRNSetCardFeedingMode")
    zBRPRNSetCleaningParam            , _ = zBRPrinter.NewProc("ZBRPRNSetCleaningParam")
    zBRPRNSetColorContrast            , _ = zBRPrinter.NewProc("ZBRPRNSetColorContrast")
    zBRPRNSetContrastIntensityLvl     , _ = zBRPrinter.NewProc("ZBRPRNSetContrastIntensityLvl")
    zBRPRNSetEncoderCoercivity        , _ = zBRPrinter.NewProc("ZBRPRNSetEncoderCoercivity")
    zBRPRNSetEncodingDir              , _ = zBRPrinter.NewProc("ZBRPRNSetEncodingDir")
    zBRPRNSetEndOfPrint               , _ = zBRPrinter.NewProc("ZBRPRNSetEndOfPrint")
    zBRPRNSetHologramIntensity        , _ = zBRPrinter.NewProc("ZBRPRNSetHologramIntensity")
    zBRPRNSetMagEncodingStd           , _ = zBRPrinter.NewProc("ZBRPRNSetMagEncodingStd")
    zBRPRNSetMonoContrast             , _ = zBRPrinter.NewProc("ZBRPRNSetMonoContrast")
    zBRPRNSetMonoIntensity            , _ = zBRPrinter.NewProc("ZBRPRNSetMonoIntensity")
    zBRPRNSetOverlayMode              , _ = zBRPrinter.NewProc("ZBRPRNSetOverlayMode")
    zBRPRNSetPrintHeadResistance      , _ = zBRPrinter.NewProc("ZBRPRNSetPrintHeadResistance")
    zBRPRNSetRelativeXOffset          , _ = zBRPrinter.NewProc("ZBRPRNSetRelativeXOffset")
    zBRPRNSetRelativeYOffset          , _ = zBRPrinter.NewProc("ZBRPRNSetRelativeYOffset")
    zBRPRNSetStartPrintSideBOffset    , _ = zBRPrinter.NewProc("ZBRPRNSetStartPrintSideBOffset")
    zBRPRNSetStartPrintSideBXOffset   , _ = zBRPrinter.NewProc("ZBRPRNSetStartPrintSideBXOffset")
    zBRPRNSetStartPrintSideBYOffset   , _ = zBRPrinter.NewProc("ZBRPRNSetStartPrintSideBYOffset")
    zBRPRNSetStartPrintXOffset        , _ = zBRPrinter.NewProc("ZBRPRNSetStartPrintXOffset")
    zBRPRNSetStartPrintYOffset        , _ = zBRPrinter.NewProc("ZBRPRNSetStartPrintYOffset")
    zBRPRNSetTrkDensity               , _ = zBRPrinter.NewProc("ZBRPRNSetTrkDensity")
    zBRPRNStartCleaningCardSeq        , _ = zBRPrinter.NewProc("ZBRPRNStartCleaningCardSeq")
    zBRPRNStartCleaningSeq            , _ = zBRPrinter.NewProc("ZBRPRNStartCleaningSeq")
    zBRPRNStartSmartCard              , _ = zBRPrinter.NewProc("ZBRPRNStartSmartCard")
    zBRPRNSuppressStatusMsgs          , _ = zBRPrinter.NewProc("ZBRPRNSuppressStatusMsgs")
    zBRPRNUpgradeFirmware             , _ = zBRPrinter.NewProc("ZBRPRNUpgradeFirmware")
    zBRPRNWriteBarCode                , _ = zBRPrinter.NewProc("ZBRPRNWriteBarCode")
    zBRPRNWriteBox                    , _ = zBRPrinter.NewProc("ZBRPRNWriteBox")
    zBRPRNWriteBoxEx                  , _ = zBRPrinter.NewProc("ZBRPRNWriteBoxEx")
    zBRPRNWriteMag                    , _ = zBRPrinter.NewProc("ZBRPRNWriteMag")
    zBRPRNWriteMagByTrk               , _ = zBRPrinter.NewProc("ZBRPRNWriteMagByTrk")
    zBRPRNWriteMagPassThru            , _ = zBRPrinter.NewProc("ZBRPRNWriteMagPassThru")
    zBRPRNWriteText                   , _ = zBRPrinter.NewProc("ZBRPRNWriteText")
    zBRPRNWriteTextEx                 , _ = zBRPrinter.NewProc("ZBRPRNWriteTextEx")
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
func ZBRCloseHandle() (ret uintptr) {
    panic("ZBRCloseHandle not implemented")
    ret, _, _ = zBRCloseHandle.Call(
        )
    return ret
}

// extern int ZBRGetHandle(out IntPtr handle, byte[] drvname, out int printertype, out int err);
func ZBRGetHandle() (ret uintptr) {
    panic("ZBRGetHandle not implemented")
    ret, _, _ = zBRGetHandle.Call(
        )
    return ret
}

// extern int ZBRPRNCheckForErrors(IntPtr hprinter, int printertype);
func ZBRPRNCheckForErrors() (ret uintptr) {
    panic("ZBRPRNCheckForErrors not implemented")
    ret, _, _ = zBRPRNCheckForErrors.Call(
        )
    return ret
}

// extern int ZBRPRNChkDueForCleaning(IntPtr hprinter, int printertype, out int imgcounter, out int cleancounter, out int cleancardcounter, out int err);
func ZBRPRNChkDueForCleaning() (ret uintptr) {
    panic("ZBRPRNChkDueForCleaning not implemented")
    ret, _, _ = zBRPRNChkDueForCleaning.Call(
        )
    return ret
}

// extern int ZBRPRNClrColorImgBuf(IntPtr hprinter, int printertype, int colorbufidx, out int err);
func ZBRPRNClrColorImgBuf() (ret uintptr) {
    panic("ZBRPRNClrColorImgBuf not implemented")
    ret, _, _ = zBRPRNClrColorImgBuf.Call(
        )
    return ret
}

// extern int ZBRPRNClrColorImgBufs(HANDLE hPrinter, int printerType, int *err)
func ZBRPRNClrColorImgBufs() (ret uintptr) {
    panic("ZBRPRNClrColorImgBufs not implemented")
    ret, _, _ = zBRPRNClrColorImgBufs.Call(
        )
    return ret
}

// extern int ZBRPRNClrErrStatusLn(IntPtr hprinter, int printertype, out int err);
func ZBRPRNClrErrStatusLn() (ret uintptr) {
    panic("ZBRPRNClrErrStatusLn not implemented")
    ret, _, _ = zBRPRNClrErrStatusLn.Call(
        )
    return ret
}

// extern int ZBRPRNClrMediaPath(IntPtr hprinter, int printertype, out int err);
func ZBRPRNClrMediaPath() (ret uintptr) {
    panic("ZBRPRNClrMediaPath not implemented")
    ret, _, _ = zBRPRNClrMediaPath.Call(
        )
    return ret
}

// extern int ZBRPRNClrMonoImgBuf(IntPtr hprinter, int printertype, int clrvarnish, out int err);
func ZBRPRNClrMonoImgBuf() (ret uintptr) {
    panic("ZBRPRNClrMonoImgBuf not implemented")
    ret, _, _ = zBRPRNClrMonoImgBuf.Call(
        )
    return ret
}

// extern int ZBRPRNClrMonoImgBufs(IntPtr hprinter, int printertype, int clrbuffer, out int err);
func ZBRPRNClrMonoImgBufs() (ret uintptr) {
    panic("ZBRPRNClrMonoImgBufs not implemented")
    ret, _, _ = zBRPRNClrMonoImgBufs.Call(
        )
    return ret
}

// extern int ZBRPRNClrSpecifiedBmp(IntPtr hprinter, int printertype, int colorbufidx, out int err);
func ZBRPRNClrSpecifiedBmp() (ret uintptr) {
    panic("ZBRPRNClrSpecifiedBmp not implemented")
    ret, _, _ = zBRPRNClrSpecifiedBmp.Call(
        )
    return ret
}

// extern int ZBRPRNEjectCard(IntPtr _handle, int prn_type, out int err);
func ZBRPRNEjectCard() (ret uintptr) {
    panic("ZBRPRNEjectCard not implemented")
    ret, _, _ = zBRPRNEjectCard.Call(
        )
    return ret
}

// extern int ZBRPRNEnableMagReadDataEncryption(IntPtr handle, int printertype, out int err);
func ZBRPRNEnableMagReadDataEncryption() (ret uintptr) {
    panic("ZBRPRNEnableMagReadDataEncryption not implemented")
    ret, _, _ = zBRPRNEnableMagReadDataEncryption.Call(
        )
    return ret
}

// extern int ZBRPRNEndSmartCard(IntPtr _handle, int printertype, int cardtype, int movement, out int err);
func ZBRPRNEndSmartCard() (ret uintptr) {
    panic("ZBRPRNEndSmartCard not implemented")
    ret, _, _ = zBRPRNEndSmartCard.Call(
        )
    return ret
}

// extern int ZBRPRNFlipCard(IntPtr hprinter, int printertype, out int err);
func ZBRPRNFlipCard() (ret uintptr) {
    panic("ZBRPRNFlipCard not implemented")
    ret, _, _ = zBRPRNFlipCard.Call(
        )
    return ret
}

// extern int ZBRPRNGetChecksum(IntPtr hprinter, int printertype, out int checksum, out int err);
func ZBRPRNGetChecksum() (ret uintptr) {
    panic("ZBRPRNGetChecksum not implemented")
    ret, _, _ = zBRPRNGetChecksum.Call(
        )
    return ret
}

// extern int ZBRPRNGetCleaningParam(IntPtr hprinter, int printertype, out int imgcounter, out int cleancounter, out int cleancardcounter, out int err);
func ZBRPRNGetCleaningParam() (ret uintptr) {
    panic("ZBRPRNGetCleaningParam not implemented")
    ret, _, _ = zBRPRNGetCleaningParam.Call(
        )
    return ret
}

// extern int ZBRPRNGetIPAddress(string PrinterName, char* ipAddress);
func ZBRPRNGetIPAddress() (ret uintptr) {
    panic("ZBRPRNGetIPAddress not implemented")
    ret, _, _ = zBRPRNGetIPAddress.Call(
        )
    return ret
}

// extern int ZBRPRNGetMsgSuppressionState(IntPtr hprinter, int printertype, out int state, out int err);
func ZBRPRNGetMsgSuppressionState() (ret uintptr) {
    panic("ZBRPRNGetMsgSuppressionState not implemented")
    ret, _, _ = zBRPRNGetMsgSuppressionState.Call(
        )
    return ret
}

// extern int ZBRPRNGetOpParam(IntPtr hprinter, int printertype, int paramidx, byte[] opparam, out int respsize, out int err);
func ZBRPRNGetOpParam() (ret uintptr) {
    panic("ZBRPRNGetOpParam not implemented")
    ret, _, _ = zBRPRNGetOpParam.Call(
        )
    return ret
}

// extern int ZBRPRNGetPanelsPrinted(IntPtr hprinter, int printertype, out int panels, out int err);
func ZBRPRNGetPanelsPrinted() (ret uintptr) {
    panic("ZBRPRNGetPanelsPrinted not implemented")
    ret, _, _ = zBRPRNGetPanelsPrinted.Call(
        )
    return ret
}

// extern int ZBRPRNGetPanelsRemaining(IntPtr hprinter, int printertype, out int panels, out int err);
func ZBRPRNGetPanelsRemaining() (ret uintptr) {
    panic("ZBRPRNGetPanelsRemaining not implemented")
    ret, _, _ = zBRPRNGetPanelsRemaining.Call(
        )
    return ret
}

// extern int ZBRPRNGetPrintCount(IntPtr hprinter, int printertype, out int printcount, out int err);
func ZBRPRNGetPrintCount() (ret uintptr) {
    panic("ZBRPRNGetPrintCount not implemented")
    ret, _, _ = zBRPRNGetPrintCount.Call(
        )
    return ret
}

// extern int ZBRPRNGetPrinterOptions(IntPtr hprinter, int printertype, byte[] options, out int respsize, out int err);
func ZBRPRNGetPrinterOptions() (ret uintptr) {
    panic("ZBRPRNGetPrinterOptions not implemented")
    ret, _, _ = zBRPRNGetPrinterOptions.Call(
        )
    return ret
}

// extern int ZBRPRNGetPrinterSerialNumb(IntPtr hprinter, int printertype, byte[] serialnum, out int respsize, out int err);
func ZBRPRNGetPrinterSerialNumb() (ret uintptr) {
    panic("ZBRPRNGetPrinterSerialNumb not implemented")
    ret, _, _ = zBRPRNGetPrinterSerialNumb.Call(
        )
    return ret
}

// extern int ZBRPRNGetPrinterSerialNumber(IntPtr hprinter, int printertype, byte[] serialnum, out int respsize, out int err);
func ZBRPRNGetPrinterSerialNumber() (ret uintptr) {
    panic("ZBRPRNGetPrinterSerialNumber not implemented")
    ret, _, _ = zBRPRNGetPrinterSerialNumber.Call(
        )
    return ret
}

// extern int ZBRPRNGetPrinterStatus(out int statuscode);
func ZBRPRNGetPrinterStatus() (ret uintptr) {
    panic("ZBRPRNGetPrinterStatus not implemented")
    ret, _, _ = zBRPRNGetPrinterStatus.Call(
        )
    return ret
}

// extern int ZBRPRNGetPrintHeadSerialNumb(IntPtr hprinter, int printertype, byte[] serialnum, out int respsize, out int err);
func ZBRPRNGetPrintHeadSerialNumb() (ret uintptr) {
    panic("ZBRPRNGetPrintHeadSerialNumb not implemented")
    ret, _, _ = zBRPRNGetPrintHeadSerialNumb.Call(
        )
    return ret
}

// extern int ZBRPRNGetPrintHeadSerialNumber(IntPtr hprinter, int printertype, byte[] serialnum, out int respsize, out int err);
func ZBRPRNGetPrintHeadSerialNumber() (ret uintptr) {
    panic("ZBRPRNGetPrintHeadSerialNumber not implemented")
    ret, _, _ = zBRPRNGetPrintHeadSerialNumber.Call(
        )
    return ret
}

// extern int ZBRPRNGetPrintStatus(IntPtr hprinter, int printertype);
func ZBRPRNGetPrintStatus() (ret uintptr) {
    panic("ZBRPRNGetPrintStatus not implemented")
    ret, _, _ = zBRPRNGetPrintStatus.Call(
        )
    return ret
}

// extern int ZBRPRNGetSDKProductVer(byte[] productversion, out int buffsize);
func ZBRPRNGetSDKProductVer() (ret uintptr) {
    panic("ZBRPRNGetSDKProductVer not implemented")
    ret, _, _ = zBRPRNGetSDKProductVer.Call(
        )
    return ret
}

// extern void ZBRPRNGetSDKVer(out int major, out int minor, out int englevel);
func ZBRPRNGetSDKVer() (ret uintptr) {
    panic("ZBRPRNGetSDKVer not implemented")
    ret, _, _ = zBRPRNGetSDKVer.Call(
        )
    return ret
}

// extern void ZBRPRNGetSDKVsn(out int major, out int minor, out int englevel);
func ZBRPRNGetSDKVsn() (ret uintptr) {
    panic("ZBRPRNGetSDKVsn not implemented")
    ret, _, _ = zBRPRNGetSDKVsn.Call(
        )
    return ret
}

// extern int ZBRPRNGetSensorStatus(IntPtr hprinter, int printertype, out byte status, out int err);
func ZBRPRNGetSensorStatus() (ret uintptr) {
    panic("ZBRPRNGetSensorStatus not implemented")
    ret, _, _ = zBRPRNGetSensorStatus.Call(
        )
    return ret
}

// extern int ZBRPRNImmediateParamSave(IntPtr hprinter, int printertype, out int err);
func ZBRPRNImmediateParamSave() (ret uintptr) {
    panic("ZBRPRNImmediateParamSave not implemented")
    ret, _, _ = zBRPRNImmediateParamSave.Call(
        )
    return ret
}

// extern int ZBRPRNIsPrinterReady(IntPtr hprinter, int printertype, out int err);
func ZBRPRNIsPrinterReady() (ret uintptr) {
    panic("ZBRPRNIsPrinterReady not implemented")
    ret, _, _ = zBRPRNIsPrinterReady.Call(
        )
    return ret
}

// extern int ZBRPRNMoveCard(IntPtr hprinter, int printertype, int steps, out int err);
func ZBRPRNMoveCard() (ret uintptr) {
    panic("ZBRPRNMoveCard not implemented")
    ret, _, _ = zBRPRNMoveCard.Call(
        )
    return ret
}

// extern int ZBRPRNMoveCardBkwd(IntPtr hprinter, int printertype, int steps, out int err);
func ZBRPRNMoveCardBkwd() (ret uintptr) {
    panic("ZBRPRNMoveCardBkwd not implemented")
    ret, _, _ = zBRPRNMoveCardBkwd.Call(
        )
    return ret
}

// extern int ZBRPRNMoveCardFwd(IntPtr hprinter, int printertype, int steps, out int err);
func ZBRPRNMoveCardFwd() (ret uintptr) {
    panic("ZBRPRNMoveCardFwd not implemented")
    ret, _, _ = zBRPRNMoveCardFwd.Call(
        )
    return ret
}

// extern int ZBRPRNMovePrintReady(IntPtr handle, int printertype, out int err);
func ZBRPRNMovePrintReady() (ret uintptr) {
    panic("ZBRPRNMovePrintReady not implemented")
    ret, _, _ = zBRPRNMovePrintReady.Call(
        )
    return ret
}

// extern int ZBRPRNPrintCardPanel(IntPtr hprinter, int printertype, int imgbufidx, out int err);
func ZBRPRNPrintCardPanel() (ret uintptr) {
    panic("ZBRPRNPrintCardPanel not implemented")
    ret, _, _ = zBRPRNPrintCardPanel.Call(
        )
    return ret
}

// extern int ZBRPRNPrintClearVarnish(IntPtr hprinter, int printertype, out int err);
func ZBRPRNPrintClearVarnish() (ret uintptr) {
    panic("ZBRPRNPrintClearVarnish not implemented")
    ret, _, _ = zBRPRNPrintClearVarnish.Call(
        )
    return ret
}

// extern int ZBRPRNPrintColorImgBuf(IntPtr hprinter, int printertype, int imgbufidx, out int err);
func ZBRPRNPrintColorImgBuf() (ret uintptr) {
    panic("ZBRPRNPrintColorImgBuf not implemented")
    ret, _, _ = zBRPRNPrintColorImgBuf.Call(
        )
    return ret
}

// extern int ZBRPRNPrintHologramOverlay(IntPtr hprinter, int printertype, int printcardcommand, out int err);
func ZBRPRNPrintHologramOverlay() (ret uintptr) {
    panic("ZBRPRNPrintHologramOverlay not implemented")
    ret, _, _ = zBRPRNPrintHologramOverlay.Call(
        )
    return ret
}

// extern int ZBRPRNPrintMonoImgBuf(IntPtr hprinter, int printertype, out int err);
func ZBRPRNPrintMonoImgBuf() (ret uintptr) {
    panic("ZBRPRNPrintMonoImgBuf not implemented")
    ret, _, _ = zBRPRNPrintMonoImgBuf.Call(
        )
    return ret
}

// extern int ZBRPRNPrintMonoImgBufEx(IntPtr hprinter, int printertype, int cardcommand, out int err);
func ZBRPRNPrintMonoImgBufEx() (ret uintptr) {
    panic("ZBRPRNPrintMonoImgBufEx not implemented")
    ret, _, _ = zBRPRNPrintMonoImgBufEx.Call(
        )
    return ret
}

// extern int ZBRPRNPrintMonoPanel(IntPtr hprinter, int printertype, out int err);
func ZBRPRNPrintMonoPanel() (ret uintptr) {
    panic("ZBRPRNPrintMonoPanel not implemented")
    ret, _, _ = zBRPRNPrintMonoPanel.Call(
        )
    return ret
}

// extern int ZBRPRNPrintPRNFile(IntPtr hprinter, int printertype, byte[] filename, out int err);
func ZBRPRNPrintPRNFile() (ret uintptr) {
    panic("ZBRPRNPrintPRNFile not implemented")
    ret, _, _ = zBRPRNPrintPRNFile.Call(
        )
    return ret
}

// extern int ZBRPRNPrintTestCard(IntPtr hprinter, int printertype, int testcardtype, out int err);
func ZBRPRNPrintTestCard() (ret uintptr) {
    panic("ZBRPRNPrintTestCard not implemented")
    ret, _, _ = zBRPRNPrintTestCard.Call(
        )
    return ret
}

// extern int ZBRPRNPrintVarnish(IntPtr hprinter, int printertype, out int err);
func ZBRPRNPrintVarnish() (ret uintptr) {
    panic("ZBRPRNPrintVarnish not implemented")
    ret, _, _ = zBRPRNPrintVarnish.Call(
        )
    return ret
}

// extern int ZBRPRNPrintVarnishEx(IntPtr hprinter, int printertype, int printcardcommand, out int err);
func ZBRPRNPrintVarnishEx() (ret uintptr) {
    panic("ZBRPRNPrintVarnishEx not implemented")
    ret, _, _ = zBRPRNPrintVarnishEx.Call(
        )
    return ret
}

// extern int ZBRPRNReadMag(IntPtr handle, int printertype, int trkstoread, byte[] trk1buf, out int trk1bytesneeded, byte[] trk2buf, out int trk2bytesneeded, byte[] trk3buf, out int trk3bytesneeded, out int err);
func ZBRPRNReadMag() (ret uintptr) {
    panic("ZBRPRNReadMag not implemented")
    ret, _, _ = zBRPRNReadMag.Call(
        )
    return ret
}

// extern int ZBRPRNReadMagByTrk(IntPtr hprinter, int printertype, int trknum, byte[] trkbuf, out int respsize, out int err);
func ZBRPRNReadMagByTrk() (ret uintptr) {
    panic("ZBRPRNReadMagByTrk not implemented")
    ret, _, _ = zBRPRNReadMagByTrk.Call(
        )
    return ret
}

// extern int ZBRPRNResetMagEncoder(IntPtr hprinter, int printertype, out int err);
func ZBRPRNResetMagEncoder() (ret uintptr) {
    panic("ZBRPRNResetMagEncoder not implemented")
    ret, _, _ = zBRPRNResetMagEncoder.Call(
        )
    return ret
}

// extern int ZBRPRNResetPrinter(IntPtr hprinter, int printertype, out int err);
func ZBRPRNResetPrinter() (ret uintptr) {
    panic("ZBRPRNResetPrinter not implemented")
    ret, _, _ = zBRPRNResetPrinter.Call(
        )
    return ret
}

// extern int ZBRPRNResync(IntPtr hprinter, int printertype, out int err);
func ZBRPRNResync() (ret uintptr) {
    panic("ZBRPRNResync not implemented")
    ret, _, _ = zBRPRNResync.Call(
        )
    return ret
}

// extern int ZBRPRNReversePrintReady(IntPtr hprinter, int printertype, out int err);
func ZBRPRNReversePrintReady() (ret uintptr) {
    panic("ZBRPRNReversePrintReady not implemented")
    ret, _, _ = zBRPRNReversePrintReady.Call(
        )
    return ret
}

// extern int ZBRPRNSelfAdj(IntPtr hprinter, int printertype, out int err);
func ZBRPRNSelfAdj() (ret uintptr) {
    panic("ZBRPRNSelfAdj not implemented")
    ret, _, _ = zBRPRNSelfAdj.Call(
        )
    return ret
}

// extern int ZBRPRNSetCardFeedingMode(IntPtr hprinter, int printertype, int mode, out int err);
func ZBRPRNSetCardFeedingMode() (ret uintptr) {
    panic("ZBRPRNSetCardFeedingMode not implemented")
    ret, _, _ = zBRPRNSetCardFeedingMode.Call(
        )
    return ret
}

// extern int ZBRPRNSetCleaningParam(IntPtr hprinter, int printertype, int ribbonpanelcounter, int cleancardpass, out int err);
func ZBRPRNSetCleaningParam() (ret uintptr) {
    panic("ZBRPRNSetCleaningParam not implemented")
    ret, _, _ = zBRPRNSetCleaningParam.Call(
        )
    return ret
}

// extern int ZBRPRNSetColorContrast(IntPtr hprinter, int printertype, int imgbufidx, int contrast, out int err);
func ZBRPRNSetColorContrast() (ret uintptr) {
    panic("ZBRPRNSetColorContrast not implemented")
    ret, _, _ = zBRPRNSetColorContrast.Call(
        )
    return ret
}

// extern int ZBRPRNSetContrastIntensityLvl(IntPtr hprinter, int printertype, int imgbufidx, int intensity, out int err);
func ZBRPRNSetContrastIntensityLvl() (ret uintptr) {
    panic("ZBRPRNSetContrastIntensityLvl not implemented")
    ret, _, _ = zBRPRNSetContrastIntensityLvl.Call(
        )
    return ret
}

// extern int ZBRPRNSetEncoderCoercivity(IntPtr hprinter, int printertype, int coercivity, out int err);
func ZBRPRNSetEncoderCoercivity() (ret uintptr) {
    panic("ZBRPRNSetEncoderCoercivity not implemented")
    ret, _, _ = zBRPRNSetEncoderCoercivity.Call(
        )
    return ret
}

// extern int ZBRPRNSetEncodingDir(IntPtr hprinter, int printertype, int dir, out int err);
func ZBRPRNSetEncodingDir() (ret uintptr) {
    panic("ZBRPRNSetEncodingDir not implemented")
    ret, _, _ = zBRPRNSetEncodingDir.Call(
        )
    return ret
}

// extern int ZBRPRNSetEndOfPrint(IntPtr hprinter, int printertype, int xwidth, out int err);
func ZBRPRNSetEndOfPrint() (ret uintptr) {
    panic("ZBRPRNSetEndOfPrint not implemented")
    ret, _, _ = zBRPRNSetEndOfPrint.Call(
        )
    return ret
}

// extern int ZBRPRNSetHologramIntensity(IntPtr hprinter, int printertype, int intensity, out int err);
func ZBRPRNSetHologramIntensity() (ret uintptr) {
    panic("ZBRPRNSetHologramIntensity not implemented")
    ret, _, _ = zBRPRNSetHologramIntensity.Call(
        )
    return ret
}

// extern int ZBRPRNSetMagEncodingStd(IntPtr hprinter, int printertype, int standard, out int err);
func ZBRPRNSetMagEncodingStd() (ret uintptr) {
    panic("ZBRPRNSetMagEncodingStd not implemented")
    ret, _, _ = zBRPRNSetMagEncodingStd.Call(
        )
    return ret
}

// extern int ZBRPRNSetMonoContrast(IntPtr hprinter, int printertype, int contrast, out int err);
func ZBRPRNSetMonoContrast() (ret uintptr) {
    panic("ZBRPRNSetMonoContrast not implemented")
    ret, _, _ = zBRPRNSetMonoContrast.Call(
        )
    return ret
}

// extern int ZBRPRNSetMonoIntensity(IntPtr hprinter, int printertype, int intensity, out int err);
func ZBRPRNSetMonoIntensity() (ret uintptr) {
    panic("ZBRPRNSetMonoIntensity not implemented")
    ret, _, _ = zBRPRNSetMonoIntensity.Call(
        )
    return ret
}

// extern int ZBRPRNSetOverlayMode(HANDLE hPrinter, int printerType, int side, OverlayType overlay, char *filename, int *err);
func ZBRPRNSetOverlayMode() (ret uintptr) {
    panic("ZBRPRNSetOverlayMode not implemented")
    ret, _, _ = zBRPRNSetOverlayMode.Call(
        )
    return ret
}

// extern int ZBRPRNSetPrintHeadResistance(IntPtr hprinter, int printertype, int resistance, out int err);
func ZBRPRNSetPrintHeadResistance() (ret uintptr) {
    panic("ZBRPRNSetPrintHeadResistance not implemented")
    ret, _, _ = zBRPRNSetPrintHeadResistance.Call(
        )
    return ret
}

// extern int ZBRPRNSetRelativeXOffset(IntPtr hprinter, int printertype, int offset, out int err);
func ZBRPRNSetRelativeXOffset() (ret uintptr) {
    panic("ZBRPRNSetRelativeXOffset not implemented")
    ret, _, _ = zBRPRNSetRelativeXOffset.Call(
        )
    return ret
}

// extern int ZBRPRNSetRelativeYOffset(IntPtr hprinter, int printertype, int offset, out int err);
func ZBRPRNSetRelativeYOffset() (ret uintptr) {
    panic("ZBRPRNSetRelativeYOffset not implemented")
    ret, _, _ = zBRPRNSetRelativeYOffset.Call(
        )
    return ret
}

// extern int ZBRPRNSetStartPrintSideBOffset(IntPtr hprinter, int printertype, int x_y, int offset, out int err);
func ZBRPRNSetStartPrintSideBOffset() (ret uintptr) {
    panic("ZBRPRNSetStartPrintSideBOffset not implemented")
    ret, _, _ = zBRPRNSetStartPrintSideBOffset.Call(
        )
    return ret
}

// extern int ZBRPRNSetStartPrintSideBXOffset(IntPtr hprinter, int printertype, int offset, out int err);
func ZBRPRNSetStartPrintSideBXOffset() (ret uintptr) {
    panic("ZBRPRNSetStartPrintSideBXOffset not implemented")
    ret, _, _ = zBRPRNSetStartPrintSideBXOffset.Call(
        )
    return ret
}

// extern int ZBRPRNSetStartPrintSideBYOffset(IntPtr hprinter, int printertype, int offset, out int err);
func ZBRPRNSetStartPrintSideBYOffset() (ret uintptr) {
    panic("ZBRPRNSetStartPrintSideBYOffset not implemented")
    ret, _, _ = zBRPRNSetStartPrintSideBYOffset.Call(
        )
    return ret
}

// extern int ZBRPRNSetStartPrintXOffset(IntPtr hprinter, int printertype, int offset, out int err);
func ZBRPRNSetStartPrintXOffset() (ret uintptr) {
    panic("ZBRPRNSetStartPrintXOffset not implemented")
    ret, _, _ = zBRPRNSetStartPrintXOffset.Call(
        )
    return ret
}

// extern int ZBRPRNSetStartPrintYOffset(IntPtr hprinter, int printertype, int offset, out int err);
func ZBRPRNSetStartPrintYOffset() (ret uintptr) {
    panic("ZBRPRNSetStartPrintYOffset not implemented")
    ret, _, _ = zBRPRNSetStartPrintYOffset.Call(
        )
    return ret
}

// extern int ZBRPRNSetTrkDensity(IntPtr hprinter, int printertype, int trknum, int density, out int err);
func ZBRPRNSetTrkDensity() (ret uintptr) {
    panic("ZBRPRNSetTrkDensity not implemented")
    ret, _, _ = zBRPRNSetTrkDensity.Call(
        )
    return ret
}

// extern int ZBRPRNStartCleaningCardSeq(IntPtr hprinter, int printertype, out int err);
func ZBRPRNStartCleaningCardSeq() (ret uintptr) {
    panic("ZBRPRNStartCleaningCardSeq not implemented")
    ret, _, _ = zBRPRNStartCleaningCardSeq.Call(
        )
    return ret
}

// extern int ZBRPRNStartCleaningSeq(IntPtr hprinter, int printertype, out int err);
func ZBRPRNStartCleaningSeq() (ret uintptr) {
    panic("ZBRPRNStartCleaningSeq not implemented")
    ret, _, _ = zBRPRNStartCleaningSeq.Call(
        )
    return ret
}

// extern int ZBRPRNStartSmartCard(IntPtr _handle, int printertype, int cardtype, out int err);
func ZBRPRNStartSmartCard() (ret uintptr) {
    panic("ZBRPRNStartSmartCard not implemented")
    ret, _, _ = zBRPRNStartSmartCard.Call(
        )
    return ret
}

// extern int ZBRPRNSuppressStatusMsgs(IntPtr hprinter, int printertype, int suppressmsgs, out int err);
func ZBRPRNSuppressStatusMsgs() (ret uintptr) {
    panic("ZBRPRNSuppressStatusMsgs not implemented")
    ret, _, _ = zBRPRNSuppressStatusMsgs.Call(
        )
    return ret
}

// extern intZBRPRNUpgradeFirmware(HANDLE hPrinter, int printerType, char *filename, int *err);
func HANDLE() (ret uintptr) {
    panic("HANDLE not implemented")
    ret, _, _ = hANDLE.Call(
        )
    return ret
}

// extern int ZBRPRNWriteBarcode(IntPtr hprinter, int printertype, int startx, int starty, int rotation, int barcodetype, int barwidthratio, int barcodemultiplier, int barcodeheight, int textunder, byte[] barcodedata, out int err);
func ZBRPRNWriteBarcode() (ret uintptr) {
    panic("ZBRPRNWriteBarcode not implemented")
    ret, _, _ = zBRPRNWriteBarcode.Call(
        )
    return ret
}

// extern int ZBRPRNWriteBox(IntPtr hprinter, int printertype, int startx, int starty, int width, int height, int thickness, out int err);
func ZBRPRNWriteBox() (ret uintptr) {
    panic("ZBRPRNWriteBox not implemented")
    ret, _, _ = zBRPRNWriteBox.Call(
        )
    return ret
}

// extern int ZBRPRNWriteBoxEx(IntPtr hprinter, int printertype, int startx, int starty, int width, int height, int thickness, int gmode, int isvarnish, out int err);
func ZBRPRNWriteBoxEx() (ret uintptr) {
    panic("ZBRPRNWriteBoxEx not implemented")
    ret, _, _ = zBRPRNWriteBoxEx.Call(
        )
    return ret
}

// extern int ZBRPRNWriteMag(IntPtr handle, int printertype, int trkstowrite, byte[] trk1data, byte[] trk2data, byte[] trk3data, out int err);
func ZBRPRNWriteMag() (ret uintptr) {
    panic("ZBRPRNWriteMag not implemented")
    ret, _, _ = zBRPRNWriteMag.Call(
        )
    return ret
}

// extern int ZBRPRNWriteMagByTrk(IntPtr hprinter, int printertype, int trknum, byte[] trkdata, out int err);
func ZBRPRNWriteMagByTrk() (ret uintptr) {
    panic("ZBRPRNWriteMagByTrk not implemented")
    ret, _, _ = zBRPRNWriteMagByTrk.Call(
        )
    return ret
}

// extern int ZBRPRNWriteMagPassThru(IntPtr hdc, int printertype, int trknum, byte[] trkdata, out int err);
func ZBRPRNWriteMagPassThru() (ret uintptr) {
    panic("ZBRPRNWriteMagPassThru not implemented")
    ret, _, _ = zBRPRNWriteMagPassThru.Call(
        )
    return ret
}

// extern int ZBRPRNWriteText(IntPtr hprinter, int printertype, int startx, int starty, int rotation, int isbold, int height, byte[] text, out int err);
func ZBRPRNWriteText() (ret uintptr) {
    panic("ZBRPRNWriteText not implemented")
    ret, _, _ = zBRPRNWriteText.Call(
        )
    return ret
}

// extern int ZBRPRNWriteTextEx(IntPtr hprinter, int printertype, int startx, int starty, int rotation, int isbold, int widht, int height, int gmode, byte[] text, int isvarnish, out int err);
func ZBRPRNWriteTextEx() (ret uintptr) {
    panic("ZBRPRNWriteTextEx not implemented")
    ret, _, _ = zBRPRNWriteTextEx.Call(
        )
    return ret
}
