package kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOrderWeight(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OrderWeight Suite")
}
