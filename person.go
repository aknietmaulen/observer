package main

import "fmt"

type Person struct {
	id string //phone number
}

func newPerson(id string) *Person {
	return &Person{id: id}
}

func (p *Person) update(message string) {
	fmt.Printf("Sending notification to %s: %s\n", p.id, message)
}

func (p *Person) getID() string {
	return p.id
}
