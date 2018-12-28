package utils

import (
	seelog "github.com/cihub/seelog"
)

/*
InitLogger initial a logger by seelog
Config file ./conf/seelog.xml
*/
func InitLogger(path string) error {

	defer seelog.Flush()

	logger, err := seelog.LoggerFromConfigAsFile(path)
	if err != nil {
		// log.Println("Parse Seelog Config File Error ...")
		return err
	}
	seelog.ReplaceLogger(logger)

	return nil

}
