package maximize_person_distance_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestMaximizePersonDistance(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "maximize_person_distance Suite")
}
