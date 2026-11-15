package main

import"fmt"
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
func main() {
	text := counter()
	fmt.Println(text())
	fmt.Println(text())
	fmt.Println(text())
}