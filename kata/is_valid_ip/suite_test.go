package is_valid_ip_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIsValidIp(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Is Valid IP Suite")
}
