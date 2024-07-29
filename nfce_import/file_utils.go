package main

import (
	"bufio"

	"log"
	"nfceimport/model"
	"os"
	"strings"
)

func getLinesFromFile(path string) []interface{} {
	// Abrir o arquivo CSV
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
		return nil
	}
	// Garantir que o arquivo será fechado ao final da função
	defer file.Close()

	// Ler o arquivo linha por linha
	scanner := bufio.NewScanner(file)

	/* 	var parametrosIni []model.RegistroINI
	   	var parametrosMon []model.RegistroMON */
	var parametros []interface{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "|")
		switch fields[0] {
		case "INI":
			parametros = append(parametros, model.GetIni(fields))
		case "MON":
			parametros = append(parametros, model.GetMon(fields))
		default:
		}
		/* 		if len(fields) > 0 {
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
		} */
	}

	// Verificar se houve algum erro na leitura do arquivo
	if err := scanner.Err(); err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	return parametros
}
