package md5_test

import (
	. "testing/md5"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Md5", func() {
	Describe("MD5 String Checksum", func() {
		Context("Return MD5 value", func() {
			It("should work correct", func() {
				Expect(MdString("Go Lang")).To(Equal("8a0190b38a11d5118efbb8b4671c55e8"))
			})
		})
	})
})
