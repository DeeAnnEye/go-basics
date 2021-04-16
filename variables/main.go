package main

import "fmt"

func main() {

	// using var
	// var name = "Thor"
	var age int32 = 30
	const isGood = true
	var size float32 = 3.4

	// Shorthand method
	name := "Thor"
	length := 2.9

	email,username := "thor@gmail.com","THORR"
	
	fmt.Println(name, age, isGood,size,email,username)
	// get variable data type
	fmt.Printf("%T\n", length)
}
