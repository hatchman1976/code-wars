package kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNumberMonths(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NumberMonths Suite")
}
