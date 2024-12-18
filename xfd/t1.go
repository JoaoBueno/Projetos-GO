package main

import (
	/*
		#include "memcpy.h"
	*/
	"C"
	"fmt"
	"unsafe"
)

/*T1 is a func to convert XFD to JSON*/
//export T1
func T1(xfdjson **C.char, length *int) {
	fmt.Println(xfdjson)
	var reta string
	reta = "fmt.Sprint(string(j))"
	retaCStr := C.CString(reta)
	xfdjson = &retaCStr
	*length = len(reta)
}

/*T2 is a func to convert XFD to JSON*/
//export T2
func T2(xfdjson *C.char, length *int) {
	fmt.Println(xfdjson)
	var reta string
	reta = "fmt.Sprint(string(j))"
	*length = len(reta)
	xfdjson = C.CString(reta)
}

/*T3 is a func to convert XFD to JSON*/
//export T3
func T3(xfdjson **C.char, length *int) {
	fmt.Println(xfdjson)
	var reta string
	reta = "fmt.Sprint(string(j))"
	*xfdjson = C.CString(reta)
	*length = len(reta)
}

/*T4 is a func to convert XFD to JSON*/
//export T4
func T4(xfdjson **C.char, length *int) {
	fmt.Println(xfdjson)
	var reta string
	reta = "fmt.Sprint(string(j))"
	retaCStr := C.CString(reta)
	C.memcpy(unsafe.Pointer(*xfdjson), unsafe.Pointer(retaCStr), C.size_t(len(reta)))
	// *xfdjson = C.CString(reta)
	*length = len(reta)
}

func main() {}
