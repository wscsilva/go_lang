package model

type RegistroPDV struct {
	Seq         int
	Campo       string
	Descricao   string
	Tipo        string
	Bytes       int
	Observacoes string
}

func GetPdv(fields []string) []RegistroPDV {
	return []RegistroPDV{
		{
			Seq:         1,
			Campo:       "TIPO",
			Descricao:   "Tipo de registro, valor fixo.",
			Tipo:        "A",
			Bytes:       3,
			Observacoes: "Fixo PDV",
		},
		{
			Seq:         2,
			Campo:       "CODTERMINAL",
			Descricao:   "Código do terminal.",
			Tipo:        "N",
			Bytes:       9,
			Observacoes: "Código do terminal no DJMonitor.",
		},
		{
			Seq:         3,
			Campo:       "CODEXTERNO",
			Descricao:   "Código externo do terminal.",
			Tipo:        "N",
			Bytes:       9,
			Observacoes: "",
		},
		{
			Seq:         4,
			Campo:       "PAF",
			Descricao:   "Versão do PAF.",
			Tipo:        "N",
			Bytes:       8,
			Observacoes: "",
		},
		{
			Seq:         5,
			Campo:       "HOMOLOGACAO",
			Descricao:   "Versão de homologação.",
			Tipo:        "N",
			Bytes:       8,
			Observacoes: "",
		},
	}
}
