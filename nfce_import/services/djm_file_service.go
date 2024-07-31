package service

import (
	model "nfceimport/models"
	repository "nfceimport/repositories"
)

func ReadFileFromDjm() {
	registrosDJM := repository.GetLinesFromFile("C:/Via/pdv/vendas/00299667.djm")

	for _, registro := range registrosDJM {
		/// Obtem o tipo de objeto da interface
		switch d := registro.(type) {
		/// Obtem o tipo DOC
		case model.GetRegistrosRV:

			if d.DENOMINACAO == "RV" {
				GravarDocRV(d)
			}

		}
	}

}
