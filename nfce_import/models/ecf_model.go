package models

type RegistroECF struct {
	TIPO           string
	CODLOJA        string
	CODTERMINAL    string
	CODTERMINALEXT string
	CODECF         string
	NUMECF         string
	NUMLOJA        string
	NUMUSUARIO     string
	CRO            string
	MFADICIONAL    string
	MARCA          string
	CNIECF         string
	VERSAOSB       string
	DATASB         string
	HORASB         string
	CODEXTERNO     string
	DATACADUSUARIO string
	HORACADUSUARIO string
	DECIMAISQTD    string
	DECIMAISPRECO  string
	ARREDONDA      string
	NUMEROSERIE    string
}

type RegistroECF2 struct {
	Seq         int    `json:"seq"`
	Campo       string `json:"campo"`
	Descricao   string `json:"descricao"`
	Tipo        string `json:"tipo"`
	Bytes       int    `json:"bytes"`
	Observacoes string `json:"observacoes"`
	Value       string
}

func GetEcf2(fields []string) []RegistroECF2 {
	return []RegistroECF2{
		{Seq: 1, Campo: "TIPO", Descricao: "Tipo de registro, valor fixo.", Tipo: "A", Bytes: 3, Observacoes: "Fixo ECF", Value: fields[0]},
		{Seq: 2, Campo: "CODLOJA", Descricao: "Código da loja no DJMonitor.", Tipo: "N", Bytes: 4, Observacoes: "", Value: fields[1]},
		{Seq: 3, Campo: "CODTERMINAL", Descricao: "Código do terminal no DJMonitor.", Tipo: "N", Bytes: 4, Observacoes: "", Value: fields[2]},
		{Seq: 4, Campo: "CODTERMINALEXT", Descricao: "Código do terminal utilizado na aplicação de retaguarda", Tipo: "AN", Bytes: 10, Observacoes: "", Value: fields[3]},
		{Seq: 5, Campo: "CODECF", Descricao: "Código do ECF no DJ Monitor.", Tipo: "N", Bytes: 4, Observacoes: "", Value: fields[4]},
		{Seq: 6, Campo: "NUM ECF", Descricao: "Número sequencial do ECF dentro da empresa.", Tipo: "N", Bytes: 4, Observacoes: "", Value: fields[5]},
		{Seq: 7, Campo: "NUM LOJA", Descricao: "Número da loja gravado no ECF.", Tipo: "N", Bytes: 4, Observacoes: "", Value: fields[6]},
		{Seq: 8, Campo: "NUM USUARIO", Descricao: "Número do usuário atual cadastrado no ECF.", Tipo: "N", Bytes: 4, Observacoes: "", Value: fields[7]},
		{Seq: 9, Campo: "CRO", Descricao: "Contador de Reinicio de Operações.", Tipo: "N", Bytes: 4, Observacoes: "", Value: fields[8]},
		{Seq: 10, Campo: "MF ADICIONAL", Descricao: "Letra indicativa da memória adicional.", Tipo: "A", Bytes: 1, Observacoes: "Ver nota 1.", Value: fields[9]},
		{Seq: 11, Campo: "MARCA", Descricao: "Marca do ECF.", Tipo: "A", Bytes: 1, Observacoes: "Ver nota 2.", Value: fields[10]},
		{Seq: 12, Campo: "CNIECF", Descricao: "Código Nacional de Identificação do ECF", Tipo: "A", Bytes: 6, Observacoes: "Ver nota 3.", Value: fields[11]},
		{Seq: 13, Campo: "VERSAOSB", Descricao: "Versão do Software Básico.", Tipo: "A", Bytes: 10, Observacoes: "", Value: fields[12]},
		{Seq: 14, Campo: "DATASB", Descricao: "Data de gravação do Software Básico.", Tipo: "D", Bytes: 8, Observacoes: "", Value: fields[13]},
		{Seq: 15, Campo: "HORASB", Descricao: "Hora de gravação do Software Básico.", Tipo: "H", Bytes: 6, Observacoes: "", Value: fields[14]},
		{Seq: 16, Campo: "CODEXTERNO", Descricao: "Código externo do ECF utilizado pela aplicação de retaguarda.", Tipo: "A", Bytes: 10, Observacoes: "", Value: fields[15]},
		{Seq: 17, Campo: "DATA CADUSUARIO", Descricao: "Data do cadastro do usuário atual.", Tipo: "D", Bytes: 8, Observacoes: "", Value: fields[16]},
		{Seq: 18, Campo: "HORA CADUSUARIO", Descricao: "Hora do cadastrado do usuário atual.", Tipo: "H", Bytes: 6, Observacoes: "", Value: fields[17]},
		{Seq: 19, Campo: "DECIMAIS QTD", Descricao: "Número máximo de casas decimais permitidas para quantidade.", Tipo: "N", Bytes: 1, Observacoes: "Ver nota 4.", Value: fields[18]},
		{Seq: 20, Campo: "DECIMAIS PRECO", Descricao: "Número máximo de casas decimais permitidas para o preço unitário.", Tipo: "N", Bytes: 1, Observacoes: "Ver nota 5.", Value: fields[19]},
		{Seq: 21, Campo: "ARREDONDA", Descricao: "Flag indicando se o ECF faz arredondamento.", Tipo: "A", Bytes: 1, Observacoes: "Ver nota 6.", Value: fields[20]},
		{Seq: 22, Campo: "NUMERO_SERIE", Descricao: "Número de série do ECF.", Tipo: "AN", Bytes: 30, Observacoes: "", Value: fields[21]},
	}
}

func GetEcf(fields []string) RegistroECF {
	return RegistroECF{
		TIPO:           fields[0],
		CODLOJA:        fields[1],
		CODTERMINAL:    fields[2],
		CODTERMINALEXT: fields[3],
		CODECF:         fields[4],
		NUMECF:         fields[5],
		NUMLOJA:        fields[6],
		NUMUSUARIO:     fields[7],
		CRO:            fields[8],
		MFADICIONAL:    fields[9],
		MARCA:          fields[10],
		CNIECF:         fields[11],
		VERSAOSB:       fields[12],
		DATASB:         fields[13],
		HORASB:         fields[14],
		CODEXTERNO:     fields[15],
		DATACADUSUARIO: fields[16],
		HORACADUSUARIO: fields[17],
		DECIMAISQTD:    fields[18],
		DECIMAISPRECO:  fields[19],
		ARREDONDA:      fields[20],
		NUMEROSERIE:    fields[21],
	}
}
