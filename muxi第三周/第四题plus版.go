package main
import(
	"fmt"
	"net/http"
	"time"
	"io"
	"sync"
	"bytes"
)
func main() {
	client := http.Client{
		Timeout : 30 * time.Second,
	}
	var wg sync.WaitGroup
	ch := make(chan string,10)
	for i := 1 ; i <= 10;i++ {
		wg.Add(1)
		go func(seq int) {
			data := []byte(fmt.Sprintf(`{"seq": %d}`,seq))
			req,err := http.NewRequest(
				http.MethodGet,
				"http://117.72.42.214:8080/api/challenge",
				bytes.NewReader(data),
			)
			if err != nil {
				fmt.Println("创建失败",err)
				return
			}
			req.Header.Set("Authorization","Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdGFnZSI6InJld2FyZCIsInNpZCI6ImY3NmZiNzFmN2FmN2Q0OTM3ZWVjMGZhOGE5NTU3ZDBlIiwiZXhwIjoxNzk1MTQwOTYzLCJpYXQiOjE3OTUxNDAzNjMsImp0aSI6IjJhMGFiMWEzYmM3MjgyNzU5NjA1NzU4YjEwZTk0MTIxIn0.54E7BeboLf5zU7UFQik0ePlvmdOKIgwo14m1Wibn-LA")
			req.Header.Set("Content-Type","application/json")
			resp,err := client.Do(req)
			if err != nil {
				fmt.Println("发送失败",err)
				return
			}
			defer resp.Body.Close()
			body,err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("读取失败",err)
				return
			}
			ch <- string(body)
			wg.Done()
        }(i)
	}
	wg.Wait()
	result := <- ch
	fmt.Println(result)
}

func main() {
	client := http.Client{
		Timeout : 30 * time.Second,
	}
	
		        req,err := http.NewRequest(
				http.MethodGet,
				"http://117.72.42.214:8080/api/reward",
				nil,
			)
			if err != nil {
				fmt.Println("创建失败",err)
				return
			}
			req.Header.Set("Authorization","Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdGFnZSI6InJld2FyZCIsInNpZCI6ImY3NmZiNzFmN2FmN2Q0OTM3ZWVjMGZhOGE5NTU3ZDBlIiwiZXhwIjoxNzk1MTQwOTYzLCJpYXQiOjE3OTUxNDAzNjMsImp0aSI6IjJhMGFiMWEzYmM3MjgyNzU5NjA1NzU4YjEwZTk0MTIxIn0.54E7BeboLf5zU7UFQik0ePlvmdOKIgwo14m1Wibn-LA")
			resp,err := client.Do(req)
			if err != nil {
				fmt.Println("发送失败",err)
				return
			}
			defer resp.Body.Close()
			body,err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("读取失败",err)
				return
			}
			fmt.Println(string(body))
	
	}