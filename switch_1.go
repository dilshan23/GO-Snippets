package main

import "fmt"

func main() {
	city := "Colombo"
        // why replace if else if
	switch city {
	case "Colombo":
		fmt.Println("going to C")
	case "Kandy":
		fmt.Println("going to K")
	default:
		fmt.Println("ggg")

	}
}
