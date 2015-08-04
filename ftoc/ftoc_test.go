package ftoc_test

import (
	. "testing/ftoc"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ftoc", func() {
	Describe("Converting Farnheit to Celcius", func() {
		Context("Checking behaviour of function", func() {
			It("should work correct", func() {
				Expect(Farnheit(32)).To(Equal(0.0))
			})
			It("should work correct", func() {
				Expect(Farnheit(5)).To(Equal(-15.0))
			})
			It("should work correct", func() {
				Expect(Farnheit(248)).To(Equal(120.0))
			})
		})
	})

})
