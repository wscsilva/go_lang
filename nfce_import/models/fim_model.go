package models

type RegistroFIM struct {
	Seq         int    `json:"seq"`
	Campo       string `json:"campo"`
	Descricao   string `json:"descricao"`
	Tipo        string `json:"tipo"`
	Bytes       int    `json:"bytes"`
	Observacoes string `json:"observacoes"`
	Value       string
}

func GetFim(fields []string) []RegistroFIM {
	return []RegistroFIM{
		{Seq: 1, Campo: "TIPO", Descricao: "Tipo de registro, valor fixo.", Tipo: "A", Bytes: 3, Observacoes: "Fixo FIM", Value: fields[0]},
		{Seq: 2, Campo: "TOTAL REG. ECF", Descricao: "Total de registros ECF neste arquivo.", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[1]},
		{Seq: 3, Campo: "TOTAL REG. DOC", Descricao: "Total de registros DOC neste arquivo.", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[2]},
		{Seq: 4, Campo: "TOTAL REG. DIT", Descricao: "Total de registros DIT neste arquivo.", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[3]},
		{Seq: 5, Campo: "TOTAL REG. DKI", Descricao: "Total de registros DKI neste arquivo.", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[4]},
		{Seq: 6, Campo: "TOTAL REG. DPG", Descricao: "Total de registros DPG neste arquivo.", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[5]},
		{Seq: 7, Campo: "TOTAL REG. DPA", Descricao: "Total de registros DPA neste arquivo.", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[6]},
		{Seq: 8, Campo: "TOTAL REG. DCH", Descricao: "Total de registros DCH neste arquivo.", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[7]},
		{Seq: 9, Campo: "TOTAL REG. TEF", Descricao: "Total de registros TEF neste arquivo.", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[8]},
		{Seq: 10, Campo: "TOTAL REG. REZ", Descricao: "Total de registros REZ neste arquivo.", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[9]},
		{Seq: 11, Campo: "TOTAL REG. RED", Descricao: "Total de registros RED neste arquivo.", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[10]},
		{Seq: 12, Campo: "TOTAL REG. TIV", Descricao: "Total de registros TIV neste arquivo.", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[11]},
		{Seq: 13, Campo: "TOTAL GERAL", Descricao: "Total de todos os registros arquivo (Incluindo INI e FIM).", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[12]},
		{Seq: 14, Campo: "DATA FIM GERAÇÃO", Descricao: "Data do término de geração do arquivo.", Tipo: "D", Bytes: 8, Observacoes: "", Value: fields[13]},
		{Seq: 15, Campo: "HORA FIM GERAÇÃO", Descricao: "Hora do término de geração do arquivo.", Tipo: "H", Bytes: 6, Observacoes: "", Value: fields[14]},
		{Seq: 16, Campo: "TOTAL REG. DFE", Descricao: "Total de registros DFE neste arquivo.", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[15]},
		{Seq: 17, Campo: "TOTAL REG. VCR", Descricao: "Total de registros VCR neste arquivo.", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[16]},
		{Seq: 18, Campo: "TOTAL REG. VIT", Descricao: "Total de registros VIT neste arquivo.", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[17]},
		{Seq: 19, Campo: "TOTAL REG. NFE", Descricao: "Total de registros NFE neste arquivo.", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[18]},
		{Seq: 20, Campo: "TOTAL REG. NEV", Descricao: "Total de registros NEV neste arquivo.", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[19]},
		{Seq: 21, Campo: "TOTAL REG. DXL", Descricao: "Total de registros DXL neste arquivo.", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[20]},
		{Seq: 22, Campo: "TOTAL REG. DOR", Descricao: "Total de registro DOR neste arquivo", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[21]},
		{Seq: 23, Campo: "TOTAL REG. DRE", Descricao: "Total de registro DRE neste arquivo", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[22]},
		{Seq: 24, Campo: "TOTAL REG. DEC", Descricao: "Total de registro DEC neste arquivo", Tipo: "N", Bytes: 9, Observacoes: "", Value: fields[23]},
	}
}
