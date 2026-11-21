package main
import(
	"fmt"
	"net/http"
	"io"
	"time"
	"encoding/json"
	"bytes"
)
// func main() {
// // 	{
// // 		"name" : "小猪",
// // 		"age"  : 18
// // 		"result" : [
// // 			"小王",
// // 			"小李",
// // 			"小张"
// // 		]
// // 	}
// // }
// {
// 	"user_name" : "小天",
// 	"user_num"  : 18
// }
// type Student struct {
// 	Name string `json:"user_name"`
// 	Num int     `json:"user_num"`
// }

// type User struct {
// 		Name string  `json:"name"`
// 		Score int    `json:"score"`
// 	}

// func main() {
// 	user := User{
// 		Name : "小猪", 
// 		Score : 18,   
// 	}
// 	data,err := json.Marshal(user)
// 	if err != nil {
// 		fmt.Println("转换错误",err)
// 		return
// 	}
// 	fmt.Println(string(data))
// }
	
// func main() {
// 	var user User
// 	data := []byte(`{"name" : "小猪","score" : 18}`)
// 	err := json.Unmarshal(data,&user)
// 	if err != nil {
// 		fmt.Println("解析错位",err)
// 		return
// 	}
// 	fmt.Println(user.Name)
// }
// type Post struct {
//     UserID int    `json:"userId"`
//     ID     int    `json:"id"`
//     Title  string `json:"title"`
//     Body   string `json:"body"`
// }
// func main() {
// 	var post Post
// 	//创建客户端
// 	client := http.Client{
// 		Timeout : 10 * time.Second,
// 	}
//     //发送get请求
// 	resp,err := client.Get(
// 		 "https://jsonplaceholder.typicode.com/posts/1",
// 	)
// 	if err != nil {
// 		fmt.Println("读取失败",err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	//查询状态码
	
// 	if resp.StatusCode <200 || resp.StatusCode >= 300 {
// 		fmt.Println("请求失败")
// 		fmt.Println(resp.StatusCode)
// 	}
// 	//读取请求体
// 	body,err := io.ReadAll(resp.Body)
//     if err != nil {
//     fmt.Println("读取错误",err)
//     return
//   }
//     err = json.Unmarshal(body,&post)
// 	fmt.Println(post.Title)
// 	fmt.Println(post.ID)
 
// } 

// type Post struct {
//     UserID int    `json:"userId"`
//     ID     int    `json:"id"`
//     Title  string `json:"title"`
//     Body   string `json:"body"`
// }
// func main() {
// 	//新建查询参数
// 	params := url.Values{}
// 	params.Set("id","1")
// 	params.Set("limit","5")
// 	params.Encode()
// 	url := "https://example.com/posts?" + params.Encode()
// 	//新建客户端
// 	client := http.Client{
// 		Timeout : 10 * time.Second,
// 	}
// 	resp,err := client.Get(url)
// 	if err != nil {
// 		fmt.Println("请求失败",err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	fmt.Println(resp.Status)
// 	body,err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("读取失败",err)
// 		return
// 	}
// 	var user User
// 	err = json.Unmarshal(body,&user)
// 	if err != nil {
// 		fmt.Println("解析失败",err)
// 		return
// 	}
// 	fmt.Println(user.UserID)

// }

// func main() {
// 	//创建客户端
// 	client := http.Client{
// 		Timeout : 10 * time.Second,
// 	}
// 	post := Post{
// 		Title:  "Go 网络请求学习",
//         Body:   "今天学习了 POST 请求",
//         UserID: 1,
// 	}
// 	data,err := json.Marshal(post)
// 	if err != nil {
// 	fmt.Println("转换失败",err)
// 	return
// 	}
// 	//创建post请求对象
// 	req,err := http.NewRequest(
// 		http.MethodPost,
// 		"https://jsonplaceholder.typicode.com/posts",
// 		bytes.NewReader(data),
//    )
//    req.Header.Set("Content-Type","application/json")
//    req.Header.Set("Accept","application/json")
//    resp,err := client.Do(req)
//    if err != nil {
// 	fmt.Println("请求失败",err)
// 	return
// 	}
// 	defer resp.Body.Close()
// 	body,err := io.ReadAll(resp.Body)
// 	if err != nil {
// 	fmt.Println("读取失败",err)
// 	return
// 	}
// 	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
// 		fmt.Println("状态出错",resp.Status)
// 		fmt.Println(string(body))
// 	}
// 	var result Post
// 	err = json.Unmarshal(body,&result)
// 	if err != nil {
// 	fmt.Println("解析失败",err)
// 	return
// 	}
// 	fmt.Println(result.Body)
// }


// func main() {
// 	data := url.Values{}
// 	data.Set("username","小猪")
// 	data.Set("id","1")
// 	resp,err := http.PostForm(
// 		"https://jsonplaceholder.typicode.com/posts",
// 		data,
// 	)
// 	ContentType := resp.Header.Get("Content-Type")
// 	fmt.Println(ContentType)
// }
http.NewRequest(

	bytes.NewReader(data)

)
req.Header.Set()
err := json.NewDecoder(resp.Body).Decode(&result)