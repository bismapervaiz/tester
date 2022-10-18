package controllers_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"tester/src/data"
	"tester/src/utilities"
	"testing"
	"time"
)

func TestGetInteger(t *testing.T) {
	var expected int64 = 4
	got, _, _ := responseTypesCntr.GetInteger()

	if got != expected {
		t.Errorf("got %d but expected %d", got, expected)
	}
}

func TestGetBinary(t *testing.T) {
	file, err := http.Get("http://localhost:3000/response/image")
	if err != nil {
		log.Fatal(err)
	}

	expected, err := ioutil.ReadAll(file.Body)
	if err != nil {
		log.Fatalln(err)
	}
	got, _, _ := responseTypesCntr.GetBinary()

	if ok := bytes.Compare(expected, got); ok != 0 {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestGetPrecision(t *testing.T) {
	var expected float64 = 4.999
	got, _, _ := responseTypesCntr.GetPrecision()

	if got != expected {
		t.Errorf("got %f but expected %f", got, expected)
	}
}

func TestGetLong(t *testing.T) {
	var expected int64 = 5147483647
	got, _, _ := responseTypesCntr.GetLong()

	if got != expected {
		t.Errorf("got %d but expected %d", got, expected)
	}
}

func TestGetBoolean(t *testing.T) {
	var expected bool = true
	got, _, _ := responseTypesCntr.GetBoolean()

	if got != expected {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

// func TestGetModel(t *testing.T) {
// 	var expected data.Person
// 	json.Unmarshal([]byte(
// 		`{"name":"Shahid Khaliq","age":5147483645,"address":"H # 531, S # 20","uid":"123321",
// 		"birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z","salary":20000,"personType":"Empl",
// 		"department":"Software Development","joiningDay":"Saturday","workingDays":["Monday","Tuesday","Friday"],
// 		"boss":{"personType":"Boss","name":"Zeeshan Ejaz","age":5147483645,"address":"H # 531, S # 20",
// 		"uid":"123321","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z","salary":20000,
// 		"department":"Software Development","joiningDay":"Saturday","workingDays":["Monday","Tuesday","Friday"],
// 		"dependents":[{"name":"Future Wife","age":5147483649,"address":"H # 531, S # 20","uid":"123412","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z"},
// 		{"name":"Future Kid","age":5147483648,"address":"H # 531, S # 20","uid":"312341","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z"}],
// 		"hiredAt":"Sun, 06 Nov 1994 08:49:37 GMT","promotedAt":1484719381},
// 		"dependents":[{"name":"Future Wife","age":5147483649,"address":"H # 531, S # 20","uid":"123412","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z"},
// 		{"name":"Future Kid","age":5147483648,"address":"H # 531, S # 20","uid":"312341","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z"}],
// 		"hiredAt":"Sun, 06 Nov 1994 08:49:37 GMT"}`), &expected)
// 	got, _, _ := responseTypesCntr.GetModel()

// 	if got != expected {
// 		t.Errorf("got %v but expected %v", got, expected)
// 	}
// }

func TestGetStringEnum(t *testing.T) {
	var expected data.Days = "Monday"
	got, _, _ := responseTypesCntr.GetStringEnum("TestString")

	if got != expected {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestGetIntEnum(t *testing.T) {
	var expected data.SuiteCode = 3
	got, _, _ := responseTypesCntr.GetIntEnum("TestString")

	if got != expected {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

// Route does not exist in the test server.
// func TestGetDynamic(t *testing.T) {
// 	expected := json.RawMessage(`{"method":"GET","body":{},"uploadCount":0}`)
// 	got, resp, _ := responseTypesCntr.GetDynamic()

// 	if resp.StatusCode != 200 {
// 		t.Errorf("got %v but expected %v", got, expected)
// 	}
// }

func TestGetHeadersAllowExtra(t *testing.T) {
	var expected http.Response = http.Response{}
	expected.Header = make(http.Header)
	expected.Header.Add("NauManAli", "")
	expected.Header.Add("WaseemHasAn", "is also awesome")
	got, _ := responseTypesCntr.GetHeaders()

	if !utilities.AreMapKeysProperSubset(expected.Header, got.Header) {
		t.Errorf("got %v but expected %v", got.Header, expected.Header)
	}
}

func TestGetContentTypeInResponse(t *testing.T) {
	var expected http.Response = http.Response{}
	expected.Header = make(http.Header)
	expected.Header.Add("Content-Type", "application/responseType")
	expected.Header.Add("Accept", "application/noTerm")
	expected.Header.Add("Accept-Encoding", "UTF-8")

	got, _ := responseTypesCntr.GetContentTypeHeaders()

	if !utilities.IsMapProperSubset(expected.Header, got.Header) {
		t.Errorf("got %v but expected %v", got.Header, expected.Header)
	}
}

func TestGetDate(t *testing.T) {
	expected, err := time.Parse(utilities.DEFAULT_DATE, "1994-02-13")
	if err != nil {
		log.Fatalln("Erroe parsing the date: ", err)
	}
	got, _, _ := responseTypesCntr.GetDate()

	if got != expected {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestGet3339DateTime(t *testing.T) {
	expected, err := time.Parse(time.RFC3339, "2016-03-13T12:52:32.123Z")
	if err != nil {
		log.Fatalln("Error parsing the date: ", err)
	}
	got, _, _ := responseTypesCntr.Get3339Datetime()

	if got != expected {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestGet1123DateTime(t *testing.T) {
	expected, err := time.Parse(time.RFC1123, "Sun, 06 Nov 1994 08:49:37 GMT")
	if err != nil {
		log.Fatalln("Error parsing the date: ", err)
	}
	got, _, _ := responseTypesCntr.Get1123DateTime()

	if !got.Equal(expected) {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestGetUnixDateTime(t *testing.T) {
	expected := time.Unix(1484719381, 0)

	got, _, _ := responseTypesCntr.GetUnixDateTime()

	if !got.Equal(expected) {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestGetDateArray(t *testing.T) {
	var arr []string = []string{
		"1994-02-13",
		"1994-02-13",
	}

	var expected []time.Time = utilities.ToTimeSlice(
		arr,
		utilities.DEFAULT_DATE)
	got, _, _ := responseTypesCntr.GetDateArray(true)

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestGet3339DateArray(t *testing.T) {
	var arr []string = []string{
		"2016-03-13T12:52:32.123Z",
		"2016-03-13T12:52:32.123Z",
	}

	var expected []time.Time = utilities.ToTimeSlice(
		arr,
		time.RFC3339)
	got, _, _ := responseTypesCntr.Get3339DatetimeArray(true)

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestGet1123DateArray(t *testing.T) {
	var arr []string = []string{
		"Sun, 06 Nov 1994 08:49:37 GMT",
		"Sun, 06 Nov 1994 08:49:37 GMT",
	}

	var expected []time.Time = utilities.ToTimeSlice(
		arr,
		time.RFC1123)
	got, _, _ := responseTypesCntr.Get1123DateTimeArray(true)

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestGetUnixDateArray(t *testing.T) {
	var arr []int64 = []int64{
		1484719381,
		1484719381,
	}

	var expected []time.Time = utilities.ToTimeSlice(
		arr,
		time.UnixDate)
	got, _, _ := responseTypesCntr.GetUnixDateTimeArray(true)

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestGetBooleanArray(t *testing.T) {
	var expected []bool = []bool{
		true,
		false,
		true,
		true,
		false,
	}

	got, _, _ := responseTypesCntr.GetBooleanArray(true)

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

// func TestGetModelArray(t *testing.T) {
// 	var expected []data.Person
// 	json.Unmarshal([]byte(`[{"name":"Shahid Khaliq","age":5147483645,"address":"H # 531, S # 20","uid":"123321",
// 	"birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z","salary":20000,"personType":"Empl",
// 	"department":"Software Development","joiningDay":"Saturday","workingDays":["Monday","Tuesday","Friday"],
// 	"boss":{"personType":"Boss","name":"Zeeshan Ejaz","age":5147483645,"address":"H # 531, S # 20",
// 	"uid":"123321","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z","salary":20000,
// 	"department":"Software Development","joiningDay":"Saturday","workingDays":["Monday","Tuesday","Friday"],
// 	"dependents":[{"name":"Future Wife","age":5147483649,"address":"H # 531, S # 20","uid":"123412","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z"},
// 	{"name":"Future Kid","age":5147483648,"address":"H # 531, S # 20","uid":"312341","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z"}],
// 	"hiredAt":"Sun, 06 Nov 1994 08:49:37 GMT","promotedAt":1484719381},
// 	"dependents":[{"name":"Future Wife","age":5147483649,"address":"H # 531, S # 20","uid":"123412","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z"},
// 	{"name":"Future Kid","age":5147483648,"address":"H # 531, S # 20","uid":"312341","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z"}],
// 	"hiredAt":"Sun, 06 Nov 1994 08:49:37 GMT"},
// 	{"name":"Shahid Khaliq","age":5147483645,"address":"H # 531, S # 20","uid":"123321",
// 	"birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z","salary":20000,"personType":"Empl",
// 	"department":"Software Development","joiningDay":"Saturday","workingDays":["Monday","Tuesday","Friday"],
// 	"boss":{"personType":"Boss","name":"Zeeshan Ejaz","age":5147483645,"address":"H # 531, S # 20",
// 	"uid":"123321","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z","salary":20000,
// 	"department":"Software Development","joiningDay":"Saturday","workingDays":["Monday","Tuesday","Friday"],
// 	"dependents":[{"name":"Future Wife","age":5147483649,"address":"H # 531, S # 20","uid":"123412","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z"},
// 	{"name":"Future Kid","age":5147483648,"address":"H # 531, S # 20","uid":"312341","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z"}],
// 	"hiredAt":"Sun, 06 Nov 1994 08:49:37 GMT","promotedAt":1484719381},
// 	"dependents":[{"name":"Future Wife","age":5147483649,"address":"H # 531, S # 20","uid":"123412","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z"},
// 	{"name":"Future Kid","age":5147483648,"address":"H # 531, S # 20","uid":"312341","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z"}],
// 	"hiredAt":"Sun, 06 Nov 1994 08:49:37 GMT"}]`),
// 		&expected)
// 	got, _, _ := responseTypesCntr.GetModelArray()

// 	if !reflect.DeepEqual(expected, got) {
// 		t.Errorf("got %v but expected %v", got, expected)
// 	}
// }

func TestGetStringEnumArray(t *testing.T) {
	var expected []data.Days
	json.Unmarshal([]byte(`["Tuesday", "Saturday", "Wednesday", "Monday", "Sunday"]`),
		&expected)
	got, _, _ := responseTypesCntr.GetStringEnumArray(true, "TestString")

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestGetIntEnumArray(t *testing.T) {
	var expected []data.SuiteCode
	json.Unmarshal([]byte(`[1, 3, 4, 2, 3]`),
		&expected)
	got, _, _ := responseTypesCntr.GetIntEnumArray(true, "TestString")

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestReturnResponseWithEnum(t *testing.T) {
	var expected data.ResponseWithEnum
	json.Unmarshal([]byte(`{"optional":false,"_type":"Long","constant":false}`), &expected)
	got, resp, _ := responseTypesCntr.ReturnResponseWithEnums()

	if resp.StatusCode != 200 {
		t.Errorf("got %v but expected %v", got, expected)
	}
}
