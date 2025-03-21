package main

import "fmt"

func main() {
	for i := range 3 {
		for j := range 5 {
			fmt.Println(j)
			if j == 2 {
				break
			}
		}
		fmt.Println(i)
	}
}
