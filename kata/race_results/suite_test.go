package race_results_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestRaceResults(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Complementary DNA Suite")
}
