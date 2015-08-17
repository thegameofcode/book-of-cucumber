package parseFeatureTest

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/thegameofcode/book-of-cucumber/lib/parseFeature"
	"testing"
	"io/ioutil"
	"os"
)

var _ = Describe("Parse feature", func() {

	AfterEach(func(){
		err := os.Remove("./input.feature")
		Expect(err).To(BeNil())

		err = os.Remove("./output.tex")
		Expect(err).To(BeNil())
	})


	It("A file should be created", func() {
		openedFile, err := os.Create("./input.feature")
		Expect(err).To(BeNil())

		_, err = openedFile.WriteString("some feature description")
		openedFile.Close()
		Expect(err).To(BeNil())

		err = parseFeature.FromFileName("./input.feature")
		Expect(err).To(BeNil())

		data, err := ioutil.ReadFile("./output.tex")
		Expect(err).To(BeNil())
		Expect(string(data)).To(Equal("some feature description"))
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parse feature")
}
