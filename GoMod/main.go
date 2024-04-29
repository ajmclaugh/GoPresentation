package main

import (
	"fmt"

	"github.com/ajmclaugh/GoMod/loopy"
)

func main() {
	helloprint := loopy.Loopit("Hello World", 5)
	fmt.Println(helloprint)
	// bob := "here"
	// switch bob {
	// case "here":
	// 	fmt.Println("Bob's here")
	// case "not here":
	// 	fmt.Println("Bob's not here")
	// }
}
