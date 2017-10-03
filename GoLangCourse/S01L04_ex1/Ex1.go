package main

import "fmt"

type square struct {
	lat int
}

type circle struct {
	rad int
}

type shape interface {
	area() float64
}

const pi float64 = 3.14

func (s square) area() float64 {
	return float64(s.lat * s.lat)
}

func (c circle) area() float64 {
	return pi * float64(c.rad*c.rad)
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	sq := square{20}
	fmt.Println(sq.area())

	cir := circle{10}
	fmt.Println(cir.area())
	fmt.Println("Printing from info")
	info(sq)
	info(cir)
}
