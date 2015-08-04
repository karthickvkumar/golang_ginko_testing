package tax_test

import (
	. "testing/tax"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tax", func() {
	Describe("Tax Calculation", func() {
		Context("People Category", func() {
			It("should work correct", func() {
				Expect(People(30)).To(Equal("Normal Citizen"))
				Expect(People(60)).To(Equal("Seniour Citizen"))
				Expect(People(82)).To(Equal("Super Seniour Citizen"))
			})
		})
		Context("Slap-1 Tax Calculation", func() {
			It("should work correct", func() {
				Expect(Taxs("GEN", 150000)).To(Equal(0.0))
				Expect(Taxs("GEN", 400000)).To(Equal(15000.0))
				Expect(Taxs("GEN", 700000)).To(Equal(85000.0))
				Expect(Taxs("GEN", 1200000)).To(Equal(295000.0))
			})
		})
		Context("Slap-2 Tax Calculation", func() {
			It("should work correct", func() {
				Expect(Taxs("SEN", 270000)).To(Equal(0.0))
				Expect(Taxs("SEN", 400000)).To(Equal(10000.0))
				Expect(Taxs("SEN", 700000)).To(Equal(80000.0))
				Expect(Taxs("SEN", 1200000)).To(Equal(290000.0))
			})
		})
		Context("Slap-3 Tax Calculation", func() {
			It("should work correct", func() {
				Expect(Taxs("SUP_SEN", 400000)).To(Equal(0.0))
				Expect(Taxs("SUP_SEN", 700000)).To(Equal(40000.0))
				Expect(Taxs("SUP_SEN", 1200000)).To(Equal(200000.0))
			})
		})
		Context("Un-Matched Category", func() {
			It("should work correct", func() {
				Expect(Taxs("NOR", 10000)).To(Equal(0.0))
			})
		})
	})
})
