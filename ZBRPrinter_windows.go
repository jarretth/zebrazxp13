package zebrazxp13

import (
    "syscall"
    "unsafe"
)

var (
    zBRPrinter                        = syscall.NewLazyDLL("ZBRPrinter.dll")
    zBRCloseHandle                    = zBRPrinter.NewProc("ZBRCloseHandle")
    zBRGetHandle                      = zBRPrinter.NewProc("ZBRGetHandle")
    zBRPRNCheckForErrors              = zBRPrinter.NewProc("ZBRPRNCheckForErrors")
    zBRPRNChkDueForCleaning           = zBRPrinter.NewProc("ZBRPRNChkDueForCleaning")
    zBRPRNClrColorImgBuf              = zBRPrinter.NewProc("ZBRPRNClrColorImgBuf")
    zBRPRNClrColorImgBufs             = zBRPrinter.NewProc("ZBRPRNClrColorImgBufs")
    zBRPRNClrErrStatusLn              = zBRPrinter.NewProc("ZBRPRNClrErrStatusLn")
    zBRPRNClrMediaPath                = zBRPrinter.NewProc("ZBRPRNClrMediaPath")
    zBRPRNClrMonoImgBuf               = zBRPrinter.NewProc("ZBRPRNClrMonoImgBuf")
    zBRPRNClrMonoImgBufs              = zBRPrinter.NewProc("ZBRPRNClrMonoImgBufs")
    zBRPRNClrSpecifiedBmp             = zBRPrinter.NewProc("ZBRPRNClrSpecifiedBmp")
    zBRPRNEjectCard                   = zBRPrinter.NewProc("ZBRPRNEjectCard")
    zBRPRNEnableMagReadDataEncryption = zBRPrinter.NewProc("ZBRPRNEnableMagReadDataEncryption")
    zBRPRNEndSmartCard                = zBRPrinter.NewProc("ZBRPRNEndSmartCard")
    zBRPRNFlipCard                    = zBRPrinter.NewProc("ZBRPRNFlipCard")
    zBRPRNGetChecksum                 = zBRPrinter.NewProc("ZBRPRNGetChecksum")
    zBRPRNGetCleaningParam            = zBRPrinter.NewProc("ZBRPRNGetCleaningParam")
    zBRPRNGetIPAddress                = zBRPrinter.NewProc("ZBRPRNGetIPAddress")
    zBRPRNGetMsgSuppressionState      = zBRPrinter.NewProc("ZBRPRNGetMsgSuppressionState")
    zBRPRNGetOpParam                  = zBRPrinter.NewProc("ZBRPRNGetOpParam")
    zBRPRNGetPanelsPrinted            = zBRPrinter.NewProc("ZBRPRNGetPanelsPrinted")
    zBRPRNGetPanelsRemaining          = zBRPrinter.NewProc("ZBRPRNGetPanelsRemaining")
    zBRPRNGetPrintCount               = zBRPrinter.NewProc("ZBRPRNGetPrintCount")
    zBRPRNGetPrinterOptions           = zBRPrinter.NewProc("ZBRPRNGetPrinterOptions")
    zBRPRNGetPrinterSerialNumb        = zBRPrinter.NewProc("ZBRPRNGetPrinterSerialNumb")
    zBRPRNGetPrinterSerialNumber      = zBRPrinter.NewProc("ZBRPRNGetPrinterSerialNumber")
    zBRPRNGetPrinterStatus            = zBRPrinter.NewProc("ZBRPRNGetPrinterStatus")
    zBRPRNGetPrintHeadSerialNumb      = zBRPrinter.NewProc("ZBRPRNGetPrintHeadSerialNumb")
    zBRPRNGetPrintHeadSerialNumber    = zBRPrinter.NewProc("ZBRPRNGetPrintHeadSerialNumber")
    zBRPRNGetPrintStatus              = zBRPrinter.NewProc("ZBRPRNGetPrintStatus")
    zBRPRNGetSDKProductVer            = zBRPrinter.NewProc("ZBRPRNGetSDKProductVer")
    zBRPRNGetSDKVer                   = zBRPrinter.NewProc("ZBRPRNGetSDKVer")
    zBRPRNGetSDKVsn                   = zBRPrinter.NewProc("ZBRPRNGetSDKVsn")
    zBRPRNGetSensorStatus             = zBRPrinter.NewProc("ZBRPRNGetSensorStatus")
    zBRPRNImmediateParamSave          = zBRPrinter.NewProc("ZBRPRNImmediateParamSave")
    zBRPRNIsPrinterReady              = zBRPrinter.NewProc("ZBRPRNIsPrinterReady")
    zBRPRNMoveCard                    = zBRPrinter.NewProc("ZBRPRNMoveCard")
    zBRPRNMoveCardBkwd                = zBRPrinter.NewProc("ZBRPRNMoveCardBkwd")
    zBRPRNMoveCardFwd                 = zBRPrinter.NewProc("ZBRPRNMoveCardFwd")
    zBRPRNMovePrintReady              = zBRPrinter.NewProc("ZBRPRNMovePrintReady")
    zBRPRNPrintCardPanel              = zBRPrinter.NewProc("ZBRPRNPrintCardPanel")
    zBRPRNPrintClearVarnish           = zBRPrinter.NewProc("ZBRPRNPrintClearVarnish")
    zBRPRNPrintColorImgBuf            = zBRPrinter.NewProc("ZBRPRNPrintColorImgBuf")
    zBRPRNPrintHologramOverlay        = zBRPrinter.NewProc("ZBRPRNPrintHologramOverlay")
    zBRPRNPrintMonoImgBuf             = zBRPrinter.NewProc("ZBRPRNPrintMonoImgBuf")
    zBRPRNPrintMonoImgBufEx           = zBRPrinter.NewProc("ZBRPRNPrintMonoImgBufEx")
    zBRPRNPrintMonoPanel              = zBRPrinter.NewProc("ZBRPRNPrintMonoPanel")
    zBRPRNPrintPrnFile                = zBRPrinter.NewProc("ZBRPRNPrintPrnFile")
    zBRPRNPrintTestCard               = zBRPrinter.NewProc("ZBRPRNPrintTestCard")
    zBRPRNPrintVarnish                = zBRPrinter.NewProc("ZBRPRNPrintVarnish")
    zBRPRNPrintVarnishEx              = zBRPrinter.NewProc("ZBRPRNPrintVarnishEx")
    zBRPRNReadMag                     = zBRPrinter.NewProc("ZBRPRNReadMag")
    zBRPRNReadMagByTrk                = zBRPrinter.NewProc("ZBRPRNReadMagByTrk")
    zBRPRNResetMagEncoder             = zBRPrinter.NewProc("ZBRPRNResetMagEncoder")
    zBRPRNResetPrinter                = zBRPrinter.NewProc("ZBRPRNResetPrinter")
    zBRPRNResync                      = zBRPrinter.NewProc("ZBRPRNResync")
    zBRPRNReversePrintReady           = zBRPrinter.NewProc("ZBRPRNReversePrintReady")
    zBRPRNSelfAdj                     = zBRPrinter.NewProc("ZBRPRNSelfAdj")
    zBRPRNSetCardFeedingMode          = zBRPrinter.NewProc("ZBRPRNSetCardFeedingMode")
    zBRPRNSetCleaningParam            = zBRPrinter.NewProc("ZBRPRNSetCleaningParam")
    zBRPRNSetColorContrast            = zBRPrinter.NewProc("ZBRPRNSetColorContrast")
    zBRPRNSetContrastIntensityLvl     = zBRPrinter.NewProc("ZBRPRNSetContrastIntensityLvl")
    zBRPRNSetEncoderCoercivity        = zBRPrinter.NewProc("ZBRPRNSetEncoderCoercivity")
    zBRPRNSetEncodingDir              = zBRPrinter.NewProc("ZBRPRNSetEncodingDir")
    zBRPRNSetEndOfPrint               = zBRPrinter.NewProc("ZBRPRNSetEndOfPrint")
    zBRPRNSetHologramIntensity        = zBRPrinter.NewProc("ZBRPRNSetHologramIntensity")
    zBRPRNSetMagEncodingStd           = zBRPrinter.NewProc("ZBRPRNSetMagEncodingStd")
    zBRPRNSetMonoContrast             = zBRPrinter.NewProc("ZBRPRNSetMonoContrast")
    zBRPRNSetMonoIntensity            = zBRPrinter.NewProc("ZBRPRNSetMonoIntensity")
    zBRPRNSetOverlayMode              = zBRPrinter.NewProc("ZBRPRNSetOverlayMode")
    zBRPRNSetPrintHeadResistance      = zBRPrinter.NewProc("ZBRPRNSetPrintHeadResistance")
    zBRPRNSetRelativeXOffset          = zBRPrinter.NewProc("ZBRPRNSetRelativeXOffset")
    zBRPRNSetRelativeYOffset          = zBRPrinter.NewProc("ZBRPRNSetRelativeYOffset")
    zBRPRNSetStartPrintSideBOffset    = zBRPrinter.NewProc("ZBRPRNSetStartPrintSideBOffset")
    zBRPRNSetStartPrintSideBXOffset   = zBRPrinter.NewProc("ZBRPRNSetStartPrintSideBXOffset")
    zBRPRNSetStartPrintSideBYOffset   = zBRPrinter.NewProc("ZBRPRNSetStartPrintSideBYOffset")
    zBRPRNSetStartPrintXOffset        = zBRPrinter.NewProc("ZBRPRNSetStartPrintXOffset")
    zBRPRNSetStartPrintYOffset        = zBRPrinter.NewProc("ZBRPRNSetStartPrintYOffset")
    zBRPRNSetTrkDensity               = zBRPrinter.NewProc("ZBRPRNSetTrkDensity")
    zBRPRNStartCleaningCardSeq        = zBRPrinter.NewProc("ZBRPRNStartCleaningCardSeq")
    zBRPRNStartCleaningSeq            = zBRPrinter.NewProc("ZBRPRNStartCleaningSeq")
    zBRPRNStartSmartCard              = zBRPrinter.NewProc("ZBRPRNStartSmartCard")
    zBRPRNSuppressStatusMsgs          = zBRPrinter.NewProc("ZBRPRNSuppressStatusMsgs")
    zBRPRNUpgradeFirmware             = zBRPrinter.NewProc("ZBRPRNUpgradeFirmware")
    zBRPRNWriteBarCode                = zBRPrinter.NewProc("ZBRPRNWriteBarCode")
    zBRPRNWriteBox                    = zBRPrinter.NewProc("ZBRPRNWriteBox")
    zBRPRNWriteBoxEx                  = zBRPrinter.NewProc("ZBRPRNWriteBoxEx")
    zBRPRNWriteMag                    = zBRPrinter.NewProc("ZBRPRNWriteMag")
    zBRPRNWriteMagByTrk               = zBRPrinter.NewProc("ZBRPRNWriteMagByTrk")
    zBRPRNWriteMagPassThru            = zBRPrinter.NewProc("ZBRPRNWriteMagPassThru")
    zBRPRNWriteText                   = zBRPrinter.NewProc("ZBRPRNWriteText")
    zBRPRNWriteTextEx                 = zBRPrinter.NewProc("ZBRPRNWriteTextEx")
)

// extern int ZBRCloseHandle(IntPtr handle, out int err);
func ZBRCloseHandle(handle Handle) (success SuccessReturn, err uint) {
    ret, _, _ := zBRCloseHandle.Call(
        uintptr(handle),
        uintptr(unsafe.Pointer(&err)),
    )
    return SuccessReturn(ret), err
}

// extern int ZBRGetHandle(out IntPtr handle, byte[] drvname, out int printertype, out int err);
func ZBRGetHandle(drvName string) (success SuccessReturn, handle Handle, prn_type uint, err uint) {
    ret, _, _ := zBRGetHandle.Call(
            uintptr(unsafe.Pointer(&handle)),
            uintptr(unsafe.Pointer(&([]byte(drvName))[0])),
            uintptr(unsafe.Pointer(&prn_type)),
            uintptr(unsafe.Pointer(&err)),
    )
    return SuccessReturn(ret), handle, prn_type, err
}

// extern int ZBRPRNCheckForErrors(IntPtr hprinter, int printertype);
func ZBRPRNCheckForErrors() (success SuccessReturn) {
    panic("ZBRPRNCheckForErrors not implemented")
    ret, _, _ := zBRPRNCheckForErrors.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNChkDueForCleaning(IntPtr hprinter, int printertype, out int imgcounter, out int cleancounter, out int cleancardcounter, out int err);
func ZBRPRNChkDueForCleaning() (success SuccessReturn) {
    panic("ZBRPRNChkDueForCleaning not implemented")
    ret, _, _ := zBRPRNChkDueForCleaning.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNClrColorImgBuf(IntPtr hprinter, int printertype, int colorbufidx, out int err);
func ZBRPRNClrColorImgBuf() (success SuccessReturn) {
    panic("ZBRPRNClrColorImgBuf not implemented")
    ret, _, _ := zBRPRNClrColorImgBuf.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNClrColorImgBufs(HANDLE hPrinter, int printerType, int *err)
func ZBRPRNClrColorImgBufs() (success SuccessReturn) {
    panic("ZBRPRNClrColorImgBufs not implemented")
    ret, _, _ := zBRPRNClrColorImgBufs.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNClrErrStatusLn(IntPtr hprinter, int printertype, out int err);
func ZBRPRNClrErrStatusLn() (success SuccessReturn) {
    panic("ZBRPRNClrErrStatusLn not implemented")
    ret, _, _ := zBRPRNClrErrStatusLn.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNClrMediaPath(IntPtr hprinter, int printertype, out int err);
func ZBRPRNClrMediaPath() (success SuccessReturn) {
    panic("ZBRPRNClrMediaPath not implemented")
    ret, _, _ := zBRPRNClrMediaPath.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNClrMonoImgBuf(IntPtr hprinter, int printertype, int clrvarnish, out int err);
func ZBRPRNClrMonoImgBuf() (success SuccessReturn) {
    panic("ZBRPRNClrMonoImgBuf not implemented")
    ret, _, _ := zBRPRNClrMonoImgBuf.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNClrMonoImgBufs(IntPtr hprinter, int printertype, int clrbuffer, out int err);
func ZBRPRNClrMonoImgBufs() (success SuccessReturn) {
    panic("ZBRPRNClrMonoImgBufs not implemented")
    ret, _, _ := zBRPRNClrMonoImgBufs.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNClrSpecifiedBmp(IntPtr hprinter, int printertype, int colorbufidx, out int err);
func ZBRPRNClrSpecifiedBmp() (success SuccessReturn) {
    panic("ZBRPRNClrSpecifiedBmp not implemented")
    ret, _, _ := zBRPRNClrSpecifiedBmp.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNEjectCard(IntPtr _handle, int prn_type, out int err);
func ZBRPRNEjectCard(handle Handle, prn_type uint) (success SuccessReturn, err uint) {
    ret, _, _ := zBRPRNEjectCard.Call(
            uintptr(handle),
            uintptr(prn_type),
            uintptr(unsafe.Pointer(&err)),
    )
    return SuccessReturn(ret), err
}

// extern int ZBRPRNEnableMagReadDataEncryption(IntPtr handle, int printertype, out int err);
func ZBRPRNEnableMagReadDataEncryption() (success SuccessReturn) {
    panic("ZBRPRNEnableMagReadDataEncryption not implemented")
    ret, _, _ := zBRPRNEnableMagReadDataEncryption.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNEndSmartCard(IntPtr _handle, int printertype, int cardtype, int movement, out int err);
func ZBRPRNEndSmartCard() (success SuccessReturn) {
    panic("ZBRPRNEndSmartCard not implemented")
    ret, _, _ := zBRPRNEndSmartCard.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNFlipCard(IntPtr hprinter, int printertype, out int err);
func ZBRPRNFlipCard() (success SuccessReturn) {
    panic("ZBRPRNFlipCard not implemented")
    ret, _, _ := zBRPRNFlipCard.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNGetChecksum(IntPtr hprinter, int printertype, out int checksum, out int err);
func ZBRPRNGetChecksum() (success SuccessReturn) {
    panic("ZBRPRNGetChecksum not implemented")
    ret, _, _ := zBRPRNGetChecksum.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNGetCleaningParam(IntPtr hprinter, int printertype, out int imgcounter, out int cleancounter, out int cleancardcounter, out int err);
func ZBRPRNGetCleaningParam() (success SuccessReturn) {
    panic("ZBRPRNGetCleaningParam not implemented")
    ret, _, _ := zBRPRNGetCleaningParam.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNGetIPAddress(string PrinterName, char* ipAddress);
func ZBRPRNGetIPAddress() (success SuccessReturn) {
    panic("ZBRPRNGetIPAddress not implemented")
    ret, _, _ := zBRPRNGetIPAddress.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNGetMsgSuppressionState(IntPtr hprinter, int printertype, out int state, out int err);
func ZBRPRNGetMsgSuppressionState() (success SuccessReturn) {
    panic("ZBRPRNGetMsgSuppressionState not implemented")
    ret, _, _ := zBRPRNGetMsgSuppressionState.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNGetOpParam(IntPtr hprinter, int printertype, int paramidx, byte[] opparam, out int respsize, out int err);
func ZBRPRNGetOpParam() (success SuccessReturn) {
    panic("ZBRPRNGetOpParam not implemented")
    ret, _, _ := zBRPRNGetOpParam.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNGetPanelsPrinted(IntPtr hprinter, int printertype, out int panels, out int err);
func ZBRPRNGetPanelsPrinted() (success SuccessReturn) {
    panic("ZBRPRNGetPanelsPrinted not implemented")
    ret, _, _ := zBRPRNGetPanelsPrinted.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNGetPanelsRemaining(IntPtr hprinter, int printertype, out int panels, out int err);
func ZBRPRNGetPanelsRemaining() (success SuccessReturn) {
    panic("ZBRPRNGetPanelsRemaining not implemented")
    ret, _, _ := zBRPRNGetPanelsRemaining.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNGetPrintCount(IntPtr hprinter, int printertype, out int printcount, out int err);
func ZBRPRNGetPrintCount() (success SuccessReturn) {
    panic("ZBRPRNGetPrintCount not implemented")
    ret, _, _ := zBRPRNGetPrintCount.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNGetPrinterOptions(IntPtr hprinter, int printertype, byte[] options, out int respsize, out int err);
func ZBRPRNGetPrinterOptions() (success SuccessReturn) {
    panic("ZBRPRNGetPrinterOptions not implemented")
    ret, _, _ := zBRPRNGetPrinterOptions.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNGetPrinterSerialNumb(IntPtr hprinter, int printertype, byte[] serialnum, out int respsize, out int err);
func ZBRPRNGetPrinterSerialNumb() (success SuccessReturn) {
    panic("ZBRPRNGetPrinterSerialNumb not implemented")
    ret, _, _ := zBRPRNGetPrinterSerialNumb.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNGetPrinterSerialNumber(IntPtr hprinter, int printertype, byte[] serialnum, out int respsize, out int err);
func ZBRPRNGetPrinterSerialNumber() (success SuccessReturn) {
    panic("ZBRPRNGetPrinterSerialNumber not implemented")
    ret, _, _ := zBRPRNGetPrinterSerialNumber.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNGetPrinterStatus(out int statuscode);
func ZBRPRNGetPrinterStatus() (success SuccessReturn) {
    panic("ZBRPRNGetPrinterStatus not implemented")
    ret, _, _ := zBRPRNGetPrinterStatus.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNGetPrintHeadSerialNumb(IntPtr hprinter, int printertype, byte[] serialnum, out int respsize, out int err);
func ZBRPRNGetPrintHeadSerialNumb() (success SuccessReturn) {
    panic("ZBRPRNGetPrintHeadSerialNumb not implemented")
    ret, _, _ := zBRPRNGetPrintHeadSerialNumb.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNGetPrintHeadSerialNumber(IntPtr hprinter, int printertype, byte[] serialnum, out int respsize, out int err);
func ZBRPRNGetPrintHeadSerialNumber() (success SuccessReturn) {
    panic("ZBRPRNGetPrintHeadSerialNumber not implemented")
    ret, _, _ := zBRPRNGetPrintHeadSerialNumber.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNGetPrintStatus(IntPtr hprinter, int printertype);
func ZBRPRNGetPrintStatus() (success SuccessReturn) {
    panic("ZBRPRNGetPrintStatus not implemented")
    ret, _, _ := zBRPRNGetPrintStatus.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNGetSDKProductVer(byte[] productversion, out int buffsize);
func ZBRPRNGetSDKProductVer() (success SuccessReturn) {
    panic("ZBRPRNGetSDKProductVer not implemented")
    ret, _, _ := zBRPRNGetSDKProductVer.Call(
        )
    return SuccessReturn(ret)
}

// extern void ZBRPRNGetSDKVer(out int major, out int minor, out int englevel);
func ZBRPRNGetSDKVer() (major, minor, engLevel int) {
    _, _, _ = zBRPRNGetSDKVer.Call(
        uintptr(unsafe.Pointer(&major)),
        uintptr(unsafe.Pointer(&minor)),
        uintptr(unsafe.Pointer(&engLevel)),
    )
    return
}

// extern void ZBRPRNGetSDKVsn(out int major, out int minor, out int englevel);
func ZBRPRNGetSDKVsn() (success SuccessReturn) {
    panic("ZBRPRNGetSDKVsn not implemented")
    ret, _, _ := zBRPRNGetSDKVsn.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNGetSensorStatus(IntPtr hprinter, int printertype, out byte status, out int err);
func ZBRPRNGetSensorStatus() (success SuccessReturn) {
    panic("ZBRPRNGetSensorStatus not implemented")
    ret, _, _ := zBRPRNGetSensorStatus.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNImmediateParamSave(IntPtr hprinter, int printertype, out int err);
func ZBRPRNImmediateParamSave() (success SuccessReturn) {
    panic("ZBRPRNImmediateParamSave not implemented")
    ret, _, _ := zBRPRNImmediateParamSave.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNIsPrinterReady(IntPtr hprinter, int printertype, out int err);
func ZBRPRNIsPrinterReady() (success SuccessReturn) {
    panic("ZBRPRNIsPrinterReady not implemented")
    ret, _, _ := zBRPRNIsPrinterReady.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNMoveCard(IntPtr hprinter, int printertype, int steps, out int err);
func ZBRPRNMoveCard() (success SuccessReturn) {
    panic("ZBRPRNMoveCard not implemented")
    ret, _, _ := zBRPRNMoveCard.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNMoveCardBkwd(IntPtr hprinter, int printertype, int steps, out int err);
func ZBRPRNMoveCardBkwd() (success SuccessReturn) {
    panic("ZBRPRNMoveCardBkwd not implemented")
    ret, _, _ := zBRPRNMoveCardBkwd.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNMoveCardFwd(IntPtr hprinter, int printertype, int steps, out int err);
func ZBRPRNMoveCardFwd() (success SuccessReturn) {
    panic("ZBRPRNMoveCardFwd not implemented")
    ret, _, _ := zBRPRNMoveCardFwd.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNMovePrintReady(IntPtr handle, int printertype, out int err);
func ZBRPRNMovePrintReady() (success SuccessReturn) {
    panic("ZBRPRNMovePrintReady not implemented")
    ret, _, _ := zBRPRNMovePrintReady.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNPrintCardPanel(IntPtr hprinter, int printertype, int imgbufidx, out int err);
func ZBRPRNPrintCardPanel() (success SuccessReturn) {
    panic("ZBRPRNPrintCardPanel not implemented")
    ret, _, _ := zBRPRNPrintCardPanel.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNPrintClearVarnish(IntPtr hprinter, int printertype, out int err);
func ZBRPRNPrintClearVarnish() (success SuccessReturn) {
    panic("ZBRPRNPrintClearVarnish not implemented")
    ret, _, _ := zBRPRNPrintClearVarnish.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNPrintColorImgBuf(IntPtr hprinter, int printertype, int imgbufidx, out int err);
func ZBRPRNPrintColorImgBuf() (success SuccessReturn) {
    panic("ZBRPRNPrintColorImgBuf not implemented")
    ret, _, _ := zBRPRNPrintColorImgBuf.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNPrintHologramOverlay(IntPtr hprinter, int printertype, int printcardcommand, out int err);
func ZBRPRNPrintHologramOverlay() (success SuccessReturn) {
    panic("ZBRPRNPrintHologramOverlay not implemented")
    ret, _, _ := zBRPRNPrintHologramOverlay.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNPrintMonoImgBuf(IntPtr hprinter, int printertype, out int err);
func ZBRPRNPrintMonoImgBuf() (success SuccessReturn) {
    panic("ZBRPRNPrintMonoImgBuf not implemented")
    ret, _, _ := zBRPRNPrintMonoImgBuf.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNPrintMonoImgBufEx(IntPtr hprinter, int printertype, int cardcommand, out int err);
func ZBRPRNPrintMonoImgBufEx() (success SuccessReturn) {
    panic("ZBRPRNPrintMonoImgBufEx not implemented")
    ret, _, _ := zBRPRNPrintMonoImgBufEx.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNPrintMonoPanel(IntPtr hprinter, int printertype, out int err);
func ZBRPRNPrintMonoPanel() (success SuccessReturn) {
    panic("ZBRPRNPrintMonoPanel not implemented")
    ret, _, _ := zBRPRNPrintMonoPanel.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNPrintPRNFile(IntPtr hprinter, int printertype, byte[] filename, out int err);
func ZBRPRNPrintPrnFile() (success SuccessReturn) {
    panic("ZBRPRNPrintPrnFile not implemented")
    ret, _, _ := zBRPRNPrintPrnFile.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNPrintTestCard(IntPtr hprinter, int printertype, int testcardtype, out int err);
func ZBRPRNPrintTestCard() (success SuccessReturn) {
    panic("ZBRPRNPrintTestCard not implemented")
    ret, _, _ := zBRPRNPrintTestCard.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNPrintVarnish(IntPtr hprinter, int printertype, out int err);
func ZBRPRNPrintVarnish() (success SuccessReturn) {
    panic("ZBRPRNPrintVarnish not implemented")
    ret, _, _ := zBRPRNPrintVarnish.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNPrintVarnishEx(IntPtr hprinter, int printertype, int printcardcommand, out int err);
func ZBRPRNPrintVarnishEx() (success SuccessReturn) {
    panic("ZBRPRNPrintVarnishEx not implemented")
    ret, _, _ := zBRPRNPrintVarnishEx.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNReadMag(IntPtr handle, int printertype, int trkstoread, byte[] trk1buf, out int trk1bytesneeded, byte[] trk2buf, out int trk2bytesneeded, byte[] trk3buf, out int trk3bytesneeded, out int err);
func ZBRPRNReadMag() (success SuccessReturn) {
    panic("ZBRPRNReadMag not implemented")
    ret, _, _ := zBRPRNReadMag.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNReadMagByTrk(IntPtr hprinter, int printertype, int trknum, byte[] trkbuf, out int respsize, out int err);
func ZBRPRNReadMagByTrk() (success SuccessReturn) {
    panic("ZBRPRNReadMagByTrk not implemented")
    ret, _, _ := zBRPRNReadMagByTrk.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNResetMagEncoder(IntPtr hprinter, int printertype, out int err);
func ZBRPRNResetMagEncoder() (success SuccessReturn) {
    panic("ZBRPRNResetMagEncoder not implemented")
    ret, _, _ := zBRPRNResetMagEncoder.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNResetPrinter(IntPtr hprinter, int printertype, out int err);
func ZBRPRNResetPrinter() (success SuccessReturn) {
    panic("ZBRPRNResetPrinter not implemented")
    ret, _, _ := zBRPRNResetPrinter.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNResync(IntPtr hprinter, int printertype, out int err);
func ZBRPRNResync() (success SuccessReturn) {
    panic("ZBRPRNResync not implemented")
    ret, _, _ := zBRPRNResync.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNReversePrintReady(IntPtr hprinter, int printertype, out int err);
func ZBRPRNReversePrintReady() (success SuccessReturn) {
    panic("ZBRPRNReversePrintReady not implemented")
    ret, _, _ := zBRPRNReversePrintReady.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSelfAdj(IntPtr hprinter, int printertype, out int err);
func ZBRPRNSelfAdj() (success SuccessReturn) {
    panic("ZBRPRNSelfAdj not implemented")
    ret, _, _ := zBRPRNSelfAdj.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetCardFeedingMode(IntPtr hprinter, int printertype, int mode, out int err);
func ZBRPRNSetCardFeedingMode() (success SuccessReturn) {
    panic("ZBRPRNSetCardFeedingMode not implemented")
    ret, _, _ := zBRPRNSetCardFeedingMode.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetCleaningParam(IntPtr hprinter, int printertype, int ribbonpanelcounter, int cleancardpass, out int err);
func ZBRPRNSetCleaningParam() (success SuccessReturn) {
    panic("ZBRPRNSetCleaningParam not implemented")
    ret, _, _ := zBRPRNSetCleaningParam.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetColorContrast(IntPtr hprinter, int printertype, int imgbufidx, int contrast, out int err);
func ZBRPRNSetColorContrast() (success SuccessReturn) {
    panic("ZBRPRNSetColorContrast not implemented")
    ret, _, _ := zBRPRNSetColorContrast.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetContrastIntensityLvl(IntPtr hprinter, int printertype, int imgbufidx, int intensity, out int err);
func ZBRPRNSetContrastIntensityLvl() (success SuccessReturn) {
    panic("ZBRPRNSetContrastIntensityLvl not implemented")
    ret, _, _ := zBRPRNSetContrastIntensityLvl.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetEncoderCoercivity(IntPtr hprinter, int printertype, int coercivity, out int err);
func ZBRPRNSetEncoderCoercivity() (success SuccessReturn) {
    panic("ZBRPRNSetEncoderCoercivity not implemented")
    ret, _, _ := zBRPRNSetEncoderCoercivity.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetEncodingDir(IntPtr hprinter, int printertype, int dir, out int err);
func ZBRPRNSetEncodingDir() (success SuccessReturn) {
    panic("ZBRPRNSetEncodingDir not implemented")
    ret, _, _ := zBRPRNSetEncodingDir.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetEndOfPrint(IntPtr hprinter, int printertype, int xwidth, out int err);
func ZBRPRNSetEndOfPrint() (success SuccessReturn) {
    panic("ZBRPRNSetEndOfPrint not implemented")
    ret, _, _ := zBRPRNSetEndOfPrint.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetHologramIntensity(IntPtr hprinter, int printertype, int intensity, out int err);
func ZBRPRNSetHologramIntensity() (success SuccessReturn) {
    panic("ZBRPRNSetHologramIntensity not implemented")
    ret, _, _ := zBRPRNSetHologramIntensity.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetMagEncodingStd(IntPtr hprinter, int printertype, int standard, out int err);
func ZBRPRNSetMagEncodingStd() (success SuccessReturn) {
    panic("ZBRPRNSetMagEncodingStd not implemented")
    ret, _, _ := zBRPRNSetMagEncodingStd.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetMonoContrast(IntPtr hprinter, int printertype, int contrast, out int err);
func ZBRPRNSetMonoContrast() (success SuccessReturn) {
    panic("ZBRPRNSetMonoContrast not implemented")
    ret, _, _ := zBRPRNSetMonoContrast.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetMonoIntensity(IntPtr hprinter, int printertype, int intensity, out int err);
func ZBRPRNSetMonoIntensity() (success SuccessReturn) {
    panic("ZBRPRNSetMonoIntensity not implemented")
    ret, _, _ := zBRPRNSetMonoIntensity.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetOverlayMode(HANDLE hPrinter, int printerType, int side, OverlayType overlay, char *filename, int *err);
func ZBRPRNSetOverlayMode() (success SuccessReturn) {
    panic("ZBRPRNSetOverlayMode not implemented")
    ret, _, _ := zBRPRNSetOverlayMode.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetPrintHeadResistance(IntPtr hprinter, int printertype, int resistance, out int err);
func ZBRPRNSetPrintHeadResistance() (success SuccessReturn) {
    panic("ZBRPRNSetPrintHeadResistance not implemented")
    ret, _, _ := zBRPRNSetPrintHeadResistance.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetRelativeXOffset(IntPtr hprinter, int printertype, int offset, out int err);
func ZBRPRNSetRelativeXOffset() (success SuccessReturn) {
    panic("ZBRPRNSetRelativeXOffset not implemented")
    ret, _, _ := zBRPRNSetRelativeXOffset.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetRelativeYOffset(IntPtr hprinter, int printertype, int offset, out int err);
func ZBRPRNSetRelativeYOffset() (success SuccessReturn) {
    panic("ZBRPRNSetRelativeYOffset not implemented")
    ret, _, _ := zBRPRNSetRelativeYOffset.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetStartPrintSideBOffset(IntPtr hprinter, int printertype, int x_y, int offset, out int err);
func ZBRPRNSetStartPrintSideBOffset() (success SuccessReturn) {
    panic("ZBRPRNSetStartPrintSideBOffset not implemented")
    ret, _, _ := zBRPRNSetStartPrintSideBOffset.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetStartPrintSideBXOffset(IntPtr hprinter, int printertype, int offset, out int err);
func ZBRPRNSetStartPrintSideBXOffset() (success SuccessReturn) {
    panic("ZBRPRNSetStartPrintSideBXOffset not implemented")
    ret, _, _ := zBRPRNSetStartPrintSideBXOffset.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetStartPrintSideBYOffset(IntPtr hprinter, int printertype, int offset, out int err);
func ZBRPRNSetStartPrintSideBYOffset() (success SuccessReturn) {
    panic("ZBRPRNSetStartPrintSideBYOffset not implemented")
    ret, _, _ := zBRPRNSetStartPrintSideBYOffset.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetStartPrintXOffset(IntPtr hprinter, int printertype, int offset, out int err);
func ZBRPRNSetStartPrintXOffset() (success SuccessReturn) {
    panic("ZBRPRNSetStartPrintXOffset not implemented")
    ret, _, _ := zBRPRNSetStartPrintXOffset.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetStartPrintYOffset(IntPtr hprinter, int printertype, int offset, out int err);
func ZBRPRNSetStartPrintYOffset() (success SuccessReturn) {
    panic("ZBRPRNSetStartPrintYOffset not implemented")
    ret, _, _ := zBRPRNSetStartPrintYOffset.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSetTrkDensity(IntPtr hprinter, int printertype, int trknum, int density, out int err);
func ZBRPRNSetTrkDensity() (success SuccessReturn) {
    panic("ZBRPRNSetTrkDensity not implemented")
    ret, _, _ := zBRPRNSetTrkDensity.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNStartCleaningCardSeq(IntPtr hprinter, int printertype, out int err);
func ZBRPRNStartCleaningCardSeq() (success SuccessReturn) {
    panic("ZBRPRNStartCleaningCardSeq not implemented")
    ret, _, _ := zBRPRNStartCleaningCardSeq.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNStartCleaningSeq(IntPtr hprinter, int printertype, out int err);
func ZBRPRNStartCleaningSeq() (success SuccessReturn) {
    panic("ZBRPRNStartCleaningSeq not implemented")
    ret, _, _ := zBRPRNStartCleaningSeq.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNStartSmartCard(IntPtr _handle, int printertype, int cardtype, out int err);
func ZBRPRNStartSmartCard() (success SuccessReturn) {
    panic("ZBRPRNStartSmartCard not implemented")
    ret, _, _ := zBRPRNStartSmartCard.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNSuppressStatusMsgs(IntPtr hprinter, int printertype, int suppressmsgs, out int err);
func ZBRPRNSuppressStatusMsgs() (success SuccessReturn) {
    panic("ZBRPRNSuppressStatusMsgs not implemented")
    ret, _, _ := zBRPRNSuppressStatusMsgs.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNUpgradeFirmware(HANDLE hPrinter, int printerType, char *filename, int *err);
func ZBRPRNUpgradeFirmware() (success SuccessReturn) {
    panic("ZBRPRNUpgradeFirmware not implemented")
    ret, _, _ := zBRPRNUpgradeFirmware.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNWriteBarCode(IntPtr hprinter, int printertype, int startx, int starty, int rotation, int barcodetype, int barwidthratio, int barcodemultiplier, int barcodeheight, int textunder, byte[] barcodedata, out int err);
func ZBRPRNWriteBarCode() (success SuccessReturn) {
    panic("ZBRPRNWriteBarCode not implemented")
    ret, _, _ := zBRPRNWriteBarCode.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNWriteBox(IntPtr hprinter, int printertype, int startx, int starty, int width, int height, int thickness, out int err);
func ZBRPRNWriteBox() (success SuccessReturn) {
    panic("ZBRPRNWriteBox not implemented")
    ret, _, _ := zBRPRNWriteBox.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNWriteBoxEx(IntPtr hprinter, int printertype, int startx, int starty, int width, int height, int thickness, int gmode, int isvarnish, out int err);
func ZBRPRNWriteBoxEx() (success SuccessReturn) {
    panic("ZBRPRNWriteBoxEx not implemented")
    ret, _, _ := zBRPRNWriteBoxEx.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNWriteMag(IntPtr handle, int printertype, int trkstowrite, byte[] trk1data, byte[] trk2data, byte[] trk3data, out int err);
func ZBRPRNWriteMag() (success SuccessReturn) {
    panic("ZBRPRNWriteMag not implemented")
    ret, _, _ := zBRPRNWriteMag.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNWriteMagByTrk(IntPtr hprinter, int printertype, int trknum, byte[] trkdata, out int err);
func ZBRPRNWriteMagByTrk() (success SuccessReturn) {
    panic("ZBRPRNWriteMagByTrk not implemented")
    ret, _, _ := zBRPRNWriteMagByTrk.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNWriteMagPassThru(IntPtr hdc, int printertype, int trknum, byte[] trkdata, out int err);
func ZBRPRNWriteMagPassThru() (success SuccessReturn) {
    panic("ZBRPRNWriteMagPassThru not implemented")
    ret, _, _ := zBRPRNWriteMagPassThru.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNWriteText(IntPtr hprinter, int printertype, int startx, int starty, int rotation, int isbold, int height, byte[] text, out int err);
func ZBRPRNWriteText() (success SuccessReturn) {
    panic("ZBRPRNWriteText not implemented")
    ret, _, _ := zBRPRNWriteText.Call(
        )
    return SuccessReturn(ret)
}

// extern int ZBRPRNWriteTextEx(IntPtr hprinter, int printertype, int startx, int starty, int rotation, int isbold, int widht, int height, int gmode, byte[] text, int isvarnish, out int err);
func ZBRPRNWriteTextEx() (success SuccessReturn) {
    panic("ZBRPRNWriteTextEx not implemented")
    ret, _, _ := zBRPRNWriteTextEx.Call(
        )
    return SuccessReturn(ret)
}
