package main

import (
	"fmt"
	"os"

	cats_api "github.com/ajmclaugh/example/cats_api"
)

func main() {
	url := os.Args[1]

	response := cats_api.ApiCall(url)

	response_json := cats_api.JsonUnmarshall(response)

	fmt.Println(response_json.Fact)

}

// "https://catfact.ninja/fact"
