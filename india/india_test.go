package india_test

import (
	. "testing/india"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("India", func() {
	Describe("Converting Indian Currency", func() {
		Context("US Doller to Indian Rupee", func() {
			It("should work correct", func() {
				Expect(CurrencyInr("USD", 1)).To(Equal(64.10))
			})
		})
		Context("Euro to Indian Rupee", func() {
			It("should work correct", func() {
				Expect(CurrencyInr("EUR", 1)).To(Equal(71.14))
			})
		})
		Context("Russian to Indian Rupee", func() {
			It("should work correct", func() {
				Expect(CurrencyInr("RUB", 1)).To(Equal(1.13))
			})
		})
		Context("Mexican to Indian Rupee", func() {
			It("should work correct", func() {
				Expect(CurrencyInr("MXN", 1)).To(Equal(4.08))
			})
		})
		Context("Saudi Riyal to Indian Rupee", func() {
			It("should work correct", func() {
				Expect(CurrencyInr("SAR", 1)).To(Equal(18.0))
			})
		})
		Context("Indian code", func() {
			It("should work correct", func() {
				Expect(CurrencyInr("INR", 1)).To(Equal(0.0))
			})
		})
	})
})
