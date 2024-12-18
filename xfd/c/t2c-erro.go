package main

// /*
// #include "t2c-h.h"
// */

/*
#include <stdio.h>
#include <string.h>
*/
import "C"

import (
	"fmt"
)

//export retornamd5
func retornamd5(s *C.char, retorno *C.char, maxLen *int) {
	fmt.Printf("Chegou = %s\n", C.GoString(s))
	fmt.Printf("Tamanho = %d\n", *maxLen)
	ms := C.GoString(s)

	// Usar a função helper para fazer a cópia e obter o tamanho real copiado
	*maxLen = CopyStringToCArray(ms, *retorno, *maxLen)

	fmt.Printf("Tamanho = %d\n", *maxLen)

	return
}

func main() {}
