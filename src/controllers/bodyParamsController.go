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

type BodyParamsController struct {
    baseController 
}

// TODO: type endpoint description here
func (b *BodyParamsController) SendModel(model data.EmployeeInterface) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("POST", "/body/model")
    req.Json(model)
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) SendDate(date time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("POST", "/body/date")
    req.Text(date.Format(utilities.DEFAULT_DATE))
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) UpdateModel(model data.EmployeeInterface) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("PUT", "/body/updateModel")
    req.Json(model)
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) SendDeletePlainText(textString string) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("DELETE", "/body/deletePlainTextBody")
    req.Text(textString)
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) SendDeleteBody(body data.DeleteBody) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("DELETE", "/body/deleteBody")
    req.Json(body)
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) SendDateArray(
    array bool,
    dates []time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("POST", "/body/date")
    req.QueryParam("array", "true")
    req.Json(utilities.TimeToStringSlice(dates, utilities.DEFAULT_DATE))
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) SendUnixDateTime(datetime time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("POST", "/body/unixdatetime")
    req.Text(strconv.FormatInt(datetime.Unix(), 10))
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) SendRfc1123DateTime(datetime time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("POST", "/body/rfc1123datetime")
    req.Text(datetime.Format(time.RFC1123))
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) SendRfc3339DateTime(datetime time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("POST", "/body/rfc3339datetime")
    req.Text(datetime.Format(time.RFC3339))
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) SendUnixDateTimeArray(
    array bool,
    datetimes []time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("POST", "/body/unixdatetime")
    req.QueryParam("array", "true")
    req.Json(utilities.TimeToStringSlice(datetimes, time.UnixDate))
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) SendRfc1123DateTimeArray(
    array bool,
    datetimes []time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("POST", "/body/rfc1123datetime")
    req.QueryParam("array", "true")
    req.Json(utilities.TimeToStringSlice(datetimes, time.RFC1123))
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) SendRfc3339DateTimeArray(
    array bool,
    datetimes []time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("POST", "/body/rfc3339datetime")
    req.QueryParam("array", "true")
    req.Json(utilities.TimeToStringSlice(datetimes, time.RFC3339))
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
    for {
    	if err := decoder.Decode(&result); err == io.EOF {
    		break
    	} else if err != nil {
    		log.Fatal(err)
    	}
    }
    
    return result, resp, err
}

// sends a string body param
func (b *BodyParamsController) SendStringArray(
    array bool,
    sarray []string) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("POST", "/body/string")
    req.QueryParam("array", "true")
    req.Json(sarray)
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) UpdateString(value string) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("PUT", "/body/updateString")
    req.Text(value)
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) SendIntegerArray(
    array bool,
    integers []int64) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("POST", "/body/number")
    req.QueryParam("array", "true")
    req.Json(integers)
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) SendModelArray(
    array bool,
    models []data.EmployeeInterface) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("POST", "/body/model")
    req.QueryParam("array", "true")
    req.Json(models)
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) SendDynamic(dynamic interface{}) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("POST", "/body/dynamic")
    req.Json(dynamic)
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) SendString(value string) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("POST", "/body/string")
    req.Text(value)
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) SendStringEnumArray(
    array bool,
    days []data.Days) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("POST", "/body/stringenum")
    req.QueryParam("array", "true")
    req.Json(days)
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) SendIntegerEnumArray(
    array bool,
    suites []data.SuiteCode) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("POST", "/body/integerenum")
    req.QueryParam("array", "true")
    req.Json(suites)
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) SendDeleteBodyWithModel(model data.EmployeeInterface) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("DELETE", "/body/deleteBody1")
    req.Json(model)
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) UpdateModelArray(
    array bool,
    models []data.EmployeeInterface) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("PUT", "/body/updateModel")
    req.QueryParam("array", "true")
    req.Json(models)
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
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
func (b *BodyParamsController) UpdateStringArray(
    array bool,
    strings []string) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := b.prepareRequest("PUT", "/body/updateString")
    req.QueryParam("array", "true")
    req.Json(strings)
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.ServerResponse
    for {
    	if err := decoder.Decode(&result); err == io.EOF {
    		break
    	} else if err != nil {
    		log.Fatal(err)
    	}
    }
    
    return result, resp, err
}

