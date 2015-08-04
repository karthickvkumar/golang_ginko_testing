package tax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTax(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tax Suite")
}
