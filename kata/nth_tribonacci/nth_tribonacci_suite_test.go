package nth_tribonacci_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNthTribonacci(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "nth_tribonacci Suite")
}
