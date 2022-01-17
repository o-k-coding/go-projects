package robot_in_circle_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/robot_in_circle"
)

// Add test cases here

var _ = Describe("Basic tests", func() {
	It("Should find the correct answer", func() {
		// Expect(RobotInCircle("GGLLGG")).To(Equal(true))
		// Expect(RobotInCircle("GG")).To(Equal(false))
		// Expect(RobotInCircle("GL")).To(Equal(true))
		// THis should put you back at the start, need to debug this specifically
		Expect(RobotInCircle("GGRGGRGGRGGR")).To(Equal(true))
	})
})
