package models

// RegistroINI represents a single row in the Registro INI table.
type RegistroINI struct {
	Seq           int    `json:"seq"`
	Campo         string `json:"campo"`
	Descricao     string `json:"descricao"`
	Tipo          string `json:"tipo"`
	Bytes         int    `json:"bytes"`
	Observacoes   string `json:"observacoes"`
	HoraInicio    string `json:"hora_inicio"`
	HoraFim       string `json:"hora_fim"`
	DataInicio    string `json:"data_inicio"`
	DataFim       string `json:"data_fim"`
	Usuario       string `json:"usuario"`
	CodExportacao string `json:"cod_exportacao"`
	CodLoja       string `json:"cod_loja"`
	Cnpj          string `json:"cnpj"`
	Value         string
}

// Create a slice of RegistroINI to represent the entire table.
var RegistroINIList []RegistroINI

func GetIni(fields []string) []RegistroINI {
	// Populate the RegistroINIList with data from the table.
	RegistroINIList = []RegistroINI{
		{Seq: 1, Campo: "Tipo", Descricao: "Tipo de registro, valor fixo.", Tipo: "A", Bytes: 3, Observacoes: "Fixo INI", Value: fields[0]},
		{Seq: 2, Campo: "Data Geração", Descricao: "Data de geração do arquivo.", Tipo: "D", Bytes: 8, Value: fields[1]},
		{Seq: 3, Campo: "Hora Geração", Descricao: "Hora de geração do arquivo.", Tipo: "H", Bytes: 6, Value: fields[2]},
		{Seq: 4, Campo: "Data de Início", Descricao: "Data de emissão do 1º documento.", Tipo: "D", Bytes: 8, Observacoes: "Data inicial da emissão de documentos.", DataInicio: "Data de emissão do 1º documento.", Value: fields[3]},
		{Seq: 5, Campo: "Hora de Início da Emissão", Descricao: "Hora de emissão do 1º documento.", Tipo: "H", Bytes: 6, Observacoes: "Hora inicial da emissão de documentos.", HoraInicio: "Hora de emissão do 1º documento.", Value: fields[4]},
		{Seq: 6, Campo: "Data de Fim da Emissão", Descricao: "Data de emissão do último documento.", Tipo: "D", Bytes: 8, Observacoes: "Data final da emissão de documentos.", DataFim: "Data de emissão do último documento.", Value: fields[5]},
		{Seq: 7, Campo: "Hora de Fim da Emissão", Descricao: "Hora de emissão do último documento.", Tipo: "H", Bytes: 6, Observacoes: "Hora final da emissão de documentos.", HoraFim: "Hora de emissão do último documento.", Value: fields[6]},
		{Seq: 8, Campo: "Usuário", Descricao: "Usuário que efetuou a exportação.", Tipo: "N", Bytes: 5, Usuario: "Usuário que efetuou a exportação.", Value: fields[7]},
		{Seq: 9, Campo: "Cód. Exportação", Descricao: "Número sequencial da exportação.", Tipo: "N", Bytes: 8, CodExportacao: "Número sequencial da exportação.", Value: fields[8]},
		{Seq: 10, Campo: "Cód. Loja", Descricao: "Código da loja no DJMonitor.", Tipo: "N", Bytes: 4, CodLoja: "Código da loja no DJMonitor.", Value: fields[9]},
		{Seq: 11, Campo: "CNPJ", Descricao: "CNPJ da loja a qual a movimentação pertence.", Tipo: "N", Bytes: 14, Cnpj: "CNPJ da loja a qual a movimentação pertence.", Value: fields[10]},
	}

	return RegistroINIList
}
