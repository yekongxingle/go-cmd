package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("what is your name?")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println("hello,", text,"have fun")
}
