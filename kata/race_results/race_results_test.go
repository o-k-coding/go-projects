package race_results_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/race_results"
)

func dotest(a1 string, exp string) {
	var ans = Stati(a1)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests Stati", func() {

	It("should handle basic cases", func() {
		dotest("01|15|59, 1|47|16, 01|17|20, 1|32|34, 2|17|17",
			"Range: 01|01|18 Average: 01|38|05 Median: 01|32|34")
		dotest("02|15|59, 2|47|16, 02|17|20, 2|32|34, 2|17|17, 2|22|00, 2|31|41",
			"Range: 00|31|17 Average: 02|26|18 Median: 02|22|00")
	})

  It("should handle random cases", func() {
    dotest("12|17|48, 01|15|59, 00|19|34, 02|32|34, 11|15|17, 12|20|14, 12|17|48, 01|32|34",
      "Range: 12|00|40 Average: 06|43|58 Median: 06|53|55")
  })

  It("should handle 0 cases", func() {
    ans := Stati("0|15|59, 0|16|16, 0|17|20, 0|22|34, 0|19|34, 0|15|0")
    Expect(ans).NotTo(Equal(""))
  })

  It("should calculateMedian properly", func() {
    times := []int{
      44268, 4559, 1174, 9154, 40517, 44414, 44268, 5554,
    }

    ans := CalculateMedian(times)
    Expect(ans).To(Equal("06|53|55"))
  })

	It("should convert time string to seconds with padding", func() {
		ans := ParseTimeIntoSeconds("01|15|59")
		Expect(ans).To(Equal(4559))
	})

	It("should convert time string to seconds without padding", func() {
		ans := ParseTimeIntoSeconds("1|47|16.14")
		Expect(ans).To(Equal(6436))
	})

	It("should convert seconds to time string", func() {
		ans := ParseSecondsIntoTime(4559)
		Expect(ans).To(Equal("01|15|59"))
	})

	It("should convert seconds to time string with padding", func() {
		ans := ParseSecondsIntoTime(6436)
		Expect(ans).To(Equal("01|47|16"))
	})
})
