package main

import (
	"fmt"
	"os"
)

type passenger struct {
	name        string
	seatNo      int
	destination string
	price       float64
}

type ticket struct {
	name       string
	passengers []passenger
}

func newTicket(name string) ticket {
	return ticket{
		name:       name,
		passengers: []passenger{},
	}
}

func (t *ticket) addPassenger(pName string, sNo int, place string, price float64) {
	p := passenger{
		name:        pName,
		seatNo:      sNo,
		destination: place,
		price:       price,
	}
	t.passengers = append(t.passengers, p)
}

func (t *ticket) format() string {
	fs := "Nomad Ticket \n"
	var total float64 = 0

	for _, p := range t.passengers {
		fs += fmt.Sprintf("%-25v ...Seat No: %v ...Destination: %v ...₹%v\n", p.name+":", p.seatNo, p.destination, p.price)
		total += p.price
	}

	fs += fmt.Sprintf("\n%-25v ...₹%v\n", "Total:", total)
	return fs
}

func (t *ticket) save() {
	data := []byte(t.format())
	err := os.WriteFile("tickets/"+t.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Ticket is saved")
}
