package main

/*
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)


// CopyStringToCArray copia um string Go para uma array de caracteres C
// e em seguida copia a array para o buffer de destino.
// Retorna o tamanho real copiado.
func CopyStringToCArray(src string, dest *C.char, maxLen int) int {
	cs := C.CString(src)
	defer C.free(unsafe.Pointer(cs)) // Lembre-se de liberar o CString após usá-lo

	// Copie o string Go para uma array temporária de caracteres C.
	temp := make([]C.char, maxLen)
	C.strncpy((*C.char)(unsafe.Pointer(&temp[0])), cs, C.size_t(len(temp)))

	// Copie a array temporária para o buffer de destino.
	C.strncpy(dest, (*C.char)(unsafe.Pointer(&temp[0])), C.size_t(len(temp)))

	return int(C.strlen(dest))
}
