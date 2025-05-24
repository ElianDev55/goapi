package main_file

import (
	"fmt"
	"os"

	"github.com/ElianDev55/goapi/libs"
)

func Create(logger *libs.LoggerP ,content string, mainPath string) error {

	f, err := os.OpenFile(fmt.Sprintf("%s/main.go", mainPath), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.ErrorLogger("ERROR OPENING THE FILE [main.go]", err)
		return err
	}

	defer f.Close()
	
	if _, err := f.WriteString(content); err != nil {
		logger.ErrorLogger("ERROR WRITING FILE [main.GO]",err)
	} else {
		logger.InfoLogger("CONTENT ADDED SUCCESSFULLY [main.go]")
	}
		return nil
}




