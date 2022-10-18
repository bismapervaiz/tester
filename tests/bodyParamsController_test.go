package controllers_test

import (
	"encoding/json"
	"log"
	"tester/src/data"
	"tester/src/utilities"
	"testing"
	"time"
)

func TestSendString(t *testing.T) {
	value := "TestString"
	got, _, _ := bodyParamsCntr.SendString(value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestUpdateString(t *testing.T) {
	value := "TestString"
	got, _, _ := bodyParamsCntr.UpdateString(value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestSendDate(t *testing.T) {
	date, err := time.Parse("2006-01-02", "1994-02-13")
	if err != nil {
		log.Fatalln("Cannot parse the time.")
	}
	got, _, _ := bodyParamsCntr.SendDate(date)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestSendUnixDateTime(t *testing.T) {
	date := time.Unix(1484719381, 0)
	got, _, _ := bodyParamsCntr.SendUnixDateTime(date)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestSendRfc3339DateTime(t *testing.T) {
	date, err := time.Parse(time.RFC3339, "1994-02-13T14:01:54.9571247Z")
	if err != nil {
		log.Fatalln("Cannot parse the time.")
	}
	got, _, _ := bodyParamsCntr.SendRfc3339DateTime(date)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestSendRfc1123DateTime(t *testing.T) {
	date, err := time.Parse(time.RFC1123, "Sun, 06 Nov 1994 08:49:37 GMT")
	if err != nil {
		log.Fatalln("Error parsing the date: ", err)
	}
	got, _, _ := bodyParamsCntr.SendRfc1123DateTime(date)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestSendRfc1123DateTimeArray(t *testing.T) {
	var arr []string = []string{
		"Sun, 06 Nov 1994 08:49:37 GMT",
		"Sun, 06 Nov 1994 08:49:37 GMT",
	}

	var dateTime []time.Time = utilities.ToTimeSlice(
		arr,
		time.RFC1123)

	got, _, _ := bodyParamsCntr.SendRfc1123DateTimeArray(true, dateTime)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestSendRfc3339DateTimeArray(t *testing.T) {
	var arr []string = []string{
		"1994-02-13T14:01:54.9571247Z",
		"1994-02-13T14:01:54.9571247Z",
	}

	var dateTime []time.Time = utilities.ToTimeSlice(
		arr,
		time.RFC3339)

	got, _, _ := bodyParamsCntr.SendRfc3339DateTimeArray(true, dateTime)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestSendUnixDateTimeArray(t *testing.T) {
	var arr []int64 = []int64{
		1484719381,
		1484719381,
	}

	var dateTime []time.Time = utilities.ToTimeSlice(
		arr,
		time.UnixDate)

	got, _, _ := bodyParamsCntr.SendUnixDateTimeArray(true, dateTime)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestSendDateArray(t *testing.T) {
	var arr []string = []string{
		"1994-02-13",
		"1994-02-13",
	}

	var dateTime []time.Time = utilities.ToTimeSlice(
		arr,
		utilities.DEFAULT_DATE)

	got, _, _ := bodyParamsCntr.SendDateArray(true, dateTime)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestSendStringArray(t *testing.T) {
	value := []string{"abc", "def"}
	got, _, _ := bodyParamsCntr.SendStringArray(true, value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestSendIntegerArray(t *testing.T) {
	value := []int64{1, 2, 3, 4, 5}
	got, _, _ := bodyParamsCntr.SendIntegerArray(true, value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestUpdateStringArrayWithBody(t *testing.T) {
	value := []string{"abc", "def"}
	got, _, _ := bodyParamsCntr.UpdateStringArray(true, value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestSendDeleteBody(t *testing.T) {
	var val data.DeleteBody
	json.Unmarshal([]byte(`{"name": "farhan", "field": "QA"}`), &val)
	got, _, _ := bodyParamsCntr.SendDeleteBody(val)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestSendDeleteBodyWithMultilinerName(t *testing.T) {
	var val data.DeleteBody
	json.Unmarshal([]byte(`{"name": "farhan\nnouman", "field": "QA"}`), &val)
	got, _, _ := bodyParamsCntr.SendDeleteBody(val)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestSendDeleteBodyWithSpecialFieldName(t *testing.T) {
	var val data.DeleteBody
	json.Unmarshal([]byte(`{"name": "farhan", "field": "&&&"}`), &val)
	got, _, _ := bodyParamsCntr.SendDeleteBody(val)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestSendDeleteBodyWithBlankField(t *testing.T) {
	var val data.DeleteBody
	json.Unmarshal([]byte(`{"name": "farhan", "field": " "}`), &val)
	got, _, _ := bodyParamsCntr.SendDeleteBody(val)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestSendDeleteBodyWithBlankName(t *testing.T) {
	var val data.DeleteBody
	json.Unmarshal([]byte(`{"name": " ", "field": "QA"}`), &val)
	got, _, _ := bodyParamsCntr.SendDeleteBody(val)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestSendDeleteBodyWithBlankNameAndField(t *testing.T) {
	var val data.DeleteBody
	json.Unmarshal([]byte(`{"name": " ", "field": " "}`), &val)
	got, _, _ := bodyParamsCntr.SendDeleteBody(val)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestSendPlainText(t *testing.T) {
	var val string = "farhan\nnouman"
	got, _, _ := bodyParamsCntr.SendDeletePlainText(val)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestStringEnumArray(t *testing.T) {
	value := []data.Days{"Tuesday", "Saturday", "Wednesday", "Monday", "Sunday"}
	got, _, _ := bodyParamsCntr.SendStringEnumArray(true, value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestIntegerEnumArray(t *testing.T) {
	value := []data.SuiteCode{1, 3, 4, 2, 3}
	got, _, _ := bodyParamsCntr.SendIntegerEnumArray(true, value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

// func TestSendModel(t *testing.T) {
// 	var val data.EmployeeInterface
// 	val, _ = data.UnmarshalEmployee([]byte(`{"name":"Shahid Khaliq","age":5147483645,"address":"H # 531, S # 20",
// 	"uid":"123321","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z","salary":20000,
// 	"department":"Software Development","joiningDay":"Saturday","workingDays":["Monday","Tuesday","Friday"],
// 	"boss":{
// 		"personType":"Boss","name":"Zeeshan Ejaz","age":5147483645,"address":"H # 531, S # 20",
// 		"uid":"123321","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z","salary":20000,
// 		"department":"Software Development","joiningDay":"Saturday","workingDays":["Monday","Tuesday","Friday"],
// 		"dependents":[
// 			{"name":"Future Wife","age":5147483649,"address":"H # 531, S # 20","uid":"123412","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z"},
// 			{"name":"Future Kid","age":5147483648,"address":"H # 531, S # 20","uid":"312341","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z"}
// 			],
// 		"hiredAt":"Sun, 06 Nov 1994 08:49:37 GMT","promotedAt":"1994-02-13"
// 	},
// 		"dependents":[
// 			{"name":"Future Wife","age":5147483649,"address":"H # 531, S # 20","uid":"123412","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z"},
// 			{"name":"Future Kid","age":5147483648,"address":"H # 531, S # 20","uid":"312341","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z"}
// 			],
// 			"hiredAt":"Sun, 06 Nov 1994 08:49:37 GMT"}`))
// 	got, _, _ := bodyParamsCntr.SendModel(val)

// 	var expected data.ServerResponse
// 	json.Unmarshal([]byte(`{"passed":true}`), &expected)

// 	if got.Passed != expected.Passed {
// 		t.Errorf("got %v but expected %v", got, expected)
// 	}
// }

// func TestSendDeleteBodyWithModel(t *testing.T) {
// 	var val data.Employee
// 	json.Unmarshal([]byte(`{"name":"Shahid Khaliq","age":5147483645,"address":"H # 531, S # 20","uid":"123321",
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
// 	"hiredAt":"Sun, 06 Nov 1994 08:49:37 GMT"}`), &val)
// 	got, _, _  := bodyParamsCntr.SendDeleteBodyWithModel(val)

// 	var expected data.ServerResponse
// 	json.Unmarshal([]byte(`{"passed":true}`), &expected)

// 	if got.Passed != expected.Passed {
// 		t.Errorf("got %v but expected %v", got, expected)
// 	}
// }
