package gphoto

// Binding sourced from: https://github.com/aqiank/go-gphoto2

// #cgo linux CFLAGS: -I/opt/shot-capture/include
// #include <gphoto2/gphoto2.h>
// #include <string.h>
import "C"
import "unsafe"

const (
	WIDGET_TYPE_TOGGLE  = C.GP_WIDGET_TOGGLE
	WIDGET_TYPE_WINDOW  = C.GP_WIDGET_WINDOW
	WIDGET_TYPE_SECTION = C.GP_WIDGET_SECTION
	WIDGET_TYPE_TEXT    = C.GP_WIDGET_TEXT
	WIDGET_TYPE_RANGE   = C.GP_WIDGET_RANGE
	WIDGET_TYPE_RADIO   = C.GP_WIDGET_RADIO
	WIDGET_TYPE_MENU    = C.GP_WIDGET_MENU
	WIDGET_TYPE_BUTTON  = C.GP_WIDGET_BUTTON
	WIDGET_TYPE_DATE    = C.GP_WIDGET_DATE
)

type CameraWidget struct {
	Ref   *C.CameraWidget
	Type  C.CameraWidgetType
	Name  string
	Value int
}

// Get the type of the widget
func (widget *CameraWidget) GetType() error {
	if ret := C.gp_widget_get_type(widget.Ref, &widget.Type); ret != PORT_RESULT_OK {
		return AsPortResult(ret).Error()
	}
	return nil
}

// Get the name of the widget
func (widget *CameraWidget) GetName() error {
	var name *C.char
	if ret := C.gp_widget_get_name(widget.Ref, &name); ret != PORT_RESULT_OK {
		return AsPortResult(ret).Error()
	}
	widget.Name = C.GoString(name)
	return nil
}

// Get the child of the widget
func (widget *CameraWidget) GetChild(key string) (*CameraWidget, error) {
	child := &CameraWidget{}
	if ret := C.gp_widget_get_child_by_name(widget.Ref, C.CString(key), &child.Ref); ret < PORT_RESULT_OK {
		// couldn't find by name, try by label instead
		ret = C.gp_widget_get_child_by_label(widget.Ref, C.CString(key), &child.Ref)
		if ret < PORT_RESULT_OK {
			return nil, AsPortResult(ret).Error()
		}
	}
	return child, nil
}

// Get the value of the widget
// This currently only works with numeric value types
func (widget *CameraWidget) GetValue() error {
	var value C.int
	if ret := C.gp_widget_get_value(widget.Ref, unsafe.Pointer(&value)); ret != PORT_RESULT_OK {
		return AsPortResult(ret).Error()
	}
	widget.Value = int(value)
	return nil
}

// Set the value of the widget
// This currently only works with numeric value types
func (widget *CameraWidget) SetValue(value int) error {
	_val := C.int(value)
	if ret := C.gp_widget_set_value(widget.Ref, unsafe.Pointer(&_val)); ret != PORT_RESULT_OK {
		return AsPortResult(ret).Error()
	}
	widget.Value = int(_val)
	return nil
}
