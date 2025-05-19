package libs

import (
	"fmt"
	"os"
)


func CreateDirectories(listD []string) error {
	for _, dir := range listD {
		errorDir := os.Mkdir(dir, 0755)
		if errorDir != nil {
			return errorDir
		}
	}
	return nil
}

func CreateFiles(listF []string) error {
	fmt.Println("Creating files")
	for _, file := range listF {
		_, err := os.Create(file)

		if err != nil {
			return err
		}
	}
	
	return nil
}


func AddPathTo(initialPath string, args []string) []string {

	newList := []string{}

	for _, file := range args {
		newPath := fmt.Sprintf("%s/%s", initialPath, file)
		newList = append(newList, newPath)
	}

	return newList
}
