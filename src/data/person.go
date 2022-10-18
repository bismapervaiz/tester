/*
tester

This file was automatically generated for Stamplay by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package data

import (
    "encoding/json"
    "log"
    "tester/src/utilities"
    "time"
)

type Person struct {
    Address    string    `json:"address"`              
    Age        int64     `json:"age"`                  
    Birthday   time.Time `json:"birthday"`             
    Birthtime  time.Time `json:"birthtime"`            
    Name       string    `json:"name"`                 
    Uid        string    `json:"uid"`                  
    PersonType string    `json:"personType,omitempty"` 
}

// Getter for address.
func (p *Person) GetAddress() (
    string) {
    return p.Address
}

// Setter for address.
func (p *Person) SetAddress(address string) {
    p.Address = address
}

// Getter for age.
func (p *Person) GetAge() (
    int64) {
    return p.Age
}

// Setter for age.
func (p *Person) SetAge(age int64) {
    p.Age = age
}

// Getter for birthday.
func (p *Person) GetBirthday() (
    time.Time) {
    return p.Birthday
}

// Setter for birthday.
func (p *Person) SetBirthday(birthday time.Time) {
    p.Birthday = birthday
}

// Getter for birthtime.
func (p *Person) GetBirthtime() (
    time.Time) {
    return p.Birthtime
}

// Setter for birthtime.
func (p *Person) SetBirthtime(birthtime time.Time) {
    p.Birthtime = birthtime
}

// Getter for name.
func (p *Person) GetName() (
    string) {
    return p.Name
}

// Setter for name.
func (p *Person) SetName(name string) {
    p.Name = name
}

// Getter for uid.
func (p *Person) GetUid() (
    string) {
    return p.Uid
}

// Setter for uid.
func (p *Person) SetUid(uid string) {
    p.Uid = uid
}

// Getter for personType.
func (p *Person) GetPersonType() (
    string) {
    return p.PersonType
}

// Setter for personType.
func (p *Person) SetPersonType(personType string) {
    p.PersonType = personType
}

func (p Person) MarshalJSON() (
    []byte,
    error) {
    p.PersonType = "Per"
    type marshallable Person
    return json.Marshal(struct {
        marshallable 
        Birthday     string `json:"birthday"`  
        Birthtime    string `json:"birthtime"` 
    }{
        Birthday:     p.Birthday.Format(utilities.DEFAULT_DATE), 
        Birthtime:    p.Birthtime.Format(time.RFC3339),          
        marshallable: marshallable(p),                           
    })
}

func (p *Person) UnmarshalJSON(input []byte) (
    error) {
    temp := &struct {
        Address    string 
        Age        int64  
        Birthday   string 
        Birthtime  string 
        Name       string 
        Uid        string 
        PersonType string 
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    p.Address = temp.Address
    p.Age = temp.Age
    p.Birthday, err = time.Parse(utilities.DEFAULT_DATE, temp.Birthday)
    if err != nil {
        log.Fatalf("Cannot Parse birthday as % s format.", utilities.DEFAULT_DATE)
    }
    p.Birthtime, err = time.Parse(time.RFC3339, temp.Birthtime)
    if err != nil {
        log.Fatalf("Cannot Parse birthtime as % s format.", time.RFC3339)
    }
    p.Name = temp.Name
    p.Uid = temp.Uid
    p.PersonType = temp.PersonType
    return nil
}

func UnmarshalPerson(input []byte) (
    PersonInterface,
    error) {
    discrim := &struct {
        PersonType *string
    }{}
    
    err := json.Unmarshal(input, &discrim)
    if err != nil {
        return nil, err
    }
    if discrim == nil {
        *discrim.PersonType = "Per"
    }
    
    var res PersonInterface
    switch *discrim.PersonType {
    case "Empl":
        res = &Employee{}
    case "Boss":
        res = &Boss{}
    default:
        res = &Person{}
    }
    json.Unmarshal(input, res)
    return res, nil
}

func ToPersonArray(person []PersonField) (
    []PersonInterface) {
    var items []PersonInterface
    for _, temp := range person {
        items = append(items, temp.PersonInterface)
    }
    return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type PersonField struct {
    PersonInterface 
}

func (p *PersonField) UnmarshalJSON(input []byte) (
    error) {
    var err error
    p.PersonInterface, err = UnmarshalPerson(input)
    return err
}

