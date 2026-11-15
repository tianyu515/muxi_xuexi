package main
import(
	"fmt"
	"math/rand"
	"sort"
	"time"
	"sync"
)

type Result struct {
	ID int
	Number int
}

func main() {
	var results [20]Result
	ch := make(chan Result)
	var wg sync.WaitGroup
	for i := 1;i <= 20;i++ {
		wg.Add(1)
		
		go func(id int) {
			defer wg.Done()
			count := rand.Intn(1001)
			time.Sleep(time.Duration(count)*time.Millisecond)
			num := rand.Intn(101)
			ch <- Result{
				ID : id,
				Number : num,
			}
        }(i)

	}
	for i := 0;i < 20;i++ {
		results[i] = <- ch 
	}
	wg.Wait()
	fmt.Println("排序前的输出:")
	for i := 0;i < 20;i++ {
		fmt.Printf("编号:%2d,数据:%d",results[i].ID,results[i].Number)
	}
    sort.Slice(results[:], func(i, j int) bool {
		return results[i].ID < results[j].ID
	})

	
	fmt.Println("\n排序后(按照编号顺序):")
	for i := 0; i < 20; i++ {
		fmt.Printf("编号：%2d,随机数:%d\n",
			results[i].ID, results[i].Number)
	}
}
	

