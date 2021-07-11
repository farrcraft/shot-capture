package ui

import "github.com/gotk3/gotk3/gtk"

type Window struct {
	Ref gtk.Window
}

func NewWindow() (*Window, error) {
	window := &Window{}
	return window, nil
}

func (window *Window) Init() error {
	gtk.Init(nil)

	var err error
	window.Ref, err = gtk.windowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return err
	}

	window.Ref.SetDefaultSize(800, 600)
	window.Ref.SetTitle("shot-capture")
	window.Ref.Connect("destroy", func() {
		gtk.MainQuit()
	}
	return nil
}

func (window *Window) Destroy() error {
	return nil
}

func (window *Window) Run() error {
	window.Ref.ShowAll()
	gtk.Main()
	return nil
}
