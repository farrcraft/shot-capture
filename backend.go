package main

import (
	"os"

	"github.com/sigsegv42/shot-capture/camera"
	"github.com/sigsegv42/shot-capture/core"
	"github.com/sigsegv42/shot-capture/gphoto"
	"github.com/sirupsen/logrus"
)

type Backend struct {
	Logger        *logrus.Logger
	Camera        *camera.Camera
	Config        *core.Config
	CameraService *camera.CameraService
}

func NewBackend(configFileName string) (*Backend, error) {
	backend := &Backend{
		Logger: logrus.New(),
	}

	// load config
	var err error
	backend.Config, err = core.NewConfig(configFileName)
	if err != nil {
		return backend, err
	}

	// setup logging
	backend.Logger.Formatter = &logrus.JSONFormatter{}
	level, err := logrus.ParseLevel(backend.Config.LogLevel)
	if err != nil {
		return backend, err
	}
	backend.Logger.Level = level

	var file *os.File
	file, err = os.OpenFile(backend.Config.LogPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0640)
	if err != nil {
		return backend, err
	}
	backend.Logger.Out = file

	// create camera
	backend.Camera, err = camera.NewCamera()
	if err != nil {
		return backend, err
	}

	return backend, nil
}

// Run the app
func (backend *Backend) Run() bool {
	backend.Logger.Info("Starting backend...")

	backend.CameraService = camera.NewCameraService()
	err := backend.CameraService.Init()
	if err != nil {
		backend.Logger.Error("Error initializing camera service - ", err)
		return false
	}

	if ok := backend.LogPorts(); !ok {
		return false
	}

	list, err := gphoto.NewList()
	if err != nil {
		backend.Logger.Error("Error creating camera list - ", err)
		return false
	}
	list.Reset()
	err = backend.Camera.Ref.Autodetect(list)
	if err != nil {
		backend.Logger.Error("Error autodetecting cameras - ", err)
		return false
	}

	backend.Logger.Info("Autodetected cameras - ", list.Count)
	for listIndex := 0; listIndex < list.Count; listIndex++ {
		name, err := list.Name(listIndex)
		if err != nil {
			backend.Logger.Error("Error getting list item name - ", err)
			return false
		}
		value, err := list.Value(listIndex)
		if err != nil {
			backend.Logger.Error("Error getting list item value - ", err)
			return false
		}
		backend.Logger.Info("Item", name, value)
	}
	list.Free()

	/*
		err := backend.Camera.Init()
		if err != nil {
			backend.Logger.Error("Error initializing camera - ", err)
			return false
		}
	*/

	return true
}

// Log all of the detected port names
func (backend *Backend) LogPorts() bool {
	backend.Logger.Info("Found ports - ", backend.CameraService.Ports.Size)

	for listIndex := 0; listIndex < backend.CameraService.Ports.Size; listIndex++ {
		info, err := backend.CameraService.Ports.GetInfo(listIndex)
		if err != nil {
			backend.Logger.Error("Error getting port info - ", err)
			return false
		}
		err = info.GetName()
		if err != nil {
			backend.Logger.Error("Error getting port info name - ", err)
			return false
		}
		backend.Logger.Info("port info name - ", info.Name)
	}
	return true
}

// Shutdown backend
func (backend *Backend) Shutdown() bool {
	var err error
	if backend.Camera != nil {
		err = backend.Camera.Free()
		if err != nil {
			backend.Logger.Error("Error cleaning up camera - ", err)
			return false
		}
	}
	if backend.CameraService != nil {
		err = backend.CameraService.Free()
		if err != nil {
			backend.Logger.Error("Error freeing camera service - ", err)
		}
	}
	return true
}
