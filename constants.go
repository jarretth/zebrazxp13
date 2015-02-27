package zebrazxp13

const (
    ZBR_SUCCESS SuccessReturn = 1
    ZBR_FAILED  SuccessReturn = 0
)

const (
    ZBR_GDI_ERROR_GENERIC_ERROR                GraphicError = 8001
    ZBR_GDI_ERROR_INVALID_PARAMETER            GraphicError = 8002
    ZBR_GDI_ERROR_OUT_OF_MEMORY                GraphicError = 8003
    ZBR_GDI_ERROR_OBJECT_BUSY                  GraphicError = 8004
    ZBR_GDI_ERROR_INSUFFICIENT_BUFFER          GraphicError = 8005
    ZBR_GDI_ERROR_NOT_IMPLEMENTED              GraphicError = 8006
    ZBR_GDI_ERROR_WIN32_ERROR                  GraphicError = 8007
    ZBR_GDI_ERROR_WRONG_STATE                  GraphicError = 8008
    ZBR_GDI_ERROR_ABORTED                      GraphicError = 8009
    ZBR_GDI_ERROR_FILE_NOT_FOUND               GraphicError = 8010
    ZBR_GDI_ERROR_VALUE_OVERFLOW               GraphicError = 8011
    ZBR_GDI_ERROR_ACCESS_DENIED                GraphicError = 8012
    ZBR_GDI_ERROR_UNKNOWN_IMAGE_FORMAT         GraphicError = 8013
    ZBR_GDI_ERROR_FONT_FAMILY_NOT_FOUND        GraphicError = 8014
    ZBR_GDI_ERROR_FONT_STYLE_NOT_FOUND         GraphicError = 8015
    ZBR_GDI_ERROR_NOT_TRUE_TYPE_FONT           GraphicError = 8016
    ZBR_GDI_ERROR_UNSUPPORTED_GDIPLUS_         GraphicError = 8017
    ZBR_GDI_ERROR_GDIPLUS_NOT_INITIALIZED      GraphicError = 8018
    ZBR_GDI_ERROR_PROPERTY_NOT_FOUND           GraphicError = 8019
    ZBR_GDI_ERROR_PROPERTY_NOT_SUPPORTED       GraphicError = 8020
    ZBR_GDI_ERROR_GRAPHICS_ALREADY_INITIALIZED GraphicError = 8021
    ZBR_GDI_ERROR_NO_GRAPHIC_DATA              GraphicError = 8022
    ZBR_GDI_ERROR_GRAPHICS_NOT_INITIALIZED     GraphicError = 8023
    ZBR_GDI_ERROR_GETTING_DEVICE_CONTEXT       GraphicError = 8024
    ZBR_DLG_ERROR_DLG_CANCELED                 GraphicError = 8025
    ZBR_DLG_ERROR_SETUP_FAILURE                GraphicError = 8026
    ZBR_DLG_ERROR_PARSE_FAILURE                GraphicError = 8027
    ZBR_DLG_ERROR_RET_DEFAULT_FAILURE          GraphicError = 8028
    ZBR_DLG_ERROR_LOAD_DRV_FAILURE             GraphicError = 8029
    ZBR_DLG_ERROR_GET_DEVMODE_FAIL             GraphicError = 8030
    ZBR_DLG_ERROR_INIT_FAILURE                 GraphicError = 8031
    ZBR_DLG_ERROR_NO_DEVICES                   GraphicError = 8032
    ZBR_DLG_ERROR_NO_DEFAULT_PRINTER           GraphicError = 8033
    ZBR_DLG_ERROR_DN_DM_MISMATCH               GraphicError = 8034
    ZBR_DLG_ERROR_CREATE_IC_FAILURE            GraphicError = 8035
    ZBR_DLG_ERROR_PRINTER_NOT_FOUND            GraphicError = 8036
    ZBR_DLG_ERROR_DEFAULT_DIFFERENT            GraphicError = 8037
)

const (
    ZBR_ERROR_PRINTER_MECHANICAL_ERROR              PrinterError = -1
    ZBR_ERROR_BROKEN_RIBBON                         PrinterError = 1
    ZBR_ERROR_TEMPERATURE                           PrinterError = 2
    ZBR_ERROR_MECHANICAL_ERROR                      PrinterError = 3
    ZBR_ERROR_OUT_OF_CARD                           PrinterError = 4
    ZBR_ERROR_CARD_IN_ENCODER                       PrinterError = 5
    ZBR_ERROR_CARD_NOT_IN_ENCODER                   PrinterError = 6
    ZBR_ERROR_PRINT_HEAD_OPEN                       PrinterError = 7
    ZBR_ERROR_OUT_OF_RIBBON                         PrinterError = 8
    ZBR_ERROR_REMOVE_RIBBON                         PrinterError = 9
    ZBR_ERROR_PARAMETERS_ERROR                      PrinterError = 10
    ZBR_ERROR_INVALID_COORDINATES                   PrinterError = 11
    ZBR_ERROR_UNKNOWN_BARCODE                       PrinterError = 12
    ZBR_ERROR_UNKNOWN_TEXT                          PrinterError = 13
    ZBR_ERROR_COMMAND_ERROR                         PrinterError = 14
    ZBR_ERROR_BARCODE_DATA_SYNTAX                   PrinterError = 20
    ZBR_ERROR_TEXT_DATA_SYNTAX                      PrinterError = 21
    ZBR_ERROR_GRAPHIC_DATA_SYNTAX                   PrinterError = 22
    ZBR_ERROR_GRAPHIC_IMAGE_INITIALIZATION          PrinterError = 30
    ZBR_ERROR_GRAPHIC_IMAGE_MAXIMUM_WIDTH_EXCEEDED  PrinterError = 31
    ZBR_ERROR_GRAPHIC_IMAGE_MAXIMUM_HEIGHT_EXCEEDED PrinterError = 32
    ZBR_ERROR_GRAPHIC_IMAGE_DATA_CHECKSUM_ERROR     PrinterError = 33
    ZBR_ERROR_DATA_TRANSFER_TIME_OUT                PrinterError = 34
    ZBR_ERROR_CHECK_RIBBON                          PrinterError = 35
    ZBR_ERROR_INVALID_MAGNETIC_DATA                 PrinterError = 40
    ZBR_ERROR_MAG_ENCODER_WRITE                     PrinterError = 41
    ZBR_ERROR_READING_ERROR                         PrinterError = 42
    ZBR_ERROR_MAG_ENCODER_MECHANICAL                PrinterError = 43
    ZBR_ERROR_MAG_ENCODER_NOT_RESPONDING            PrinterError = 44
    ZBR_ERROR_MAG_ENCODER_MISSING_OR_CARD_          PrinterError = 45
    ZBR_ERROR_ROTATION_ERROR                        PrinterError = 47
    ZBR_ERROR_COVER_OPEN                            PrinterError = 48
    ZBR_ERROR_ENCODING_ERROR                        PrinterError = 49
    ZBR_ERROR_MAGNETIC_ERROR                        PrinterError = 50
    ZBR_ERROR_BLANK_TRACK                           PrinterError = 51
    ZBR_ERROR_FLASH_ERROR                           PrinterError = 52
    ZBR_ERROR_NO_ACCESS                             PrinterError = 53
    ZBR_ERROR_SEQUENCE_ERROR                        PrinterError = 54
    ZBR_ERROR_PROX_ERROR                            PrinterError = 55
    ZBR_ERROR_CONTACT_DATA_ERROR                    PrinterError = 56
    ZBR_ERROR_PROX_DATA_ERROR                       PrinterError = 57
    ZBR_SDK_ERROR_PRINTER_NOT_SUPPORTED             PrinterError = 60
    ZBR_SDK_ERROR_CANNOT_GET_PRINTER_HANDLE         PrinterError = 61
    ZBR_SDK_ERROR_CANNOT_GET_PRINTER_DRIVER         PrinterError = 62
    ZBR_SDK_ERROR_INVALID_PARAMETER                 PrinterError = 63
    ZBR_SDK_ERROR_PRINTER_BUSY                      PrinterError = 64
    ZBR_SDK_ERROR_INVALID_PRINTER_HANDLE            PrinterError = 65
    ZBR_SDK_ERROR_CLOSE_HANDLE_ERROR                PrinterError = 66
    ZBR_SDK_ERROR_COMMUNICATION_ERROR               PrinterError = 67
    ZBR_SDK_ERROR_BUFFER_OVERFLOW                   PrinterError = 68
    ZBR_SDK_ERROR_READ_DATA_ERROR                   PrinterError = 69
    ZBR_SDK_ERROR_WRITE_DATA_ERROR                  PrinterError = 70
    ZBR_SDK_ERROR_LOAD_LIBRARY_ERROR                PrinterError = 71
    ZBR_SDK_ERROR_INVALID_STRUCT_ALIGNMENT          PrinterError = 72
    ZBR_SDK_ERROR_GETTING_DEVICE_CONTEXT            PrinterError = 73
    ZBR_SDK_ERROR_SPOOLER_ERROR                     PrinterError = 74
    ZBR_SDK_ERROR_OUT_OF_MEMORY                     PrinterError = 75
    ZBR_SDK_ERROR_OUT_OF_DISK_SPACE                 PrinterError = 76
    ZBR_SDK_ERROR_USER_ABORT                        PrinterError = 77
    ZBR_SDK_ERROR_APPLICATION_ABORT                 PrinterError = 78
    ZBR_SDK_ERROR_CREATE_FILE_ERROR                 PrinterError = 79
    ZBR_SDK_ERROR_WRITE_FILE_ERROR                  PrinterError = 80
    ZBR_SDK_ERROR_READ_FILE_ERROR                   PrinterError = 81
    ZBR_SDK_ERROR_INVALID_MEDIA                     PrinterError = 82
    ZBR_SDK_ERROR_MEMORY_ALLOCATION                 PrinterError = 83
    ZBR_SDK_ERROR_UNKNOWN_ERROR                     PrinterError = 255
)


const (
    TEXT_FONT_STYLE_BOLD      TextFontStyle = 1
    TEXT_FONT_STYLE_ITALIC    TextFontStyle = 2
    TEXT_FONT_STYLE_UNDERLINE TextFontStyle = 4
    TEXT_FONT_STYLE_STRIKE    TextFontStyle = 8
)

const (
    TEXT_ALIGN_CENTER TextAlignment = 4
    TEXT_ALIGN_LEFT   TextAlignment = 5
    TEXT_ALIGN_RIGHT  TextAlignment = 6
)

const (
    TEXT_FONT_ARIAL = "Arial"
)

const (
    BARCODE_ROTATION_LOWER_LEFT_0   BarCodeRotation = 0
    BARCODE_ROTATION_LOWER_LEFT_90  BarCodeRotation = 1
    BARCODE_ROTATION_LOWER_LEFT_180 BarCodeRotation = 2
    BARCODE_ROTATION_LOWER_LEFT_270 BarCodeRotation = 3
    BARCODE_ROTATION_CENTER_0       BarCodeRotation = 4
    BARCODE_ROTATION_CENTER_90      BarCodeRotation = 5
    BARCODE_ROTATION_CENTER_180     BarCodeRotation = 6
    BARCODE_ROTATION_CENTER_270     BarCodeRotation = 7
)

const (
    BARCODE_TYPE_CODE_39           BarCodeType = 0
    BARCODE_TYPE_25_INTERLEAVE     BarCodeType = 1
    BARCODE_TYPE_25_INDUSTRIAL     BarCodeType = 2
    BARCODE_TYPE_EAN_8             BarCodeType = 3
    BARCODE_TYPE_EAN_13            BarCodeType = 4
    BARCODE_TYPE_UPCA              BarCodeType = 5
    BARCODE_TYPE_CODE_128C_NOCHECK BarCodeType = 7
    BARCODE_TYPE_CODE_128B_NOCHECK BarCodeType = 8
    BARCODE_TYPE_CODE_128C_CHECK   BarCodeType = 107
    BARCODE_TYPE_CODE_128B_CHECK   BarCodeType = 108
)

const (
    BARCODE_WIDTH_1_2 BarCodeWidth = 0
    BARCODE_WIDTH_1_3 BarCodeWidth = 1
    BARCODE_WIDTH_2_5 BarCodeWidth = 2
)

const (
    BARCODE_TEXT_UNDER_YES BarCodeTextUnder = 1
    BARCODE_TEXT_UNDER_NO  BarCodeTextUnder = 0
)
