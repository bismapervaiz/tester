package controllers_test

import (
	"encoding/json"
	"log"
	"tester/src/data"
	"tester/src/utilities"
	"testing"
	"time"
)

// COMMENTED OUT UNTIL FILE FUNCTIONALITY IMPLEMENTED
// func TestSendFile(t *testing.T) {
// 	got, _, _ := formParamsCntr.SendFile(utilities.GetFile("http://localhost:3000/response/image"))

// 	var expected data.ServerResponse
// 	json.Unmarshal([]byte(`{"passed":true}`), &expected)

// 	if got.Passed != expected.Passed {
// 		t.Errorf("got %v but expected %v", got, expected)
// 	}
// }

func TestFormSendString(t *testing.T) {
	value := "TestString"
	got, _, _ := formParamsCntr.SendString(value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestUpdateStringWithForm(t *testing.T) {
	value := "TestString"
	got, _, _ := formParamsCntr.UpdateStringWithForm(value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestUpdateStringArrayWithForm(t *testing.T) {
	value := []string{"abc", "def"}
	got, _, _ := formParamsCntr.UpdateStringArrayWithForm(true, value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}
func TestFormSendDate(t *testing.T) {
	date, err := time.Parse("2006-01-02", "1994-02-13")
	if err != nil {
		log.Fatalln("Cannot parse the time.")
	}

	got, _, _ := formParamsCntr.SendDate(date)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestFormSendUnixDateTime(t *testing.T) {
	date := time.Unix(1484719381, 0)

	got, _, _ := formParamsCntr.SendUnixDateTime(date)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestFormSendRfc3339DateTime(t *testing.T) {
	date, err := time.Parse(time.RFC3339, "1994-02-13T14:01:54.9571247Z")
	if err != nil {
		log.Fatalln("Cannot parse the time.")
	}

	got, _, _ := formParamsCntr.SendRfc3339DateTime(date)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestFormSendRfc1123DateTime(t *testing.T) {
	date, err := time.Parse(time.RFC1123, "Sun, 06 Nov 1994 08:49:37 GMT")
	if err != nil {
		log.Fatalln("Error parsing the date: ", err)
	}

	got, _, _ := formParamsCntr.SendRfc1123DateTime(date)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestFormSendRfc1123DateTimeArray(t *testing.T) {
	var arr []string = []string{
		"Sun, 06 Nov 1994 08:49:37 GMT",
		"Sun, 06 Nov 1994 08:49:37 GMT",
	}

	var dateTime []time.Time = utilities.ToTimeSlice(
		arr,
		time.RFC1123)

	got, _, _ := formParamsCntr.SendRfc1123DateTimeArray(true, dateTime)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestFormSendRfc3339DateTimeArray(t *testing.T) {
	var arr []string = []string{
		"1994-02-13T14:01:54.9571247Z",
		"1994-02-13T14:01:54.9571247Z",
	}

	var dateTime []time.Time = utilities.ToTimeSlice(
		arr,
		time.RFC3339)

	got, _, _ := formParamsCntr.SendRfc3339DateTimeArray(true, dateTime)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestFormSendUnixDateTimeArray(t *testing.T) {
	var arr []int64 = []int64{
		1484719381,
		1484719381,
	}

	var dateTime []time.Time = utilities.ToTimeSlice(
		arr,
		time.UnixDate)

	got, _, _ := formParamsCntr.SendUnixDateTimeArray(true, dateTime)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestFormSendDateArray(t *testing.T) {
	var arr []string = []string{
		"1994-02-13",
		"1994-02-13",
	}

	var dateTime []time.Time = utilities.ToTimeSlice(
		arr,
		utilities.DEFAULT_DATE)

	got, _, _ := formParamsCntr.SendDateArray(true, dateTime)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestFormSendStringArray(t *testing.T) {
	value := []string{"abc", "def"}
	got, _, _ := formParamsCntr.SendStringArray(true, value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestFormSendIntegerArray(t *testing.T) {
	value := []int64{1, 2, 3, 4, 5}
	got, _, _ := formParamsCntr.SendIntegerArray(true, value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestFormSendLong(t *testing.T) {
	var value int64 = 5147483647
	got, _, _ := formParamsCntr.SendLong(value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestFormIntegerEnumArray(t *testing.T) {
	value := []data.SuiteCode{1, 3, 4, 2, 3}
	got, _, _ := formParamsCntr.SendIntegerEnumArray(true, value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestFormStringEnumArray(t *testing.T) {
	value := []data.Days{"Tuesday", "Saturday", "Wednesday", "Monday", "Sunday"}
	got, _, _ := formParamsCntr.SendStringEnumArray(true, value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestFormDeleteForm(t *testing.T) {
	value := data.DeleteBody{Field: "QA", Name: "farhan"}
	got, _, _ := formParamsCntr.SendDeleteForm(value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestFormDeleteFormWithBlankField(t *testing.T) {
	value := data.DeleteBody{Field: "QA", Name: " "}
	got, _, _ := formParamsCntr.SendDeleteForm(value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestFormDeleteFormWithBlankNameField(t *testing.T) {
	value := data.DeleteBody{Field: " ", Name: " "}
	got, _, _ := formParamsCntr.SendDeleteForm(value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestFormDeleteFormWithBlankName(t *testing.T) {
	value := data.DeleteBody{Field: " ", Name: "farhan"}
	got, _, _ := formParamsCntr.SendDeleteForm(value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestFormDeleteFormWithSpecialCharactersField(t *testing.T) {
	value := data.DeleteBody{Field: "&&&", Name: "farhan"}
	got, _, _ := formParamsCntr.SendDeleteForm(value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestFormDeleteFormWithMultilinerName(t *testing.T) {
	value := data.DeleteBody{Field: "QA", Name: "farhan\nnouman"}
	got, _, _ := formParamsCntr.SendDeleteForm(value)

	var expected data.ServerResponse
	json.Unmarshal([]byte(`{"passed":true}`), &expected)

	if got.Passed != expected.Passed {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

// func TestFormDeleteFormWithModel(t *testing.T) {
// 	var val data.Employee
// 	json.Unmarshal([]byte(`{"name":"Shahid Khaliq","age":5147483645,"address":"H # 531, S # 20",
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
// 		"hiredAt":"Sun, 06 Nov 1994 08:49:37 GMT","promotedAt":1484719381
// 	},
// 		"dependents":[
// 			{"name":"Future Wife","age":5147483649,"address":"H # 531, S # 20","uid":"123412","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z"},
// 			{"name":"Future Kid","age":5147483648,"address":"H # 531, S # 20","uid":"312341","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z"}
// 			],
// 			"hiredAt":"Sun, 06 Nov 1994 08:49:37 GMT"}`), &val)
// 	got, _, _ := formParamsCntr.SendDeleteForm1(val)

// 	var expected data.ServerResponse
// 	json.Unmarshal([]byte(`{"passed":true}`), &expected)

// 	if got.Passed != expected.Passed {
// 		t.Errorf("got %v but expected %v", got, expected)
// 	}
// }

// func TestFormSendModel(t *testing.T) {
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
// 		"hiredAt":"Sun, 06 Nov 1994 08:49:37 GMT","promotedAt":1484719381
// 	},
// 		"dependents":[
// 			{"name":"Future Wife","age":5147483649,"address":"H # 531, S # 20","uid":"123412","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z"},
// 			{"name":"Future Kid","age":5147483648,"address":"H # 531, S # 20","uid":"312341","birthday":"1994-02-13","birthtime":"1994-02-13T14:01:54.9571247Z"}
// 			],
// 			"hiredAt":"Sun, 06 Nov 1994 08:49:37 GMT"}`))
// 	got, _, _ := formParamsCntr.SendModel(val)

// 	var expected data.ServerResponse
// 	json.Unmarshal([]byte(`{"passed":true}`), &expected)

// 	if got.Passed != expected.Passed {
// 		t.Errorf("got %v but expected %v", got, expected)
// 	}
// }
