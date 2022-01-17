package duplicate_encode_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestDuplicateEncode(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "duplicate_encode Suite")
}
