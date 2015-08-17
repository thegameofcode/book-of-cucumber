package createTexFile

import (
	"os"
)

func WithData (inputData string) (error){
	fileName := "./output.tex"

	openedFile, err := os.Create(fileName)
	if err != nil {
		return err
	}

	_, err = openedFile.WriteString(inputData)

	openedFile.Close()
	return err
}
