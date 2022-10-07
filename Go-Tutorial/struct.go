package main

import "fmt"

type customer struct {
	name    string
	address string
	balance float64
}

func getCustomerInfo(c customer) {
	fmt.Println("The customer owes: ", c.balance)
	fmt.Println("The customer lives on: ", c.address)
}

func newCustAdd(c *customer, address string) {
	c.address = address
}

func main() {

	var tS customer
	tS.name = "Ahmed Kidwai"
	tS.address = "123 fake street"
	tS.balance = 234.56
	fmt.Println(tS.name)
	getCustomerInfo(tS)
	newCustAdd(&tS, "33 not fake street")
	getCustomerInfo(tS)

}
