package main

import (
	"fmt"

	"first_prog/exercices"
	"runtime/debug"

)

func main() {
	fmt.Println("enter a number")
	var number int 
	fmt.Scanf("%d",&number)
	exercices.Fibo(number)
	

}
func namefield(nom *string, prenom *string) {
	defer recoverFullName()
	if nom == nil {
		panic("runtime error: first name cannot be null")
	}
	if prenom == nil {
		panic("runtime error: prenom cannot be null")
	}
}
func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
		debug.PrintStack()
	}
}
