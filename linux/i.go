package main

// #include <stdio.h>
// #include <stdlib.h>

import (
	"C"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

//export ImgutilGetImageSize
func ImgutilGetImageSize(path *C.char, retorno **C.char, w *int, h *int) (r int) {
	fmt.Println(C.GoString(path))
	fmt.Println(*C.GoString(retorno))

	file, err := os.Open(C.GoString(path))
	if err != nil {
		fmt.Println(-1, err.Error())
		r = int(-2)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(-2, err.Error())
		r = int(-2)
		return
	}

	rect := img.Bounds()
	*w = int(rect.Dx())
	*h = int(rect.Dy())
	*path = *C.CString("Reste")
	//	retorno = C.CString(fmt.Sprintf("From DLL: Hello, %s!\n", C.GoString(path)))
	*retorno = C.CString(fmt.Sprintf("From DLL: Hello, %s!\n", C.GoString(path)))
	// defer C.free(unsafe.Pointer(retorno))
	r = int(200)
	return
}

func main() {}
