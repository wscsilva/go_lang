package model

type RegistroMON struct {
	Seq         int    `json:"seq"`
	Campo       string `json:"campo"`
	Descricao   string `json:"descricao"`
	Tipo        string `json:"tipo"`
	Bytes       int    `json:"bytes"`
	Observacoes string `json:"observacoes"`
	Value       string
}

func GetMon(fields []string) []RegistroMON {
	return []RegistroMON{
		{Seq: 1, Campo: "Tipo", Descricao: "Tipo de registro, valor fixo.", Tipo: "A", Bytes: 3, Observacoes: "Fixo MON", Value: fields[0]},
		{Seq: 2, Campo: "Versão", Descricao: "Versão do DJ Monitor.", Tipo: "A", Bytes: 8, Observacoes: "", Value: fields[1]},
		{Seq: 3, Campo: "Tabela", Descricao: "Versão das tabelas do banco de dados.", Tipo: "N", Bytes: 4, Observacoes: "", Value: fields[2]},
		{Seq: 4, Campo: "Firebird", Descricao: "Versão do Firebird instalado.", Tipo: "A", Bytes: 4, Observacoes: "", Value: fields[3]},
		{Seq: 5, Campo: "ODS", Descricao: "Versão do ODS instalado.", Tipo: "A", Bytes: 4, Observacoes: "", Value: fields[4]},
	}
}
