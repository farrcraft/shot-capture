package gphoto

// Binding sourced from: https://github.com/aqiank/go-gphoto2

// #cgo linux CFLAGS: -I/opt/shot-capture/include
// #include <gphoto2/gphoto2.h>
// #include <string.h>
import "C"
import "errors"

func e(ret C.int) error {
	return errors.New(C.GoString(C.gp_result_as_string(ret)))
}
