package testing_Ginkgo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	testgo "cursos-go/testingGinkgo"
	"testing"
)

func TestTestingGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestingGinkgo Suite")
}

var _ = Describe("Person.IsChild()", func() {
	Context("when the person is child", func() {
		It("returns True", func() {
			person := testgo.Person{Age: 10}

			response := person.IsChild()

			Expect(response).To(BeTrue())
		})
	})
})
