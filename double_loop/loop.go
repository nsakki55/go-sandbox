package main

import "fmt"

func main() {
	items1 := []int{1, 2, 3, 4, 5}
	items2 := []int{1, 2, 3, 4, 5}
loop:
	for _, i := range items1 {
		for _, l := range items2 {
			if l == 3 {
				continue loop
			}
			fmt.Println(i, l)
		}
	}
}
