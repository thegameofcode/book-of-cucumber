package parseFeature
import (
	"io/ioutil"
	"os"
)

func FromFileName(fileName string) (err error) {
	inputFileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	openedFile, err := os.Create("./output.tex")
	_, err = openedFile.WriteString(assembleTexFile(string(inputFileContent)))
	openedFile.Close()

	return err
}

func assembleTexFile (inputFileContent string) (string) {
	return "\\section{" + inputFileContent + "}"
}
