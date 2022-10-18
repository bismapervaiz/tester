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

type FormParamsController struct {
    baseController 
}

// TODO: type endpoint description here
func (f *FormParamsController) SendFile(file []byte) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("POST", "/form/file")
    req.FormParam("file", file)
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
func (f *FormParamsController) SendModel(model data.EmployeeInterface) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("POST", "/form/model")
    req.FormParam("model", model)
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
func (f *FormParamsController) SendDate(date time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("POST", "/form/date")
    req.FormParam("date", date.Format(utilities.DEFAULT_DATE))
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
func (f *FormParamsController) SendDeleteForm(body data.DeleteBody) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("DELETE", "/form/deleteForm")
    req.FormParam("body", body)
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
func (f *FormParamsController) SendDeleteMultipart(file []byte) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("DELETE", "/form/deleteMultipart")
    req.FormParam("file", file)
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
func (f *FormParamsController) SendDateArray(
    array bool,
    dates []time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("POST", "/form/date")
    req.QueryParam("array", "true")
    req.FormParam("dates", utilities.TimeToStringSlice(dates, utilities.DEFAULT_DATE))
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
func (f *FormParamsController) SendUnixDateTime(datetime time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("POST", "/form/unixdatetime")
    req.FormParam("datetime", strconv.FormatInt(datetime.Unix(), 10))
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
func (f *FormParamsController) SendRfc1123DateTime(datetime time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("POST", "/form/rfc1123datetime")
    req.FormParam("datetime", datetime.Format(time.RFC1123))
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
func (f *FormParamsController) SendRfc3339DateTime(datetime time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("POST", "/form/rfc3339datetime")
    req.FormParam("datetime", datetime.Format(time.RFC3339))
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
func (f *FormParamsController) SendUnixDateTimeArray(
    array bool,
    datetimes []time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("POST", "/form/unixdatetime")
    req.QueryParam("array", "true")
    req.FormParam("datetimes", utilities.TimeToStringSlice(datetimes, time.UnixDate))
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
func (f *FormParamsController) SendRfc1123DateTimeArray(
    array bool,
    datetimes []time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("POST", "/form/rfc1123datetime")
    req.QueryParam("array", "true")
    req.FormParam("datetimes", utilities.TimeToStringSlice(datetimes, time.RFC1123))
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
func (f *FormParamsController) SendLong(value int64) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("POST", "/form/number")
    req.FormParam("value", value)
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
func (f *FormParamsController) SendIntegerArray(
    array bool,
    integers []int64) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("POST", "/form/number")
    req.QueryParam("array", "true")
    req.FormParam("integers", integers)
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
func (f *FormParamsController) SendStringArray(
    array bool,
    strings []string) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("POST", "/form/string")
    req.QueryParam("array", "true")
    req.FormParam("strings", strings)
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
func (f *FormParamsController) SendModelArray(
    array bool,
    models []data.EmployeeInterface) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("POST", "/form/model")
    req.QueryParam("array", "true")
    req.FormParam("models", models)
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
func (f *FormParamsController) SendString(value string) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("POST", "/form/string")
    req.FormParam("value", value)
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
func (f *FormParamsController) SendRfc3339DateTimeArray(
    array bool,
    datetimes []time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("POST", "/form/rfc3339datetime")
    req.QueryParam("array", "true")
    req.FormParam("datetimes", utilities.TimeToStringSlice(datetimes, time.RFC3339))
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

// Send a variety for form params. Returns file count and body params
func (f *FormParamsController) SendMixedArray(
    file []byte,
    integers []int64,
    models []data.EmployeeInterface,
    strings []string) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("POST", "/form/mixed")
    req.FormParam("file", file)
    req.FormParam("integers", integers)
    req.FormParam("models", models)
    req.FormParam("strings", strings)
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
func (f *FormParamsController) UpdateModelWithForm(model data.EmployeeInterface) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("PUT", "/form/updateModel")
    req.FormParam("model", model)
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
func (f *FormParamsController) SendDeleteForm1(model data.EmployeeInterface) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("DELETE", "/form/deleteForm1")
    req.FormParam("model", model)
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
func (f *FormParamsController) SendDeleteFormWithModelArray(
    array bool,
    models []data.EmployeeInterface) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("DELETE", "/form/deleteForm1")
    req.QueryParam("array", "true")
    req.FormParam("models", models)
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
func (f *FormParamsController) UpdateModelArrayWithForm(
    array bool,
    models []data.EmployeeInterface) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("PUT", "/form/updateModel")
    req.QueryParam("array", "true")
    req.FormParam("models", models)
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
func (f *FormParamsController) UpdateStringWithForm(value string) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("PUT", "/form/updateString")
    req.FormParam("value", value)
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
func (f *FormParamsController) UpdateStringArrayWithForm(
    array bool,
    strings []string) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("PUT", "/form/updateString")
    req.QueryParam("array", "true")
    req.FormParam("strings", strings)
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
func (f *FormParamsController) SendIntegerEnumArray(
    array bool,
    suites []data.SuiteCode) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("POST", "/form/integerenum")
    req.QueryParam("array", "true")
    req.FormParam("suites", suites)
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
func (f *FormParamsController) SendStringEnumArray(
    array bool,
    days []data.Days) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := f.prepareRequest("POST", "/form/stringenum")
    req.QueryParam("array", "true")
    req.FormParam("days", days)
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

