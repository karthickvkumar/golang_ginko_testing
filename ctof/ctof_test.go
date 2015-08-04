package ctof_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "testing/ctof"
)

var _ = Describe("Ctof", func() {
	Describe("Converting Celcius to Farnhet", func() {
		Context("Checking behaviour of function", func() {
			It("should work correct", func() {
				Expect(Celcius(0)).To(Equal(32.0))
			})
			It("should work correct", func() {
				Expect(Celcius(-15)).To(Equal(5.0))
			})
			It("should work correct", func() {
				Expect(Celcius(120)).To(Equal(248.0))
			})
		})
	})
})
