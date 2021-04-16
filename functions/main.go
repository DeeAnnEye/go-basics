package main

import "fmt"

func messages(name string) string {

	return "Hello " + name

}

func main(){
	fmt.Println(messages("world"))
}