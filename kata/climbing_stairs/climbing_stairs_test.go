package climbing_stairs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/climbing_stairs"
)

// Add test cases here

var _ = Describe("Basic tests", func() {
	It("Should find the correct answer", func() {
		Expect(ClimbingStairs(2)).To(Equal(2))
		Expect(ClimbingStairs(3)).To(Equal(3))
	})
})
