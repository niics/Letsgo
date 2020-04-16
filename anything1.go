package main

import "fmt"

func main() {
	fmt.Println("Hello World, it's awksome.")
	foo()
	for i := 0; i < 100; i++ {
		if i % 8 == 0{
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm in foo.")
}
