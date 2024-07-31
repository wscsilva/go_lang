package service

import (
	"fmt"

	model "nfceimport/models"
)

func GravarDocRV(DOC model.GetRegistrosRV) {

	fmt.Printf("Data: %s COO: %s | Série: %s\n", DOC.DATA_INICIO, DOC.COO, DOC.SERIE)
}
