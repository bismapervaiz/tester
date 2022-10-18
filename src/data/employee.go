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

type Employee struct {
    Person      
    Department  string            `json:"department"`     
    Dependents  []PersonInterface `json:"dependents"`     
    HiredAt     time.Time         `json:"hiredAt"`        
    JoiningDay  Days              `json:"joiningDay"`     
    Salary      int64             `json:"salary"`         
    WorkingDays []Days            `json:"workingDays"`    
    Boss        PersonInterface   `json:"boss,omitempty"` 
}

// Getter for department.
func (e *Employee) GetDepartment() (
    string) {
    return e.Department
}

// Setter for department.
func (e *Employee) SetDepartment(department string) {
    e.Department = department
}

// Getter for dependents.
func (e *Employee) GetDependents() (
    []PersonInterface) {
    return e.Dependents
}

// Setter for dependents.
func (e *Employee) SetDependents(dependents []PersonInterface) {
    e.Dependents = dependents
}

// Getter for hiredAt.
func (e *Employee) GetHiredAt() (
    time.Time) {
    return e.HiredAt
}

// Setter for hiredAt.
func (e *Employee) SetHiredAt(hiredAt time.Time) {
    e.HiredAt = hiredAt
}

// Getter for joiningDay.
func (e *Employee) GetJoiningDay() (
    Days) {
    return e.JoiningDay
}

// Setter for joiningDay.
func (e *Employee) SetJoiningDay(joiningDay Days) {
    e.JoiningDay = joiningDay
}

// Getter for salary.
func (e *Employee) GetSalary() (
    int64) {
    return e.Salary
}

// Setter for salary.
func (e *Employee) SetSalary(salary int64) {
    e.Salary = salary
}

// Getter for workingDays.
func (e *Employee) GetWorkingDays() (
    []Days) {
    return e.WorkingDays
}

// Setter for workingDays.
func (e *Employee) SetWorkingDays(workingDays []Days) {
    e.WorkingDays = workingDays
}

// Getter for boss.
func (e *Employee) GetBoss() (
    PersonInterface) {
    return e.Boss
}

// Setter for boss.
func (e *Employee) SetBoss(boss PersonInterface) {
    e.Boss = boss
}

func (e Employee) MarshalJSON() (
    []byte,
    error) {
    e.PersonType = "Empl"
    type marshallable Employee
    return json.Marshal(struct {
        marshallable 
        Birthday     string `json:"birthday"`  
        Birthtime    string `json:"birthtime"` 
        HiredAt      string `json:"hiredAt"`   
    }{
        Birthday:     e.Birthday.Format(utilities.DEFAULT_DATE), 
        Birthtime:    e.Birthtime.Format(time.RFC3339),          
        HiredAt:      e.HiredAt.Format(time.RFC1123),            
        marshallable: marshallable(e),                           
    })
}

func (e *Employee) UnmarshalJSON(input []byte) (
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
        Salary      int64         
        Uid         string        
        WorkingDays []Days        
        Boss        PersonField   
        PersonType  string        
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    e.Address = temp.Address
    e.Age = temp.Age
    e.Birthday, err = time.Parse(utilities.DEFAULT_DATE, temp.Birthday)
    if err != nil {
        log.Fatalf("Cannot Parse birthday as % s format.", utilities.DEFAULT_DATE)
    }
    e.Birthtime, err = time.Parse(time.RFC3339, temp.Birthtime)
    if err != nil {
        log.Fatalf("Cannot Parse birthtime as % s format.", time.RFC3339)
    }
    e.Department = temp.Department
    e.Dependents = ToPersonArray(temp.Dependents)
    e.HiredAt, err = time.Parse(time.RFC1123, temp.HiredAt)
    if err != nil {
        log.Fatalf("Cannot Parse hiredAt as % s format.", time.RFC1123)
    }
    e.JoiningDay = temp.JoiningDay
    e.Name = temp.Name
    e.Salary = temp.Salary
    e.Uid = temp.Uid
    e.WorkingDays = temp.WorkingDays
    e.Boss = temp.Boss
    e.PersonType = temp.PersonType
    return nil
}

func UnmarshalEmployee(input []byte) (
    EmployeeInterface,
    error) {
    person, err := UnmarshalPerson(input)
    if err != nil {
        return nil, err
    }
    
    if employee, ok := person.(EmployeeInterface); ok {
        return employee, nil
    }
    return nil, fmt.Errorf("Cannot unmarshal employee %v", person)
}

func ToEmployeeArray(employee []EmployeeField) (
    []EmployeeInterface) {
    var items []EmployeeInterface
    for _, temp := range employee {
        items = append(items, temp.EmployeeInterface)
    }
    return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type EmployeeField struct {
    EmployeeInterface 
}

func (e *EmployeeField) UnmarshalJSON(input []byte) (
    error) {
    var err error
    e.EmployeeInterface, err = UnmarshalEmployee(input)
    return err
}

