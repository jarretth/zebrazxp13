// +build !windows
package zebrazxp13

func ZBRGDIGetSDKVer() (major, minor, engLevel int) { panic("ZBRGDIGetSDKVer Cannot be called on a non-windows platform"); return }
func ZBRGDIIsPrinterReady(strPrinterName string) (ret uintptr,err uint) { panic("ZBRGDIIsPrinterReady Cannot be called on a non-windows platform"); return 0,0 }
func ZBRGDIGetSDKVsn() (ret uintptr) { panic("ZBRGDIGetSDKVsn Cannot be called on a non-windows platform"); return  }
func ZBRGDIGetSDKProductVer() (ret uintptr) { panic("ZBRGDIGetSDKProductVer Cannot be called on a non-windows platform"); return }
func ZBRGDIInitGraphics(strPrinterName string) (ret uintptr, handle Handle, err uint) { panic("ZBRGDIInitGraphics Cannot be called on a non-windows platform"); return 0,0,0 }
func ZBRGDIInitGraphicsEx() (ret uintptr) { panic("ZBRGDIInitGraphicsEx Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIInitGraphicsFromPrintDlg() (ret uintptr) { panic("ZBRGDIInitGraphicsFromPrintDlg Cannot be called on a non-windows platform"); return 0 }
func ZBRGDICloseGraphics(handle Handle) (ret uintptr, err uint){ panic("ZBRGDICloseGraphics Cannot be called on a non-windows platform"); return 0,0 }
func ZBRGDIClearGraphics() (ret uintptr, err uint) { panic("ZBRGDIClearGraphics Cannot be called on a non-windows platform"); return 0,0 }
func ZBRGDIStartPage() (ret uintptr) { panic("ZBRGDIStartPage Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIEndPage() (ret uintptr) { panic("ZBRGDIEndPage Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIPreviewGraphics() (ret uintptr) { panic("ZBRGDIPreviewGraphics Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIPrintGraphics(handle Handle) (ret uintptr, err uint) { panic("ZBRGDIPrintGraphics Cannot be called on a non-windows platform"); return 0,0 }
func ZBRGDIPrintGraphicsEx() (ret uintptr) { panic("ZBRGDIPrintGraphicsEx Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIPrintFilePos() (ret uintptr) { panic("ZBRGDIPrintFilePos Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIPrintImagePos() (ret uintptr) { panic("ZBRGDIPrintImagePos Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIPrintImageRect() (ret uintptr) { panic("ZBRGDIPrintImageRect Cannot be called on a non-windows platform"); return 0 }
func ZBRGDICancelJob() (ret uintptr) { panic("ZBRGDICancelJob Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIDrawText(x uint, y uint, text string, font string, fontsize uint, fontstyle TextFontStyle, color uint32) (ret uintptr, err uint) { panic("ZBRGDIDrawText Cannot be called on a non-windows platform"); return 0,0 }
func ZBRGDIDrawTextEx() (ret uintptr) { panic("ZBRGDIDrawTextEx Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIDrawTextUnicode(x uint, y uint, text string, font string, fontsize uint, fontstyle TextFontStyle, color uint32) (ret uintptr, err uint) { panic("ZBRGDIDrawTextUnicode Cannot be called on a non-windows platform"); return 0,0 }
func ZBRGDIDrawTextUnicodeEx() (ret uintptr) { panic("ZBRGDIDrawTextUnicodeEx Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIDrawTextRect(x uint, y uint, width uint, height uint, alignment TextAlignment, text string, font string, fontsize uint, fontstyle TextFontStyle, color uint32) (ret uintptr, err uint) { panic("ZBRGDIDrawTextRect Cannot be called on a non-windows platform"); return 0,0 }
func ZBRGDIDrawTextRectEx() (ret uintptr) { panic("ZBRGDIDrawTextRectEx Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIDrawLine(x1 uint, y1 uint, x2 uint, y2 uint, color uint32, thickness float32) (ret uintptr, err uint) { panic("ZBRGDIDrawLine Cannot be called on a non-windows platform"); return 0,0 }
func ZBRGDIDrawImage() (ret uintptr) { panic("ZBRGDIDrawImage Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIDrawImageEx() (ret uintptr) { panic("ZBRGDIDrawImageEx Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIDrawImagePos() (ret uintptr) { panic("ZBRGDIDrawImagePos Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIDrawImagePosEx() (ret uintptr) { panic("ZBRGDIDrawImagePosEx Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIDrawImageRect() (ret uintptr) { panic("ZBRGDIDrawImageRect Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIDrawImageRectEx() (ret uintptr) { panic("ZBRGDIDrawImageRectEx Cannot be called on a non-windows platform"); return 0 }
func ZBRGDIDrawRectangle(x uint, y uint, width uint, height uint, thickness float32, color uint32) (ret uintptr, err uint) { panic("ZBRGDIDrawRectangle Cannot be called on a non-windows platform"); return 0,0 }
func ZBRGDIDrawEllipse(x uint, y uint, width uint, height uint, thickness float32, color uint32) (ret uintptr, err uint) { panic("ZBRGDIDrawEllipse Cannot be called on a non-windows platform"); return 0,0 }
func ZBRGDIDrawBarCode(x uint, y uint, rotation BarCodeRotation, barcodetype BarCodeType, barwidthratio BarCodeWidth, barcodemultiplier uint, barcodeheight uint, textunder BarCodeTextUnder, barcodedata string) (ret uintptr,err uint) { panic("ZBRGDIDrawBarCode Cannot be called on a non-windows platform"); return 0,0 }
