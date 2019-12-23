package windev
/*
#cgo CFLAGS: -I./cdep
#cgo LDFLAGS: -L./lib -lwindevice
#include "windevice.h"
*/
import "C"

func KeyDownUp(e C.int) C.int{
	return C.keyDownUp(C.int(e))
}

func SetCursorPos(x, y C.int) C.bool{
	return C.setCursorPos(C.int(x), C.int(y))
}

func MouseEvent(dwFlags, dx, dy, dwData C.int, dwExtraInfo C.longlong) {
	C.mouseEvent(C.int(dwFlags), C.int(dx), C.int(dy), C.int(dwData), C.longlong(dwExtraInfo))
}