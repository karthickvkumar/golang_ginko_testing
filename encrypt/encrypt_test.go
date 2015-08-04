package encrypt_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "testing/encrypt"
)

var _ = Describe("Encrypt", func() {
	Describe("Tax Calculation", func() {
		Context("People Category", func() {
			It("should work correct", func() {
				Expect(EncryptAes("Go Lang")).To(Equal("Go Lang"))

			})
		})
	})
})
