package climbing_stairs_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestClimbingStairs(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "climbing_stairs Suite")
}
