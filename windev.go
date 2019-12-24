package windev
/*
#cgo CFLAGS: -I./cdep
#cgo LDFLAGS: -L./lib -lwindevice
#include "windevice.h"
*/
import "C"

// KeyDownUp ..
func KeyDownUp(e int) C.int{
	return C.keyDownUp(C.int(e))
}

// SetCursorPos ..
func SetCursorPos(x, y int) C.bool{
	return C.setCursorPos(C.int(x), C.int(y))
}

// MouseEvent ..
func MouseEvent(dwFlags, dx, dy, dwData int, dwExtraInfo int64) {
	C.mouseEvent(C.int(dwFlags), C.int(dx), C.int(dy), C.int(dwData), C.longlong(dwExtraInfo))
}
