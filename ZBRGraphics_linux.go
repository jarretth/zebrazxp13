package zebrazxp13

func ZBRGDIGetSDKVer() (major, minor, engLevel int) { panic("ZBRGDIGetSDKVer Cannot be called on a non-windows platform"); return }
func ZBRGDIIsPrinterReady(strPrinterName string) (ret SuccessReturn,err GraphicError) { panic("ZBRGDIIsPrinterReady Cannot be called on a non-windows platform"); return 0,0 }
func ZBRGDIGetSDKVsn() (ret SuccessReturn) { panic("ZBRGDIGetSDKVsn Cannot be called on a non-windows platform"); return  }
func ZBRGDIGetSDKProductVer() (ret SuccessReturn) { panic("ZBRGDIGetSDKProductVer Cannot be called on a non-windows platform"); return }
func ZBRGDIInitGraphics(strPrinterName string) (ret SuccessReturn, handle Handle, err GraphicError) { panic("ZBRGDIInitGraphics Cannot be called on a non-windows platform"); return 0,0,0 }
func ZBRGDIInitGraphicsEx() (ret SuccessReturn) { panic("ZBRGDIInitGraphicsEx Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIInitGraphicsFromPrintDlg() (ret SuccessReturn) { panic("ZBRGDIInitGraphicsFromPrintDlg Cannot be called on a non-windows platform"); return 0 }
func ZBRGDICloseGraphics(handle Handle) (ret SuccessReturn, err GraphicError){ panic("ZBRGDICloseGraphics Cannot be called on a non-windows platform"); return 0,0 }
func ZBRGDIClearGraphics() (ret SuccessReturn, err GraphicError) { panic("ZBRGDIClearGraphics Cannot be called on a non-windows platform"); return 0,0 }
func ZBRGDIStartPage() (ret SuccessReturn) { panic("ZBRGDIStartPage Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIEndPage() (ret SuccessReturn) { panic("ZBRGDIEndPage Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIPreviewGraphics() (ret SuccessReturn) { panic("ZBRGDIPreviewGraphics Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIPrintGraphics(handle Handle) (ret SuccessReturn, err GraphicError) { panic("ZBRGDIPrintGraphics Cannot be called on a non-windows platform"); return 0,0 }
func ZBRGDIPrintGraphicsEx() (ret SuccessReturn) { panic("ZBRGDIPrintGraphicsEx Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIPrintFilePos() (ret SuccessReturn) { panic("ZBRGDIPrintFilePos Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIPrintImagePos() (ret SuccessReturn) { panic("ZBRGDIPrintImagePos Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIPrintImageRect() (ret SuccessReturn) { panic("ZBRGDIPrintImageRect Cannot be called on a non-windows platform"); return 0 }
func ZBRGDICancelJob() (ret SuccessReturn) { panic("ZBRGDICancelJob Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIDrawText(x uint, y uint, text string, font string, fontsize uint, fontstyle TextFontStyle, color uint32) (ret SuccessReturn, err GraphicError) { panic("ZBRGDIDrawText Cannot be called on a non-windows platform"); return 0,0 }
func ZBRGDIDrawTextEx() (ret SuccessReturn) { panic("ZBRGDIDrawTextEx Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIDrawTextUnicode(x uint, y uint, text string, font string, fontsize uint, fontstyle TextFontStyle, color uint32) (ret SuccessReturn, err GraphicError) { panic("ZBRGDIDrawTextUnicode Cannot be called on a non-windows platform"); return 0,0 }
func ZBRGDIDrawTextUnicodeEx() (ret SuccessReturn) { panic("ZBRGDIDrawTextUnicodeEx Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIDrawTextRect(x uint, y uint, width uint, height uint, alignment TextAlignment, text string, font string, fontsize uint, fontstyle TextFontStyle, color uint32) (ret SuccessReturn, err GraphicError) { panic("ZBRGDIDrawTextRect Cannot be called on a non-windows platform"); return 0,0 }
func ZBRGDIDrawTextRectEx() (ret SuccessReturn) { panic("ZBRGDIDrawTextRectEx Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIDrawLine(x1 uint, y1 uint, x2 uint, y2 uint, color uint32, thickness float32) (ret SuccessReturn, err GraphicError) { panic("ZBRGDIDrawLine Cannot be called on a non-windows platform"); return 0,0 }
func ZBRGDIDrawImage() (ret SuccessReturn) { panic("ZBRGDIDrawImage Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIDrawImageEx() (ret SuccessReturn) { panic("ZBRGDIDrawImageEx Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIDrawImagePos() (ret SuccessReturn) { panic("ZBRGDIDrawImagePos Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIDrawImagePosEx() (ret SuccessReturn) { panic("ZBRGDIDrawImagePosEx Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIDrawImageRect() (ret SuccessReturn) { panic("ZBRGDIDrawImageRect Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIDrawImageRectEx() (ret SuccessReturn) { panic("ZBRGDIDrawImageRectEx Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIDrawRectangle(x uint, y uint, width uint, height uint, thickness float32, color uint32) (ret SuccessReturn, err GraphicError) { panic("ZBRGDIDrawRectangle Cannot be called on a non-windows platform"); return 0,0 }
func ZBRGDIDrawEllipse(x uint, y uint, width uint, height uint, thickness float32, color uint32) (ret SuccessReturn, err GraphicError) { panic("ZBRGDIDrawEllipse Cannot be called on a non-windows platform"); return 0,0 }
func ZBRGDIDrawBarCode(x uint, y uint, rotation BarCodeRotation, barcodetype BarCodeType, barwidthratio BarCodeWidth, barcodemultiplier uint, barcodeheight uint, textunder BarCodeTextUnder, barcodedata string) (ret SuccessReturn,err GraphicError) { panic("ZBRGDIDrawBarCode Cannot be called on a non-windows platform"); return 0,0 }
