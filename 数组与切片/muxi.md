# go数组与切片
## 数组
**定义**：一种数据结构，能够把多个相同类型的数据放在一起管理
- 数组包含索引和值，索引从0开始；
- go中数组属于值类型，数组赋值会复制整个数组，修改副本不会影响原数组。
### 数组遍历
```go
scores := [5]int{80, 85, 90, 95, 88}

for i := 0; i < len(scores); i++ {
    fmt.Println(scores[i])
}
```
或者
```go
for index, value := range scores {
    fmt.Println(index, value)
}
```
## 切片
为什么需要切片？实际生活中，储存的数据是无限的，是没有长度的
### 1.切片创建
```go
numbers := make([]int, 3)
```
创建一个长度为3的切片，基本形式：
```go
make([]数据类型, 长度)
```
### 2.长度与容量
- `len`：当前切片中实际拥有多少个元素。
- `cap`：当前切片从起始位置开始，最多能够延伸到多大范围而不需要更换底层数组。
 ```go
numbers := make([]int, 3, 10)
fmt.Println(len(numbers))
fmt.Println(cap(numbers))
 ```
**容量不等于长度**
### 3.切片表达式s[i:j]
```go
numbers := []int{10, 20, 30, 40, 50}
result := numbers[1:4]
[20 30 40]
```
读切片表达式采用左闭右开，不包含右边的索引
### 4.切片共享底层数组
代码：
```go
numbers := []int{10, 20, 30, 40, 50}

result := numbers[1:4]

result[0] = 999

fmt.Println(numbers)
fmt.Println(result)
```
结果：
```go
[10 999 30 40 50]
[999 30 40]
```
原因：通过切片表达式得到的新切片，通常与原切片共享同一个底层数组。
result修改了底层数组
### 5. append：追加元素
```go
numbers := []int{1, 2, 3}

numbers = append(numbers, 4)
```
结果·
```go
[1 2 3 4]
```
append函数会返回一个值
### 6.copy复制切片
**切片和数组很不同的一点： 比如 b := a ,对于数组，把a直接复制给b，修改b不会影响a
。切片不是这样，a和b可能共享底层数组，所以该怎么修改b而不影响到a呢**
```go
a := []int{1, 2, 3}

b := make([]int, len(a))

copy(b, a)
b[0] = 100

fmt.Println(a)
fmt.Println(b)
```
结果
```go
[1 2 3]
[100 2 3]
```
### 7,两个拓展的小点
- nil和空切片：
  ```go
  var a []int

  b := []int{}
  ```
  其中a是nil,二者长度都为0
- 多维切片
 ```go
 result := [][]int{
 {15,16}
 {17,18}
 } 
 ```
前一个[]表示这是一个切片，后面的[]int表示切片元素是整数切片