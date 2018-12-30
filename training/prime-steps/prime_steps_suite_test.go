package kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPrimeSteps(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PrimeSteps Suite")
}
