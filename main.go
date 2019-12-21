package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

// wailsinit is a struct who's sole purpose is to be bound by the Wails runtime
type wailsinit struct {
	log     *wails.CustomLogger
	runtime *wails.Runtime
}

// WailsInit is automatically called by the Wails runtime during startup
// The callback is used to startup the processing code.
func (w *wailsinit) WailsInit(runtime *wails.Runtime) error {
	w.log = runtime.Log.New("WailsKiosk")
	w.runtime = runtime
	runtime.Events.On("wails:ready", func(data ...interface{}) {
		runtime.Window.Fullscreen()
	})
	return nil
}

func (w *wailsinit) WailsShutdown() {

}

func (w *wailsinit) Shutdown() {
	w.runtime.Window.Close()
}

func basic() string {
	return "Hello World!"
}

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1024,
		Height:    768,
		Resizable: true,
		Title:     "WailsKiosk",
		JS:        js,
		CSS:       css,
		Colour:    "#131313",
	})
	app.Bind(basic)
	app.Bind(&wailsinit{})
	app.Run()
}
