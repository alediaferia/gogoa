package gogoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
//#include "g_window.h"
import "C"
import "unsafe"

type Window struct {
	ptr unsafe.Pointer
}

func NewWindow(x, y, width, height int) (*Window) {
	w := new(Window)
	w.ptr = C.NewWindow(C.int(x), C.int(y), C.int(width), C.int(height))

	return w
}

func (self *Window) MakeKeyAndOrderFront() {
	C.MakeKeyAndOrderFront(self.ptr)
}

func (self *Window) SetTitle(title string) {
	C.SetTitle(self.ptr, C.CString(title))
}
