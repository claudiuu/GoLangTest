package main

import "fmt"

type person struct {
	lname string
	fname string
}

type secretAgent struct {
	agent     person
	topSecret bool
}

type human interface {
	walk()
}

func (p person) pSpeak() {
	fmt.Println("My name is", p.fname, p.lname)
}

func (sa secretAgent) saSpeak() {
	fmt.Println("I am a secret agent")
}

func (p person) walk() {
	fmt.Println("walking around as a person")
}

func (sa secretAgent) walk() {
	fmt.Println("walking around as a secret agent")
}

func move(h human) {
	h.walk()
}

func main() {
	p := person{
		"Mateias",
		"Claudiu",
	}

	sa := secretAgent{
		person{
			"Bond",
			"James",
		},
		true,
	}

	fmt.Println(p.lname)
	p.pSpeak()

	fmt.Println(sa.topSecret)
	sa.saSpeak()
	sa.agent.pSpeak()

	move(p)
	move(sa)

}
