package parseFeature
import (
	"io/ioutil"
	"os"
)

func FromFileName(fileName string) (err error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	openedFile, err := os.Create("./output.tex")
	_, err = openedFile.WriteString(string(data))
	openedFile.Close()

	return err
}
