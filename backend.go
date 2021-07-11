package main

import (
	"os"

	"github.com/sigsegv42/shot-capture/camera"
	"github.com/sigsegv42/shot-capture/core"
	"github.com/sirupsen/logrus"
)

type Backend struct {
	Logger *logrus.Logger
	Camera *camera.Camera
	Config *core.Config
}

func NewBackend(configFileName string) (*Backend, error) {
	backend := &Backend{
		Logger: logrus.New(),
	}

	// load config
	var err error
	backend.Config, err = core.NewConfig(configFileName)
	if err != nil {
		return nil, err
	}

	// setup logging
	backend.Logger.Formatter = &logrus.JSONFormatter{}
	level, err := logrus.ParseLevel(backend.Config.LogLevel)
	if err != nil {
		return nil, err
	}
	backend.Logger.Level = level

	var file *os.File
	file, err = os.OpenFile(backend.Config.LogPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0640)
	if err != nil {
		return nil, err
	}
	backend.Logger.Out = file

	// create camera
	backend.Camera = camera.NewCamera()

	return backend, nil
}

func (backend *Backend) Run() bool {
	backend.Logger.Info("Starting backend...")
	err := backend.Camera.Init()
	if err != nil {
		backend.Logger.Error("Error initializing camera - ", err)
		return false
	}
	return true
}

func (backend *Backend) Shutdown() bool {
	err := backend.Camera.Free()
	if err != nil {
		backend.Logger.Error("Error cleaning up camera - ", err)
		return false
	}
	return true
}
