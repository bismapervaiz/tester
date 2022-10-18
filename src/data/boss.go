/*
tester

This file was automatically generated for Stamplay by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package data

import (
    "encoding/json"
    "fmt"
    "log"
    "tester/src/utilities"
    "time"
)

// Testing circular reference.
type Boss struct {
    Employee   
    PromotedAt time.Time         `json:"promotedAt"`          
    Assistant  EmployeeInterface `json:"assistant,omitempty"` 
}

// Getter for promotedAt.
func (b *Boss) GetPromotedAt() (
    time.Time) {
    return b.PromotedAt
}

// Setter for promotedAt.
func (b *Boss) SetPromotedAt(promotedAt time.Time) {
    b.PromotedAt = promotedAt
}

// Getter for assistant.
func (b *Boss) GetAssistant() (
    EmployeeInterface) {
    return b.Assistant
}

// Setter for assistant.
func (b *Boss) SetAssistant(assistant EmployeeInterface) {
    b.Assistant = assistant
}

func (b Boss) MarshalJSON() (
    []byte,
    error) {
    b.PersonType = "Boss"
    type marshallable Boss
    return json.Marshal(struct {
        marshallable 
        Birthday     string `json:"birthday"`   
        Birthtime    string `json:"birthtime"`  
        HiredAt      string `json:"hiredAt"`    
        PromotedAt   int64  `json:"promotedAt"` 
    }{
        Birthday:     b.Birthday.Format(utilities.DEFAULT_DATE), 
        Birthtime:    b.Birthtime.Format(time.RFC3339),          
        HiredAt:      b.HiredAt.Format(time.RFC1123),            
        PromotedAt:   b.PromotedAt.Unix(),                       
        marshallable: marshallable(b),                           
    })
}

func (b *Boss) UnmarshalJSON(input []byte) (
    error) {
    temp := &struct {
        Address     string        
        Age         int64         
        Birthday    string        
        Birthtime   string        
        Department  string        
        Dependents  []PersonField 
        HiredAt     string        
        JoiningDay  Days          
        Name        string        
        PromotedAt  int64         
        Salary      int64         
        Uid         string        
        WorkingDays []Days        
        Assistant   EmployeeField 
        Boss        PersonField   
        PersonType  string        
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    b.Address = temp.Address
    b.Age = temp.Age
    b.Birthday, err = time.Parse(utilities.DEFAULT_DATE, temp.Birthday)
    if err != nil {
        log.Fatalf("Cannot Parse birthday as % s format.", utilities.DEFAULT_DATE)
    }
    b.Birthtime, err = time.Parse(time.RFC3339, temp.Birthtime)
    if err != nil {
        log.Fatalf("Cannot Parse birthtime as % s format.", time.RFC3339)
    }
    b.Department = temp.Department
    b.HiredAt, err = time.Parse(time.RFC1123, temp.HiredAt)
    if err != nil {
        log.Fatalf("Cannot Parse hiredAt as % s format.", time.RFC1123)
    }
    b.JoiningDay = temp.JoiningDay
    b.Name = temp.Name
    b.PromotedAt = time.Unix(temp.PromotedAt, 0)
    b.Salary = temp.Salary
    b.Uid = temp.Uid
    b.WorkingDays = temp.WorkingDays
    b.Assistant = temp.Assistant
    b.Boss = temp.Boss
    b.PersonType = temp.PersonType
    return nil
}

func UnmarshalBoss(input []byte) (
    BossInterface,
    error) {
    person, err := UnmarshalPerson(input)
    if err != nil {
        return nil, err
    }
    
    if boss, ok := person.(BossInterface); ok {
        return boss, nil
    }
    return nil, fmt.Errorf("Cannot unmarshal boss %v", person)
}

func ToBossArray(boss []BossField) (
    []BossInterface) {
    var items []BossInterface
    for _, temp := range boss {
        items = append(items, temp.BossInterface)
    }
    return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type BossField struct {
    BossInterface 
}

func (b *BossField) UnmarshalJSON(input []byte) (
    error) {
    var err error
    b.BossInterface, err = UnmarshalBoss(input)
    return err
}

