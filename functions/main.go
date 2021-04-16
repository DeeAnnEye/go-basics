package main

import "fmt"

func messages(name string) string {

	return "Hello " + name

}

func findSum(num1 int ,num2 int) int {

	return num1+num2

}

func main(){
	fmt.Println(messages("world"))
	fmt.Println(findSum(23,45))
}