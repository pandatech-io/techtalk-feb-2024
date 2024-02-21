package panichandler

import "fmt"

func RunDefer() {
	defer func() {
		fmt.Println("fifth")
	}()

	fmt.Println("first")
	fmt.Println("second")

	defer func() {
		fmt.Println("fourth")
	}()

	fmt.Println("third")
}
