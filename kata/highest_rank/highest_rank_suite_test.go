package highest_rank_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHighestRank(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "highest_rank Suite")
}
