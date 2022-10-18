package controllers_test

import (
	"encoding/json"
	"log"
	"tester/src/data"
	"tester/src/utilities"
	"testing"
	"time"
)

func TestOptionalDynamicQueryParam(t *testing.T) {
	value := "farhan"
	queryParam := map[string]interface{}{}
	queryParam["field"] = "QA"
	got, _, _ := queryParamsCntr.OptionalDynamicQueryParam(value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestNoParams(t *testing.T) {
	got, _, _ := queryParamsCntr.NoParams()

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestDate(t *testing.T) {
	date, err := time.Parse("2006-01-02", "1994-02-13")
	if err != nil {
		log.Fatalln("Cannot parse the time.")
	}
	got, _, _ := queryParamsCntr.Date(date)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestUnixDateTime(t *testing.T) {
	date := time.Unix(1484719381, 0)

	got, _, _ := queryParamsCntr.UnixDateTime(date)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestRfc3339DateTime(t *testing.T) {
	date, err := time.Parse(time.RFC3339, "1994-02-13T14:01:54.9571247Z")
	if err != nil {
		log.Fatalln("Cannot parse the time.")
	}

	got, _, _ := queryParamsCntr.Rfc1123DateTime(date)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestRfc1123DateTime(t *testing.T) {
	date, err := time.Parse(time.RFC1123, "Sun, 06 Nov 1994 08:49:37 GMT")
	if err != nil {
		log.Fatalln("Error parsing the date: ", err)
	}

	got, _, _ := queryParamsCntr.Rfc1123DateTime(date)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestStringParam(t *testing.T) {
	value := "l;asd;asdwe[2304&&;'.d??\\a\\\\\\;sd//"
	got, _, _ := queryParamsCntr.StringParam(value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestUrlParam(t *testing.T) {
	value := "https://www.shahidisawesome.com/and/also/a/narcissist?thisis=aparameter&another=one"
	got, _, _ := queryParamsCntr.UrlParam(value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestSimpleQuery(t *testing.T) {
	boolean := true
	var number int64 = 4
	str := "TestString"
	got, _, _ := queryParamsCntr.SimpleQuery(boolean, number, str)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestMultipleParams(t *testing.T) {
	precision := 1112.34
	var number int64 = 123412312
	str := "\"\"test./;\";12&&3asl\"\";\"qw1&34\"///..//."
	url := "http://www.abc.com/test?a=b&c=\"http://lolol.com?param=no&another=lol\""
	got, _, _ := queryParamsCntr.MultipleParams(number, precision, str, url)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestStringArray(t *testing.T) {
	value := []string{"abc", "def"}
	got, _, _ := queryParamsCntr.StringArray(value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestIntegerArray(t *testing.T) {
	value := []int64{1, 2, 3, 4, 5}
	got, _, _ := queryParamsCntr.NumberArray(value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestQueryIntegerEnumArray(t *testing.T) {
	value := []data.SuiteCode{1, 3, 4, 2, 3}
	got, _, _ := queryParamsCntr.IntegerEnumArray(value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestQueryStringEnumArray(t *testing.T) {
	value := []data.Days{"Tuesday", "Saturday", "Wednesday", "Monday", "Sunday"}
	got, _, _ := queryParamsCntr.StringEnumArray(value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestRfc1123DateTimeArray(t *testing.T) {
	var arr []string = []string{
		"Sun, 06 Nov 1994 08:49:37 GMT",
		"Sun, 06 Nov 1994 08:49:37 GMT",
	}

	var dateTime []time.Time = utilities.ToTimeSlice(
		arr,
		time.RFC1123)

	got, _, _ := queryParamsCntr.Rfc1123DateTimeArray(dateTime)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestRfc3339DateTimeArray(t *testing.T) {
	var arr []string = []string{
		"1994-02-13T14:01:54.9571247Z",
		"1994-02-13T14:01:54.9571247Z",
	}

	var dateTime []time.Time = utilities.ToTimeSlice(
		arr,
		time.RFC3339)

	got, _, _ := queryParamsCntr.Rfc3339DateTimeArray(dateTime)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestUnixDateTimeArray(t *testing.T) {
	var arr []int64 = []int64{
		1484719381,
		1484719381,
	}

	var dateTime []time.Time = utilities.ToTimeSlice(
		arr,
		time.UnixDate)

	got, _, _ := queryParamsCntr.UnixDateTimeArray(dateTime)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestDateArray(t *testing.T) {
	var arr []string = []string{
		"1994-02-13",
		"1994-02-13",
	}

	var dateTime []time.Time = utilities.ToTimeSlice(
		arr,
		utilities.DEFAULT_DATE)

	got, _, _ := queryParamsCntr.DateArray(dateTime)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}
