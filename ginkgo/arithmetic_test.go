package arithmetic_test

import (
	"github.com/kaleo211/golang-demo/ginkgo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Arithmetic", func() {
	Describe("Double", func() {
		Context("when input is a integer", func() {
			It("returns a value 2 times the input without error", func() {
				input := 4
				result, err := arithmetic.Double(input)
				Expect(err).ToNot(HaveOccurred())
				Expect(result).To(Equal(input * 2))
			})
		})

		Context("when input is a string", func() {
			It("returns a string that has 2 copies of origin without error", func() {
				input := "4"
				result, err := arithmetic.Double(input)
				Expect(err).ToNot(HaveOccurred())
				Expect(result).To(Equal(input + input))
			})
		})
	})
})
