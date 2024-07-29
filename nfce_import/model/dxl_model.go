package model

type RegistroDXL struct {
	Seq         int    `json:"seq"`
	Campo       string `json:"campo"`
	Descricao   string `json:"descricao"`
	Tipo        string `json:"tipo"`
	Bytes       string `json:"bytes"`
	Observacoes string `json:"observacoes"`
	Value       string
}

func GetDxl(fields []string) []RegistroDXL {
	return []RegistroDXL{
		{Seq: 1, Campo: "TIPO", Descricao: "Tipo de registro, valor fixo.", Tipo: "A", Bytes: "3", Observacoes: "Fixo DXL", Value: fields[0]},
		{Seq: 2, Campo: "TERMINAL", Descricao: "Código do terminal no DJMonitor.", Tipo: "N", Bytes: "4", Observacoes: "", Value: fields[1]},
		{Seq: 3, Campo: "TURNO", Descricao: "Código do turno.", Tipo: "N", Bytes: "4", Observacoes: "", Value: fields[2]},
		{Seq: 4, Campo: "COD ECF", Descricao: "Código do ECF no DJMonitor.", Tipo: "N", Bytes: "4", Observacoes: "", Value: fields[3]},
		{Seq: 5, Campo: "COO", Descricao: "COO do documento.", Tipo: "N", Bytes: "9", Observacoes: "", Value: fields[4]},
		{Seq: 6, Campo: "CHAVE DFE", Descricao: "Chave do documento.", Tipo: "N", Bytes: "50", Observacoes: "", Value: fields[5]},
		{Seq: 7, Campo: "DATA EMIS", Descricao: "Data emissão do documento.", Tipo: "D", Bytes: "8", Observacoes: "", Value: fields[6]},
		{Seq: 8, Campo: "HORA EMIS", Descricao: "Hora de emissão do documento.", Tipo: "H", Bytes: "6", Observacoes: "", Value: fields[7]},
		{Seq: 9, Campo: "DATA AUTOR", Descricao: "Data autorização do documento.", Tipo: "D", Bytes: "8", Observacoes: "", Value: fields[8]},
		{Seq: 10, Campo: "HORA AUTOR", Descricao: "Hora autorização do documento.", Tipo: "H", Bytes: "6", Observacoes: "", Value: fields[9]},
		{Seq: 11, Campo: "TIPO DOCTO", Descricao: "Tipo do documento.", Tipo: "A", Bytes: "1", Observacoes: "Ver nota 1.", Value: fields[10]},
		{Seq: 12, Campo: "STATUS", Descricao: "Status do documento.", Tipo: "A", Bytes: "1", Observacoes: "Ver nota 2.", Value: fields[11]},
		{Seq: 13, Campo: "CSTAT", Descricao: "Código do status da resposta", Tipo: "N", Bytes: "4", Observacoes: "", Value: fields[12]},
		{Seq: 14, Campo: "MOTIVO", Descricao: "Motivo.", Tipo: "AN", Bytes: "255", Observacoes: "", Value: fields[13]},
		{Seq: 15, Campo: "PROTOCOLO", Descricao: "Protocolo.", Tipo: "N", Bytes: "20", Observacoes: "", Value: fields[14]},
		{Seq: 16, Campo: "MODELO", Descricao: "Modelo do documento.", Tipo: "N", Bytes: "2", Observacoes: "", Value: fields[15]},
		{Seq: 17, Campo: "XML", Descricao: "Conteúdo do XML.", Tipo: "AN", Bytes: "-", Observacoes: "", Value: fields[16]},
	}
}
