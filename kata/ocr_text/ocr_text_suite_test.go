package ocr_text_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestOcrText(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "ocr_text Suite")
}
