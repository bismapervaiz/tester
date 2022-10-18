/*
tester

This file was automatically generated for Stamplay by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package data

import (
    "time"
)

type DeleteBodyInterface interface {
    GetName() (
        string)
    SetName(name string)
    GetField() (
        string)
    SetField(field string)
}

// Testing circular reference.
type BossInterface interface {
    EmployeeInterface
    GetPromotedAt() (
        time.Time)
    SetPromotedAt(promotedAt time.Time)
    GetAssistant() (
        EmployeeInterface)
    SetAssistant(assistant EmployeeInterface)
}

// Query parameter key value pair echoed back from the server.
type QueryParameterInterface interface {
    GetKey() (
        string)
    SetKey(key string)
}

// To test specific local exceptions.
type LocalTestExceptionInterface interface {
    GlobalTestExceptionInterface
    GetSecretMessageForEndpoint() (
        string)
    SetSecretMessageForEndpoint(secretMessageForEndpoint string)
}

type EmployeeInterface interface {
    PersonInterface
    GetDepartment() (
        string)
    SetDepartment(department string)
    GetDependents() (
        []PersonInterface)
    SetDependents(dependents []PersonInterface)
    GetHiredAt() (
        time.Time)
    SetHiredAt(hiredAt time.Time)
    GetJoiningDay() (
        Days)
    SetJoiningDay(joiningDay Days)
    GetSalary() (
        int64)
    SetSalary(salary int64)
    GetWorkingDays() (
        []Days)
    SetWorkingDays(workingDays []Days)
    GetBoss() (
        PersonInterface)
    SetBoss(boss PersonInterface)
}

// A string enum representing days of the week
type DaysInterface interface {
    GetSunday() (
        string)
    SetSunday(sunday string)
    GetMonday() (
        string)
    SetMonday(monday string)
    GetTuesday() (
        string)
    SetTuesday(tuesday string)
    GetWednesday() (
        string)
    SetWednesday(wednesday string)
    GetThursday() (
        string)
    SetThursday(thursday string)
    GetFriDay() (
        string)
    SetFriDay(friDay string)
    GetSaturday() (
        string)
    SetSaturday(saturday string)
}

// To test specific global exceptions.
type GlobalTestExceptionInterface interface {
    GetServerMessage() (
        string)
    SetServerMessage(serverMessage string)
    GetServerCode() (
        int64)
    SetServerCode(serverCode int64)
}

// Raw http Request info
type EchoResponseInterface interface {
    GetBody() (
        map[string]interface{})
    SetBody(body map[string]interface{})
    GetHeaders() (
        map[string]string)
    SetHeaders(headers map[string]string)
    GetMethod() (
        string)
    SetMethod(method string)
    GetPath() (
        string)
    SetPath(path string)
    GetQuery() (
        map[string]QueryParameter)
    SetQuery(query map[string]QueryParameter)
    GetUploadCount() (
        int64)
    SetUploadCount(uploadCount int64)
}

// A integer based enum representing a Suite in a game of cards
type SuiteCodeInterface interface {
    GetHearts() (
        int64)
    SetHearts(hearts int64)
    GetSpades() (
        int64)
    SetSpades(spades int64)
    GetClubs() (
        int64)
    SetClubs(clubs int64)
    GetDiamonds() (
        int64)
    SetDiamonds(diamonds int64)
}

type PersonInterface interface {
    GetAddress() (
        string)
    SetAddress(address string)
    GetAge() (
        int64)
    SetAge(age int64)
    GetBirthday() (
        time.Time)
    SetBirthday(birthday time.Time)
    GetBirthtime() (
        time.Time)
    SetBirthtime(birthtime time.Time)
    GetName() (
        string)
    SetName(name string)
    GetUid() (
        string)
    SetUid(uid string)
    GetPersonType() (
        string)
    SetPersonType(personType string)
}

type ServerResponseInterface interface {
    GetPassed() (
        bool)
    SetPassed(passed bool)
    GetMessage() (
        string)
    SetMessage(message string)
    GetInput() (
        map[string]interface{})
    SetInput(input map[string]interface{})
}

type NestedModelExceptionInterface interface {
    GetServerMessage() (
        string)
    SetServerMessage(serverMessage string)
    GetServerCode() (
        string)
    SetServerCode(serverCode string)
    GetModel() (
        Validate)
    SetModel(model Validate)
}

type ValidateInterface interface {
    GetField() (
        string)
    SetField(field string)
    GetName() (
        string)
    SetName(name string)
    GetAddress() (
        string)
    SetAddress(address string)
}

type ParamFormatInterface interface {
    GetTemplate() (
        string)
    SetTemplate(template string)
    GetForm() (
        string)
    SetForm(form string)
    GetBody() (
        string)
    SetBody(body string)
    GetHeader() (
        string)
    SetHeader(header string)
    GetQuery() (
        string)
    SetQuery(query string)
}

type AttributesInterface interface {
    GetExclusiveMaximum() (
        bool)
    SetExclusiveMaximum(exclusiveMaximum bool)
    GetExclusiveMinimum() (
        bool)
    SetExclusiveMinimum(exclusiveMinimum bool)
    GetId() (
        string)
    SetId(id string)
}

type TypeInterface interface {
    GetLong() (
        string)
    SetLong(long string)
    GetNumber() (
        string)
    SetNumber(number string)
    GetPrecision() (
        string)
    SetPrecision(precision string)
    GetBoolean() (
        string)
    SetBoolean(boolean string)
    GetDateTime() (
        string)
    SetDateTime(dateTime string)
    GetDate() (
        string)
    SetDate(date string)
    GetString() (
        string)
    SetString(mString string)
}

type ResponseWithEnumInterface interface {
    GetParamFormat() (
        ParamFormat)
    SetParamFormat(paramFormat ParamFormat)
    GetOptional() (
        bool)
    SetOptional(optional bool)
    GetType() (
        Type)
    SetType(mType Type)
    GetConstant() (
        bool)
    SetConstant(constant bool)
    GetIsArray() (
        bool)
    SetIsArray(isArray bool)
    GetIsStream() (
        bool)
    SetIsStream(isStream bool)
    GetIsAttribute() (
        bool)
    SetIsAttribute(isAttribute bool)
    GetIsMap() (
        bool)
    SetIsMap(isMap bool)
    GetAttributes() (
        Attributes)
    SetAttributes(attributes Attributes)
    GetNullable() (
        bool)
    SetNullable(nullable bool)
    GetId() (
        string)
    SetId(id string)
    GetName() (
        string)
    SetName(name string)
    GetDescription() (
        string)
    SetDescription(description string)
}

