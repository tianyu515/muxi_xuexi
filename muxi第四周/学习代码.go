// package main
// import(
// 	"fmt"
// 	"net/http"
// )

// func hello(w http.ResponseWriter,r *http.Request) {
// 	fmt.Fprintf(w,"你好，小猪")
// 	fmt.Println(r.Method)
// }
// func main() {
// 	http.HandleFunc("/hello",hello)
// 	err := http.ListenAndServe(":8080",nil)
// 	if err != nil {
// 		fmt.Println("创建失败",err)
// 		return
// 	}
// } 

// package main
// import(
// 	"fmt"
// 	"net/http"
// )
// func hello(w http.ResponseWriter,r *http.Request) {
// 	fmt.Fprintf(w,"这是一个web服务")
// 	fmt.Println(r.URL.Path)
// 	fmt.Println(r.Method)
// }
// func main() {
// 	http.HandleFunc("/hello",hello)
// 	err := http.ListenAndServe(":8080",nil)
// 	if err != nil {
// 		fmt.Println("创建失败")
// 	}
// }


package main
import(
	"fmt"
	"encoding/json"
	"net/http"
	"sync"
)
type User struct{
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Passport string `json:"passport"`
}
type Rep struct{
	Username string `json:"username"`
	Passport string `json:"passport"`
} 
type ChangePassport struct{
	Oldpass string `json:"oldpass"`
	Newpass string `json:"newpass"`
}
type UpdateUserRequest struct {
    Nickname string `json:"nickname"`
}
var sessions = make(map[string]string)

var sessionsMu sync.Mutex
var users = make(map[string]User)
var usersMu sync.Mutex
func register(w http.ResponseWriter,r *http.Request) {
	//检查请求方法
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type","application/json,charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "请使用POST",
		})
		 return
	}
	//检查json是否合规
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Header().Set("Content-Type","application/json,charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "JSON解析错误",
		})
		return
	}
	//检查用户名字和密码是否为空
	if user.Username == "" || user.Passport == "" {
		w.Header().Set("Content-Type","application/json,charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "用户名和密码不能为空",
		})
		return
	}
	//上锁
	  usersMu.Lock()
	  _,exists := users[user.Username]
	  if !exists {
		users[user.Username] = user
	  }
	  usersMu.Unlock()
    if exists {
		w.Header().Set("Content-Type","application/json,charset=utf-8")
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "用户名已经存在",
		})
		return
	}

	//返回JSON
	w.Header().Set("Content-Type","application/json,charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "注册成功",
			"username" : user.Username,
			"nickname" : user.Nickname,
		})
}
// func main() {
// 	http.HandleFunc("/register",register)
// 	err := http.ListenAndServe(":8081",nil)
// 	if err != nil {
// 		fmt.Println("创建失败")
// 	}
// }


//登录函数


func loginRequest(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json,charset=utf-8")
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type","application/json,charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "请使用POST",
		})
		 return
    }
	var rep Rep
	err := json.NewDecoder(r.Body).Decode(&rep)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "JSON解析失败",
		})
		return
	}
	usersMu.Lock()
	user,exists := users[rep.Username]
	usersMu.Unlock()
	//验证用户名和密码
	if !exists || rep.Passport != user.Passport   {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "用户名或密码错误",
		})
		return
	}
	//设置session
	sessionId := rep.Username + "_sessiom"
	//保存session
	sessionsMu.Lock()
	sessions[sessionId] = user.Username
	sessionsMu.Unlock()
	//设置cookie
	http.SetCookie(w,&http.Cookie{
		Name : "session_id",
		Path : "/",
		Value : sessionId,
	})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
        "message": "登录成功",
        "username": user.Username,
    })
}


//查询用户信息
func me(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	//从请求中查找Cookie
	cookie,err := r.Cookie("session_id")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "请先登录",
		})
		return
	}

	//查找session
	sessionsMu.Lock()
	username,exists := sessions[cookie.Value]
	sessionsMu.Unlock()
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "登录状态失效",
		})
		return
	}

	//查找用户
	usersMu.Lock()
	user,exists := users[username]
	usersMu.Unlock()
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "用户不存在",
		})
		return
	}

	//发送用户信息
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
        "Nickname": user.Nickname,
        "username": user.Username,
    })

}
func Change(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "请使用POST请求",
		})
    return
	}
	var rep ChangePassport
	err := json.NewDecoder(r.Body).Decode(&rep)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "JSON解析错误",
		})
    return
	}
	//确认登录情况
	cookie,err := r.Cookie("session_id")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "请先登录",
		})
    return
	}
	sessionsMu.Lock()
	username,exists := sessions[cookie.Value]
	sessionsMu.Unlock()
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "登录状态失效",
		})
    return
	}
    //确认用户信息
	usersMu.Lock()
	user,exists := users[username]
	usersMu.Unlock()
		if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "用户不存在",
		})
    return
	}
	//检查旧密码
	if user.Passport != rep.Oldpass {
    w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "密码错误",
		})
    return
	}
	//修改密码
	user.Passport = rep.Newpass
	users[user.Username] = user
	//返回信息
	w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "密码修改成功",
			"passport" : rep.Newpass,
		})
}

func UpdateUser(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "请使用POST请求",
		})
    return
	}
	var rep UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&rep)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "JSON解析错误",
		})
    return
	}
	//确认登录情况
	cookie,err := r.Cookie("session_id")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "请先登录",
		})
    return
	}
	sessionsMu.Lock()
	username,exists := sessions[cookie.Value]
	sessionsMu.Unlock()
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "登录状态失效",
		})
    return
	}
    //确认用户信息
	usersMu.Lock()
	user,exists := users[username]
	usersMu.Unlock()
		if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "用户不存在",
		})
    return
	}
	//检查新昵称
	if rep.Nickname == "" {
    w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "昵称不能为空",
		})
    return
	}
	//修改密码
	user.Nickname = rep.Nickname
	users[user.Username] = user
	//返回信息
	w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "昵称修改成功",
		})
}
func main() {
	http.HandleFunc("/register",register)
	http.HandleFunc("/loginRequest",loginRequest)
	http.HandleFunc("/me",me)
	http.HandleFunc("/change",Change)
	http.HandleFunc("/updateUser",UpdateUser)
	err := http.ListenAndServe(":8082",nil)
	if err != nil {
		fmt.Println("请求失败")
		return
	}
}
	
  



