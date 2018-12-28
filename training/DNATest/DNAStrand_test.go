package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	  
)

import "./kata"

var _ = Describe("DNAStrand", func() {
	It("Basic Tests", func() {
		Expect(kata.DNAStrand("AAAA")).To(Equal("TTTT"))
		Expect(kata.DNAStrand("ATTGC")).To(Equal("TAACG"))
		Expect(kata.DNAStrand("GTAT")).To(Equal("CATA"))
	  })
})
