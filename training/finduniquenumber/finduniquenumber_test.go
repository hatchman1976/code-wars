package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"."
)

var _ = Describe("FindUniq", func() {
	It("should work for some basic cases", func() {
	  Expect(kata.FindUniq([]float32{ 1.0, 1.0, 1.0, 2.0, 1.0, 1.0 })).To(Equal(float32(2)))
	  Expect(kata.FindUniq([]float32{ 0, 0, 0.55, 0, 0  })).To(Equal(float32(0.55)))
	  Expect(kata.FindUniq([]float32{ 0, 1, 1, 1, 1  })).To(Equal(float32(0)))
	  Expect(kata.FindUniq([]float32{ 1, 0, 1, 1, 1  })).To(Equal(float32(0)))
	  Expect(kata.FindUniq([]float32{ 1, 1, 1, 1, 1, 0, 1 })).To(Equal(float32(0)))
	  Expect(kata.FindUniq([]float32{ 1, 1, 1, 1, 1, 1, 0 })).To(Equal(float32(0)))
	})
  })