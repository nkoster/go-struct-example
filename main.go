package main

import "fmt"

type contactStruct struct {
	email string
	phone string
}

type person struct {
	firstName string
	lastName  string
	contact   contactStruct
}

func main() {
	// dangerous assignment, order is important
	p1 := person{"Niels", "Koster", contactStruct{"n@w.l", "06"}}
	fmt.Println(p1)
	// safe assignment
	p2 := person{firstName: "Gekke", lastName: "Henkie"}
	fmt.Println(p2)
	fmt.Println(p2.firstName, p2.lastName)
	fmt.Printf("%+v\n", p2)
	p3 := person{
		firstName: "Jan",
		lastName:  "Snel",
		contact: contactStruct{
			email: "jan@w3b.net",
			phone: "061111111",
		}, // <-- those last  ^  commas almost killed me, crazy!
	}
	fmt.Printf("%+v\n", p3)
	fmt.Println(p3.firstName, p3.contact.email)
}
