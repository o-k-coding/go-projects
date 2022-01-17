package in_array_sorted_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/in_array_sorted"
)

func dotest(array1 []string, array2 []string, exp []string) {
    var ans = InArray(array1, array2)
    Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests Tour", func() {

    It("should handle basic cases", func() {
        var a1 = []string{"live", "arp", "strong"}
        var a2 = []string{"lively", "alive", "harp", "sharp", "armstrong"}
        var r = []string{"arp", "live", "strong"}
        dotest(a1, a2, r)

        a1 = []string{"cod", "code", "wars", "ewar", "ar"}
        a2 = []string{}
        r = []string{}
        dotest(a1, a2, r)

    })
})
