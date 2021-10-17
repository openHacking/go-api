package main

import (
	"log"

	"github.com/openHacking/search"
)

func main() {
	log.SetPrefix("search:")
	log.SetFlags(0)

	// names := []string{"Alex", "Elaine", "Sum"}
	// message, err := search.Searchs(names)

	search.SearchByText()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(message)
}
