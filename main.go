package main

import "hello/customer"

func main() {
	x := customer.Person{}
	x.SetName("Arm")

	println(x.GetName())

}
