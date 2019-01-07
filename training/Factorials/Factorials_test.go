package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"."
)

func dotest(n int, exp float64) {
    var ans = kata.Going(n)
    Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

    It("should handle basic cases", func() {
        dotest(5, 1.275)
        dotest(6, 1.2125)
        dotest(7, 1.173214)
        dotest(8, 1.146651)
        dotest(200, 1.005025)
    })
})