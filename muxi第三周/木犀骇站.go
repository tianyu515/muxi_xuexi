package main
import(
    //"encoding/base64"
    "fmt"
	"github.com/weiji6/hacker-support/httptool"
	//"bytes"
	//"time"
	//"net/http"
	//"io"
	//"encoding/json"
)
	//"net/http"
	

//发送请求，在响应头中得到密钥
// func main() {
// 	client := http.Client{
// 		Timeout : 30 * time.Second,
// 	}
// 	req,err := http.NewRequest(
// 		http.MethodGet,
// 		"https://gtainmuxi.muxixyz.com/api/v1/organization/code",
// 		nil,
// 	)
// 	if err != nil {
// 		fmt.Println("创建错误",err)
// 		return
// 	}
// 	req.Header.Set("Code","0")
// 	resp,err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("发送失败",err)
// 		return
// 	}
// 	fmt.Println(resp.Header)

// }


//解密
// func main() {
// 	data := "eyJjb2RlIjoiMCIsImlhdCI6MTc5NTIyNDczNiwibmJmIjoxNzk1MjI0NzM2fQ"
// 	decoded, err := base64.RawURLEncoding.DecodeString(data)
// 	code := string(decoded)
//     if err != nil {
// 		fmt.Println("解码失败",err)
// 		return
// 	}
// 		client := http.Client{
// 		Timeout : 30 * time.Second,
// 	}
	//passport := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiMCIsImlhdCI6MTc5NTIyNDczNiwibmJmIjoxNzk1MjI0NzM2fQ.cNsoB3TsvXKkCieLNgDzD9KheBCDxecyA3y87nAh3uE"
// 	req,err := http.NewRequest(
// 		http.MethodGet,
// 		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/secret_key",
// 		nil,
// 	)
// 	if err != nil {
// 		fmt.Println("创建错误",err)
// 		return
// 	}
// 	req.Header.Set("Code","0")
//     req.Header.Set("passport",passport)
// 	resp,err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("发送失败",err)
// 		return
// 	}
// 	body,err := io.ReadAll(resp.Body)
// 	fmt.Println(string(body))
// }

//二次解密
// func main() {
// 	 codenum := "c2VjcmV0X2tleTpNdXhpU3R1ZGlvMjAzMzA0LCBlcnJvcl9jb2RlOmZvciB7Z28gZnVuYygpe3RpbWUuU2xlZXAoMSp0aW1lLkhvdXIpfSgpfQ=="
// 	 errorCode, err := encrypt.Base64Decode(codenum)
// 	 if err != nil {
// 		fmt.Println("解密失败",err)
// 		return
// 	 }
// 	 fmt.Println(errorCode)


//{"text":"恭喜你拿到了 passport，现在你可以着手准备骇入银行。\r\n银行的第一道门是代码安全门，我们计划将错误代码写入此门来破解它。
//  type Code struct {
// 	Data string `json:"content"`
//  }
// func main() {
// 	passport := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiMCIsImlhdCI6MTc5NTIyNDczNiwibmJmIjoxNzk1MjI0NzM2fQ.cNsoB3TsvXKkCieLNgDzD9KheBCDxecyA3y87nAh3uE"
// 	errorCode := "for {go func(){time.Sleep(1*time.Hour)}()}"
// 	key := "MuxiStudio203304"
// 	data, err := encrypt.AESEncryptOutInBase64([]byte(errorCode), []byte(key))
// 	if err != nil {
// 				fmt.Println("解密失败",err)
// 				return
// 			}
// 	text := string(data)
// 	code := Code{
// 		Data : text,
// 	}
// 	code1,err := json.Marshal(code)
//     if err != nil {
// 				fmt.Println("转换失败",err)
// 				return
// 			}
//  	client := http.Client{
// 		Timeout : 30 * time.Second,
//     }
// 	req,err := http.NewRequest(
// 				http.MethodGet,
// 				"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate",
// 				bytes.NewReader(code1),
// 			)
// 			if err != nil {
// 				fmt.Println("创建失败",err)
// 				return
// 			}
// 			req.Header.Set("Code","0")
//             req.Header.Set("passport",passport)
// 			req.Header.Set("Content-Type", "application/json")
// 			resp,err := client.Do(req)
// 			if err != nil {
// 				fmt.Println("发送失败",err)
// 				return
// 			}
// 			defer resp.Body.Close()
// 			body,err := io.ReadAll(resp.Body)
// 			if err != nil {
// 				fmt.Println("读取失败",err)
// 				return
// 			}
// 			fmt.Println(string(body))
// }
//{"text":"干的漂亮！你已经突破了第一扇门，请继续访问 http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate 。","extra_info":""}}
//{"text":"你现在已经到第二扇门了，是虹膜识别安全门。\r\n你需要向组织请求已准备好的虹膜样本，
// func main() {
// 	passport := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiMCIsImlhdCI6MTc5NTIyNDczNiwibmJmIjoxNzk1MjI0NzM2fQ.cNsoB3TsvXKkCieLNgDzD9KheBCDxecyA3y87nAh3uE"
// 	request, err := httptool.NewRequest(
// 		httptool.GETMETHOD,
// 		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/iris_sample",
// 		"",
// 		httptool.DEFAULT)
// 		if err != nil {
// 			fmt.Println("创建失败",err)
// 			return
// 		}
//     request.AddHeader("code","0")
//     request.AddHeader("passport",passport)
// 	resp, err := request.SendRequest()
// 	if err != nil {
// 			fmt.Println("发送失败",err)
// 			return
// 		}
//     err = resp.Save(`D:\01_学习资料\muxi\muxi第三周\iris.jpg`)
// 	if err != nil {
// 			fmt.Println("下载失败",err)
// 			return
// 		}
// 		resp.ShowBody()
// }


//已经得到虹膜照片
// func main() {
// 	passport := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiMCIsImlhdCI6MTc5NTIyNDczNiwibmJmIjoxNzk1MjI0NzM2fQ.cNsoB3TsvXKkCieLNgDzD9KheBCDxecyA3y87nAh3uE"
// 	request, err := httptool.NewRequest(
// 		httptool.POSTMETHOD,
// 		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination",
// 		`D:\01_学习资料\muxi\muxi第三周\iris.jpg`,
// 		httptool.FILE)
// 		if err != nil {
// 			fmt.Println("创建失败",err)
// 			return
// 		}
// 	    request.AddHeader("code","0")
//         request.AddHeader("passport",passport)
// 		resp, err := request.SendRequest()
// 		if err != nil {
// 			fmt.Println("发送失败",err)
// 			return
// 		}
// 		resp.ShowBody()
// 		resp.ShowHeader()
// 	}
	//还剩最后一道门了！我们需要银行结构图碎片，这些碎片就隐藏在前面某四个路由的响应头中，位于 map-fragments 字段。它们用"/"拼起来就是最后一道门的所在位置！注意response的信息
			

	//http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination


	//全排列
	func permute(nums []int) [][]int {
		
	}
	func main() {
		testSlice := make([]int, n)
		fmt.Scanf("%d", &testSlice)
		res := permute(testSlice)
		fmt.Println(res)
	}
		