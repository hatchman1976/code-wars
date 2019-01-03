package kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestValidBraces(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ValidBraces Suite")
}
