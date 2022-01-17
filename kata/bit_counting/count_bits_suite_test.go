package bit_counting_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestAdder(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Bit Counting Suite")
}
