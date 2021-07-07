package util

import (
	. "bou.ke/monkey"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/xszyh/gotest/app/model"
	"github.com/xszyh/gotest/app/service"
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	Convey("test has digit", t, func() {
		Convey("for succ", func() {
			outputExpect := 3
			guard := Patch(
				Add,
				func(first int, slic ...int) int {
					return 3
				})
			defer guard.Unpatch()
			output := Add(1, 2)
			So(output, ShouldEqual, outputExpect)
		})
	})
}

func TestService(t *testing.T) {
	Convey("--", t, func() {
		Convey("++", func() {
			var e *service.UserService
			e = service.NewUserService()

			guard := PatchInstanceMethod(
				reflect.TypeOf(e),
				"GetUserInfo",
				func(_ *service.UserService, _ string) *model.UserInfo {
					return &model.UserInfo{"name", "id", 101}
				})
			defer guard.Unpatch()

			output := e.GetUserInfo("1")
			So(output.UserId, ShouldEqual, "id")
		})
	})
}
