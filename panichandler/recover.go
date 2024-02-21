package panichandler

import "fmt"

func RunRecover() {
	fmt.Println("First function called")
	panicRaiser()
	fmt.Println("First function finished") // This will not get called
}

func panicRaiser() {
	fmt.Println("Second function called")
	panic("Panic happens")                  // Go library for panic
	fmt.Println("Second function finished") // This will not get called
}
