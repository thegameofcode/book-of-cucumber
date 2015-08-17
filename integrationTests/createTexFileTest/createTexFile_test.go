package createTexFileTest

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/thegameofcode/book-of-cucumber/lib/createTexFile"
	"testing"
	"io/ioutil"
	"os"
)

var _ = Describe("Create tex file", func() {
	AfterEach(func(){
		err := os.Remove("./output.tex")
		Expect(err).To(BeNil())
	})

	It("A file should be created", func() {
		err := createTexFile.WithData("abc")
		Expect(err).To(BeNil())

		data, err := ioutil.ReadFile("./output.tex")

		Expect(err).To(BeNil())
		Expect(data).ToNot(BeNil())
	})

	It("Should write into the file", func() {
		err := createTexFile.WithData("abc")
		Expect(err).To(BeNil())

		data, err := ioutil.ReadFile("./output.tex")
		Expect(err).To(BeNil())
		Expect(string(data)).To(Equal("abc"))
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tex file creation")
}
