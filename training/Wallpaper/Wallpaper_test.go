package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
"."
)
func dotest(l, w, h float64, exp string) {
    var ans = kata.WallPaper(l, w, h)
    Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests WallPaper", func() {

    It("should handle basic cases", func() {
        dotest(6.3, 4.5, 3.29, "sixteen")        
        dotest(6.3, 5.8, 3.13, "seventeen")
        
    })
})
