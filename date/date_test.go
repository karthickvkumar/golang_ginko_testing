package date_test

import (
	. "testing/date"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Date", func() {
	Describe("Converting Date Format", func() {
		Context("Checking behaviour of function", func() {
			It("should work correct", func() {
				Expect(Years()).To(Equal(2009))
			})
		})
		Context("Checking behaviour of function", func() {
			It("should work correct", func() {
				Expect(Format()).To(Equal("12 PM"))
			})
		})
	})
})
