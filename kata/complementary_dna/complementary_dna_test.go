package complementary_dna_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/complementary_dna"
)


var _ = Describe("DNAStrand", func() {
  It("Basic Tests", func() {
    Expect(DNAStrand("AAAA")).To(Equal("TTTT"))
    Expect(DNAStrand("ATTGC")).To(Equal("TAACG"))
    Expect(DNAStrand("GTAT")).To(Equal("CATA"))
  })
})
