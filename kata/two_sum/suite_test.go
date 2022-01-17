package two_sum_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestTwoSum(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Two Sum Suite")
}
