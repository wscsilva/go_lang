package service

import "nfceimport/util"

func SaveDOC(path string) {
	util.GetLinesFromFile(path)
}
