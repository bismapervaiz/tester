/*
tester

This file was automatically generated for Stamplay by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package data

// A string enum representing days of the week
type Days string

const (
    Days_SUNDAY    Days = "Sunday"    
    Days_MONDAY    Days = "Monday"    
    Days_TUESDAY   Days = "Tuesday"   
    Days_WEDNESDAY Days = "Wednesday" 
    Days_THURSDAY  Days = "Thursday"  
    Days_FRIDAY    Days = "Friday"    
    Days_SATURDAY  Days = "Saturday"  
)

// A integer based enum representing a Suite in a game of cards
type SuiteCode int64

const (
    SuiteCode_HEARTS   SuiteCode = 1 
    SuiteCode_SPADES   SuiteCode = 2 
    SuiteCode_CLUBS    SuiteCode = 3 
    SuiteCode_DIAMONDS SuiteCode = 4 
)

type ParamFormat string

const (
    ParamFormat_TEMPLATE ParamFormat = "Template" 
    ParamFormat_FORM     ParamFormat = "Form"     
    ParamFormat_BODY     ParamFormat = "Body"     
    ParamFormat_HEADER   ParamFormat = "Header"   
    ParamFormat_QUERY    ParamFormat = "Query"    
)

type Type string

const (
    Type_LONG      Type = "Long"      
    Type_NUMBER    Type = "Number"    
    Type_PRECISION Type = "Precision" 
    Type_BOOLEAN   Type = "Boolean"   
    Type_DATETIME  Type = "DateTime"  
    Type_DATE      Type = "Date"      
    Type_STRING    Type = "String"    
)

