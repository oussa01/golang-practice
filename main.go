package main

import ("fmt"
	"bufio"
	"os"
	"runtime/debug")
func main() {
	fmt.Println("enter a user name")
	reader := bufio.NewReader(os.Stdin)
	nom, _ := reader.ReadString('\n')
	
	fmt.Println("enterer prenom")
	prenom, _ := reader.ReadString('\n')
	fullname:= nom+prenom
	namefield(nil,nil)
	fmt.Print("your name is ",fullname)
	
}	
func namefield(nom *string, prenom *string ){
	defer recoverFullName()
	if nom == nil{
		panic("runtime error: first name cannot be null")
	}
	if prenom== nil{
		panic("runtime error: prenom cannot be null")
	}
}
func recoverFullName(){
	if r:= recover();r!=nil{
		fmt.Println("recovered from",r)
		debug.PrintStack()
	}
}