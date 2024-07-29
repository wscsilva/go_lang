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
	for _, parametro := range parametros {
		switch p := parametro.(type) {
		case []model.RegistroINI:
			//fmt.Println(p[0].Value)
			for _, registro := range p {
				fmt.Printf("%s:  %s\n", registro.Campo, registro.Value)
			}
			// acesso aos campos específicos do RegistroINI
		case []model.RegistroMON:
			// acesso aos campos específicos do RegistroMON
			for _, registro := range p {
				fmt.Printf("%s:  %s\n", registro.Nome, registro.Value)
			}
		case []model.RegistroDXL:
			// acesso aos campos específicos do RegistroMON
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

/*
	fmt.Println("Registro PDV - Informações sobre o Terminal")
	fmt.Println("---------------------------------------")
	fmt.Println("SEQ | CAMPO | DESCRIÇÃO | TIPO | BYTES | OBSERVAÇÕES")
	fmt.Println("---------------------------------------")
	for _, r := range registroPDV {
		fmt.Printf("%02d | %s | %s | %s | %d | %s\n", r.Seq, r.Campo, r.Descricao, r.Tipo, r.Bytes, r.Observacoes)
	}
	fmt.Println("---------------------------------------")
*/
