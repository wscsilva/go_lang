package model

type RegistroDFE struct {
	Seq         int    `json:"seq"`
	Campo       string `json:"campo"`
	Descricao   string `json:"descricao"`
	Tipo        string `json:"tipo"`
	Bytes       string `json:"bytes"`
	Observacoes string `json:"observacoes"`
	Value       string
}

func GetDfe(fields []string) []RegistroDFE {
	if len(fields) < 5 {
		return []RegistroDFE{}
	}

	return []RegistroDFE{
		{Seq: 1, Campo: "TIPO", Descricao: "Tipo de registro, valor fixo.", Tipo: "A", Bytes: "3", Observacoes: "Fixo DFE", Value: fields[0]},
		{Seq: 2, Campo: "DFE_CHAVE", Descricao: "Chave do XML.", Tipo: "N", Bytes: "X", Observacoes: "", Value: fields[1]},
		{Seq: 3, Campo: "DFE_PROTOCOLO", Descricao: "Protocolo.", Tipo: "N", Bytes: "15", Observacoes: "", Value: fields[2]},
		{Seq: 4, Campo: "DFE_MODELO", Descricao: "Modelo do documento.", Tipo: "N", Bytes: "2", Observacoes: "", Value: fields[3]},
		{Seq: 5, Campo: "DFE_XML", Descricao: "XML.", Tipo: "AN", Bytes: "X", Observacoes: "", Value: fields[4]},
	}
}
