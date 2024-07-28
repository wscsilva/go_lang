package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func lerArquivo(path string) (*os.File, error) {
	// Abrir o arquivo CSV
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
		return nil, err
	}
	defer file.Close()
	return file, nil
}

// Move o arquivo processado para o diretório de processados
func moveFileToProcessedDir(sourcePath string) {
	processedDir := "C:/Via/pdv/vendas/"
	destinationPath := filepath.Join(processedDir, filepath.Base(sourcePath))

	if err := copyFile(sourcePath, destinationPath); err != nil {
		log.Printf("Erro ao mover o arquivo: %v\n", err)
	} else {
		fmt.Printf("Arquivo movido para: %s\n", destinationPath)
		if err := os.Remove(sourcePath); err != nil {
			log.Printf("Erro ao remover o arquivo original: %v\n", err)
		}
	}
}

// Copia um arquivo de um local para outro
func copyFile(sourcePath, destinationPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	return err
}
