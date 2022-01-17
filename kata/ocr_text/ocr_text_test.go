package ocr_text_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/ocr_text"
)

// Add test cases here
var _ = Describe("Basic tests", func() {
	It("Should find the correct answer", func() {
		Expect(OcrText("A2Le", "2PL1")).To(Equal(true))
		Expect(OcrText("a10", "10a")).To(Equal(true))
		Expect(OcrText("ba1", "1Ad")).To(Equal(false))
		Expect(OcrText("3x2x", "8")).To(Equal(false))
	})
})
