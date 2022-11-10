package main

import (
	"fmt"
	"sort"
	"strings"
)

var size uint8
var cp uint8

func main() {

	var slist []string
	fmt.Println("Enter size")
	fmt.Scanln(&size)

	for cp = 0; cp < size; cp++ {
		var input string
		fmt.Println("Enter word", cp+1)
		fmt.Scanln(&input)
		slist = append(slist, strings.ToLower(input))
	}

	sort.Strings(slist)

	fmt.Println(slist)

}
