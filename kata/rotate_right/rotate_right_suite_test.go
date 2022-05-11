package rotate_right_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestRotateRight(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "rotate_right Suite")
}
