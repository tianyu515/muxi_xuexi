Package:每个.go文件都会有package,它是一个包，同一个目录下的go源文件通常是一个包
Module:模块，一个module是一个大的项目，下方包含很多包
## go.mod:项目的身份证和依赖清单
1. 怎么创建这个文件：
打开vs进入到终端：输入`go mod init example.com/hello`
2. 有什么：
```go
module example.com/hello

go 1.xx//使用go的语言版本

require github.com/gin-gonic/gin v1.10.0//记录项目需要哪些第三方模块以及对应的版本。
```
有时候需要的库不是go标准库自带的，而是其他开发者开发的，这个时候就需要在go.mod文件中引入第三方依赖
## go.sum:
go.mod确定需要什么，go.sum确定需要的内容是否正确，是否被修改

## go mod 命令：平时怎么管理依赖？
1. go mod init:
新建项目时初始化并创建go.mod
2. go get:添加一个第三方工具
如果我想使用Gin
```go
go get github.com/gin-gonic/gin
```
3. go mod tidy:刚开始我引入gin，后来删掉代码，go.mod还存在不需要的依赖
作用：
- 缺少的依赖会被补充
- 不再需要的依赖会被清理
- 同时整理 go.sum

## goroot与gopath
![[Pasted image 20261113191155.png]]
## go.work：多个 Go 项目怎么一起开发？
1. 一整个大项目里有多个Module，想让这一个直接使用另一个修改后的代码，就可以用到Woekplace![[Pasted image 20261113191938.png]] 
把两个本地模块加入到一个工作区里面去
2. 怎么创建go.work文件：
进入到包含两个模块的总文件夹，输入：
```
go work init ./user-service ./common
```
**总结一下：多个.go文件在一个package里，很多个package在一个Module里，go.work又把相关的模块连在一起，最终组成一个大的**go项目