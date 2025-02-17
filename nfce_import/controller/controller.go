package controller

import (
	"fmt"
	model "nfceimport/models"
	repository "nfceimport/repositories"
	"nfceimport/services"
)

func ReadFileFromDjm() {
	registrosDJM := repository.GetRegistryFromDJMFile("C:/Via/pdv/vendas/00299667.djm")

	for _, registro := range registrosDJM {
		ecf := registrosDJM[2]
		/// Obtem o tipo de objeto da interface
		switch d := registro.(type) {

		/// Obtem o tipo DOC
		case model.GetRegistrosRV:
			r_ecf := ecf.(model.RegistroECF)

			fmt.Println(r_ecf.NUMEROSERIE)
			if d.DENOMINACAO == "RV" {
				services.GravarDocRV(d)
			}

		}
	}

}
