package date_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Date Suite")
}
