## defer
- defer用于将某个行为延迟到函数结束后执行
- 用于资源清理
- 遵循后出的原理，例如
```go
defer fmt.Println(a)
defer fmt.Println(b)
```
最后打印的时候，是先打印b
## 闭包 
- 匿名函数可以被赋值给变量，把变量当成普通函数那样来调用
- Go 中的函数字面量（匿名函数）是闭包，它可以引用外围函数中的变量，并且与外围函数共享这些变量。
- 闭包引用的外部变量不是简单复制一份，而是可以与外部函数共享同一个变量，当闭包在函数内部修改了变量，外部的变量也会进行修改
- 相当于外部函数定义的局部变量，形成闭包的这个匿名函数也拥有了这个变量
## 串行、并发、并行
### 串行：
一个任务执行完再执行下一个任务
### 并发：
多个任务在一段时间内交替运行
**注意：并发并不意味着多个任务一定在同一时间进行**
### 并行：
多个任务一定在同一时间进行
**并行是并发任务的一种特殊情况**

## Go 并发编程的三个核心工具
### Goroutine:
Go 语言提供的一种轻量级并发执行单元，可以让多个任务同时推进
```go
package main
import"fmt"
func sayHello() {
    fmt.Println("Hello")
}

func main() {
    go sayHello()

    fmt.Println("main")
}
```
由于go,Go 就会启动一个新的 Goroutine 来执行这个函数。此时 `main()` 不需要等待 `sayHello()` 执行结束，就可以继续向下运行。但这会导致一个问题，最终只会输出main，我们希望主函数要等待其他goroutine执行完毕
### sync:
Go 标准库中的一个包，提供并发协调工具![[Pasted image 20261113160706.png]]
```go
package main

import(

    "fmt"
    "sync"
)

func main(){
    var wg sync.WaitGroup

    wg.Add(2)

    go func() {

        defer wg.Done()

        fmt.Println("任务a已完成")

    }()

  

    go func() {

        defer wg.Done()

        fmt.Println("任务b已完成")

    }()

  

    wg.Wait()

    fmt.Println("任务已全部完成")

  

}
```
**先不用过多思考这个类型的变量为什么可以这样做，记住sync可以解决主程序已结束而Gotoutine还未解决的问题**
wg.Add()是给WaitGroup加任务，wg.Done()是完成任务后减一，wg.Wait()当任务数=0时程序才接着往下走
### Channel:
#### 没有缓冲区的Channel
一个真正的 Goroutine + Channel 示例
```go
package main

import "fmt"

func main() {

    ch := make(chan int)//创建通道的方式，chan是go的关键字

    go func() {
        ch <- 100//向ch传递数据
    }()

    num := <-ch//把数据从通道中拿出

    fmt.Println(num)
}
```
- 没有缓冲区的Channel会造成堵塞，如果只有一个Goroutine传递信息，而没有另一个Goroutine接受信息，就会导致程序等待一直堵塞
- main函数往下走本身就是创建了一个主Goroutine
#### 有缓冲区的**Channel：
```go
ch := make(chan int,3)
```
创建了一个可以暂存3个数据的通道，存的数据超过了暂存数才会堵塞
#### Channel的关闭：
```go
close(ch)
```
通过close函数来关闭通道
**注意点：1.Channel 关闭后，里面已经存在的数据仍然可以取出来(类似闭包)
          2.已经关闭的 Channel 不能再发送数据，否则会发生 panic（运行时异常）**
### sync.Mutex:互斥锁
用途：当多个routine想同时去改同一个变量，或做不可分割的一个任务时，需要避免冲突
```go
package main

import(

    "fmt"

    "sync"

)

  

func main() {

    count := 0

    var wg sync.WaitGroup

    var mu sync.Mutex

    wg.Add(2)

    go func() {

        defer wg.Done()

        for i:=0;i<1000;i++{

            mu.Lock()//上锁

            count++

            mu.Unlock()//解锁

        }

    }()

    go func() {

        defer wg.Done()

        for i:=0;i<1000;i++{

            mu.Lock()

            count++

            mu.Unlock()

        }

    }()

    wg.Wait()

    fmt.Println("任务已完成")

    fmt.Println(count)

}
```