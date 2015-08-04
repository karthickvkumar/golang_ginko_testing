package hello_test

import (
	. "testing/hello"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hello", func() {
	It("collects min information", func() {
		Expect(Add(2, 2)).To(Equal(4))
	})
})
