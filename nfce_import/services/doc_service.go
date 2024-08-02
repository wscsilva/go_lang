package services

import (
	"errors"
	"fmt"
	"log"

	"nfceimport/db"
	model "nfceimport/models"
)

type Cupom struct {
	Cupom      string
	Caixa      string
	Data       string
	Cancelado  string
	NfceNumero string
	NfceSerie  string
}

func GravarDocRV(DOC model.GetRegistrosRV) {
	//data, err := util.ConvertStringToDate(DOC.DATA_INICIO, constants.FORMATO_DATA)
	/* 	if err != nil {
		fmt.Println(err)
	} */
	cupom, err := getCupom(DOC)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(cupom.NfceNumero)

	//dataFormatada := data.Format(constants.FORMATO_DATA_SAIDA)
	//fmt.Printf("Data: %s COO: %s | Série: %s\n", dataFormatada, DOC.COO, DOC.SERIE)
}

func getCupom(DOC model.GetRegistrosRV) (Cupom, error) {
	sql := `
        select 
            c.cupom ,
            c.caixa ,
            c."data" ,
            c.cancelado ,
            c.nfce_numero ,
            c.nfce_serie 
        from wscupom c
        where c.nfce_numero = $1
        and c.nfce_serie = $2    
    `
	connection := getConection()
	defer db.FecharConexao(connection)

	rows, err := connection.Query(sql, DOC.NUMERO, DOC.SERIE)
	if err != nil {
		return Cupom{}, err
	}
	defer rows.Close()

	var cupom Cupom
	if rows.Next() {
		err := rows.Scan(
			&cupom.Cupom,
			&cupom.Caixa,
			&cupom.Data,
			&cupom.Cancelado,
			&cupom.NfceNumero,
			&cupom.NfceSerie,
		)
		if err != nil {
			return Cupom{}, err
		}
	} else {
		return Cupom{}, errors.New("nenhum resultado encontrado")
	}

	return cupom, nil
}
