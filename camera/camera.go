package camera

import (
	"github.com/sigsegv42/shot-capture/gphoto"
)

type Camera struct {
	Ref *gphoto.Camera
}

func NewCamera() (*Camera, error) {
	camera := &Camera{}
	var err error
	camera.Ref, err = gphoto.NewCamera()
	if err != nil {
		return nil, err
	}
	return camera, nil
}

/*
func (camera *Camera) Autodetect() error {

}
*/

// Retrieve all of the information about ports and connected cameras
func (camera *Camera) SystemStats() error {
	// get port info

	return nil
}

func (camera *Camera) Free() error {
	if camera.Ref != nil {
		err := camera.Ref.Free()
		if err != nil {
			return err
		}
	}
	return nil
}
