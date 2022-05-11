package binary_tree_complete_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestBinaryTreeComplete(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "binary_tree_complete Suite")
}
