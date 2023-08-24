package main

import (
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country,
	}
}

type Person struct{
	Name string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
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