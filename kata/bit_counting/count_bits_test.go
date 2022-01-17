package bit_counting_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/bit_counting"
)

var _ = Describe("CountBits", func() {
    It("should count bits", func() {
			Expect(CountBits(0)).To(Equal(0))
			Expect(CountBits(4)).To(Equal(1))
			Expect(CountBits(7)         ).To(Equal(3))
			Expect(CountBits(9)         ).To(Equal(2))
			Expect(CountBits(10)        ).To(Equal(2))
    })
})
