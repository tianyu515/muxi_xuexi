## 定义：
结构体方法类似于函数，在外部去声明，也是需要在main函数里调用，在方法里通过语句可以实现一些功能
例如：以下是一个函数
```go
func Average() {
}
```
以下是一个结构体方法：
```go
func (s Student) Average() float64{
}
```
**这就是一个完整方法的框架**:(s Student)是一个接受者，s是接受者变量名，可任意取，Student是结构体类型，表示这个方法属于 Student 类型，调用时由一个 Student 对象来担任接收者。方法的名称就叫做Average(**s是一个笼统的称呼，在实际的方法中指代的就是具体的变量，比如student**)
```go
type Student struct{

Name string
Score int
Age int
}

student := Student{
Name : "小王",
Score : 90,
Age : 18,
}
```
## 使用：
```go
student.Average()
```
将student这个已经赋值的结构体变量与方法绑定，方法里面的s的值就相当于是student的值
## 值接受者与指针接受者
```go
func (s Student) Average() float64{
   s.Name = 99//由于是结构体，不会改变student的值
}
```
这就是一个值接受者，接受者是结构体,方法拿到的是结构体，用于复制数据
```go
func (s *Student) Average() float64{
}
```
这就是一个指针接受者，方法拿到的是指向结构体的地址，用于修改数据
对于指针接受者调用方法，也是直接用`student.Average()`，go的语言特性
----

任务三：![[Pasted image 20261114103742.png]]
1. Average方法
```go
package main

import"fmt"
type Student struct {

    Name   string

    Scores []int // 记录若干项成绩

}

func (s Student) Average() float64{

    if len(s.Scores) == 0 {

        return 0

    }

    sum := 0

    for i := 0;i < len(s.Scores);i++ {

        sum += s.Scores[i]

    }

    return  float64(sum) / float64(len(s.Scores))

}

```
**我错的地方**：为了最后得到一个float64值，就不能直接用 sun / len了，要将两个值，都进行强制类型转换为float64,方法`float64(int)`
2. Addscore方法：
```go
func (s Student) Addscore() []int {

    scores := append(s.Scores,88)

    return scores

}
```

3. Top方法：
- 冒泡排序：
在一轮中，通过相邻两个数的比较，把最大的数逐渐移到最右边，再进行n轮，把数按照从小到大的顺序依次排开，即需要两个for循环
```go
for i := 0;i < len(a);i++ {

        for j :=0;j < len(a)-1-i;j++ {

            if a[j] < a[j+1] {

                temp := a[j+1]

                a[j+1] = a[j]//把小的值赋给右边

                a[j] = temp//把大的值移到左边

            }

        }

    }
```

整体：
```go
func (s Student) Top() []int {

    if len(s.Scores) <= 0 {

        return nil

    }

    a := make([]int,len(s.Scores))

    copy(a,s.Scores)

    for i := 0;i < len(a);i++ {

        for j :=0;j < len(a)-1-i;j++ {

            if a[j] < a[j+1] {

                temp := a[j+1]

                a[j+1] = a[j]

                a[j] = temp

            }

        }

    }

    var n int

    fmt.Scanf("%d",&n)

    if n >= len(a) {

        for i := 0;i < len(a);i++ {

            fmt.Println(a[i])

        }

    } else {

        for i := 0;i < n;i++ {

            fmt.Println(a[i])

        }

    }

    return a
    }
```
- Mergr方法：
~~我的错误写法：~~
```go
func (s Student) Merge() Student {

    student1 := Student{

        Name : "小王",

        Scores : []int{85,90,94},

    }

    student2 := Student{

        Name : "小张",

        Scores : []int{88,100,75,99},

    }

    b := len(student1.Scores)

    c := len(student2.Scores)

    a := make([]int,b+c)

    for i := 0;i < b ;i++ {

        a[i] = student1.Scores[i]

    }

    for j :=0 ;j <c;j++ {

        a[b+j] = student2.Scores[j]

  

    }

    s.Name = "s"

    for i :=0; i < b+c;i++ {

        s.Scores[i] = a[i]

    }

    return s
}
func main() {
var student3 Student

    fmt.Println(student3.Merge())
}
```
**错误点：**
1. 对题意理解错误：s只是实际接受者变量名的一个在方法里的替代，不是说s是Name
2. **在main中只是去声明了一个Student,其中的切片是nil,长度为0，在方法中应该去声明一个长度**
学习到的：
- 对于一个nil切片，可以直接用 b = make([]int,6)来创建长度
- 结构体中如果含有切片，直接用  b = a,把a的值赋给b

正确写法
```go
func (s Student) Merge(other Student) Student {

  

    // 计算两个切片的长度

    b := len(s.Scores)

    c := len(other.Scores)

  

    // 创建一个全新的切片

    a := make([]int, b+c)

  

    // 把第一个人的成绩放进去

    for i := 0; i < b; i++ {

        a[i] = s.Scores[i]

    }

  

    // 把第二个人的成绩接在后面

    for j := 0; j < c; j++ {

        a[b+j] = other.Scores[j]

    }

  

    // 返回一个新的 Student

    result := Student{

        Name:   s.Name,

        Scores: a,

    }

  

    return result

}

  

func main() {

  

    student1 := Student{

        Name:   "小王",

        Scores: []int{85, 90, 94},

    }

  

    student2 := Student{

        Name:   "小张",

        Scores: []int{88, 100, 75, 99},

    }

  

    merged := student1.Merge(student2)

  

    fmt.Println(merged)

}
```