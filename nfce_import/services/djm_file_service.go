package service

import (
	"fmt"
	"nfceimport/repositories"
)

func ReadFileFromDjm() {
	djm := repositories.GetLinesFromFile("C:/Via/pdv/vendas/00299667.djm")
	fmt.Println(djm)
}
