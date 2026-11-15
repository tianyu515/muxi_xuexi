## map:
- map创建的两种方式
 ```go
scores := make(map[string]int)//使用make函数
```
```go
scores := map[string]int {
  "小明" ：90，
  "小红" :85,
}
```
----
- 想要知道是否有这样一个key,可以用以下方法
```go
  score,ok := scores["小王"]
  fmt.Println(score,ok)//看返回的bool是true or false
```
- nil map与空map
 ```go
var scores mao[string]int
  ```
这只是创建了一个变量，而没有进行初始化，**可以读取，但不能写入（往map加数据）**
 ```go
scores := map[string]int{}//更多是去创建空map
  ``` 
## struct
- 创建struct
  ```go
  type Student struct{//命名一个新的数据类型Student,它是一个结构体型

  Name string
Score int
Age int
}

  s := Student{
  Name : "小王",
 Score : 90,
Age : 18,
}
  ```
  **struct和数组都是单独的，给另一个变量赋值时只是把值复制一份
  slice和map给另一个变量赋值时，它们是共用一个底层，修改一个，另一个一起修改**
## map与struct的结合
```go
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
```
**结构体可以通过.操作符直接访问字段**
~~Go 不允许直接修改 map 中存储的结构体的字段~~：
通过 map 的键取出结构体时，得到的是结构体的一份值，而不是一个可以直接修改字段的变量。
那该怎么修改呢：
```go
student := students["001"]

student.Age = 19

students["001"] = student
```