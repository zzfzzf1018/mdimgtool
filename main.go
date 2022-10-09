package main

import (
	"fmt"
	"mdimgtool/command"
	"os"
)

func main() {
	params := os.Args[1:]
	fmt.Println(params)
	// fmt.Println("test")
	// result := command.Command("ls")
	// fmt.Println(result)

	// result1 := command.Command("df", "-h")
	// fmt.Println(result1)

	// result2 := command.Command("df -h")
	// fmt.Println(result2)
	cl := command.New()
	cl.Parse(os.Args)
}
