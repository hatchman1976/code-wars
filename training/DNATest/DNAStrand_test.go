package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	  
)

import "./kata"

var _ = Describe("DNAStrand", func() {
	It("Basic Tests", func() {
		Expect(kata.DNAStrand("AAAA")).To(Equal("AAAA"))
	  })
})
