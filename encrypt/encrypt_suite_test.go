package encrypt_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEncrypt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Encrypt Suite")
}
