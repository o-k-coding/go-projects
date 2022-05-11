package min_remove_valid_parens_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestMinRemoveValidParens(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "min_remove_valid_parens Suite")
}
