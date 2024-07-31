package models

type RegistroDoc2 struct {
	Seq         int    `json:"seq"`
	Campo       string `json:"campo"`
	Descricao   string `json:"descricao"`
	Tipo        string `json:"tipo"`
	Bytes       int    `json:"bytes"`
	Observacoes string `json:"observacoes"`
	Value       string
	// Added for the second table
}

type RegistroDoc struct {
	Documento map[string]string
}

func GetDoc2(fields []string) []RegistroDoc2 {
	return []RegistroDoc2{
		{Seq: 1, Campo: "TIPO", Descricao: "Tipo de registro, valor fixo.", Tipo: "A", Bytes: 3, Observacoes: "Fixo DOC", Value: fields[0]},
		{Seq: 2, Campo: "CODTERMINAL", Descricao: "Código do terminal no DJMonitor", Tipo: "N", Bytes: 4, Observacoes: "",
			Value: fields[1]},
		{Seq: 3, Campo: "CODTURNO", Descricao: "Código do turno", Tipo: "N", Bytes: 4, Observacoes: "",
			Value: fields[2]},
		{Seq: 4, Campo: "CODECF", Descricao: "Código do ECF", Tipo: "N", Bytes: 4, Observacoes: "",
			Value: fields[3]},
		{Seq: 5, Campo: "COO", Descricao: "COO do documento", Tipo: "N", Bytes: 9, Observacoes: "",
			Value: fields[4]},
		{Seq: 6, Campo: "DENOMINACAO", Descricao: "Conforme tabela da CAT 60/07", Tipo: "A", Bytes: 2, Observacoes: "Ver nota 1.",
			Value: fields[5]},
		{Seq: 7, Campo: "CONTADOR", Descricao: "Valor do contador relacionado ao documento emitido", Tipo: "N", Bytes: 9, Observacoes: "Ver nota 2.",
			Value: fields[6]},
		{Seq: 8, Campo: "DATA_INICIO", Descricao: "Data do início da emissão do documento", Tipo: "D", Bytes: 8, Observacoes: "",
			Value: fields[7]},
		{Seq: 9, Campo: "HORA_INICIO", Descricao: "Hora do inicio da emissão do documento", Tipo: "H", Bytes: 6, Observacoes: "",
			Value: fields[8]},
		{Seq: 10, Campo: "DATA_FIM", Descricao: "Data do fim da emissão do documento", Tipo: "D", Bytes: 8, Observacoes: "",
			Value: fields[9]},
		{Seq: 11, Campo: "HORA_FIM", Descricao: "Hora do fim da emissão do documento", Tipo: "H", Bytes: 6, Observacoes: "",
			Value: fields[10]},
		{Seq: 12, Campo: "DESCRICAO", Descricao: "Breve descrição sobre o documento", Tipo: "A", Bytes: 60, Observacoes: "",
			Value: fields[11]},
		{Seq: 13, Campo: "CODCLIENTEEXT", Descricao: "Código do cliente utilizado pela aplicação de retaguarda", Tipo: "AN", Bytes: 20, Observacoes: "",
			Value: fields[12]},
		{Seq: 14, Campo: "CODCLIENTE", Descricao: "Código do cliente no DJMonitor", Tipo: "N", Bytes: 9, Observacoes: "",
			Value: fields[13]},
		{Seq: 15, Campo: "CODAUTORIZADO", Descricao: "Código do autorizado do DJMonitor", Tipo: "N", Bytes: 9, Observacoes: "",
			Value: fields[14]},
		{Seq: 16, Campo: "DOCUMENTO_CLI", Descricao: "Documento de identificação do", Tipo: "A", Bytes: 20, Observacoes: "",
			Value: fields[15]},
		{Seq: 17, Campo: "NOME_CLI", Descricao: "Nome do cliente", Tipo: "A", Bytes: 50, Observacoes: "",
			Value: fields[16]},
		{Seq: 18, Campo: "SUBTOTAL", Descricao: "Subtotal do documento, soma dos itens, ainda sem descontos ou acréscimos", Tipo: "N", Bytes: 12, Observacoes: "Com 2 casas decimais.",
			Value: fields[17]},
		{Seq: 19, Campo: "DESCONTO_ACRESCIMO", Descricao: "Desconto ou acréscimo aplicado no documento", Tipo: "N", Bytes: 10, Observacoes: "Com 2 casas decimais.",
			Value: fields[18]},
		{Seq: 20, Campo: "TOTAL_PAGO", Descricao: "Total dos pagamentos efetuados", Tipo: "N", Bytes: 12, Observacoes: "Com 2 casas decimais.",
			Value: fields[19]},
		{Seq: 21, Campo: "TROCO", Descricao: "", Tipo: "N", Bytes: 12, Observacoes: "Com 2 casas decimais.",
			Value: fields[20]},
		{Seq: 22, Campo: "CODPLANOPGEXT", Descricao: "Código do plano de pagamento na aplicação de retaguarda", Tipo: "AN", Bytes: 10, Observacoes: "",
			Value: fields[21]},
		{Seq: 23, Campo: "PLANOPAGTO", Descricao: "Código do plano de pagamento no DJMonitor", Tipo: "A", Bytes: 2, Observacoes: "",
			Value: fields[22]},
		{Seq: 24, Campo: "CANCELADO", Descricao: "Documento cancelado", Tipo: "A", Bytes: 1, Observacoes: "S - Sim\nN - Não",
			Value: fields[23]},
		{Seq: 25, Campo: "QTD_TOTAL_ITENS", Descricao: "Quantidade total de itens da venda", Tipo: "N", Bytes: 12, Observacoes: "",
			Value: fields[24]},
		{Seq: 26, Campo: "CD_EXT_PREVENDA", Descricao: "Código externo da pré-venda/ DAV (Caso tenha sido gerada através da importação de pré-vendas/ DAV)", Tipo: "AN", Bytes: 20, Observacoes: "",
			Value: fields[25]},
		{Seq: 27, Campo: "VENDEDOR_PRE", Descricao: "Código externo do operador que cadastrou a pré-venda", Tipo: "AN", Bytes: 10, Observacoes: "Ver nota 3.",
			Value: fields[26]},
		{Seq: 28, Campo: "ID_MOTIVO_CANC", Descricao: "Código externo do motivo de cancelamento da venda", Tipo: "N", Bytes: 6, Observacoes: "Ver nota 4.",
			Value: fields[27]},
		{Seq: 29, Campo: "ID_MOTIVO_DESC", Descricao: "Código externo do motivo de desconto na venda", Tipo: "N", Bytes: 6, Observacoes: "",
			Value: fields[28]},
		{Seq: 30, Campo: "OBSERVACAO", Descricao: "Observações impressas no rodapé do cupom fiscal", Tipo: "A", Bytes: 60, Observacoes: "",
			Value: fields[29]},
		{Seq: 31, Campo: "TIPO_ORCAMENTO", Descricao: "Campo para diferenciar se o orçamento veio da pré-venda DAV", Tipo: "AN", Bytes: 1, Observacoes: "Ver nota 5.",
			Value: fields[30]},
		{Seq: 32, Campo: "MODELO", Descricao: "Campo exibirá o modelo do documento emitido", Tipo: "AN", Bytes: 2, Observacoes: "Ver nota 6.",
			Value: fields[31]},
		{Seq: 33, Campo: "TIPO_EMISSAO", Descricao: "Campo exibirá forma de emissão do documento", Tipo: "N", Bytes: 1, Observacoes: "Ver nota 7.",
			Value: fields[32]},
		{Seq: 34, Campo: "NUMERO", Descricao: "Campo exibirá o número do cupom fiscal eletrônico (NFC-e/ SAT/ MFE)", Tipo: "N", Bytes: 6, Observacoes: "Ver nota 8. 6 ou 9",
			Value: fields[33]},
		{Seq: 35, Campo: "SERIE", Descricao: "Campo exibirá a série do cupom fiscal eletrônico (NFC-e/ SAT/ MFE)", Tipo: "N", Bytes: 3, Observacoes: "",
			Value: fields[34]},
		{Seq: 36, Campo: "CHAVE_DFE", Descricao: "Chave do Documento", Tipo: "N", Bytes: 50, Observacoes: "",
			Value: fields[35]},
	}
}

func GetDoc1(fields []string) RegistroDoc {
	resultado := make(map[string]string)
	resultado["TIPO"] = fields[0]
	resultado["CODTERMINAL"] = fields[1]
	resultado["CODTURNO"] = fields[2]
	resultado["CODECF"] = fields[3]
	resultado["COO"] = fields[4]
	resultado["DENOMINACAO"] = fields[5]
	resultado["CONTADOR"] = fields[6]
	resultado["DATA_INICIO"] = fields[7]
	resultado["HORA_INICIO"] = fields[8]
	resultado["DATA_FIM"] = fields[9]
	resultado["HORA_FIM"] = fields[10]
	resultado["DESCRICAO"] = fields[11]
	resultado["CODCLIENTEEXT"] = fields[12]
	resultado["CODCLIENTE"] = fields[13]
	resultado["CODAUTORIZADO"] = fields[14]
	resultado["DOCUMENTO_CLI"] = fields[15]
	resultado["NOME_CLI"] = fields[16]
	resultado["SUBTOTAL"] = fields[17]
	resultado["DESCONTO_ACRESCIMO"] = fields[18]
	resultado["TOTAL_PAGO"] = fields[19]
	resultado["TROCO"] = fields[20]
	resultado["CODPLANOPGEXT"] = fields[21]
	resultado["PLANOPAGTO"] = fields[22]
	resultado["CANCELADO"] = fields[23]
	resultado["QTD_TOTAL_ITENS"] = fields[24]
	resultado["CD_EXT_PREVENDA"] = fields[25]
	resultado["VENDEDOR_PRE"] = fields[26]
	resultado["ID_MOTIVO_CANC"] = fields[27]
	resultado["ID_MOTIVO_DESC"] = fields[28]
	resultado["OBSERVACAO"] = fields[29]
	resultado["TIPO_ORCAMENTO"] = fields[30]
	resultado["MODELO"] = fields[31]
	resultado["TIPO_EMISSAO"] = fields[32]
	resultado["NUMERO"] = fields[33]
	resultado["SERIE"] = fields[34]
	resultado["CHAVE_DFE"] = fields[35]
	return RegistroDoc{Documento: resultado}
}

type GetRegistrosRV struct {
	TIPO               string
	CODTERMINAL        string
	CODTURNO           string
	CODECF             string
	COO                string
	DENOMINACAO        string
	CONTADOR           string
	DATA_INICIO        string
	HORA_INICIO        string
	DATA_FIM           string
	HORA_FIM           string
	DESCRICAO          string
	CODCLIENTEEXT      string
	CODCLIENTE         string
	CODAUTORIZADO      string
	DOCUMENTO_CLI      string
	NOME_CLI           string
	SUBTOTAL           string
	DESCONTO_ACRESCIMO string
	TOTAL_PAGO         string
	TROCO              string
	CODPLANOPGEXT      string
	PLANOPAGTO         string
	CANCELADO          string
	QTD_TOTAL_ITENS    string
	CD_EXT_PREVENDA    string
	VENDEDOR_PRE       string
	ID_MOTIVO_CANC     string
	ID_MOTIVO_DESC     string
	OBSERVACAO         string
	TIPO_ORCAMENTO     string
	MODELO             string
	TIPO_EMISSAO       string
	NUMERO             string
	SERIE              string
	CHAVE_DFE          string
}

func GetDoc(fields []string) GetRegistrosRV {
	return GetRegistrosRV{
		TIPO:               fields[0],
		CODTERMINAL:        fields[1],
		CODTURNO:           fields[2],
		CODECF:             fields[3],
		COO:                fields[4],
		DENOMINACAO:        fields[5],
		CONTADOR:           fields[6],
		DATA_INICIO:        fields[7],
		HORA_INICIO:        fields[8],
		DATA_FIM:           fields[9],
		HORA_FIM:           fields[10],
		DESCRICAO:          fields[11],
		CODCLIENTEEXT:      fields[12],
		CODCLIENTE:         fields[13],
		CODAUTORIZADO:      fields[14],
		DOCUMENTO_CLI:      fields[15],
		NOME_CLI:           fields[16],
		SUBTOTAL:           fields[17],
		DESCONTO_ACRESCIMO: fields[18],
		TOTAL_PAGO:         fields[19],
		TROCO:              fields[20],
		CODPLANOPGEXT:      fields[21],
		PLANOPAGTO:         fields[22],
		CANCELADO:          fields[23],
		QTD_TOTAL_ITENS:    fields[24],
		CD_EXT_PREVENDA:    fields[25],
		VENDEDOR_PRE:       fields[26],
		ID_MOTIVO_CANC:     fields[27],
		ID_MOTIVO_DESC:     fields[28],
		OBSERVACAO:         fields[29],
		TIPO_ORCAMENTO:     fields[30],
		MODELO:             fields[31],
		TIPO_EMISSAO:       fields[32],
		NUMERO:             fields[33],
		SERIE:              fields[34],
		CHAVE_DFE:          fields[35],
	}
}
