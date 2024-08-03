package services

import (
	"fmt"
)

type Loja struct {
	LojCodigo                 string
	LojRazSocial              string
	LojEmiteNfce              string
	CupomBaixaEstoque         string
	CupomBaixaEstoquePedido   string
	LojCupomReprocessarFinanc string
	LojCnpj                   string
}

// RecuperarLojaNFCE recupera informações de uma loja específica que emite NFCE.
func GetLoja() (Loja, error) {
	sqlGetParams := `
        select 
            l.loj_codigo ,
            l.loj_raz_social ,
            l.loj_emite_nfce ,
            l.cupom_baixa_estoque_pedido ,
            l.cupom_baixa_estoque ,
            l.loj_cupom_reprocessar_financ ,
            l.loj_cnpj 
        from wsloja l
        where l.loj_emite_nfce = 'S'
    `

	fmt.Println(sqlGetParams)

	var param Loja
	/* 	for rows.Next() {
		err := rows.Scan(
			&param.LojCodigo,
			&param.LojRazSocial,
			&param.LojEmiteNfce,
			&param.CupomBaixaEstoquePedido,
			&param.CupomBaixaEstoque,
			&param.LojCupomReprocessarFinanc,
			&param.LojCnpj,
		)
		if err != nil {
			return Loja{}, err
		}
	} */

	return param, nil
}

func GetLoja2() {
	sqlGetParams := `
		select 
			l.loj_codigo ,
			l.loj_raz_social ,
			l.loj_emite_nfce ,
			l.cupom_baixa_estoque_pedido ,
			l.cupom_baixa_estoque ,
			l.loj_cupom_reprocessar_financ ,
			l.loj_cnpj 
		from wsloja l
		where l.loj_emite_nfce = 'S'
	`
	/* 	connection := getConection()
	   	rows, err := connection.Query(sqlGetParams)
	   	if err != nil {
	   		log.Fatalf("Erro ao abrir a consulta: %v", err)
	   	}
	   	defer rows.Close() */
	fmt.Println(sqlGetParams)
}
