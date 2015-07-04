package gogoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
//#include "g_application.h"
import "C"
import "unsafe"


type Application struct {
	ptr unsafe.Pointer
}

func SharedApplication() (*Application) {
	app := new(Application)
	app.ptr = C.SharedApplication()

	return app
}

func (this *Application) Run() {
	C.Run(this.ptr)
}
