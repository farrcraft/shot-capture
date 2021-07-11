package gphoto

// Binding sourced from: https://github.com/aqiank/go-gphoto2
// We set -rpath in LDFLAGS so that we don't need to use LD_LIBRARY_PATH at runtime

// #cgo linux CFLAGS: -I/opt/shot-capture/include
// #cgo linux LDFLAGS: -L/opt/shot-capture/lib -Wl,-rpath -Wl,/opt/shot-capture/lib -lgphoto2 -lgphoto2_port
// #include <gphoto2/gphoto2.h>
// #include <string.h>
import "C"
import "unsafe"

const (
	CAPTURE_IMAGE = C.GP_CAPTURE_IMAGE
	CAPTURE_MOVIE = C.GP_CAPTURE_MOVIE
	CAPTURE_SOUND = C.GP_CAPTURE_SOUND
)

type Camera C.Camera
type CameraCaptureType int

func NewCamera() (*Camera, error) {
	var _cam *C.Camera

	if ret := C.gp_camera_new(&_cam); ret != 0 {
		return nil, AsPortResult(ret).Error()
	}

	return (*Camera)(_cam), nil
}

func (camera *Camera) Init(ctx *Context) error {
	if ret := C.gp_camera_init(camera.c(), ctx.c()); ret != 0 {
		return AsPortResult(ret).Error()
	}

	return nil
}

func (camera *Camera) Capture(captureType CameraCaptureType, ctx *Context) (CameraFilePath, error) {
	var path CameraFilePath
	var _path C.CameraFilePath

	_captureType := C.CameraCaptureType(captureType)
	if ret := C.gp_camera_capture(camera.c(), _captureType, &_path, ctx.c()); ret != 0 {
		return CameraFilePath{"", ""}, AsPortResult(ret).Error()
	}

	path.Name = C.GoString(&_path.name[0])
	path.Folder = C.GoString(&_path.folder[0])
	return path, nil
}

func (camera *Camera) File(folder, name string, filetype CameraFileType, context *Context) (*CameraFile, error) {
	var _file *C.CameraFile
	C.gp_file_new(&_file)

	_camera := (*C.Camera)(unsafe.Pointer(camera))
	_folder := C.CString(folder)
	_name := C.CString(name)
	_context := (*C.GPContext)(unsafe.Pointer(context))
	_filetype := (C.CameraFileType)(filetype)
	if ret := C.gp_camera_file_get(_camera, _folder, _name, _filetype, _file, _context); ret != 0 {
		return nil, AsPortResult(ret).Error()
	}
	return (*CameraFile)(unsafe.Pointer(_file)), nil
}

func (camera *Camera) Free() error {
	if ret := C.gp_camera_free(camera.c()); ret != 0 {
		return AsPortResult(ret).Error()
	}
	return nil
}

func (camera *Camera) c() *C.Camera {
	return (*C.Camera)(camera)
}

func (path *CameraFilePath) c() *C.CameraFilePath {
	return (*C.CameraFilePath)(unsafe.Pointer(path))
}
