package service

import (
	"fmt"

	model "nfceimport/models"
)

func GravarDocRV(DOC model.RegistroDoc) {

	fmt.Println(DOC.Documento["SERIE"])
}
