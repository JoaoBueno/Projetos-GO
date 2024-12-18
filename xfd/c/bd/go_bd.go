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
	"database/sql"
	"fmt"
	"os"
	"bd/bd"
	"unsafe"
	"time"
	"log"

	"github.com/joho/godotenv"
)

type MyError struct {
	Code    int
	Message string
}

// Error makes MyError implement the error interface.
func (e *MyError) Error() string {
    return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

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

func formatarDados(fco, nome, jurfis, cnpj, cidade, uf string) string {
	// Formatando cada campo para ter o tamanho desejado
	// %014d - formata um número inteiro para ter 14 dígitos, preenchendo com zeros à esquerda
	// %-55s - formata uma string para ter 55 caracteres, alinhando à esquerda
	// %-50s - formata uma string para ter 50 caracteres, alinhando à esquerda
	// %-2s  - formata uma string para ter 2 caracteres, alinhando à esquerda
	fcoFormatado := fmt.Sprintf("%014s", fco)
	nomeFormatado := fmt.Sprintf("%-55s", nome)
	jurfisFormatado := fmt.Sprintf("%-1s", jurfis)
	cnpjFormatado := fmt.Sprintf("%014s", cnpj)
	cidadeFormatada := fmt.Sprintf("%-50s", cidade)
	ufFormatada := fmt.Sprintf("%-2s", uf)

	// Concatenando todos os campos formatados
	dadosConcatenados := fcoFormatado + nomeFormatado + jurfisFormatado + cnpjFormatado + cidadeFormatada + ufFormatada

	return dadosConcatenados
}

var db *sql.DB
var rows *sql.Rows
var err error

//export fco_bd
func fco_bd(c_cnpj *C.char, c_retorno *C.char, ret *int64) int64 {
	godotenv.Load(".env")

	retorno := ""

	if db == nil {
		db, err = bd.Connect()
		if err != nil {
			db = nil
			if os.Getenv("DEBUG") == "true" {
				fmt.Println("RETORNAR PARA O COBOL - ❌ Erro ao conectar no banco de dados: " + err.Error())
			}
			retorno = err.Error()
			*ret = copyStringToCArray(retorno, c_retorno, int64(len(retorno)))
			*ret = -1001
			return 0
		}
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Println("Conectado ao banco de dados")
		time.Sleep(3 * time.Second)
	}

	rows, err = bd.Buscar(db)

	if err != nil {
		myErr, _ := err.(*MyError)
		if os.Getenv("DEBUG") == "true" {
			fmt.Println("RETORNAR PARA O COBOL - ❌ Erro na consulta no banco de dados: ", myErr.Code, myErr.Message)
		}
		retorno = myErr.Message
		*ret = copyStringToCArray(retorno, c_retorno, int64(len(retorno)))
		*ret = int64(myErr.Code)
		return 0
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Println("Lido")
		// time.Sleep(3 * time.Second)
	}

	// Iterando sobre os resultados
	for rows.Next() {
		var fco, nome, jurfis, cnpj, cidade, uf string
		if err := rows.Scan(&fco, &nome, &jurfis, &cnpj, &cidade, &uf); err != nil {
			log.Fatal(err)
		}
		
		if os.Getenv("DEBUG") == "true" {
			fmt.Println("fco => ", fco)
			fmt.Println("nome => ", nome)
			fmt.Println("jurfis => ", jurfis)
			fmt.Println("cnpj => ", cnpj)
			fmt.Println("cidade => ", cidade)
			fmt.Println("uf => ", uf)
			time.Sleep(3 * time.Second)
		}

		retorno = formatarDados(fco, nome, jurfis, cnpj, cidade, uf)
		*ret = copyStringToCArray(retorno, c_retorno, int64(len(retorno)))
		return int64(*ret * -1)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}	
	return int64(*ret * -1)
}

//export fco_next
func fco_next(c_cnpj *C.char, c_retorno *C.char, ret *int64) int64 {
	retorno := ""

	for rows.Next() {
		var fco, nome, jurfis, cnpj, cidade, uf string
		if err := rows.Scan(&fco, &nome, &jurfis, &cnpj, &cidade, &uf); err != nil {
			log.Fatal(err)
		}
		// Faça algo com os dados aqui
		if os.Getenv("DEBUG") == "true" {
			fmt.Println("fco => ", fco)
			fmt.Println("nome => ", nome)
			fmt.Println("jurfis => ", jurfis)
			fmt.Println("cnpj => ", cnpj)
			fmt.Println("cidade => ", cidade)
			fmt.Println("uf => ", uf)
			time.Sleep(3 * time.Second)
		}

		retorno = formatarDados(fco, nome, jurfis, cnpj, cidade, uf)
		*ret = copyStringToCArray(retorno, c_retorno, int64(len(retorno)))
		return int64(0)
	}

	retorno = "<<<fim>>>"
	*ret = copyStringToCArray(retorno, c_retorno, int64(len(retorno)))

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}	
	return int64(0)
}

//export closeRows
func closeRows() {
    if rows != nil {
        rows.Close()
    }
	if db != nil {
		db.Close()
	}
}

func main() {}
