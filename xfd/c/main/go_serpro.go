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
	"encoding/json"
	"fmt"
	"os"
	"serpro/apiSerpro"
	"serpro/bd"
	"serpro/md5"
	"serpro/tipos"
	"strconv"
	"strings"
	"time"
	"unsafe"
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

	fmt.Printf("Chegou = %s\n", C.GoString(cs))

	// Copie o string Go para uma array temporária de caracteres C.
	temp := make([]C.char, maxLen)
	fmt.Printf("Tamanho = %d\n", len(temp))
	C.strncpy((*C.char)(unsafe.Pointer(&temp[0])), cs, C.size_t(len(temp)))
	fmt.Printf("Tamanho = %d\n", len(temp))

	// Copie a array temporária para o buffer de destino.
	C.strncpy(dest, (*C.char)(unsafe.Pointer(&temp[0])), C.size_t(len(temp)))
	fmt.Printf(	"Chegou = %s\n", C.GoString(dest))

	return int64(C.strlen(dest))
}

var db *sql.DB
var err error

func getTipoPorte(porte string) string {
	switch porte {
	case "01":
		return "ME"
	case "03":
		return "PPA"
	default:
		return "MED"
	}
}

func concatenaDados(dadosEmpresa interface{}) string {
	var empresa tipos.Empresa

	err := json.Unmarshal([]byte(dadosEmpresa.(string)), &empresa)
	if err != nil {
		fmt.Println("Erro")
		return ""
	}

	var dataAbertura = empresa.DataAbertura[0:4] + empresa.DataAbertura[5:7] + empresa.DataAbertura[8:10]
	var dadosConcatenados = empresa.NomeEmpresarial + "|" +
		empresa.NomeFantasia + "|" +
		dataAbertura + "|" +
		empresa.CorreioEletronico + "|" +
		getTipoPorte(empresa.Porte) + "|" +
		empresa.Endereco.TipoLogradouro + "|" +
		empresa.Endereco.Logradouro + "|" +
		empresa.Endereco.Bairro + "|" +
		empresa.Endereco.Municipio.Descricao + "|" +
		empresa.Endereco.UF + "|" +
		empresa.Endereco.CEP

	for i := 0; i < 4; i++ {
		dadosConcatenados += "|"
		if i < len(empresa.Telefones) {
			dadosConcatenados += fmt.Sprintf("%s%s", empresa.Telefones[i].DDD, empresa.Telefones[i].Numero)
		}
	}

	dadosConcatenados += "|" + empresa.CnaePrincipal.Id

	for i := 0; i < len(empresa.CnaeSecundarias); i++ {
		dadosConcatenados += fmt.Sprintf("%s", empresa.CnaeSecundarias[i].Id)
	}

	if len(dadosConcatenados) > 0 && dadosConcatenados[len(dadosConcatenados)-1] == '|' {
		dadosConcatenados = dadosConcatenados[:len(dadosConcatenados)-1]
	}

	return dadosConcatenados
}

//export serpro
func serpro(c_cnpj *C.char, c_retorno *C.char, ret *int64) int64 {
	cnpj := C.GoString(c_cnpj)
	cnpj = strings.TrimSpace(cnpj)

	// fmt.Println("CNPJ: " + cnpj + " => " + strconv.Itoa(len(cnpj)))

	// if os.Getenv("DEBUG") == "true" {
		fmt.Println("CNPJ: " + cnpj + " - "+ strconv.Itoa(len(cnpj)))
		fmt.Printf("Tamanho = %d\n", *ret)
		time.Sleep(3 * time.Second)
	// }
	
	retorno := ""

	diaHoraAtual := time.Now()
	dataAtual := diaHoraAtual.Format("2006-01-02")
	dataData, _ := time.Parse("2006-01-02", dataAtual)

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

	defer db.Close()

	if os.Getenv("DEBUG") == "true" {
		fmt.Println("Conectado ao banco de dados")
		// time.Sleep(3 * time.Second)
	}

	dadosEmpresa, err := bd.Buscar(db, cnpj)

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
		fmt.Println("Lido: " + dadosEmpresa.Cnpj)
		// time.Sleep(3 * time.Second)
	}

	// Caso não haja o CNPJ no banco de dados
	if dadosEmpresa.Cnpj == "" {
		// Caso 1: CNPJ não encontrado na nossa base de dados. Função: buscar e inserir
		consultaSerpro, err := apiSerpro.BuscarCnpjSerpro(cnpj)

		if err != nil {
			if os.Getenv("DEBUG") == "true" {
				fmt.Println(consultaSerpro + "RETORNAR PARA O COBOL - ❌ Erro ao buscar CNPJ no SERPRO:", err.Error())
			}
			return 0 
		}

		md5Api := ""
		tamanho := 0
		md5.MD5String(consultaSerpro, &md5Api, &tamanho)

		bd.Inserir(db, bd.RegistroSerpro{Cnpj: cnpj, Data: dataData, Md5: md5Api, Dados: consultaSerpro})

		if err != nil {
			// Devolver o dado apesar da falha no insert
			if os.Getenv("DEBUG") == "true" {
				fmt.Println("RETORNAR PARA O COBOL - ❌ Erro ao inserir os dados da empresa no banco de dados:", err.Error())
			}
			return 0
		}

		if os.Getenv("DEBUG") == "true" {
			fmt.Println("RETORNAR DADOS DA EMPRESA PARA O COBOL - ✅ CNPJ encontrado e inserido na base de dados")
		}

		dadosEmpresa, err = bd.Buscar(db, cnpj)

		if err != nil {
			myErr, _ := err.(*MyError)
			if os.Getenv("DEBUG") == "true" {
				fmt.Println("❌ Erro na consulta no banco de dados(inconsistente): ", myErr.Code, myErr.Message)
			}
			retorno = myErr.Message
			*ret = copyStringToCArray(retorno, c_retorno, int64(len(retorno)))
			*ret = int64(myErr.Code)
			return 0
		}

		if dadosEmpresa.Cnpj == "" {
			if os.Getenv("DEBUG") == "true" {
				fmt.Println("❌ Inconsistencia na leitura do banco de dados:")
			}
			retorno = "Inconsistencia na leitura do banco de dados"

// usar len(retorno) em todas as chamadas ao copyStringToCArray

			*ret = copyStringToCArray(retorno, c_retorno, int64(len(retorno)))
			*ret = -3001
			return 0
		}
	}
	
	var empresa tipos.Empresa
	var dadosConcatenados = concatenaDados(dadosEmpresa.Dados)
	if os.Getenv("DEBUG") == "true" {
		fmt.Println(dadosConcatenados)
	}
	err = json.Unmarshal([]byte(dadosEmpresa.Dados), &empresa)

	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			fmt.Println("Erro na decodificação dos dados da empresa: ", err.Error())
		}
		retorno = "Erro na decodificação dos dados da empresa: " + err.Error()
		*ret = copyStringToCArray(retorno, c_retorno, int64(len(retorno)))
		*ret = -3002
		return 0
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Println("Eba chegamos até aqui")
	}

	retorno = dadosConcatenados
	*ret = copyStringToCArray(retorno, c_retorno, int64(len(dadosConcatenados)))
	return int64(*ret * -1)

// retirado para futura implementações	
	// diasAtualizacao, _ := strconv.Atoi(os.Getenv("DIAS_ATUALIZACAO"))
	// diasAtualizacaoDuracao := time.Duration(diasAtualizacao) * 24 * time.Hour

	// err = bd.InserirCnae(db, bd.Cnae{Id: empresa.CnaePrincipal.Id, Descricao: empresa.CnaePrincipal.Descricao})

	// if err != nil {
	// 	fmt.Println("Erro na inserção da CNAE: ", err.Error())
	// 	return
	// }

	// var cnaes string
	// cnaes = empresa.CnaePrincipal.Id

	// for _, cnae := range empresa.CnaeSecundarias {
	// 	bd.InserirCnae(db, bd.Cnae{Id: cnae.Id, Descricao: cnae.Descricao})

	// 	if err != nil {
	// 		fmt.Println("Erro na inserção da CNAE: ", err.Error())
	// 	}
	// 	cnaes = cnaes + cnae.Id
	// }

	// duracao := time.Since(dadosEmpresa.Data)
	// if duracao > diasAtualizacaoDuracao {
	// 	// Caso 2: CNPJ encontrado na base de dados, mas desatualizado. Função: buscar e atualizar
	// 	// fmt.Println("Última atualização do dado superior a " + strconv.Itoa(diasAtualizacao) + " dias. Chamando API SERPRO.")

	// 	consultaSerpro, err := apiSerpro.BuscarCnpjSerpro(cnpj)

	// 	if err != nil {
	// 		fmt.Println("RETORNAR PARA O COBOL - ❌ Erro ao buscar CNPJ no SERPRO:", err.Error())
	// 		return
	// 	}

	// 	md5Api := ""
	// 	tamanho := 0
	// 	md5.MD5String(consultaSerpro, &md5Api, &tamanho)

	// 	dadoAtual, err := bd.Buscar(db, cnpj)

	// 	if dadoAtual.Md5 == md5Api {
	// 		// CNPJ encontrado (já atualizado na base de dados)
	// 		err = bd.AtualizarData(db, bd.RegistroSerpro{Cnpj: cnpj, Data: dataData})
	// 	} else {
	// 		// CNPJ encontrado (precisou ser atualizado na base de dados)
	// 		err = bd.Atualizar(db, bd.RegistroSerpro{Cnpj: cnpj, Data: dataData, Md5: md5Api, Dados: consultaSerpro})
	// 	}

	// 	if err != nil {
	// 		fmt.Println("RETORNAR PARA O COBOL - ❌ Erro ao atualizar os dados da empresa no banco de dados:", err.Error())
	// 		return
	// 	}

	// 	fmt.Println("RETORNAR PARA O COBOL - ✅ CNPJ encontrado")
	// } else {
	// 	// Caso 3: CNPJ encontrado na base de dados e atualizado. Função: retornar
	// 	fmt.Println("RETORNAR PARA O COBOL - ✅ CNPJ encontrado")
	// }
// retirado para futura implementações	

}

func main() {}
