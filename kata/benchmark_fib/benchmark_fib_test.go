package benchmark_fib_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/benchmark_fib"
)


// Add test cases here

var _ = Describe("Basic tests", func() {
	It("Should find the correct answer", func() {
		Expect(BenchmarkFib("")).To(Equal(""))
	})
})
