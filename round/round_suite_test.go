package round_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRound(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Round Suite")
}
