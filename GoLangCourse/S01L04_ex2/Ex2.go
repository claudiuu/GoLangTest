package main

import "fmt"

type person struct {
	lname   string
	fname   string
	favFood []string
}

type secretAgent struct {
	agent     person
	topSecret bool
}

type human interface {
	walk()
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle   vehicle
	fourWheel bool
}

type sedan struct {
	vehicle vehicle
	luxury  bool
}

type transportation interface {
	transportationDevice() string
}

type gator int

type swan bool

type swampCreature interface {
	greeting()
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

func (t truck) transportationDevice() string {
	return "I am a truck"
}

func (s sedan) transportationDevice() string {
	return "I am a sedan"
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
}

func (g gator) greeting() {
	fmt.Println("I am a gator")
}

func (s swan) greeting() {
	fmt.Println("I am a beautiful swamp")
}

func whatAreYou(sc swampCreature) {
	sc.greeting()
}

func main() {
	p := person{
		"Mateias",
		"Claudiu",
		[]string{"pizza", "broccoli"},
	}

	sa := secretAgent{
		person{
			"Bond",
			"James",
			[]string{},
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

	a := []int{11, 22, 33, 44}[:]
	for i, v := range a {
		fmt.Println(i, "-", v)
	}

	m := map[string]int{
		"one": 1,
		"two": 2,
	}
	// go over the map, use only the keys
	for k := range m {
		fmt.Println(k)
	}

	// go over the map, use both key and value
	for k, v := range m {
		fmt.Println(k, "=", v)
	}

	t1 := truck{
		vehicle: vehicle{
			4,
			"blue",
		},
		fourWheel: true,
	}

	fmt.Println(t1)
	fmt.Printf("Vehicle with %d doors;%s\n", t1.vehicle.doors, t1.transportationDevice())
	report(t1)

	s1 := sedan{
		vehicle{
			2,
			"red",
		},
		true,
	}

	fmt.Println(s1)
	fmt.Printf("%s vehicle with luxury flag %v: %+v; %s\n", s1.vehicle.color, s1.luxury, s1, s1.transportationDevice())
	report(s1)

	var g1 gator
	g1 = 12
	fmt.Printf("%T = %v\n", g1, g1)

	var x int
	x = 23
	fmt.Printf("%T = %v\n", x, x)
	x = int(g1)
	g1.greeting()

	var s2 swan
	s2 = false

	whatAreYou(s2)

	s := "i'm sory dave i can't do that"
	fmt.Println(s)
	//s = s[:]
	fmt.Println(s[:13])
	fmt.Println(s[14:])
	for _, w := range []byte(s) {
		fmt.Println(string(w))
	}
}
