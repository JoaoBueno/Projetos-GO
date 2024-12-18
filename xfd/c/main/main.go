package main

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
	"time"
)

var db *sql.DB
var err error

func main() {
	cnpj := os.Args[1]

	diaHoraAtual := time.Now()
	dataAtual := diaHoraAtual.Format("2006-01-02")
	dataData, _ := time.Parse("2006-01-02", dataAtual)

	if db == nil {
		db, err = bd.Connect()
		if err != nil {
			db = nil
			fmt.Println("RETORNAR PARA O COBOL - ❌ Erro ao conectar no banco de dados: " + err.Error())
			return
		}
	}

	fmt.Println("Conectado ao banco de dados")

	dadosEmpresa, err := bd.Buscar(db, cnpj)

	if err != nil {
		fmt.Println("RETORNAR PARA O COBOL - ❌ Erro na consulta no banco de dados: ", err.Error())
		return
	}

	if dadosEmpresa.Cnpj == "" {
		// Caso 1: CNPJ não encontrado na nossa base de dados. Função: buscar e inserir
		consultaSerpro, err := apiSerpro.BuscarCnpjSerpro(cnpj)

		if err != nil {
			fmt.Println("RETORNAR PARA O COBOL - ❌ Erro ao buscar CNPJ no SERPRO:", err.Error())
			return
		}

		md5Api := ""
		tamanho := 0
		md5.MD5String(consultaSerpro, &md5Api, &tamanho)

		bd.Inserir(db, bd.RegistroSerpro{Cnpj: cnpj, Data: dataData, Md5: md5Api, Dados: consultaSerpro})

		if err != nil {
			// Devolver o dado apesar da falha no insert
			fmt.Println("RETORNAR PARA O COBOL - ❌ Erro ao inserir os dados da empresa no banco de dados:", err.Error())
			return
		}

		fmt.Println("RETORNAR DADOS DA EMPRESA PARA O COBOL - ✅ CNPJ encontrado e inserido na base de dados")

	} else {
		diasAtualizacao, _ := strconv.Atoi(os.Getenv("DIAS_ATUALIZACAO"))
		diasAtualizacaoDuracao := time.Duration(diasAtualizacao) * 24 * time.Hour

		var empresa tipos.Empresa
		var dadosConcatenados = concatenaDados(dadosEmpresa.Dados)
		fmt.Println(dadosConcatenados)
		err = json.Unmarshal([]byte(dadosEmpresa.Dados), &empresa)

		if err != nil {
			fmt.Println("Erro na decodificação dos dados da empresa: ", err.Error())
			return
		}

		err = bd.InserirCnae(db, bd.Cnae{Id: empresa.CnaePrincipal.Id, Descricao: empresa.CnaePrincipal.Descricao})

		if err != nil {
			fmt.Println("Erro na inserção da CNAE: ", err.Error())
			return
		}

		var cnaes string
		cnaes = empresa.CnaePrincipal.Id

		for _, cnae := range empresa.CnaeSecundarias {
			bd.InserirCnae(db, bd.Cnae{Id: cnae.Id, Descricao: cnae.Descricao})

			if err != nil {
				fmt.Println("Erro na inserção da CNAE: ", err.Error())
			}
			cnaes = cnaes + cnae.Id
		}

		duracao := time.Since(dadosEmpresa.Data)
		if duracao > diasAtualizacaoDuracao {
			// Caso 2: CNPJ encontrado na base de dados, mas desatualizado. Função: buscar e atualizar
			fmt.Println("Última atualização do dado superior a " + strconv.Itoa(diasAtualizacao) + " dias. Chamando API SERPRO.")

			consultaSerpro, err := apiSerpro.BuscarCnpjSerpro(cnpj)

			if err != nil {
				fmt.Println("RETORNAR PARA O COBOL - ❌ Erro ao buscar CNPJ no SERPRO:", err.Error())
				return
			}

			md5Api := ""
			tamanho := 0
			md5.MD5String(consultaSerpro, &md5Api, &tamanho)

			dadoAtual, err := bd.Buscar(db, cnpj)

			if dadoAtual.Md5 == md5Api {
				// CNPJ encontrado (já atualizado na base de dados)
				err = bd.AtualizarData(db, bd.RegistroSerpro{Cnpj: cnpj, Data: dataData})
			} else {
				// CNPJ encontrado (precisou ser atualizado na base de dados)
				err = bd.Atualizar(db, bd.RegistroSerpro{Cnpj: cnpj, Data: dataData, Md5: md5Api, Dados: consultaSerpro})
			}

			if err != nil {
				fmt.Println("RETORNAR PARA O COBOL - ❌ Erro ao atualizar os dados da empresa no banco de dados:", err.Error())
				return
			}

			fmt.Println("RETORNAR PARA O COBOL - ✅ CNPJ encontrado")
		} else {
			// Caso 3: CNPJ encontrado na base de dados e atualizado. Função: retornar
			fmt.Println("RETORNAR PARA O COBOL - ✅ CNPJ encontrado")
		}
	}

	defer db.Close()
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
