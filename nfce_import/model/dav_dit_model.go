package model

type RegistroDAV_DIT struct {
	Seq       int    `json:"seq"`
	Campo     string `json:"campo"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
	Value     string `json:"value,omitempty"`
}

func GetDavDit(fields []string) []RegistroDAV_DIT {
	return []RegistroDAV_DIT{
		{Seq: 1, Campo: "Tipo de Registro", Tipo: "AN", Descricao: "Fixo DIT", Value: fields[0]},
		{Seq: 2, Campo: "Sequência", Tipo: "SI", Descricao: "Número de sequência do item no DAV.", Value: fields[1]},
		{Seq: 3, Campo: "Código do Produto", Tipo: "AN", Descricao: "Código do produto pode ser:\nExterno: o código do produto no programa retaguarda.\nBarras EAN14. Ver Notas\nBarras Alternativo. Ver Notas.\nTamanho máximo de 20 caracteres.", Value: fields[2]},
		{Seq: 4, Campo: "Quantidade", Tipo: "AN", Descricao: "Quantidade a ser vendida Com 3 casas decimais.", Value: fields[3]},
		{Seq: 5, Campo: "Preço Unitário", Tipo: "N", Descricao: "Preço unitário do produto. Com 3 casas decimais.", Value: fields[4]},
		{Seq: 6, Campo: "Valor Desconto", Tipo: "N", Descricao: "Valor do desconto no produto. Com 2 casas decimais.", Value: fields[5]},
		{Seq: 7, Campo: "Valor Acréscimo", Tipo: "N", Descricao: "Valor do acréscimo no produto. Com 2 casas decimais.", Value: fields[6]},
		{Seq: 8, Campo: "Total Líquido", Tipo: "N", Descricao: "Valor total dos itens levando em consideração os descontos e acréscimos. Com\n2 casas de cimais.", Value: fields[7]},
		{Seq: 9, Campo: "Código de Barras", Tipo: "AN", Descricao: "Código de barras do produto (não devem existir produtos diferentes com o mesmo\ncódigo de barras). Tamanho máximo de 20 caracteres.", Value: fields[8]},
		{Seq: 10, Campo: "Descrição", Tipo: "AN", Descricao: "Descrição do produto. Tamanho máximo de 120 caracteres.", Value: fields[9]},
		{Seq: 11, Campo: "Complemento", Tipo: "AN", Descricao: "Complemento da descrição do produto. Tamanho máximo de 500 caracteres.", Value: fields[10]},
		{Seq: 12, Campo: "Unidade de Medida", Tipo: "AN", Descricao: "Unidade de medida do produto (Usar no máximo 4 caracteres).", Value: fields[11]},
		{Seq: 13, Campo: "Situação Tributária", Tipo: "A", Descricao: "Situação tributária do produto. Valores Válidos:\nF- Substituição Tributaria 1;\nG- Substituição Tributaria 2;\nH- Substituição Tributaria 3;\n1 - Isento 1;\nJ\n- Isento 2:\nK - Isento 3:\nN- Não Tributado 1:\nO- Não Tributado 2:\nP - Não Tributado 3:\nQ- Isento ISSQN;\nS- Tributada pelo ISSQN;\nT- Tributada pelo ICMS.", Value: fields[12]},
		{Seq: 14, Campo: "ICMS", Tipo: "N", Descricao: "Alíquota do ICMS ou ISS, caso não haja incidência destes impostos, informar 0.\nCom 2 casas decimais.", Value: fields[13]},
		{Seq: 15, Campo: "Calcula Quantidade", Tipo: "AN", Descricao: "'S' ou 'N' - Use 'S' para produtos \"pesáveis\" onde a etiqueta contem o código e\na quantidade ou o valor total. Nesse caso o campo Código de Barras deve ter\napenas 4 dígitos. Use 'N' para produtos não \"pesáveis\".", Value: fields[14]},
		{Seq: 16, Campo: "Bloqueia Quantidade Fracionária", Tipo: "AN", Descricao: "'S' ou 'N' - Se 'S' não permite a digitação de decimais nesse produto. (Ex: 1,5 coca\ncola). Se Calcula Qtd = 'S' ou Bloq Qtd = 'S assume se 'S'\nassume s'neste campo'. Use 'N'\npara não bloquear a digitação da quantidade fracionária..", Value: fields[15]},
		{Seq: 17, Campo: "Bloqueia Quantidade", Tipo: "AN", Descricao: "'S' ou 'N' - Se 'S' não permite a digitação a quantidade para este produto, vende\nsempre 1. Se Calcula Qtd = 'S' ou Bloq Qtd = 'S'. Use 'N' para não bloquear a\ndigitação da quantidade.", Value: fields[16]},
		{Seq: 18, Campo: "Produção Própria", Tipo: "AN", Descricao: "Se o produto é de produção própria. S - Sim, N - Não", Value: fields[17]},
		{Seq: 19, Campo: "Quantidade Estoque", Tipo: "N", Descricao: "Quantidade em estoque do produto. Com 3 casas decimais", Value: fields[18]},
		{Seq: 20, Campo: "Descrição Adicional", Tipo: "AN", Descricao: "Descrição adicional para o produto. O texto informado aqui é impresso no cupom\nfical. Pode ser usado para número de série do item, por exemplo. Máximo de 40\ncaracteres.", Value: fields[19]},
		{Seq: 21, Campo: "Código do Vendedor", Tipo: "AN", Descricao: "Código do vendedor que incluiu o item. (Este código é o mesmo da tabela de\nvendedores do DJMonitor, não é utilizado o código externo, por isso a tabela de\nvendedores deve ser cadastrada com os mesmos códigos utilizados pelo\nretaguarda. Máximo de 6 caracteres).", Value: fields[20]},
		{Seq: 22, Campo: "Nome do Vendedor", Tipo: "AN", Descricao: "Nome do vendedor, caso o mesmo não seja encontrado na tabela de vendedores,\nele será incluído, caso este campo seja informado. Máximo de 30 caracteres.", Value: fields[21]},
		{Seq: 23, Campo: "Gtin Contábil", Tipo: "AN", Descricao: "Tamanho máximo de 20 caracteres.", Value: fields[22]},
		{Seq: 24, Campo: "EX TIPI", Tipo: "AN", Descricao: "Tamanho máximo de 20 caracteres.", Value: fields[23]},
		{Seq: 25, Campo: "Gtin Tributável", Tipo: "AN", Descricao: "Tamanho máximo de 20 caracteres.", Value: fields[24]},
		{Seq: 26, Campo: "ID ICMS", Tipo: "N", Descricao: "Chave estrangeira para a tabela ICMS. Máximo de 6 dígitos.", Value: fields[25]},
		{Seq: 27, Campo: "ID IPI", Tipo: "N", Descricao: "Chave estrangeira para a tabela IPI. Máximo de 6 dígitos.", Value: fields[26]},
		{Seq: 28, Campo: "ID ISSQN", Tipo: "N", Descricao: "Chave estrangeira para a tabela ISSQN. Máximo de 6 dígitos.", Value: fields[27]},
		{Seq: 29, Campo: "ID II", Tipo: "N", Descricao: "Chave estrangeira para a tabela II. Máximo de 6 dígitos.", Value: fields[28]},
		{Seq: 30, Campo: "ID PIS", Tipo: "N", Descricao: "Chave estrangeira para a tabela PIS. Máximo de 6 dígitos.", Value: fields[29]},
		{Seq: 31, Campo: "ID PIS ST", Tipo: "N", Descricao: "Chave estrangeira para a tabela PIS ST. Máximo de 6 dígitos.", Value: fields[30]},
		{Seq: 32, Campo: "ID COFINS", Tipo: "N", Descricao: "Chave estrangeira para a tabela COFINS. Máximo de 6 dígitos.", Value: fields[31]},
		{Seq: 33, Campo: "ID COFINS ST", Tipo: "N", Descricao: "Chave estrangeira para a tabela COFINS ST. Máximo de 6 dígitos.", Value: fields[32]},
		{Seq: 34, Campo: "NCM", Tipo: "AN", Descricao: "Este campo se torna obrigatório para a lei dos impostos 12.741/2012. Tamanho\nmáximo de 20 caracteres. Ver Notas.", Value: fields[33]},
		{Seq: 35, Campo: "KIT", Tipo: "AN", Descricao: "Informe 'S' caso este produto seja um Kit, do contrário informe \"N\".", Value: fields[34]},
		{Seq: 36, Campo: "Prazo de Devolução", Tipo: "AN", Descricao: "Prazo, em dias, para devolução de um item específico. '0' para sem prazo de\ndevolução, '-1' para não permitir a devolução.", Value: fields[35]},
		{Seq: 37, Campo: "CEST", Tipo: "N", Descricao: "Obrigatório para produtos com substituição tributária a partir de 01/04/2016.", Value: fields[36]},
		{Seq: 38, Campo: "Controle de Estoque", Tipo: "AN", Descricao: "'S' ou ' '(vazio) para controlar o estoque e 'N' para não controlar.", Value: fields[37]},
		{Seq: 39, Campo: "ANP", Tipo: "N", Descricao: "Código ANP (Agência Nacional de Petróleo)", Value: fields[38]},
		{Seq: 40, Campo: "Dupla Pesagem", Tipo: "AN", Descricao: "Informe 'S' para habilitar a checagem de peso, caso seja um produto pesável.", Value: fields[39]},
		{Seq: 41, Campo: "Margem de Segurança", Tipo: "N", Descricao: "Porcentagem da margem de segurança para a liberação do item na checagem de\npeso. Com 2 casa decimais.", Value: fields[40]},
		{Seq: 42, Campo: "Indicador de Escala", Tipo: "A", Descricao: "'S' - Escala Relevante\n'N' - Escala Não Relevante.\nInformar'' (vazio) se não for utilizar.", Value: fields[41]},
		{Seq: 43, Campo: "CNPJ do Fabricante", Tipo: "AN", Descricao: "CNPJ do Fabricante da Mercadoria", Value: fields[42]},
		{Seq: 44, Campo: "Código do Benefício Fiscal", Tipo: "AN", Descricao: "", Value: fields[43]},
		{Seq: 45, Campo: "Percentual de GLP", Tipo: "N", Descricao: "Percentual do GLP derivado de petróleo do produto GLP. Com 4 casas\ndecimais.", Value: fields[44]},
		{Seq: 46, Campo: "Percentual de GNN", Tipo: "N", Descricao: "Percentual de Gás Natural Nacional. Com 4 casas decimais.", Value: fields[45]},
		{Seq: 47, Campo: "Percentual de GNI", Tipo: "N", Descricao: "Percentual de Gás Natural Importado. Com 4 casas de cimais.", Value: fields[46]},
		{Seq: 48, Campo: "Valor de Partilha", Tipo: "N", Descricao: "Valor do quilograma sem ICMS. Com até 13 casas inteiras e 2 casas\ndecimais.", Value: fields[47]},
		{Seq: 49, Campo: "Tipo Desconto", Tipo: "A", Descricao: "Informa qual é o tipo do desconto do produto informado no campo 50.\n'V' - Desconto em Valor\n'P' - Desconto em Porcentagem.", Value: fields[48]},
		{Seq: 50, Campo: "Desconto", Tipo: "N", Descricao: "Desconto para este produto, com 3 casas decimais. (Este é\ndo cadastro do produto e será aplicado em todas as vendas com\neste produto, não apenas na DAV.)\nUtilizar o campo 49 para informar se o desconto é por valor ou por porcentagem.\nPor padrão é porcentagem.\nUtiliza o desconto apenas se configurado nos Planos de Pagamento.", Value: fields[49]},
		{Seq: 51, Campo: "Unidade Tributável", Tipo: "A", Descricao: "Unidade tributável do produto. (Usar no máximo 4 caracteres).", Value: fields[50]},
		{Seq: 52, Campo: "Quantidade Tributável", Tipo: "N", Descricao: "Quantidade Tributável. Com até 11 casas inteiras e 4 casas decimais.", Value: fields[51]},
		{Seq: 53, Campo: "Cód. Externo Grupo Impressão", Tipo: "S", Descricao: "Código externo da impressora utilizada para integração com o retaguarda.", Value: fields[52]},
		{Seq: 54, Campo: "Desconto Máximo", Tipo: "N", Descricao: "Valor máximo de desconto permitdo para o produto em uma venda, com 3 cassa\ndecimais.\nVer notas.", Value: fields[53]},
	}

}
