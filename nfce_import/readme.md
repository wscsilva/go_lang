# Anotações sobre Go lang
=====================================

### Compilando o Programa
-------------------------

Para compilar o programa, utilize o comando:
```go
go build -o nfce_import.exe
```
### Execute o programa:
```shell
.\nfce_import.exe
```

### Script de Lote (.bat):
```batch
@echo off
cd C:\DevPrograms\go-projetos\nfce_import
start nfce_import.exe
```

### Packages
- Go exige que todos os arquivos em um diretório pertençam ao mesmo pacote. No seu caso, parece que você tem arquivos main.go e bdados.go no mesmo diretório, mas com pacotes diferentes (main e db).

### defer 
```go 
defer file.Close()
```
- A palavra-chave defer em Go é usada para adiar a execução de uma função até que a função em que ela foi chamada termine. A sintaxe defer file.Close() é comumente usada para garantir que os recursos sejam liberados de forma adequada, independentemente de como a função termine, seja por uma execução normal ou por um erro.