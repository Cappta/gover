package gover

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVersionAdapter(t *testing.T) {
	Convey("Given a new Version strucutre", t, func() {
		version := &Version{}
		Convey("Given a Version string", func() {
			versionString := "1.2"
			Convey("When Scanning", func() {
				err := version.Scan(versionString)
				Convey("Then error should be nil", func() {
					So(err, ShouldBeNil)
				})
				Convey("When retrieving driver's Value", func() {
					value, err := version.Value()
					Convey("Then error should be nil", func() {
						So(err, ShouldBeNil)
					})
					Convey("Then value should equal version string", func() {
						So(value, ShouldEqual, versionString)
					})
				})
			})
		})
		Convey("Given an integer slice", func() {
			intSlice := []int{1, 2}
			Convey("When Scanning", func() {
				err := version.Scan(intSlice)
				Convey("Then error should not be nil", func() {
					So(err, ShouldNotBeNil)
				})
				Convey("When retrieving driver's Value", func() {
					value, err := version.Value()
					Convey("Then error should be nil", func() {
						So(err, ShouldBeNil)
					})
					Convey("Then value should equal an empty string", func() {
						So(value, ShouldEqual, "")
					})
				})
			})
		})
	})
	Convey("Given a nil Version pointer", t, func() {
		var version *Version
		Convey("Given a Version string", func() {
			versionString := "1.2"
			Convey("When Scanning", func() {
				err := version.Scan(versionString)
				Convey("Then error should not be nil", func() {
					So(err, ShouldNotBeNil)
				})
				Convey("When retrieving driver's Value", func() {
					value, err := version.Value()
					Convey("Then error should be nil", func() {
						So(err, ShouldBeNil)
					})
					Convey("Then value should be nil", func() {
						So(value, ShouldBeNil)
					})
				})
			})
		})
	})
}
