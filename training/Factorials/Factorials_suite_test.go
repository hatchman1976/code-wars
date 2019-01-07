package kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFactorials(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Factorials Suite")
}
