package example_env

import (
	"fmt"
	"os"

	"github.com/ElianDev55/goapi/libs"
)




func Create(logger *libs.LoggerP ,content string, mainPath string) error {

	f, err := os.OpenFile(fmt.Sprintf("%s/.env.example", mainPath), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.ErrorLogger("ERROR OPENING THE FILE [.env.example]", err)
		return err
	}

	defer f.Close()
	
	if _, err := f.WriteString(content); err != nil {
		logger.ErrorLogger("ERROR WRITING TO THE FILE [.env.example]", err)
	} else {
		logger.InfoLogger("CONTENT ADDED SUCCESSFULLY [.env.example]")
	}
		return nil
}

func CreateGoModFile(logger *libs.LoggerP ,content string, mainPath string) error {

	f, err := os.OpenFile(fmt.Sprintf("%s/go.mod", mainPath), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.ErrorLogger("ERROR OPENING THE FILE [go.mod]", err)
		return err
	}

	defer f.Close()
	
	if _, err := f.WriteString(content); err != nil {
		logger.ErrorLogger("ERROR WRITING TO THE FILE [go.mod]", err)
	} else {
		logger.InfoLogger("CONTENT ADDED SUCCESSFULLY [go.mod]")
	}
		return nil
}


func CreateGoSumFile(logger *libs.LoggerP ,content string, mainPath string) error {

	f, err := os.OpenFile(fmt.Sprintf("%s/go.sum", mainPath), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.ErrorLogger("ERROR OPENING THE FILE [go.sum]", err)
		return err
	}

	defer f.Close()
	
	if _, err := f.WriteString(content); err != nil {
		logger.ErrorLogger("ERROR WRITING TO THE FILE [go.sum]", err)
	} else {
		logger.InfoLogger("CONTENT ADDED SUCCESSFULLY.")
	}
		return nil
}
