package db

import (
	"fmt"
	"os"

	"github.com/ElianDev55/goapi/libs"
)

func Create(logger *libs.LoggerP ,content string, mainPath string) error {

	f, err := os.OpenFile(fmt.Sprintf("%s/pkg/db/db.go", mainPath), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.ErrorLogger("ERROR OPENING THE FILE [db.go]", err)
		return err
	}

	defer f.Close()
	
	if _, err := f.WriteString(content); err != nil {
		logger.ErrorLogger("ERROR WRITING FILE [DB.GO]",err)
	} else {
		logger.InfoLogger("CONTENT ADDED SUCCESSFULLY [db.go]")
	}
		return nil
}




