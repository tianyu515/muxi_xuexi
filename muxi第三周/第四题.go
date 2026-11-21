package main
import(
	"fmt"
	"net/http"
	"time"
	"io"
)

func main() {
	//var wg sync.WaitGroup
	//wg.Add(10)
    
 	client := http.Client{
		Timeout : 30 * time.Second,
    }
	for i := 2 ; i <= 10; i++ {
	data1 := []byte(fmt.Sprintf(`{"seq" : %d}`,i))
	req1,err := http.NewRequest(
		http.MethodPost,
		"http://117.72.42.214:8080/api/challenge",
        bytes.NewReader(data1),
	)
	if err != nil {
		fmt.Println("创建失败",err)
		return
	}
	req1.Header.Set("Authorization","Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdGFnZSI6ImNoYWxsZW5nZSIsImV4cCI6MTc5NTA5NDkwNiwiaWF0IjoxNzk1MDkzMTA2LCJqdGkiOiJhMGY0ZjA5MWU0Nzg3YWY1NDUyMGZiZjg3MmViZjRiYSJ9.U4R8UPI50IYfgmwwBMD0bz0rbOxF_-nooziTuzqSTks")
	req1.Header.Set("Content-Type","application/json")
	resp1,err := client.Do(req1)
	if err != nil {
		fmt.Println("发送失败",err)
		return
	}
	fmt.Println(resp1.StatusCode)
}
// 	data := []byte(`{"seq" : 1}`)
	req,err := http.NewRequest(
		http.MethodGet,
		"http://117.72.42.214:8080/api/reward",
        nil,
	)
	if err != nil {
		fmt.Println("创建失败",err)
		return
	}
	req.Header.Set("Authorization","Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdGFnZSI6InJld2FyZCIsInNpZCI6ImEwZjRmMDkxZTQ3ODdhZjU0NTIwZmJmODcyZWJmNGJhIiwiZXhwIjoxNzk1MDk0MjEyLCJpYXQiOjE3OTUwOTM2MTIsImp0aSI6IjI4ZDQ5YzdkNTgwNWU1YWJlYzY3ZGI5N2I0Y2FkMDZiIn0.6nfhbjIHBVPXoPy1JqpFE8Az7OBn1Y5GehbIIr6uB4s")
	resp,err := client.Do(req)
	if err != nil {
		fmt.Println("发送失败",err)
		return
	}
	
	body,err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取失败",err)
		return
	}
	fmt.Println(string(body))

}