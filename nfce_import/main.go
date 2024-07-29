package main

import (
	"fmt"

	logger "nfceimport/log"
	"nfceimport/model"

	services "nfceimport/service"
)

func main() {

	teste := services.GetLoja()

	fmt.Println(teste[0].LojCnpj)
	logInstance, err := logger.NewZapLogger()
	logInstanceLogrus := logger.NewLogger()
	if err != nil {
		logInstance.Fatalf("Não foi possível inicializar o logger: %v", err)
	}

	logInstance.Infof("Aplicação iniciada")
	logInstanceLogrus.Fatalf("dkfjdaskfjskfjds")

	//OpenFile()

}

func OpenFile() {
	parametros := getLinesFromFile("C:/Via/pdv/vendas/00301645.djm")
	for _, parametro := range parametros {
		switch p := parametro.(type) {
		case []model.RegistroINI:

			for _, registro := range p {
				fmt.Printf("%s:  %s\n", registro.Campo, registro.Value)
			}

		case []model.RegistroMON:

			for _, registro := range p {
				fmt.Printf("%s:  %s\n", registro.Nome, registro.Value)
			}
		case []model.RegistroDXL:

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
