package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"kata"
)



var _ = Describe("ValidIp", func() {
	It("should test that 12.255.56.1 is correct", func() {
		Expect(kata.Is_valid_ip("12.255.56.1")).To(Equal(false))
	  })
})
