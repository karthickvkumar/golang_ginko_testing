package ftoc_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFtoc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ftoc Suite")
}
