package main

// /*
// #include "t2c-h.h"
// */

/*
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)


// copyStringToCArray copia um string Go para uma array de caracteres C
// e em seguida copia a array para o buffer de destino.
// Retorna o tamanho real copiado.
func copyStringToCArray(src string, dest *C.char, maxLen int64) int64 {
	cs := C.CString(src)
	defer C.free(unsafe.Pointer(cs)) // Lembre-se de liberar o CString após usá-lo

	// Copie o string Go para uma array temporária de caracteres C.
	temp := make([]C.char, maxLen)
	C.strncpy((*C.char)(unsafe.Pointer(&temp[0])), cs, C.size_t(len(temp)))

	// Copie a array temporária para o buffer de destino.
	C.strncpy(dest, (*C.char)(unsafe.Pointer(&temp[0])), C.size_t(len(temp)))

	return int64(C.strlen(dest))
}

//export retornamd5
func retornamd5(s *C.char, retorno *C.char, len1 *int64) *C.char {
	fmt.Printf("Chegou = %s\n", C.GoString(s))
	fmt.Printf("Tamanho = %d\n", *len1)
	ms := C.GoString(s)

	// Usar a função helper para fazer a cópia e obter o tamanho real copiado
	*len1 = copyStringToCArray(ms, retorno, *len1)

	fmt.Printf("Tamanho = %d\n", *len1)

	return retorno
}

func main() {}
