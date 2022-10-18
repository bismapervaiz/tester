/*
tester

This file was automatically generated for Stamplay by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package controllers

import (
    "net/http"
    "tester/src/exceptions"
)

type ErrorCodesController struct {
    baseController 
}

// TODO: type endpoint description here
func (e *ErrorCodesController) Get500() (
    interface{},
    *http.Response,
    error) {
    req := e.prepareRequest("GET", "/error/500")
    stream, resp, err := req.CallAsStream()
    return stream, resp, err
}

// TODO: type endpoint description here
func (e *ErrorCodesController) Get400() (
    interface{},
    *http.Response,
    error) {
    req := e.prepareRequest("GET", "/error/400")
    stream, resp, err := req.CallAsStream()
    return stream, resp, err
}

// TODO: type endpoint description here
func (e *ErrorCodesController) Get501() (
    interface{},
    *http.Response,
    error) {
    req := e.prepareRequest("GET", "/error/501")
    stream, resp, err := req.CallAsStream()
    if resp.StatusCode == 501 {
        		err = exceptions.NewNestedModelException(501, "error 501")
    }
    return stream, resp, err
}

// TODO: type endpoint description here
func (e *ErrorCodesController) Catch412GlobalError() (
    interface{},
    *http.Response,
    error) {
    req := e.prepareRequest("GET", "/error/412")
    stream, resp, err := req.CallAsStream()
    return stream, resp, err
}

// TODO: type endpoint description here
func (e *ErrorCodesController) Get401() (
    interface{},
    *http.Response,
    error) {
    req := e.prepareRequest("GET", "/error/401")
    stream, resp, err := req.CallAsStream()
    if resp.StatusCode == 401 {
        		err = exceptions.NewLocalTestException(401, "401 Local")
    }
    if resp.StatusCode == 421 {
        		err = exceptions.NewLocalTestException(421, "Default")
    }
    if resp.StatusCode == 431 {
        		err = exceptions.NewLocalTestException(431, "Default")
    }
    if resp.StatusCode == 432 {
        		err = exceptions.NewLocalTestException(432, "Default")
    }
    if resp.StatusCode == 441 {
        		err = exceptions.NewLocalTestException(441, "Default")
    }
    if resp.StatusCode == 0 {
        		err = exceptions.NewLocalTestException(0, "Invalid response.")
    }
    return stream, resp, err
}

