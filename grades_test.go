package moira

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetEventGrades(testing *testing.T) {
	Convey("Progress should contains progress grade", testing, func() {
		event := NotificationEvent{
			State:    "OK",
			OldState: "WARN",
		}
		expected := []string{"PROGRESS"}
		actual := event.GetEventGrades()
		So(actual, ShouldResemble, expected)
	})

	Convey("Degradation should contains degradation grade", testing, func() {
		Convey("WARN -> OK", func() {
			event := NotificationEvent{
				State:    "WARN",
				OldState: "OK",
			}
			expected := []string{"DEGRADATION"}
			actual := event.GetEventGrades()
			So(actual, ShouldResemble, expected)
		})

		Convey("ERROR -> WARN", func() {
			event := NotificationEvent{
				State:    "ERROR",
				OldState: "WARN",
			}
			expected := []string{"DEGRADATION"}
			actual := event.GetEventGrades()
			So(actual, ShouldResemble, expected)
		})
	})

	Convey("High degradation should contains HIGH DEGRADATION grade", testing, func() {
		Convey("ERROR -> OK", func() {
			event := NotificationEvent{
				State:    "ERROR",
				OldState: "OK",
			}
			expected := []string{"HIGH DEGRADATION", "DEGRADATION"}
			actual := event.GetEventGrades()
			So(actual, ShouldResemble, expected)
		})

		Convey("NODATA -> ERROR", func() {
			event := NotificationEvent{
				State:    "NODATA",
				OldState: "ERROR",
			}
			expected := []string{"HIGH DEGRADATION", "DEGRADATION"}
			actual := event.GetEventGrades()
			So(actual, ShouldResemble, expected)
		})
	})
}
