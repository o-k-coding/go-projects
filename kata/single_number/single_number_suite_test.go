package single_number_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestSingleNumber(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "single_number Suite")
}
