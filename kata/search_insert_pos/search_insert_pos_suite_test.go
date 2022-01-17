package search_insert_pos_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestSearchInsertPos(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "search_insert_pos Suite")
}
