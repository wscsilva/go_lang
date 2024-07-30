package model

type RegistroPDV struct {
	Seq         int    `json:"seq"`
	Campo       string `json:"campo"`
	Descricao   string `json:"descricao"`
	Tipo        string `json:"tipo"`
	Bytes       int    `json:"bytes"`
	Observacoes string `json:"observacoes"`
	Value       string
}

func GetPdv(fields []string) []RegistroPDV {
	return []RegistroPDV{
		{Seq: 1, Campo: "TIPO", Descricao: "Tipo de registro, valor fixo.", Tipo: "A", Bytes: 3, Observacoes: "Fixo PDV", Value: fields[0]},
		{Seq: 2, Campo: "CODTERMINAL", Descricao: "Código do terminal.", Tipo: "N", Bytes: 9, Observacoes: "Código do terminal no DJMonitor.", Value: fields[1]},
		{Seq: 3, Campo: "CODEXTERNO", Descricao: "Código externo do terminal.", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[2]},
		{Seq: 4, Campo: "PAF", Descricao: "Versão do PAF.", Tipo: "N", Bytes: 8, Observacoes: "", Value: fields[3]},
		{Seq: 5, Campo: "HOMOLOGACAO", Descricao: "Versão de homologação.", Tipo: "N", Bytes: 8, Observacoes: "", Value: fields[4]},
	}
}
