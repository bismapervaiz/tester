/*
tester

This file was automatically generated for Stamplay by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package controllers

import (
    "io"
    "log"
    "net/http"
    "strconv"
    "tester/src/data"
    "tester/src/utilities"
    "time"
)

type ResponseTypesController struct {
    baseController 
}

// Gets a integer response
func (r *ResponseTypesController) GetInteger() (
    int64,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/integer")
    str, resp, err := req.CallAsText()
    var result int64
    result, err = strconv.ParseInt(str, 10, 64)
    if err != nil {
        //error in parsing
		log.Fatalln(err)
    }

    return result, resp, err
}

// gets a binary object
func (r *ResponseTypesController) GetBinary() (
    []byte,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/binary")
    stream, resp, err := req.CallAsStream()
    return stream, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) GetPrecision() (
    float64,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/precision")
    str, resp, err := req.CallAsText()
    var result float64
    result, err = strconv.ParseFloat(str, 64)
    if err != nil {
        //error in parsing
		log.Fatalln(err)
    }

    return result, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) GetModelArray(array bool) (
    []data.PersonInterface,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/model")
    req.QueryParam("array", "true")
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result []data.PersonInterface
    for {
    	if err := decoder.Decode(&result); err == io.EOF {
    		break
    	} else if err != nil {
    		log.Fatal(err)
    	}
    }
    
    return result, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) GetModel() (
    data.PersonInterface,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/model")
    decoder, resp, err := req.CallAsJson()
    
    var result data.PersonInterface
    for {
    	if err := decoder.Decode(&result); err == io.EOF {
    		break
    	} else if err != nil {
    		log.Fatal(err)
    	}
    }
    
    return result, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) GetLong() (
    int64,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/long")
    req.BaseUrl("default")
    str, resp, err := req.CallAsText()
    var result int64
    result, err = strconv.ParseInt(str, 10, 64)
    if err != nil {
        //error in parsing
		log.Fatalln(err)
    }

    return result, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) GetDateArray(array bool) (
    []time.Time,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/date")
    req.QueryParam("array", "true")
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    result := utilities.ToTimeSlice(
        utilities.JsonDecoderToStringSlice(decoder),
        utilities.DEFAULT_DATE)
    return result, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) GetDate() (
    time.Time,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/date")
    str, resp, err := req.CallAsText()
    var result time.Time
    result, err = time.Parse(utilities.DEFAULT_DATE, str)
    if err != nil {
        //error in parsing
		log.Fatalln(err)
    }

    return result, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) GetContentTypeHeaders() (
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/getContentType")
    context, err := req.Call()
    return context.Response, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) GetUnixDateTime() (
    time.Time,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/unixdatetime")
    str, resp, err := req.CallAsText()
    var result time.Time
    timestamp, err := strconv.ParseInt(str, 10, 64)
    if err != nil {
        //error in parsing
		log.Fatalln(err)
    }
    result = time.Unix(timestamp, 0)

    return result, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) Get1123DateTime() (
    time.Time,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/1123datetime")
    str, resp, err := req.CallAsText()
    var result time.Time
    result, err = time.Parse(time.RFC1123, str)
    if err != nil {
        //error in parsing
		log.Fatalln(err)
    }

    return result, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) GetHeaders() (
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/headers")
    context, err := req.Call()
    return context.Response, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) GetBoolean() (
    bool,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/boolean")
    str, resp, err := req.CallAsText()
    var result bool
    result, err = strconv.ParseBool(str)
    if err != nil {
        //error in parsing
		log.Fatalln(err)
    }

    return result, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) Get3339Datetime() (
    time.Time,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/3339datetime")
    str, resp, err := req.CallAsText()
    var result time.Time
    result, err = time.Parse(time.RFC3339, str)
    if err != nil {
        //error in parsing
		log.Fatalln(err)
    }

    return result, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) GetDynamic() (
    interface{},
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/dynamic")
    stream, resp, err := req.CallAsStream()
    return stream, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) ReturnResponseWithEnums() (
    data.ResponseWithEnum,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/responseWitEnum")
    decoder, resp, err := req.CallAsJson()
    
    var result data.ResponseWithEnum
    for {
    	if err := decoder.Decode(&result); err == io.EOF {
    		break
    	} else if err != nil {
    		log.Fatal(err)
    	}
    }
    
    return result, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) GetStringEnumArray(
    array bool,
    mType string) (
    []data.Days,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/enum")
    req.QueryParam("array", "true")
    req.QueryParam("type", "string")
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result []data.Days
    for {
    	if err := decoder.Decode(&result); err == io.EOF {
    		break
    	} else if err != nil {
    		log.Fatal(err)
    	}
    }
    
    return result, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) GetStringEnum(mType string) (
    data.Days,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/enum")
    req.QueryParam("type", "string")
    req.Authenticate(true)
    str, resp, err := req.CallAsText()
    var result data.Days
    i, err := strconv.ParseInt(str, 10, 64)
    if err != nil {
        //error in parsing
		log.Fatalln(err)
    }
    result = data.Days(i)

    return result, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) GetIntEnum(mType string) (
    data.SuiteCode,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/enum")
    req.QueryParam("type", "int")
    req.Authenticate(true)
    str, resp, err := req.CallAsText()
    var result data.SuiteCode
    i, err := strconv.ParseInt(str, 10, 64)
    if err != nil {
        //error in parsing
		log.Fatalln(err)
    }
    result = data.SuiteCode(i)

    return result, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) GetIntEnumArray(
    array bool,
    mType string) (
    []data.SuiteCode,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/enum")
    req.QueryParam("array", "true")
    req.QueryParam("type", "int")
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result []data.SuiteCode
    for {
    	if err := decoder.Decode(&result); err == io.EOF {
    		break
    	} else if err != nil {
    		log.Fatal(err)
    	}
    }
    
    return result, resp, err
}

// Get an array of integers.
func (r *ResponseTypesController) GetIntegerArray(array bool) (
    []int64,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/integer")
    req.QueryParam("array", "true")
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    result := utilities.JsonDecoderToIntSlice(decoder)
    return result, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) GetDynamicArray(array bool) (
    interface{},
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/dynamic")
    req.QueryParam("array", "true")
    req.Authenticate(true)
    stream, resp, err := req.CallAsStream()
    return stream, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) Get3339DatetimeArray(array bool) (
    []time.Time,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/3339datetime")
    req.QueryParam("array", "true")
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    result := utilities.ToTimeSlice(
        utilities.JsonDecoderToStringSlice(decoder),
        time.RFC3339)
    return result, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) GetBooleanArray(array bool) (
    []bool,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/boolean")
    req.QueryParam("array", "true")
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    result := utilities.JsonDecoderToBooleanSlice(decoder)
    return result, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) Get1123DateTimeArray(array bool) (
    []time.Time,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/1123datetime")
    req.QueryParam("array", "true")
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    result := utilities.ToTimeSlice(
        utilities.JsonDecoderToStringSlice(decoder),
        time.RFC1123)
    return result, resp, err
}

// TODO: type endpoint description here
func (r *ResponseTypesController) GetUnixDateTimeArray(array bool) (
    []time.Time,
    *http.Response,
    error) {
    req := r.prepareRequest("GET", "/response/unixdatetime")
    req.QueryParam("array", "true")
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    result := utilities.ToTimeSlice(
        utilities.JsonDecoderToStringSlice(decoder),
        time.UnixDate)
    return result, resp, err
}

