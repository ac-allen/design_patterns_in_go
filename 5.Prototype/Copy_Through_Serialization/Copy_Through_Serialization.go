package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct{
	Name string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println("\n\n\n", string(b.Bytes()), "\n\n\n")

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)
	return &result
}

func (p *Person) print(){
	fmt.Println(p, p.Address)
}

func main(){
	john := Person{"John", 
	&Address{"Street 1", "City 1", "Country 1"},
	[]string{"Chris", "Matt"}}

	john.print()

	jane := john.DeepCopy()

	jane.Name = "Jane"
	jane.Address.StreetAddress = "Street 2"
	jane.Friends = append(jane.Friends, "New friend")

	john.print()
	jane.print()
}