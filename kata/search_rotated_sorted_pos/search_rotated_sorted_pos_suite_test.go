package search_rotated_sorted_pos_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestSearchRotatedSortedPos(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "search_rotated_sorted_pos Suite")
}
