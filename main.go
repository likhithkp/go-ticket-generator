package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, i *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := i.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createTicket() ticket {
	fmt.Println("Welcome to Nomad")

	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a ticket", reader)

	t := newTicket(name)
	fmt.Println("Created a new ticket for: ", t.name)
	return t
}

func promptOptions(tkt ticket) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("b (book a ticket), c (check destination and price), p (print ticket)", reader)

	switch opt {
	case "b":
		personName, _ := getInput("Enter the passenger name: ", reader)
		seatNo, _ := getInput("Enter the seat number: ", reader)
		place, _ := getInput("Enter the destination: ", reader)
		price, _ := getInput("Enter the price: ", reader)

		sNo, err := strconv.Atoi(seatNo)
		if err != nil {
			fmt.Println("Please enter a valid seat number")
			promptOptions(tkt)
		}

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Please enter a valid price")
			promptOptions(tkt)
		}

		tkt.addPassenger(personName, sNo, place, p)

		fmt.Printf("Ticket: %v\n", tkt)
		promptOptions(tkt)
	case "c":
		fmt.Printf("Check destination and price")
	case "p":
		tkt.save()
	default:
		fmt.Println("Not a valid option...")
		promptOptions(tkt)
	}
}

func main() {
	ticket := createTicket()
	promptOptions(ticket)
}
