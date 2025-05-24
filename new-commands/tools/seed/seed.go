package seed

import (
	"fmt"
	"os"

	"github.com/ElianDev55/goapi/libs"
)

func Create(logger *libs.LoggerP ,content string, mainPath string) error {

	f, err := os.OpenFile(fmt.Sprintf("%s/tools/seed/seed-db-main.go", mainPath), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.ErrorLogger("ERROR OPENING THE FILE [seed-db-main.go]", err)
		return err
	}

	defer f.Close()
	
	if _, err := f.WriteString(content); err != nil {
		logger.ErrorLogger("ERROR WRITING TO THE FILE [seed-db-main.go]", err)
	} else {
		logger.InfoLogger("CONTENT ADDED SUCCESSFULLY [seed-db-main.go]")
	}
		return nil
}
