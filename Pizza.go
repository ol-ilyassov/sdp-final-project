package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	var age int
	fmt.Print("Введите имя: ")
	fmt.Fscan(os.Stdin, &name)

	fmt.Print("Введите возраст: ")
	fmt.Fscan(os.Stdin, &age)

	fmt.Println(name, age)

	var flag int
	for {

		fmt.Fscan(os.Stdin, &flag)
		if flag == 0 {
			break
		}
	}
}
