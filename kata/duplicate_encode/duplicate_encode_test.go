package duplicate_encode_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/duplicate_encode"
)

// Add test cases here

var _ = Describe("Test Example", func() {
  It("should return the correct value", func() {
    Expect(DuplicateEncode("din")).To(Equal("((("))
    Expect(DuplicateEncode("recede")).To(Equal("()()()"))
    Expect(DuplicateEncode("(( @")).To(Equal("))(("))
  })

  It("should ignore case", func() {
    Expect(DuplicateEncode("Success")).To(Equal(")())())"))
  })
})
