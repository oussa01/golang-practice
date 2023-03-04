package main

import ("fmt"
	"bufio"
	"os")
func main() {
	fmt.Println("enter a user name")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Print("your name is ",name)

}	