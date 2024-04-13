package main

import "fmt"

func Hello(name string) string { //Domain code
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("Chris")) //fmt.print is side effect
	//so better to seperate domain from side effect
}
