package search_2d_matrix_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestSearch2dMatrix(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "search_2d_matrix Suite")
}
