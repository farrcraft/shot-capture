package main

import (
	"os"

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
		Usage: "capture images from camera",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
			},
			&cli.BoolFlag{
				Name:  "ui",
				Usage: "Run in GUI mode",
			},
			&cli.BoolFlag{
				Name:  "list-ports",
				Usage: "Display available ports",
			},
			&cli.BoolFlag{
				Name:  "autodetect",
				Usage: "Autodetect available cameras",
			},
			&cli.StringFlag{
				Name:  "camera",
				Usage: "Select `CAMERA`",
			},
			&cli.BoolFlag{
				Name:  "dump-camera-config",
				Usage: "Print camera configuration",
			},
		},
		Action: runApp,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Error("Couldn't run CLI - ", err)
		os.Exit(-1)
	}
	/*
		path, err := camera.Capture(gphoto.CAPTURE_IMAGE, context)
		if err != nil {
			log.Error("Error capturing image - ", err)
			os.Exit(-1)
		}
		fmt.Println(path.Name)
	*/
}

func runApp(ctx *cli.Context) error {
	configFileName := ctx.String("config")
	if configFileName == "" {
		return cli.Exit("Missing config!", -1)
	}
	backend, err := NewBackend(configFileName)
	if err != nil {
		return err
	}
	exitCode := 1
	if ctx.Bool("ui") {
		if ok := backend.runUI(); !ok {
			exitCode = -1
		}
	} else {
		// set backend options
		backend.Options.Autodetect = ctx.Bool("autodetect")
		backend.Options.ListPorts = ctx.Bool("list-ports")
		backend.Options.DisplayCameraConfig = ctx.Bool("dump-camera-config")

		if ok := backend.Run(); !ok {
			exitCode = -1
		}
	}

	if ok := backend.Shutdown(); !ok {
		exitCode = -1
	}

	if exitCode != 1 {
		os.Exit(exitCode)
	}

	return nil
}
