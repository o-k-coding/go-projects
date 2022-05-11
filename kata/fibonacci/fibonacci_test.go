package fibonacci_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/fibonacci"
)

// Add test cases here

var _ = Describe("Basic tests", func() {
	It("Should find the correct answer", func() {
		Expect(Fibonacci(2)).To(Equal(1))
		Expect(Fibonacci(3)).To(Equal(2))
		Expect(Fibonacci(4)).To(Equal(3))
	})
})
