package main

import (
	"fmt"
	"os"

	"github.com/sigsegv42/shot-capture/gphoto"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	// setup logging
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)

	// setup CLI
	app := &cli.App{
		Name:  "shot-capture",
		Usage: "capture shots from camera",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Error("Couldn't setup CLI - ", err)
		os.Exit(-1)
	}

	// initialize camera
	context := gphoto.NewContext()
	defer context.Free()

	camera, err := gphoto.NewCamera()
	if err != nil {
		log.Error("Couldn't create new camera - ", err)
		os.Exit(-1)
	}

	err = camera.Init(context)
	if err != nil {
		log.Error("Failed to initialize camera - ", err)
		os.Exit(-1)
	}

	path, err := camera.Capture(gphoto.CAPTURE_IMAGE, context)
	if err != nil {
		log.Error("Error capturing image - ", err)
		os.Exit(-1)
	}
	fmt.Println(path.Name)

	err = camera.Free()
	if err != nil {
		log.Error("Couldn't free camera - ", err)
		os.Exit(-1)
	}

}
