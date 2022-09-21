package main

import (
	"math"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestFoo28(t *testing.T) {
	for i := 0; i < 10; i++ {
		convey.Convey("integer should smaller than return val", t, func() {
			convey.So(0x7, convey.ShouldBeLessThan, math.MinInt32)
		})
	}
}
