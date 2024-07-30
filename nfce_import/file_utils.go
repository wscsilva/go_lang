package main

import (
	"bufio"
	"nfceimport/model"

	"log"

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

	var parametros []interface{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "|")
		switch fields[0] {
		case "INI":
			parametros = append(parametros, model.GetIni(fields))
		case "MON":
			parametros = append(parametros, model.GetMon(fields))
		case "ECF":
			parametros = append(parametros, model.GetEcf(fields))
		case "PDV":
			parametros = append(parametros, model.GetPdv(fields))
		case "DOC":
			parametros = append(parametros, model.GetDoc(fields))
		case "DFE":
			parametros = append(parametros, model.GetDfe(fields))
		case "DIT":
			parametros = append(parametros, model.GetDit(fields))
		case "DXL":
			parametros = append(parametros, model.GetDxl(fields))
		case "FIM":
			parametros = append(parametros, model.GetFim(fields))
		default:
		}

	}

	// Verificar se houve algum erro na leitura do arquivo
	if err := scanner.Err(); err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	return parametros
}
