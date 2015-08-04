package usd_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUsd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Usd Suite")
}
