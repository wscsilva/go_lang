package models

type RegistroDIT struct {
	Seq         int    `json:"seq"`
	Campo       string `json:"campo"`
	Descricao   string `json:"descricao"`
	Tipo        string `json:"tipo"`
	Bytes       int    `json:"bytes"`
	Observacoes string `json:"observacoes"`
}

func GetDit(fields []string) []RegistroDIT {
	return []RegistroDIT{
		{Seq: 1, Campo: "TIPO", Descricao: "Tipo de registro, valor fixo.", Tipo: "A", Bytes: 3, Observacoes: "Fixo DIT"},
		{Seq: 2, Campo: "CODLOJA", Descricao: "Código da loja no DJMonitor.", Tipo: "N", Bytes: 4, Observacoes: ""},
		{Seq: 3, Campo: "CODTERMINAL", Descricao: "Código do terminal no DJMonitor.", Tipo: "N", Bytes: 4, Observacoes: ""},
		{Seq: 4, Campo: "CODTURNO", Descricao: "Código do turno.", Tipo: "N", Bytes: 4, Observacoes: ""},
		{Seq: 5, Campo: "COO", Descricao: "COO do documento.", Tipo: "N", Bytes: 9, Observacoes: ""},
		{Seq: 6, Campo: "SEQUENCIA", Descricao: "Sequência de ocorrência dos itens.", Tipo: "N", Bytes: 4, Observacoes: ""},
		{Seq: 7, Campo: "CODBARRAS", Descricao: "Código de barras do produto.", Tipo: "AN", Bytes: 20, Observacoes: ""},
		{Seq: 8, Campo: "CODPRODUTOEXT", Descricao: "Código do produto utilizado na aplicação de retaguarda", Tipo: "AN", Bytes: 20, Observacoes: ""},
		{Seq: 9, Campo: "QUANTIDADE", Descricao: "Quantidade vendida", Tipo: "N", Bytes: 9, Observacoes: "Com 3 casas decimais."},
		{Seq: 10, Campo: "PRECO UNIT", Descricao: "Preço unitário do produto.", Tipo: "N", Bytes: 12, Observacoes: "Com 3 casas decimais."},
		{Seq: 11, Campo: "DESCONTO_ACRESCIMO", Descricao: "Desconto ou acréscimo aplicado neste produto. (Se houve um desconto o valor grava será negativo).", Tipo: "N", Bytes: 10, Observacoes: "Com 2 casas decimais."},
		{Seq: 12, Campo: "SITU TRIBUTARIA", Descricao: "Situação tributária", Tipo: "A", Bytes: 1, Observacoes: "Conforme nota"},
		{Seq: 13, Campo: "ICMS", Descricao: "Valor da alíquota de ICMS ou ISS para este produto.", Tipo: "N", Bytes: 4, Observacoes: "Com 2 casas decimais."},
		{Seq: 14, Campo: "CANCELADO", Descricao: "Se o item foi cancelado.", Tipo: "A", Bytes: 1, Observacoes: "S - Sim\nN - Não."},
		{Seq: 15, Campo: "DESC ADICIONAL", Descricao: "Descrição adicional para o item.", Tipo: "AN", Bytes: 40, Observacoes: ""},
		{Seq: 16, Campo: "COD VENDEDOR", Descricao: "Código do vendedor que vendeu o item.", Tipo: "AN", Bytes: 6, Observacoes: ""},
		{Seq: 17, Campo: "TOTAL LIQ ITEM", Descricao: "Valor total do item (Incluindo os descontos ou acréscimos).", Tipo: "N", Bytes: 12, Observacoes: "Com 2 casas decimais."},
		{Seq: 18, Campo: "ID MOTIVO_CANC", Descricao: "Código externo do motivo de desconto do item.", Tipo: "N", Bytes: 6, Observacoes: ""},
		{Seq: 19, Campo: "ID MOTIVO_DESC", Descricao: "Código externo do motivo de desconto do item.", Tipo: "N", Bytes: 6, Observacoes: ""},
		{Seq: 20, Campo: "COD TOTALIZADOR", Descricao: "Código do totalizador de tributação do produto no ECF.", Tipo: "AN", Bytes: 7, Observacoes: "Conforme leiaute exigido no SPED Fiscal."},
		{Seq: 21, Campo: "CONTROLA_ESTOQUE", Descricao: "Produto controla estoque.", Tipo: "AN", Bytes: 1, Observacoes: "S - Sim\nN - Não."},
		{Seq: 22, Campo: "CODLIDOVENDA", Descricao: "Código utilizado na hora da venda", Tipo: "AN", Bytes: 0, Observacoes: ""},
		{Seq: 23, Campo: "NUMVALEPRESENTE", Descricao: "Número do vale presente vendido.", Tipo: "AN", Bytes: 40, Observacoes: ""},
	}
}
