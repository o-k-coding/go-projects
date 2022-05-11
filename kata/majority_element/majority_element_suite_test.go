package majority_element_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestMajorityElement(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "majority_element Suite")
}
