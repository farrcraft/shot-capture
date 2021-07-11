package gphoto

// Binding sourced from: https://github.com/aqiank/go-gphoto2

// #cgo linux CFLAGS: -I/opt/shot-capture/include
// #include <gphoto2/gphoto2.h>
// #include <string.h>
import "C"
import "errors"

const (
	PORT_RESULT_OK                         = C.GP_OK
	PORT_RESULT_ERROR                      = C.GP_ERROR
	PORT_RESULT_ERROR_BAD_PARAMETERS       = C.GP_ERROR_BAD_PARAMETERS
	PORT_RESULT_ERROR_NO_MEMORY            = C.GP_ERROR_NO_MEMORY
	PORT_RESULT_ERROR_LIBRARY              = C.GP_ERROR_LIBRARY
	PORT_RESULT_ERROR_UNKNOWN_PORT         = C.GP_ERROR_UNKNOWN_PORT
	PORT_RESULT_ERROR_NOT_SUPPORTED        = C.GP_ERROR_NOT_SUPPORTED
	PORT_RESULT_ERROR_IO                   = C.GP_ERROR_IO
	PORT_RESULT_ERROR_FIXED_LIMIT_EXCEEDED = C.GP_ERROR_FIXED_LIMIT_EXCEEDED
	PORT_RESULT_ERROR_TIMEOUT              = C.GP_ERROR_TIMEOUT
	PORT_RESULT_ERROR_IO_SUPPORTED_SERIAL  = C.GP_ERROR_IO_SUPPORTED_SERIAL
	PORT_RESULT_ERROR_IO_SUPPORTED_USB     = C.GP_ERROR_IO_SUPPORTED_USB
	PORT_RESULT_ERROR_IO_INIT              = C.GP_ERROR_IO_INIT
	PORT_RESULT_ERROR_IO_READ              = C.GP_ERROR_IO_READ
	PORT_RESULT_ERROR_IO_WRITE             = C.GP_ERROR_IO_WRITE
	PORT_RESULT_ERROR_IO_UPDATE            = C.GP_ERROR_IO_UPDATE
	PORT_RESULT_ERROR_IO_SERIAL_SPEED      = C.GP_ERROR_IO_SERIAL_SPEED
	PORT_RESULT_ERROR_IO_USB_CLEAR_HALT    = C.GP_ERROR_IO_USB_CLEAR_HALT
	PORT_RESULT_ERROR_IO_USB_FIND          = C.GP_ERROR_IO_USB_FIND
	PORT_RESULT_ERROR_IO_USB_CLAIM         = C.GP_ERROR_IO_USB_CLAIM
	PORT_RESULT_ERROR_IO_LOCK              = C.GP_ERROR_IO_LOCK
	PORT_RESULT_ERROR_HAL                  = C.GP_ERROR_HAL
)

type PortResult C.int

func AsPortResult(ret C.int) PortResult {
	return PortResult(ret)
}

func (result PortResult) String() string {
	return C.GoString(C.gp_port_result_as_string(C.int(result)))
}

func (result PortResult) Error() error {
	return errors.New(C.GoString(C.gp_result_as_string(C.int(result))))
}
