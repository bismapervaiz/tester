/*
tester

This file was automatically generated for Stamplay by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package data

type ResponseWithEnum struct {
    ParamFormat ParamFormat `json:"paramFormat"` 
    Optional    bool        `json:"optional"`    
    Type        Type        `json:"type"`        
    Constant    bool        `json:"constant"`    
    IsArray     bool        `json:"isArray"`     
    IsStream    bool        `json:"isStream"`    
    IsAttribute bool        `json:"isAttribute"` 
    IsMap       bool        `json:"isMap"`       
    Attributes  Attributes  `json:"attributes"`  
    Nullable    bool        `json:"nullable"`    
    Id          string      `json:"id"`          
    Name        string      `json:"name"`        
    Description string      `json:"description"` 
}

