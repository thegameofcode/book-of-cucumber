package parseFeatureTest

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/thegameofcode/book-of-cucumber/lib/parseFeature"
	"testing"
	"io/ioutil"
	"os"
	"strings"
)

var _ = Describe("Parse feature", func() {

	AfterEach(func() {
		if _, err := os.Stat("./input.feature"); err == nil {
			err = os.Remove("./input.feature")
			Expect(err).To(BeNil())
		}

		if _, err := os.Stat("./output.tex"); err == nil {
			err = os.Remove("./output.tex")
			Expect(err).To(BeNil())
		}
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

	FIt("Feature name is displayed as section in the tex file", func() {
		err := parseFeature.FromFileName("../../integrationTests/testFiles/feature/login.feature")
		Expect(err).To(BeNil())

		outputTexFileData, err := ioutil.ReadFile("./output.tex")
		Expect(err).To(BeNil())

		Expect(strings.Contains(string(outputTexFileData), "\\section{Feature: Login}")).To(BeTrue())
	})
})


func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parse feature")
}
