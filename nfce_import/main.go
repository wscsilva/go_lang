package main

import (
	"fmt"
	"log"
	"nfceimport/model"

	"nfceimport/db"
)

func main() {
	// Abrir a conexão com o banco de dados
	conexao, err := db.ConectarBancoDeDados()
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão: %v", err)
	}
	defer db.FecharConexao(conexao)

	parametros := getLinesFromFile("C:/Via/pdv/vendas/00301645.djm")
	//fmt.Println(parametros[0])
	for _, parametro := range parametros {
		//fmt.Println(parametro)
		switch p := parametro.(type) {
		case []model.RegistroINI:
			fmt.Println(p[0].Value)
			/* 			for _, registro := range p {
				fmt.Printf("%s:  %s\n", registro.Campo, registro.Value)
			} */
			// acesso aos campos específicos do RegistroINI
		case []model.RegistroMON:
			// acesso aos campos específicos do RegistroMON
			//fmt.Println(p[0].Value)
			for _, registro := range p {
				fmt.Printf("%s:  %s\n", registro.Nome, registro.Value)
			}

		default:
		}
	}

}
