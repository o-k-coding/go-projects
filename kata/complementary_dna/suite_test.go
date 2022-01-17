package complementary_dna_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestComplementaryDna(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Complementary DNA Suite")
}
