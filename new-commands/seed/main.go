package seed

import (
	"fmt"
	"os"
)

func Create(content string) error {
	
	f, err := os.OpenFile("pkg/db/seed-db-main.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()
	
	if _, err := f.WriteString(content); err != nil {
		fmt.Println("Error al escribir:", err)
		} else {
			fmt.Println("Contenido agregado correctamente.")
		}
		return nil
}
