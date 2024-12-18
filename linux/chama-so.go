package main

import (
	"fmt"
)

// #cgo LDFLAGS: -llibimgutil32
//
// #include <libimgutil32.h>

import "C"

func main() {
	i := C.libimgutil32.ImgutilGetImageSize("m.jpg", "", 0, 0)
	fmt.Printf("Hello dll function returns %d\n", i)
}
