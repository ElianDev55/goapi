package example_env

import (
	"fmt"
	"os"

	"github.com/ElianDev55/goapi/libs"
)



func Create(content string, mainPath string) error {

	f, err := os.OpenFile(fmt.Sprintf("%s/.env.example", mainPath), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		libs.Error("ERROR OPENING THE FILE")
		return err
	}

	defer f.Close()
	
	if _, err := f.WriteString(content); err != nil {
		fmt.Println("ERROR WRITING TO THE FILE", err)
	} else {
		fmt.Println("CONTENT ADDED SUCCESSFULLY.")
	}
		return nil
}

func CreateGoModFile(content string, mainPath string) error {

	f, err := os.OpenFile(fmt.Sprintf("%s/go.mod", mainPath), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		libs.Error("ERROR OPENING THE FILE")
		return err
	}

	defer f.Close()
	
	if _, err := f.WriteString(content); err != nil {
		fmt.Println("ERROR WRITING TO THE FILE", err)
	} else {
		fmt.Println("CONTENT ADDED SUCCESSFULLY.")
	}
		return nil
}


func CreateGoSumFile(content string, mainPath string) error {

	f, err := os.OpenFile(fmt.Sprintf("%s/go.sum", mainPath), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		libs.Error("ERROR OPENING THE FILE")
		return err
	}

	defer f.Close()
	
	if _, err := f.WriteString(content); err != nil {
		fmt.Println("ERROR WRITING TO THE FILE", err)
	} else {
		fmt.Println("CONTENT ADDED SUCCESSFULLY.")
	}
		return nil
}
