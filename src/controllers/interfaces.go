/*
tester

This file was automatically generated for Stamplay by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package controllers

import (
    "tester/src/data"
    "tester/src/https"
    "time"
)

type ResponseTypesControllerInterface interface {
    GetInteger() (
        int64,
        *https.HttpContext,
        error)
    GetBinary() (
        []byte,
        *https.HttpContext,
        error)
    GetPrecision() (
        float64,
        *https.HttpContext,
        error)
    GetModelArray() (
        []data.PersonInterface,
        *https.HttpContext,
        error)
    GetModel() (
        data.PersonInterface,
        *https.HttpContext,
        error)
    GetLong() (
        int64,
        *https.HttpContext,
        error)
    GetDateArray() (
        []time.Time,
        *https.HttpContext,
        error)
    GetDate() (
        time.Time,
        *https.HttpContext,
        error)
    GetContentTypeHeaders() (
        *https.HttpContext,
        error)
    GetUnixDateTime() (
        time.Time,
        *https.HttpContext,
        error)
    Get1123DateTime() (
        time.Time,
        *https.HttpContext,
        error)
    GetHeaders() (
        *https.HttpContext,
        error)
    GetBoolean() (
        bool,
        *https.HttpContext,
        error)
    Get3339Datetime() (
        time.Time,
        *https.HttpContext,
        error)
    GetDynamic() (
        interface{},
        *https.HttpContext,
        error)
    ReturnResponseWithEnums() (
        data.ResponseWithEnum,
        *https.HttpContext,
        error)
    GetStringEnumArray() (
        []data.Days,
        *https.HttpContext,
        error)
    GetStringEnum() (
        data.Days,
        *https.HttpContext,
        error)
    GetIntEnum() (
        data.SuiteCode,
        *https.HttpContext,
        error)
    GetIntEnumArray() (
        []data.SuiteCode,
        *https.HttpContext,
        error)
    GetIntegerArray() (
        []int64,
        *https.HttpContext,
        error)
    GetDynamicArray() (
        interface{},
        *https.HttpContext,
        error)
    Get3339DatetimeArray() (
        []time.Time,
        *https.HttpContext,
        error)
    GetBooleanArray() (
        []bool,
        *https.HttpContext,
        error)
    Get1123DateTimeArray() (
        []time.Time,
        *https.HttpContext,
        error)
    GetUnixDateTimeArray() (
        []time.Time,
        *https.HttpContext,
        error)
}

func NewResponseTypesControllerInterface(baseController baseController) (
    *ResponseTypesController) {
    client := ResponseTypesController{baseController: baseController}
    return &client
}

type FormParamsControllerInterface interface {
    SendFile(file []byte) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendModel(model data.EmployeeInterface) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendDate(date time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendDeleteForm(body data.DeleteBody) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendDeleteMultipart(file []byte) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendDateArray(dates []time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendUnixDateTime(datetime time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendRfc1123DateTime(datetime time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendRfc3339DateTime(datetime time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendUnixDateTimeArray(datetimes []time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendRfc1123DateTimeArray(datetimes []time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendLong(value int64) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendIntegerArray(integers []int64) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendStringArray(strings []string) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendModelArray(models []data.EmployeeInterface) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendString(value string) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendRfc3339DateTimeArray(datetimes []time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendMixedArray(
        file []byte,
        integers []int64,
        models []data.EmployeeInterface,
        strings []string) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    UpdateModelWithForm(model data.EmployeeInterface) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendDeleteForm1(model data.EmployeeInterface) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendDeleteFormWithModelArray(models []data.EmployeeInterface) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    UpdateModelArrayWithForm(models []data.EmployeeInterface) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    UpdateStringWithForm(value string) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    UpdateStringArrayWithForm(strings []string) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendIntegerEnumArray(suites []data.SuiteCode) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendStringEnumArray(days []data.Days) (
        data.ServerResponse,
        *https.HttpContext,
        error)
}

func NewFormParamsControllerInterface(baseController baseController) (
    *FormParamsController) {
    client := FormParamsController{baseController: baseController}
    return &client
}

type BodyParamsControllerInterface interface {
    SendModel(model data.EmployeeInterface) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendDate(date time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    UpdateModel(model data.EmployeeInterface) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendDeletePlainText(textString string) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendDeleteBody(body data.DeleteBody) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendDateArray(dates []time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendUnixDateTime(datetime time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendRfc1123DateTime(datetime time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendRfc3339DateTime(datetime time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendUnixDateTimeArray(datetimes []time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendRfc1123DateTimeArray(datetimes []time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendRfc3339DateTimeArray(datetimes []time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendStringArray(sarray []string) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    UpdateString(value string) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendIntegerArray(integers []int64) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendModelArray(models []data.EmployeeInterface) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendDynamic(dynamic interface{}) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendString(value string) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendStringEnumArray(days []data.Days) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendIntegerEnumArray(suites []data.SuiteCode) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SendDeleteBodyWithModel(model data.EmployeeInterface) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    UpdateModelArray(models []data.EmployeeInterface) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    UpdateStringArray(strings []string) (
        data.ServerResponse,
        *https.HttpContext,
        error)
}

func NewBodyParamsControllerInterface(baseController baseController) (
    *BodyParamsController) {
    client := BodyParamsController{baseController: baseController}
    return &client
}

type ErrorCodesControllerInterface interface {
    Get500() (
        interface{},
        *https.HttpContext,
        error)
    Get400() (
        interface{},
        *https.HttpContext,
        error)
    Get501() (
        interface{},
        *https.HttpContext,
        error)
    Catch412GlobalError() (
        interface{},
        *https.HttpContext,
        error)
    Get401() (
        interface{},
        *https.HttpContext,
        error)
}

func NewErrorCodesControllerInterface(baseController baseController) (
    *ErrorCodesController) {
    client := ErrorCodesController{baseController: baseController}
    return &client
}

type QueryParamControllerInterface interface {
    Date(date time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    DateArray(dates []time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    NoParams() (
        data.ServerResponse,
        *https.HttpContext,
        error)
    OptionalDynamicQueryParam(name string) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    UnixDateTimeArray(datetimes []time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    UnixDateTime(datetime time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    Rfc1123DateTime(datetime time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    Rfc1123DateTimeArray(datetimes []time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    Rfc3339DateTimeArray(datetimes []time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    Rfc3339DateTime(datetime time.Time) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    StringParam(mString string) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    UrlParam(url string) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    NumberArray(integers []int64) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    StringArray(strings []string) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    SimpleQuery(
        boolean bool,
        number int64,
        mString string) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    StringEnumArray(days []data.Days) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    MultipleParams(
        number int64,
        precision float64,
        mString string,
        url string) (
        data.ServerResponse,
        *https.HttpContext,
        error)
    IntegerEnumArray(suites []data.SuiteCode) (
        data.ServerResponse,
        *https.HttpContext,
        error)
}

func NewQueryParamControllerInterface(baseController baseController) (
    *QueryParamController) {
    client := QueryParamController{baseController: baseController}
    return &client
}

type EchoControllerInterface interface {
    QueryEcho() (
        data.EchoResponse,
        *https.HttpContext,
        error)
    JsonEcho(input interface{}) (
        interface{},
        *https.HttpContext,
        error)
    FormEcho(input interface{}) (
        interface{},
        *https.HttpContext,
        error)
}

func NewEchoControllerInterface(baseController baseController) (
    *EchoController) {
    client := EchoController{baseController: baseController}
    return &client
}

type HeaderControllerInterface interface {
    SendHeaders(
        customHeader string,
        value string) (
        data.ServerResponse,
        *https.HttpContext,
        error)
}

func NewHeaderControllerInterface(baseController baseController) (
    *HeaderController) {
    client := HeaderController{baseController: baseController}
    return &client
}

type TemplateParamsControllerInterface interface {
    SendStringArray(strings []string) (
        data.EchoResponse,
        *https.HttpContext,
        error)
    SendIntegerArray(integers []int64) (
        data.EchoResponse,
        *https.HttpContext,
        error)
}

func NewTemplateParamsControllerInterface(baseController baseController) (
    *TemplateParamsController) {
    client := TemplateParamsController{baseController: baseController}
    return &client
}

