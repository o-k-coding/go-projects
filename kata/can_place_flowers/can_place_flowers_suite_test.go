package can_place_flowers_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestCanPlaceFlowers(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "can_place_flowers Suite")
}
