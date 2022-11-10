package main

import (
	"fmt"
)

var nbTx int
var time string
var nbTime int

func main() {

	fmt.Print("Enter the number of Tx: ")
	fmt.Scanln(&nbTx)
	fmt.Print("Enter secound, minute, hour : ")
	fmt.Scanln(&time)
	fmt.Print("Enter number of ", time, " : ")
	fmt.Scanln(&nbTime)

	fmt.Print("The volume is :", nbTx/nbTime, " per ", time)

}