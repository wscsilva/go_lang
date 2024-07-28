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
