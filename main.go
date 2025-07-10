package main

import "fmt"

func main() {
	var mainOption int
	mainSreen() // Using the function from interface.go
	fmt.Print(": ")
	fmt.Scanln(&mainOption)

	if mainOption == 1 {
		AddNotes() // Using the function from interface.go
	} else if mainOption == 2 {
		SavedNotes() // Using the function from interface.go
	} else if mainOption == 3 {
		fmt.Println("Goodbye!")
	} else {
		fmt.Println("Invalid option")
	}
}
