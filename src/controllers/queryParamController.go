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

type QueryParamController struct {
    baseController 
}

// TODO: type endpoint description here
func (q *QueryParamController) Date(date time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := q.prepareRequest("GET", "/query/date")
    req.QueryParam("date", date.Format(utilities.DEFAULT_DATE))
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
func (q *QueryParamController) DateArray(dates []time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := q.prepareRequest("GET", "/query/datearray")
    req.QueryParam("dates", utilities.TimeToStringSlice(dates, utilities.DEFAULT_DATE))
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
func (q *QueryParamController) NoParams() (
    data.ServerResponse,
    *http.Response,
    error) {
    req := q.prepareRequest("GET", "/query/noparams")
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

// get optional dynamic query parameter
func (q *QueryParamController) OptionalDynamicQueryParam(name string) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := q.prepareRequest("GET", "/query/optionalQueryParam")
    req.QueryParam("name", name)
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
func (q *QueryParamController) UnixDateTimeArray(datetimes []time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := q.prepareRequest("GET", "/query/unixdatetimearray")
    req.QueryParam("datetimes", utilities.TimeToStringSlice(datetimes, time.UnixDate))
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
func (q *QueryParamController) UnixDateTime(datetime time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := q.prepareRequest("GET", "/query/unixdatetime")
    req.QueryParam("datetime", strconv.FormatInt(datetime.Unix(), 10))
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
func (q *QueryParamController) Rfc1123DateTime(datetime time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := q.prepareRequest("GET", "/query/rfc1123datetime")
    req.QueryParam("datetime", datetime.Format(time.RFC1123))
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
func (q *QueryParamController) Rfc1123DateTimeArray(datetimes []time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := q.prepareRequest("GET", "/query/rfc1123datetimearray")
    req.QueryParam("datetimes", utilities.TimeToStringSlice(datetimes, time.RFC1123))
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
func (q *QueryParamController) Rfc3339DateTimeArray(datetimes []time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := q.prepareRequest("GET", "/query/rfc3339datetimearray")
    req.QueryParam("datetimes", utilities.TimeToStringSlice(datetimes, time.RFC3339))
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
func (q *QueryParamController) Rfc3339DateTime(datetime time.Time) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := q.prepareRequest("GET", "/query/rfc3339datetime")
    req.QueryParam("datetime", datetime.Format(time.RFC3339))
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
func (q *QueryParamController) StringParam(mString string) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := q.prepareRequest("GET", "/query/stringparam")
    req.QueryParam("string", mString)
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
func (q *QueryParamController) UrlParam(url string) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := q.prepareRequest("GET", "/query/urlparam")
    req.QueryParam("url", url)
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
func (q *QueryParamController) NumberArray(integers []int64) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := q.prepareRequest("GET", "/query/numberarray")
    req.QueryParam("integers", integers)
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
func (q *QueryParamController) StringArray(strings []string) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := q.prepareRequest("GET", "/query/stringarray")
    req.QueryParam("strings", strings)
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
func (q *QueryParamController) SimpleQuery(
    boolean bool,
    number int64,
    mString string) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := q.prepareRequest("GET", "/query")
    req.QueryParam("boolean", boolean)
    req.QueryParam("number", number)
    req.QueryParam("string", mString)
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
func (q *QueryParamController) StringEnumArray(days []data.Days) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := q.prepareRequest("GET", "/query/stringenumarray")
    req.QueryParam("days", days)
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
func (q *QueryParamController) MultipleParams(
    number int64,
    precision float64,
    mString string,
    url string) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := q.prepareRequest("GET", "/query/multipleparams")
    req.QueryParam("number", number)
    req.QueryParam("precision", precision)
    req.QueryParam("string", mString)
    req.QueryParam("url", url)
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
func (q *QueryParamController) IntegerEnumArray(suites []data.SuiteCode) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := q.prepareRequest("GET", "/query/integerenumarray")
    req.QueryParam("suites", suites)
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

