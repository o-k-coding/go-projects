package counting_bits_dyn_p_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestCountingBitsDynP(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "counting_bits_dyn_p Suite")
}
