/*
tester

This file was automatically generated for Stamplay by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package configuration

import (
    "tester/src/data"
)

type Configuration struct {
    Environment Environment    
    Port        string         
    Suites      data.SuiteCode 
}

// This is a type declaration.
func DefaultConfig() (
    Configuration) {
    return Configuration{
        Environment: TESTING,
        Port: "80",
        Suites: 1,
    }
}

// Available Servers
type Server string

const (
    ENUMDEFAULT Server = "default"     
    AUTHSERVER  Server = "auth server" 
)

// Available Environments
type Environment string

const (
    PRODUCTION Environment = "production" 
    TESTING    Environment = "testing"    
)

