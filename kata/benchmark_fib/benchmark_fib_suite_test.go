package benchmark_fib_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestBenchmarkFib(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "benchmark_fib Suite")
}
