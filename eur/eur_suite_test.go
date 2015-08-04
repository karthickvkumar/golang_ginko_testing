package eur_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEur(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Eur Suite")
}
