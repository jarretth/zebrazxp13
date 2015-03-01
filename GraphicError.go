package zebrazxp13;

func (e GraphicError) Error() string {
    switch e {
        case ZBR_GDI_ERROR_GENERIC_ERROR:
            return "ZBR_GDI_ERROR_GENERIC_ERROR"
        case ZBR_GDI_ERROR_INVALID_PARAMETER:
            return "ZBR_GDI_ERROR_INVALID_PARAMETER"
        case ZBR_GDI_ERROR_OUT_OF_MEMORY:
            return "ZBR_GDI_ERROR_OUT_OF_MEMORY"
        case ZBR_GDI_ERROR_OBJECT_BUSY:
            return "ZBR_GDI_ERROR_OBJECT_BUSY"
        case ZBR_GDI_ERROR_INSUFFICIENT_BUFFER:
            return "ZBR_GDI_ERROR_INSUFFICIENT_BUFFER"
        case ZBR_GDI_ERROR_NOT_IMPLEMENTED:
            return "ZBR_GDI_ERROR_NOT_IMPLEMENTED"
        case ZBR_GDI_ERROR_WIN32_ERROR:
            return "ZBR_GDI_ERROR_WIN32_ERROR"
        case ZBR_GDI_ERROR_WRONG_STATE:
            return "ZBR_GDI_ERROR_WRONG_STATE"
        case ZBR_GDI_ERROR_ABORTED:
            return "ZBR_GDI_ERROR_ABORTED"
        case ZBR_GDI_ERROR_FILE_NOT_FOUND:
            return "ZBR_GDI_ERROR_FILE_NOT_FOUND"
        case ZBR_GDI_ERROR_VALUE_OVERFLOW:
            return "ZBR_GDI_ERROR_VALUE_OVERFLOW"
        case ZBR_GDI_ERROR_ACCESS_DENIED:
            return "ZBR_GDI_ERROR_ACCESS_DENIED"
        case ZBR_GDI_ERROR_UNKNOWN_IMAGE_FORMAT:
            return "ZBR_GDI_ERROR_UNKNOWN_IMAGE_FORMAT"
        case ZBR_GDI_ERROR_FONT_FAMILY_NOT_FOUND:
            return "ZBR_GDI_ERROR_FONT_FAMILY_NOT_FOUND"
        case ZBR_GDI_ERROR_FONT_STYLE_NOT_FOUND:
            return "ZBR_GDI_ERROR_FONT_STYLE_NOT_FOUND"
        case ZBR_GDI_ERROR_NOT_TRUE_TYPE_FONT:
            return "ZBR_GDI_ERROR_NOT_TRUE_TYPE_FONT"
        case ZBR_GDI_ERROR_UNSUPPORTED_GDIPLUS_:
            return "ZBR_GDI_ERROR_UNSUPPORTED_GDIPLUS_"
        case ZBR_GDI_ERROR_GDIPLUS_NOT_INITIALIZED:
            return "ZBR_GDI_ERROR_GDIPLUS_NOT_INITIALIZED"
        case ZBR_GDI_ERROR_PROPERTY_NOT_FOUND:
            return "ZBR_GDI_ERROR_PROPERTY_NOT_FOUND"
        case ZBR_GDI_ERROR_PROPERTY_NOT_SUPPORTED:
            return "ZBR_GDI_ERROR_PROPERTY_NOT_SUPPORTED"
        case ZBR_GDI_ERROR_GRAPHICS_ALREADY_INITIALIZED:
            return "ZBR_GDI_ERROR_GRAPHICS_ALREADY_INITIALIZED"
        case ZBR_GDI_ERROR_NO_GRAPHIC_DATA:
            return "ZBR_GDI_ERROR_NO_GRAPHIC_DATA"
        case ZBR_GDI_ERROR_GRAPHICS_NOT_INITIALIZED:
            return "ZBR_GDI_ERROR_GRAPHICS_NOT_INITIALIZED"
        case ZBR_GDI_ERROR_GETTING_DEVICE_CONTEXT:
            return "ZBR_GDI_ERROR_GETTING_DEVICE_CONTEXT"
        case ZBR_DLG_ERROR_DLG_CANCELED:
            return "ZBR_DLG_ERROR_DLG_CANCELED"
        case ZBR_DLG_ERROR_SETUP_FAILURE:
            return "ZBR_DLG_ERROR_SETUP_FAILURE"
        case ZBR_DLG_ERROR_PARSE_FAILURE:
            return "ZBR_DLG_ERROR_PARSE_FAILURE"
        case ZBR_DLG_ERROR_RET_DEFAULT_FAILURE:
            return "ZBR_DLG_ERROR_RET_DEFAULT_FAILURE"
        case ZBR_DLG_ERROR_LOAD_DRV_FAILURE:
            return "ZBR_DLG_ERROR_LOAD_DRV_FAILURE"
        case ZBR_DLG_ERROR_GET_DEVMODE_FAIL:
            return "ZBR_DLG_ERROR_GET_DEVMODE_FAIL"
        case ZBR_DLG_ERROR_INIT_FAILURE:
            return "ZBR_DLG_ERROR_INIT_FAILURE"
        case ZBR_DLG_ERROR_NO_DEVICES:
            return "ZBR_DLG_ERROR_NO_DEVICES"
        case ZBR_DLG_ERROR_NO_DEFAULT_PRINTER:
            return "ZBR_DLG_ERROR_NO_DEFAULT_PRINTER"
        case ZBR_DLG_ERROR_DN_DM_MISMATCH:
            return "ZBR_DLG_ERROR_DN_DM_MISMATCH"
        case ZBR_DLG_ERROR_CREATE_IC_FAILURE:
            return "ZBR_DLG_ERROR_CREATE_IC_FAILURE"
        case ZBR_DLG_ERROR_PRINTER_NOT_FOUND:
            return "ZBR_DLG_ERROR_PRINTER_NOT_FOUND"
        case ZBR_DLG_ERROR_DEFAULT_DIFFERENT:
            return "ZBR_DLG_ERROR_DEFAULT_DIFFERENT"
        case 0:
            return "Success"
        default:
            return "Unknown Error"
    }
}
