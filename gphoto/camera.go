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

type Camera struct {
	Ref     *C.Camera
	Context *Context
}

type CameraCaptureType int

// Create a new camera instance
func NewCamera() (*Camera, error) {
	cam := &Camera{}

	if ret := C.gp_camera_new(&cam.Ref); ret != 0 {
		return nil, AsPortResult(ret).Error()
	}

	return cam, nil
}

// Initialize a camera
func (camera *Camera) Init(ctx *Context) error {
	camera.Context = ctx
	if ret := C.gp_camera_init(camera.c(), ctx.c()); ret != PORT_RESULT_OK {
		return AsPortResult(ret).Error()
	}

	return nil
}

// Autodetect all connected cameras
func (camera *Camera) Autodetect(list *CameraList) PortResult {
	list.Reset()
	ret := C.gp_camera_autodetect(list.c(), camera.Context.c())
	if ret < 0 {
		return AsPortResult(ret)
	}
	list.Count = int(ret)
	return PORT_RESULT_OK
}

// Set the abilities of the camera
// This should be called before initializing the camera, otherwise it will be autodetected
// This tells the library which model and driver to use for the camera
func (camera *Camera) SetAbilities(abilities *CameraAbilities) error {
	if ret := C.gp_camera_set_abilities(camera.c(), abilities.Ref); ret != PORT_RESULT_OK {
		return AsPortResult(ret).Error()
	}
	return nil
}

// Set the port info for the camera
func (camera *Camera) SetPortInfo(info *PortInfo) error {
	if ret := C.gp_camera_set_port_info(camera.c(), info.Ref); ret != PORT_RESULT_OK {
		return AsPortResult(ret).Error()
	}
	return nil
}

func (camera *Camera) Capture(captureType CameraCaptureType) (CameraFilePath, error) {
	var path CameraFilePath
	var _path C.CameraFilePath

	_captureType := C.CameraCaptureType(captureType)
	if ret := C.gp_camera_capture(camera.c(), _captureType, &_path, camera.Context.c()); ret != PORT_RESULT_OK {
		return CameraFilePath{"", ""}, AsPortResult(ret).Error()
	}

	path.Name = C.GoString(&_path.name[0])
	path.Folder = C.GoString(&_path.folder[0])
	return path, nil
}

func (camera *Camera) File(folder, name string, filetype CameraFileType) (*CameraFile, error) {
	var _file *C.CameraFile
	C.gp_file_new(&_file)

	_camera := (*C.Camera)(unsafe.Pointer(camera))
	_folder := C.CString(folder)
	_name := C.CString(name)
	_context := (*C.GPContext)(unsafe.Pointer(camera.Context))
	_filetype := (C.CameraFileType)(filetype)
	if ret := C.gp_camera_file_get(_camera, _folder, _name, _filetype, _file, _context); ret != PORT_RESULT_OK {
		return nil, AsPortResult(ret).Error()
	}
	return (*CameraFile)(unsafe.Pointer(_file)), nil
}

// Free camera resources
func (camera *Camera) Free() error {
	if ret := C.gp_camera_free(camera.c()); ret != PORT_RESULT_OK {
		return AsPortResult(ret).Error()
	}
	return nil
}

// c camera pointer
func (camera *Camera) c() *C.Camera {
	return (*C.Camera)(camera.Ref)
}

// c camera file path
func (path *CameraFilePath) c() *C.CameraFilePath {
	return (*C.CameraFilePath)(unsafe.Pointer(path))
}
