package main

import "fmt"

func main(){
	
	// arrays
	var fruitArr [2]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// declare and assign array
	names := [2]string{"Mark","Ben"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])
	fmt.Println(names)

	// slices
	name := []string{"Mark","Ben","Tom","Ginger"}

	fmt.Println(name)
	// length of array
	fmt.Println(len(name))
	// range of array 
	fmt.Println(name[1:3])
}