package model

type RegistroMON struct {
	Codigo      string
	Nome        string
	Descricao   string
	Tipo        string
	Tamanho     string
	ValorPadrao string
	Value       string
}

func GetMon(fields []string) []RegistroMON {
	return []RegistroMON{
		{
			Codigo:      "01",
			Nome:        "Tipo",
			Descricao:   "Tipo de registro, valor fixo.",
			Tipo:        "A",
			Tamanho:     "3",
			ValorPadrao: "Fixo MON",
			Value:       fields[0],
		},
		{
			Codigo:      "02",
			Nome:        "Versão",
			Descricao:   "Versão do DJ Monitor.",
			Tipo:        "A",
			Tamanho:     "8",
			ValorPadrao: "",
			Value:       fields[1],
		},
		{
			Codigo:      "03",
			Nome:        "Tabela",
			Descricao:   "Versão das tabelas do banco de dados.",
			Tipo:        "N",
			Tamanho:     "4",
			ValorPadrao: "",
			Value:       fields[2],
		},
		{
			Codigo:      "04",
			Nome:        "Firebird",
			Descricao:   "Versão do Firebird instalado.",
			Tipo:        "A",
			Tamanho:     "4",
			ValorPadrao: "",
			Value:       fields[3],
		},
		{
			Codigo:      "05",
			Nome:        "ODS",
			Descricao:   "Versão do ODS instalado.",
			Tipo:        "A",
			Tamanho:     "4",
			ValorPadrao: "",
			Value:       fields[4],
		},
	}
}
