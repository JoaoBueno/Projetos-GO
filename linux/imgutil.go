package main

import (
	"C"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)
import "fmt"

//export ImgutilGetImageSize
func ImgutilGetImageSize(path *C.char, w *uint, h *uint) *C.char {
	file, err := os.Open(C.GoString(path))
	if err != nil {
		return C.CString(err.Error())
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return C.CString(err.Error())
	}

	rect := img.Bounds()
	*w = uint(rect.Dx())
	*h = uint(rect.Dy())
	path = C.CString("teste")
	fmt.Println(*w, *h)
	fmt.Println(C.GoString(path))
	*path = *path
	fmt.Println()
	fmt.Println()

	return nil
}

func main() {}
