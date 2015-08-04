package usd_test

import (
	. "testing/usd"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Usd", func() {
	Describe("Converting USD Currency", func() {
		Context("Indian Rupee to US Doller", func() {
			It("should work correct", func() {
				Expect(CurrencyUsd("INR", 1)).To(Equal(0.016))
			})
		})
		Context("Euro to US Doller", func() {
			It("should work correct", func() {
				Expect(CurrencyUsd("EUR", 1)).To(Equal(1.11))
			})
		})
		Context("Russian to US Doller", func() {
			It("should work correct", func() {
				Expect(CurrencyUsd("RUB", 1)).To(Equal(62.56))
			})
		})
		Context("Mexican to US Doller", func() {
			It("should work correct", func() {
				Expect(CurrencyUsd("MXN", 1)).To(Equal(17.46))
			})
		})
		Context("Saudi Riyal to US Doller", func() {
			It("should work correct", func() {
				Expect(CurrencyUsd("SAR", 1)).To(Equal(4.17))
			})
		})
		Context("Un-Mateched currency", func() {
			It("should work correct", func() {
				Expect(CurrencyUsd("USD", 1)).To(Equal(0.0))
			})
		})
	})
})
