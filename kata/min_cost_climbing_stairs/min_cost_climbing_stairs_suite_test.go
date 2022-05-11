package min_cost_climbing_stairs_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "min_cost_climbing_stairs Suite")
}
