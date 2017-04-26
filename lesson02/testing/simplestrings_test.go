package simplestring

import "testing"

const weekdays = "Monday Tuesday Wednesday Thursday Friday"

// test that Tuesday is a weekday
func TestDayOfTheWeek(t *testing.T)  {
	if !Contains(weekdays, "Tuesday"){
		t.Fail()
	}
}
// test that Sunday is not a weekday
func TestNotDayOfTheWeek(t *testing.T){
	if Contains(weekdays, "Sunday"){
		t.Fail()
	}
}

// test that an empty search string returns 0
func TestEmptySearchString(t *testing.T)  {
	if Index(weekdays, "") != 0 {
		t.Fail()
	}
}
// test that the string Monday is not found in the empty string
func TestNoDayInEmptyWeek(t *testing.T)  {
	if Contains("", "Monday"){
		t.Fail()
	}
}