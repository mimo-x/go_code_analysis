package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("this Odd number ！")
		fmt.Printf("this val is %v\n", i)
	}

}
