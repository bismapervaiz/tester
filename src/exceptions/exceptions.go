/*
tester

This file was automatically generated for Stamplay by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package exceptions

import (
    "fmt"
    "net/http"
    "tester/src/data"
)

// This is the base struct for all exceptions that represent an error response from the server.
type ApiError struct {
    Request    http.Request      `json:"Request"`    
    StatusCode int64             `json:"StatusCode"` 
    Headers    map[string]string `json:"Headers"`    
    Body       string            `json:"Body"`       
}

// Constructor for ApiError.
func NewApiError(
    statusCode int64,
    body string) (
    *ApiError) {
    return &ApiError{
    	StatusCode: statusCode,
    	Body:       body,
    }
}

// Implementing the Error method for the error interface.
func (a *ApiError) Error() (
    string) {
    return fmt.Sprintf("ApiError occured %v", a.Body)
}

// To test specific local exceptions.
type LocalTestException struct {
    GlobalTestException      
    SecretMessageForEndpoint string `json:"SecretMessageForEndpoint"` 
}

// Constructor for LocalTestException.
func NewLocalTestException(
    statusCode int64,
    body string) (
    *LocalTestException) {
    return &LocalTestException{
    	GlobalTestException: GlobalTestException{
    	    ApiError: ApiError{
    		    StatusCode: statusCode,
    		    Body:       body,
    	    },
    	},
    }
}

// Implementing the Error method for the error interface.
func (l *LocalTestException) Error() (
    string) {
    return fmt.Sprintf("LocalTestException occured %v", l.Body)
}

// To test specific global exceptions.
type GlobalTestException struct {
    ApiError      
    ServerMessage string `json:"ServerMessage"` 
    ServerCode    int64  `json:"ServerCode"`    
}

// Constructor for GlobalTestException.
func NewGlobalTestException(
    statusCode int64,
    body string) (
    *GlobalTestException) {
    return &GlobalTestException{
    	ApiError: ApiError{
    		StatusCode: statusCode,
    		Body:       body,
    	},
    }
}

// Implementing the Error method for the error interface.
func (g *GlobalTestException) Error() (
    string) {
    return fmt.Sprintf("GlobalTestException occured %v", g.Body)
}

type NestedModelException struct {
    ApiError      
    ServerMessage string        `json:"ServerMessage"` 
    ServerCode    string        `json:"ServerCode"`    
    Model         data.Validate `json:"model"`         
}

// Constructor for NestedModelException.
func NewNestedModelException(
    statusCode int64,
    body string) (
    *NestedModelException) {
    return &NestedModelException{
    	ApiError: ApiError{
    		StatusCode: statusCode,
    		Body:       body,
    	},
    }
}

// Implementing the Error method for the error interface.
func (n *NestedModelException) Error() (
    string) {
    return fmt.Sprintf("NestedModelException occured %v", n.Body)
}

