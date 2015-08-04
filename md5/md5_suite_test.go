package md5_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMd5(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Md5 Suite")
}
