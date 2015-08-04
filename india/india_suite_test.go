package india_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestIndia(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "India Suite")
}
