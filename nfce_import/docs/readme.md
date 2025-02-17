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

## blackbox.ia
### Prompt Converter imagem em tabela
- leia o arquivo e crie uma tabela de todos os campos.

- crie uma estrutura em golang para gravar todos os dados da tabela, crie uma função com retorno, os dados devem ser preenchido com os valores da tabela.

- com a tabela criada, crie um arquivo em golang da google para guardar todas as colunas


- crie uma tabela do arquivo, enviar a imagem
- com a tabela gerada, crie um arquivo em golang da google para guardar todas as colunas, o retorno tem que ser um conjuto json de chave e valor de acordo com as colunas da tabela

- crie uma tabela do arquivo
- com a tabela gerada, crie um arquivo em golang da google para guardar todas as colunas e linhas, o retorno deve ser um slice contendo as chaves e valores de acordo com as colunas e linhas da tabela gerada.

- o GetEcf deve conter no slice chave e valor de acordo com a tabela da imagem.

- Crie uma chave no struct chamada Value

- Complete o GetDfe, fazer a primeira linha de exmplo

### Prompt gerar os objetos struct

- reescreva o retorno RegistroDoc para um map contendo a chave com o Campo e o valor com o Valuetransforme o retorno RegistroDoc em um objeto map contendo a chave com o Campo e o valor com o Value

- crie um struct usando a função GetEcf, usando os valores de Campo
- substitua a função GetEcf para returnar um struct do tipo ECF

### Consulta no banco de dados
- Erro: Query é utilizado para executar consultas SQL que retornam múltiplos resultados, enquanto o método Scan é utilizado para ler os resultados de uma consulta que retorna apenas um resultado. No entanto, no seu código, você está utilizando o método Query para executar a consulta e, em seguida, está tentando ler os resultados com o método Scan.

- Solução: Para resolver esse problema, você pode utilizar o método QueryRow em vez de Query. O método QueryRow é utilizado para executar consultas SQL que retornam apenas um resultado e é compatível com o método Scan.

### Ciclo de inicialização do GO
- O init() é uma função especial em Go que é executada automaticamente quando o pacote é carregado. Isso significa que o init() é executado antes de qualquer outra função do pacote ser chamada.

No caso do seu código, o init() é executado quando o pacote db é carregado, o que ocorre quando o programa Go é iniciado.

Aqui está a ordem em que as coisas acontecem:

O programa Go é iniciado.
O pacote db é carregado.
O init() do pacote db é executado.
O init() estabelece a conexão com o banco de dados.
O programa Go continua a executar, e as funções do pacote db podem ser chamadas.
É importante notar que o init() é executado apenas uma vez, quando o pacote é carregado. Se você tiver várias funções init() em diferentes pacotes, elas serão executadas em ordem, começando com o pacote principal.