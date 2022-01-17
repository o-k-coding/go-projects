package smallest_int_not_in_a_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestSmallestIntNotInA(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "smallest_int_not_in_a Suite")
}
