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
func copyStringToCArray(src string, dest *C.char, maxLen int) int {
	cs := C.CString(src)
	defer C.free(unsafe.Pointer(cs)) // Lembre-se de liberar o CString após usá-lo

	// Copie o string Go para uma array temporária de caracteres C.
	temp := make([]C.char, maxLen)
	C.strncpy((*C.char)(unsafe.Pointer(&temp[0])), cs, C.size_t(len(temp)))

	// Copie a array temporária para o buffer de destino.
	C.strncpy(dest, (*C.char)(unsafe.Pointer(&temp[0])), C.size_t(len(temp)))

	return int(C.strlen(dest))
}

//export retornamd5
func retornamd5(s *C.char, retorno *C.char, len1 *int) {
	fmt.Printf("Chegou = %s\n", C.GoString(s))
	fmt.Printf("Tamanho = %d\n", *len1)
	ms := C.GoString(s)

//	destSize := 50

	// Usar a função helper para fazer a cópia e obter o tamanho real copiado
	*len1 = copyStringToCArray(ms, retorno, *len1)

	fmt.Printf("Tamanho = %d\n", *len1)

	return
}


// //export retornamd5
// func retornamd5(s *C.char, retorno *C.char, len1 *int) {
// 	fmt.Printf("Chegou = %s\n", C.GoString(s))
// 	fmt.Printf("Tamanho = %d\n", len1)

// 	dest := make([]C.char, 50)
// 	// str1 := C.CString("quick")
// 	// str2 := C.CString("brown")
// 	// str3 := C.CString("lazy")
// 	// C.snprintf(
// 	// 	(*C.char)(unsafe.Pointer(&dest[0])),
// 	// 	C.size_t(len1(dest)),
// 	// 	C.CString("Isto %s é um %s teste %s"),
// 	// 	str1,
// 	// 	str2,
// 	// 	str3,
// 	// )
// 	C.strcpy((*C.char)(unsafe.Pointer(&dest[0])), C.CString("tua porra de saco de merda"))

// 	C.strncpy(retorno, (*C.char)(unsafe.Pointer(&dest[0])), C.size_t(len(dest)))

// 	// *len1 = len(retorno)
// 	*len1 = int(C.strlen(retorno))

// 	return
// }

// // func retornamd5(s *C.char, retorno *C.char, len1 *C.int) C.int {
// // 	fmt.Printf("Chegou = %s\n", C.GoString(s))
// // 	fmt.Printf("Tamanho = %d\n", len1)

// // 	// C.snprintf(retorno, C.size_t(50), C.CString("tua porra de saco de merda"))
// // 	C.snprintf(retorno, C.size_t(50), (*C.char)(C.CString("tua porra de saco de merda")))
// // 	*len1 = C.int(C.strlen(retorno))

// // 	return C.int(333)
// // }

func main() {}
