package camera

import "github.com/sigsegv42/shot-capture/gphoto"

type CameraService struct {
	Context   *gphoto.Context
	Abilities *gphoto.CameraAbilitiesList
	Ports     *gphoto.PortInfoList
}

func NewCameraService() *CameraService {
	service := &CameraService{}
	service.Context = gphoto.NewContext()

	return service
}

func (service *CameraService) Init() error {
	var err error
	service.Abilities, err = gphoto.NewAbilitiesList(service.Context)
	if err != nil {
		return err
	}
	err = service.Abilities.Load()
	if err != nil {
		return err
	}

	service.Ports, err = gphoto.NewPortInfoList()
	if err != nil {
		return err
	}
	err = service.Ports.Load()
	if err != nil {
		return err
	}
	err = service.Ports.Count()
	if err != nil {
		return err
	}

	return nil
}

func (service *CameraService) Free() error {
	if service.Context != nil {
		service.Context.Free()
	}
	if service.Ports != nil {
		err := service.Ports.Free()
		if err != nil {
			return err
		}
	}
	if service.Abilities != nil {
		err := service.Abilities.Free()
		if err != nil {
			return err
		}
	}

	return nil
}

func (service *CameraService) InitializeCamera(camera *Camera) error {
	err := camera.Ref.Init(service.Context)
	if err != nil {
		return err
	}

	return nil
}

func (service *CameraService) Open(model string, port string) (*Camera, error) {
	camera, err := NewCamera()
	if err != nil {
		return nil, err
	}
	modelIndex, err := service.Abilities.LookupModel(model)
	if err != nil {
		return camera, err
	}
	abilities, err := service.Abilities.GetAbilities(modelIndex)
	if err != nil {
		return camera, err
	}
	err = camera.Ref.SetAbilities(abilities)
	if err != nil {
		return camera, err
	}

	portIndex, err := service.Ports.LookupPath(port)
	if err != nil {
		return camera, err
	}
	info, err := service.Ports.GetInfo(portIndex)
	if err != nil {
		return camera, nil
	}
	err = camera.Ref.SetPortInfo(info)
	if err != nil {
		return camera, err
	}
	return camera, nil
}
