// package main

// import (
//     "fmt"
// )

// func sayHello() {
//     fmt.Println("Hello")
// }

// func main() {
//      sayHello()

//     fmt.Println("main")
// }




// package main
// import(
// 	"fmt"
// 	"sync"
// )
// func main(){
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("任务a已完成")
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("任务b已完成")
// 	}()

// 	wg.Wait()
// 	fmt.Println("任务已全部完成")

// }




// package main
// import"fmt"

// func main(){
// 	ch := make(chan int,3)
// 	ch <- 100
// 	ch <- 200
// 	close(ch)
// 	for num := range ch{
// 		fmt.Println(num)
// 	}
// }


// package main
// import(
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	count := 0
// 	var wg sync.WaitGroup
// 	var mu sync.Mutex
// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		for i:=0;i<1000;i++{
// 			mu.Lock()
// 			count++
// 			mu.Unlock()
// 		}
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		for i:=0;i<1000;i++{
// 			mu.Lock()
// 			count++
// 			mu.Unlock()
// 		}
// 	}()
// 	wg.Wait()
// 	fmt.Println("任务已完成")
// 	fmt.Println(count)
// }


// package main
// import"fmt"

// func main() {
// 	// scores := make(map[string]int)
// 	// scores["小明"] = 90
// 	// score,ok := scores["小明"]
// 	// fmt.Println(score,ok)
// 	scores := map[string]int {
// 		"小明" : 90,
// 		"小王" : 85,
// 	}
// 	for _,score := range scores{
// 		fmt.Println(score)
// 	}
// }

package main
import"fmt"
func main() {
	type Student struct{
		Name string
		Score int
		Age int
	}
	scores := make(map[string]Student)
	scores["2026215615"] = Student{
		Name : "小王",
		Score : 90,
		Age : 18,
	}
	score,ok := scores["2026215615"]
	fmt.Println(score.Score,ok)
	
}