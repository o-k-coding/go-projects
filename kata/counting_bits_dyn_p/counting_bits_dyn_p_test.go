package counting_bits_dyn_p_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/counting_bits_dyn_p"
)

// Add test cases here

var _ = Describe("Basic tests", func() {
	It("ConvertBinary should convert to binary", func() {
		Expect(OneCountBinary(5, make(map[int]int))).To(Equal(2))
	})

	It("CountingBitsDynP should count the correct number of ones", func() {
		Expect(CountingBitsDynP(5)).To(Equal([]int{0,1,1,2,1,2}))
	})
})
