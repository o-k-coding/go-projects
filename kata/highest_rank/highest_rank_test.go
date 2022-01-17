package highest_rank_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/highest_rank"
)



var _ = Describe("Tests", func() {
    Describe("Sample tests", func() {
        It("Sample test 1: 12, 10, 8, 12, 7, 6, 4, 10, 12", func() {
            Expect(HighestRank([]int{12, 10, 8, 12, 7, 6, 4, 10, 12 })).To(Equal(12))
        })
    })
})
