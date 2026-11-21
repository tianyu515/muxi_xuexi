package main

import"fmt"
func main() {
	numbers := []int{1,2,3,4}
	result := make([]int,len(numbers))
	copy(result,numbers)
	result[2] = 10086
	fmt.Println(numbers,result)
	

}