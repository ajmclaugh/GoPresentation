package main

import (
	"fmt"

	"github.com/ajmclaugh/HelloGo/world"
	"github.com/softchris/math"
)

func main() {
	hello := "Hello"
	Hello_World := HelloGo.AddWorld(hello)
	fmt.Println(Hello_World)

	HelloGo.AddWorldPointer(&hello)
	fmt.Println(hello)

	fmt.Println(math.Add(1, 1))
}

// go cache "C:\Users\AnthonyMcLaughlin\go\pkg\mod\cache\download"
