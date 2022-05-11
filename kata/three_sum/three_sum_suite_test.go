package three_sum_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestThreeSum(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "three_sum Suite")
}
