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

	// Verificar se o cupom ja foi importado
	cupom, err := getCupom(DOC)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(cupom.NfceNumero)

}

// Retorna o cupom se existir
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
	/* 	connection, err := db.ConectarBancoDeDados()
	   	if err != nil {
	   		return Cupom{}, err
	   	}
	   	defer db.FecharConexao(connection)



	   	// Execute a consulta
	   	rows, err := connection.Query(sql, DOC.NUMERO, DOC.SERIE) */

	// Verifique se os parâmetros são válidos
	if DOC.NUMERO == "" || DOC.SERIE == "" {
		return Cupom{}, errors.New("parâmetros inválidos")
	}
	rows, err := db.ExecuteQuery(sql, DOC.NUMERO, DOC.SERIE)
	if err != nil {
		return Cupom{}, err
	}
	defer rows.Close()

	if rows.Next() {
		return Cupom{}, fmt.Errorf("NFCe número: %s com a Série: %s já cadastrado no sistema", DOC.NUMERO, DOC.SERIE)

	}

	// Retorne o cupom
	var cupom Cupom

	err = rows.Scan(
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

	return cupom, nil
}
