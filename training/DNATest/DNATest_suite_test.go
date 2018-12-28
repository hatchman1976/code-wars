package kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDNATest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DNATest Suite")
}
