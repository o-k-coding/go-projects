package search_range_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestSearchRange(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "search_range Suite")
}
