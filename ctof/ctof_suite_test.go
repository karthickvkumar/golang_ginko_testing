package ctof_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCtof(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ctof Suite")
}
