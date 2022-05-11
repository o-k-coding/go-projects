package binary_search_tree_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestBinarySearchTree(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "binary_search_tree Suite")
}
