package camera

import (
	"github.com/sigsegv42/shot-capture/gphoto"
)

type Camera struct {
	Context *gphoto.Context
	Ref     *gphoto.Camera
}

func NewCamera() *Camera {
	camera := &Camera{}
	return camera
}

func (camera *Camera) Init() error {
	// initialize camera
	camera.Context = gphoto.NewContext()

	var err error
	camera.Ref, err = gphoto.NewCamera()
	if err != nil {
		return err
	}

	err = camera.Ref.Init(camera.Context)
	if err != nil {
		return err
	}

	return nil
}

func (camera *Camera) Free() error {
	camera.Context.Free()

	err := camera.Ref.Free()
	if err != nil {
		return err
	}
	return nil
}
