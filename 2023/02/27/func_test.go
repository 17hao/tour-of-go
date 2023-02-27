package main

import (
	"testing"

	"github.com/bytedance/mockey"
	"github.com/magiconair/properties/assert"
)

func TestDoSomething(t *testing.T) {
	mockey.PatchConvey("", t, func() {
		mockey.Mock(GetMyInterface).Return(&MyInterfaceImpl{}).Build()
		mockey.Mock((*MyInterfaceImpl).DoSomething).Return("mock").Build()

		res := GetMyInterface().DoSomething()
		assert.Equal(t, "mock", res)
	})
}
