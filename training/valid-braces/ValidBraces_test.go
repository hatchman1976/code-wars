package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
"fmt"
	"."
)

func singleTest(str string, res bool) {
    It(fmt.Sprintf("should return %v for \"%v\"",res,str), func() {
      Expect(kata.ValidBraces(str)).To(Equal(res))
    })
}

var _ = Describe("Valid Braces", func() { 
    singleTest("(){}[]",true)
    singleTest("([{}])",true)
    singleTest("(}",false)
    singleTest("[(])",false)
    singleTest("[({)](]",false)
	singleTest("[(){]}",false)
	singleTest("(({{[[]]}}))",true)
	
	
})