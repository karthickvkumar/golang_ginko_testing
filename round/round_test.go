package round_test

import (
	. "testing/round"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Round", func() {
	Describe("Rounding Numbers", func() {
		Context("Floor Method", func() {
			It("should work correct", func() {
				Expect(Rounding("Floor", 12.9999)).To(Equal(12.0))
			})
		})
		Context("Ceil Method", func() {
			It("should work correct", func() {
				Expect(Rounding("Ceil", 15.00001)).To(Equal(16.0))
			})
		})
		Context("Un-Matched Method", func() {
			It("should work correct", func() {
				Expect(Rounding("Round", 1)).To(Equal(0.0))
			})
		})

	})
})
