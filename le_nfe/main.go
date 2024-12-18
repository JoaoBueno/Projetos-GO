package main

import (
	"fmt"
	"os"

	"github.com/clbanning/mxj/v2"
)

func main() {
	// Caminho para o arquivo XML
	filePath := "53240624907602000195550010007207531338721034-procNFe.xml"

	// Ler o conteúdo do arquivo XML
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo: %v\n", err)
		os.Exit(1)
	}

	// Remover a tag Signature
	// re := regexp.MustCompile(`(?s)<Signature.*?</Signature>`)
	// dataCleaned := re.ReplaceAll(data, []byte(""))
	// // fmt.Println(string(dataCleaned))

	// Carregar o XML em um mapa
	mv, err := mxj.NewMapXml(data)
	if err != nil {
		fmt.Printf("Erro ao decodificar XML: %v\n", err)
		os.Exit(1)
	}

	// Exibir o mapa completo
	// fmt.Println(mv)

	// Determinar a base do caminho de acordo com a presença de nfeProc
	basePath := "NFe"
	if _, err := mv.ValueForPath("nfeProc"); err == nil {
		basePath = "nfeProc.NFe"
	}

	// Acessar elementos dinamicamente usando a basePath
	cUF, _ := mv.ValueForPathString(fmt.Sprintf("%s.infNFe.ide.cUF", basePath))
	cNF, _ := mv.ValueForPathString(fmt.Sprintf("%s.infNFe.ide.cNF", basePath))
	emitCNPJ, _ := mv.ValueForPathString(fmt.Sprintf("%s.infNFe.emit.CNPJ", basePath))
	emitXNome, _ := mv.ValueForPathString(fmt.Sprintf("%s.infNFe.emit.xNome", basePath))
	destCNPJ, _ := mv.ValueForPathString(fmt.Sprintf("%s.infNFe.dest.CNPJ", basePath))
	destXNome, _ := mv.ValueForPathString(fmt.Sprintf("%s.infNFe.dest.xNome", basePath))

	// Acessar atributos de infNFe
	infNFeId, _ := mv.ValueForPathString(fmt.Sprintf("%s.infNFe.-Id", basePath))
	infNFeVersao, _ := mv.ValueForPathString(fmt.Sprintf("%s.infNFe.-versao", basePath))

	// Exibir os valores
	fmt.Printf("Código UF: %s\n", cUF)
	fmt.Printf("Código NF: %s\n", cNF)
	fmt.Printf("Emitente CNPJ: %s\n", emitCNPJ)
	fmt.Printf("Emitente Nome: %s\n", emitXNome)
	fmt.Printf("Destinatário CNPJ: %s\n", destCNPJ)
	fmt.Printf("Destinatário Nome: %s\n", destXNome)
	fmt.Printf("infNFe Id: %s\n", infNFeId)
	fmt.Printf("infNFe Versão: %s\n", infNFeVersao)

	// Acessar itens (det)
	items, err := mv.ValuesForPath(fmt.Sprintf("%s.infNFe.det", basePath))
	if err != nil {
		fmt.Printf("Erro ao acessar itens: %v\n", err)
		os.Exit(1)
	}

	// Percorrer e exibir detalhes dos itens
	for i, item := range items {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			fmt.Printf("Item %d não é um mapa válido\n", i)
			continue
		}

		prod, ok := itemMap["prod"].(map[string]interface{})
		if !ok {
			fmt.Printf("Produto no item %d não é um mapa válido\n", i)
			continue
		}

		cProd := prod["cProd"]
		xProd := prod["xProd"]
		qCom := prod["qCom"]
		vUnCom := prod["vUnCom"]
		vProd := prod["vProd"]

		fmt.Printf("Item %d:\n", i+1)
		fmt.Printf("  Código do Produto: %v\n", cProd)
		fmt.Printf("  Descrição do Produto: %v\n", xProd)
		fmt.Printf("  Quantidade Comercial: %v\n", qCom)
		fmt.Printf("  Valor Unitário Comercial: %v\n", vUnCom)
		fmt.Printf("  Valor Total do Produto: %v\n", vProd)
	}
}
