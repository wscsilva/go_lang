package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	"nfceimport/db"
)

type Parametros struct {
	Parametro string
	Dados     map[string]string
}

func main() {
	// Abrir a conexão com o banco de dados
	conexao, err := db.ConectarBancoDeDados()
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão: %v", err)
	}
	defer db.FecharConexao(conexao)

	file, err := lerArquivo("C:/Via/pdv/vendas/00301645.djm")

	// Abrir o arquivo CSV
	/* 	file, err := os.Open("C:/Via/pdv/vendas/00301645.djm")
	   	if err != nil {
	   		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	   	}
	   	defer file.Close()
	*/
	// Ler o arquivo linha por linha
	scanner := bufio.NewScanner(file)

	var parametros []Parametros

	linha := 1
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "|")
		if len(fields) > 0 {
			parametro := Parametros{
				Parametro: fields[0],
				Dados:     make(map[string]string),
			}
			for i, field := range fields[1:] {
				if field != "" {
					parametro.Dados[fmt.Sprintf("campo_%d", i+1)] = field
				}
			}
			parametros = append(parametros, parametro)
		}
		//fmt.Printf("Linha %d: %s\n", linha, line)
		linha++
	}

	//jsonData, err := json.Marshal(parametros)
	if err != nil {
		log.Fatalf("Erro ao criar JSON: %v", err)
	}
	// fmt.Println(string(jsonData))
	fmt.Println(parametros[1].Dados)
	fmt.Println(parametros[1].Dados["campo_1"])
}
