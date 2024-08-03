package services

import (
	"database/sql"
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

	// Consultar WSCUPOM
	rows, err := consultarWSCupom(DOC)
	if err != nil {
		log.Fatalf("Erro na consulta: %s", err)
	}
	defer rows.Close()

	// Verificar se existem registros
	if db.ExistemRegistros(rows) {
		log.Println("Existem registros")
		var cupom Cupom
		err := rows.Scan(
			&cupom.Cupom,
			&cupom.Caixa,
			&cupom.Data,
			&cupom.Cancelado,
			&cupom.NfceNumero,
			&cupom.NfceSerie,
		)
		if err != nil {
			log.Println(err)
		}
	} else {
		importarNFce(DOC)
	}

}

func importarNFce(DOC model.GetRegistrosRV) (int64, error) {
	sql := `
		INSERT INTO rows (
			cupom, 
			caixa, 
			"data", 
			valor, 
			cpf_cnpj, 
			nome_cliente, 
			cancelado, 
			data_cancelamento, 
			loja, 
			desconto, 
			dinheiro, 
			acrescimo, 
			cli_codigo, 
			perc_desc, 
			arquivo, 
			perc_acrescimo, 
			orc_codigo, 
			itens_conferidos, 
			observacao, 
			nfce_modelo, 
			nfce_tipo_emissao, 
			nfce_numero, 
			nfce_serie, 
			nfce_chave_dfe, 
			nfce_xml, 
			coo, 
			num_ecf, 
			cod_terminal, 
			cod_ecf, 
			protocolo, 
			vbc, 
			vicms, 
			vnf, 
			vtottrib, 
			val_base_icms, 
			val_icms, 
			baixou_estoque, 
			tip_cancelamento
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38
		);
	`
	registros, err := db.ExecutarComandos(sql,
		DOC.COO,
		DOC.CODECF,
		DOC.DATA_FIM,
		DOC.SUBTOTAL,
		DOC.DOCUMENTO_CLI,
		DOC.NOME_CLI,
		DOC.CANCELADO,
		DOC.DATA_FIM,
	)
	if err != nil {
		return registros, err
	}
	fmt.Printf("NFCe número: %s com a Série: %s importado com sucesso.", DOC.NUMERO, DOC.SERIE)
	return registros, nil
}

// Consulta WSCUPOM
// Retorna um ResultSet
func consultarWSCupom(DOC model.GetRegistrosRV) (*sql.Rows, error) {
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
	// Verifique se os parâmetros são válidos
	if DOC.NUMERO == "" || DOC.SERIE == "" {
		return nil, errors.New("parâmetros inválidos")
	}

	rows, err := db.AbrirConsulta(sql, DOC.NUMERO, DOC.SERIE)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
