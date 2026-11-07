package main

import"fmt"
func main() {
	
	//scores := [5]int{11,13,15,26,27}
	//for i := 0; i < len(scores); i++ {
	//	fmt.Println(scores[i])
	//}
	//for index,value  := range scores {
	//	fmt.Println(index,value)
	//}
	a := "Hello,世界"
    fmt.Println(len(a))
	for b := range a{
		fmt.Println(b)
	}



}