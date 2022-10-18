/*
tester

This file was automatically generated for Stamplay by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package controllers

import (
    "io"
    "log"
    "net/http"
    "tester/src/data"
)

type HeaderController struct {
    baseController 
}

// Sends a single header params
func (h *HeaderController) SendHeaders(
    customHeader string,
    value string) (
    data.ServerResponse,
    *http.Response,
    error) {
    req := h.prepareRequest("POST", "/header")
    req.Header("custom-header", customHeader)
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

