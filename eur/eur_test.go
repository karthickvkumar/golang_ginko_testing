package eur_test

import (
	. "testing/eur"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Eur", func() {
	Describe("Converting EUR Currency", func() {
		Context("Indian Rupee to Euro", func() {
			It("should work correct", func() {
				Expect(CurrencyEur("INR", 1)).To(Equal(0.014))
			})
		})
		Context("US Doller to Euro", func() {
			It("should work correct", func() {
				Expect(CurrencyEur("USD", 1)).To(Equal(0.90))
			})
		})
		Context("Russian to Euro", func() {
			It("should work correct", func() {
				Expect(CurrencyEur("RUB", 1)).To(Equal(0.016))
			})
		})
		Context("Mexican to Euro", func() {
			It("should work correct", func() {
				Expect(CurrencyEur("MXN", 1)).To(Equal(0.057))
			})
		})
		Context("Saudi Riyal to Euro", func() {
			It("should work correct", func() {
				Expect(CurrencyEur("SAR", 1)).To(Equal(0.24))
			})
		})
		Context("Un-Mateched currency", func() {
			It("should work correct", func() {
				Expect(CurrencyEur("EUR", 1)).To(Equal(0.0))
			})
		})
	})
})
