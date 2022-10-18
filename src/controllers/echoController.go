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

type EchoController struct {
    baseController 
}

// TODO: type endpoint description here
func (e *EchoController) QueryEcho() (
    data.EchoResponse,
    *http.Response,
    error) {
    req := e.prepareRequest("GET", "/")
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    
    var result data.EchoResponse
    for {
    	if err := decoder.Decode(&result); err == io.EOF {
    		break
    	} else if err != nil {
    		log.Fatal(err)
    	}
    }
    
    return result, resp, err
}

// Echo's back the request
func (e *EchoController) JsonEcho(input interface{}) (
    interface{},
    *http.Response,
    error) {
    req := e.prepareRequest("POST", "/")
    req.Json(input)
    req.Authenticate(true)
    stream, resp, err := req.CallAsStream()
    return stream, resp, err
}

// Sends the request including any form params as JSON
func (e *EchoController) FormEcho(input interface{}) (
    interface{},
    *http.Response,
    error) {
    req := e.prepareRequest("POST", "/")
    req.FormParam("input", input)
    req.Authenticate(true)
    stream, resp, err := req.CallAsStream()
    return stream, resp, err
}

