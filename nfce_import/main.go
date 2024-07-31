package main

import (
	"fmt"
	"nfceimport/models"
	service "nfceimport/services"
	"nfceimport/util"
)

func main() {

	loja := service.GetLoja()

	fmt.Println(loja.LojCnpj)

	service.ReadFileFromDjm()

	//saveDOC()

	//openFile()

}

func saveDOC() {
	/// Retorna todas as linhas do arquivo em um objeto
	registros := readRegistros()

	for _, doc := range registros {
		/// Obtem o tipo de objeto da interface
		switch d := doc.(type) {
		/// Obtem o tipo DOC
		case models.RegistroDoc2:
			if d.Doc["DENOMINACAO"] == "RV" {
				fmt.Println(d.Doc["CHAVE_DFE"])
			}

		}
	}
}

func readRegistros() []interface{} {
	registros := util.GetLinesFromFile("C:/Via/pdv/vendas/00299667.djm")
	return registros
}

func openFile() {
	parametros := util.GetLinesFromFile("C:/Via/pdv/vendas/00299667.djm")
	for _, parametro := range parametros {
		switch p := parametro.(type) {
		case []models.RegistroINI:

			for _, registro := range p {
				fmt.Printf("%s:  %s\n", registro.Campo, registro.Value)
			}

		case []models.RegistroMON:

			for _, registro := range p {
				fmt.Printf("%s:  %s\n", registro.Campo, registro.Value)
			}
		case []models.RegistroDXL:

			fmt.Println("Registro PDV - Informações sobre o Terminal")
			fmt.Println("---------------------------------------")
			fmt.Println("SEQ | CAMPO | DESCRIÇÃO | TIPO | BYTES | OBSERVAÇÕES | VALUE")
			fmt.Println("---------------------------------------")
			for _, r := range p {
				fmt.Printf(
					"%s | %s\n",
					r.Campo,
					r.Value,
				)
			}
			fmt.Println("---------------------------------------")
		default:
		}
	}
}
