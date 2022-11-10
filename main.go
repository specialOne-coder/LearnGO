package main

import (
	"fmt"
)

type Ville struct {
	nbHabitants int
	codepostal  int
	superficie  int
}

var villes map[int]Ville

var nbInput int
var cpInput int
var spInput int

func main() {
	fmt.Scanf("%d", &nbInput)
	fmt.Scanf("%d", &cpInput)
	fmt.Scanf("%d", &spInput)
	villes = make(map[int]Ville)
	villes[0] = Ville{nbInput, cpInput, spInput}
	fmt.Println(villes[0])
}
