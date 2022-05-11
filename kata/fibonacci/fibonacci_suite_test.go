package fibonacci_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestFibonacci(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "fibonacci Suite")
}
