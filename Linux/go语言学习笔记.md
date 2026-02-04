# 结构

```go
package main

import"fmt"

func main(){
    fmt.Println("Hello,world")
}
```

# 变量

```go
var d = 4 //外部声明只能这样
func main(){
    var b = 2

    a := 2 //声明的另一种

    var c int  //不带内容需声明变量类型
    c = 3

    var e,f = 1,2

    fmt.Println(a)
    fmt.Println(b)
}
```

# 常量

```go
const var a int = 1

```

# 数据类型

```go
int
int8
int16
int32
int64
uint
uint16

float32
float64

var Mybool bool= false //默认 or
var Mybool bool= true 



```

# 字符串

```go
var myString string ="hello"+" "+"world"

```



# 函数

```go
len( "test")//获取字符串长度 >>4

```

# 定义函数

```go
func printMe(/*输入的值*/){

    fmt.Println("hello")

}


func(num1 int,num2 int)int/*函数输出类型*/{
    x := num1 + num2
    return x
}


func(num1 int,num2 int)(int, int /*返回多个值*/){
    x := num1 / num2
    y := num1 % num2
    return x,y
}
```

# 输出函数

```go
fmt.Printf("ok it's string print,a num like %v %v",x,y)
```

# 错误返回

```go
import ("errors")

/******************************/

var err error//初始化为 nil
if x == 0{
    err = errors.New("it cannot be zero")
}
if err != nil{
    fmt.Printf("ERROR")
}
```

**Go语言中switch中隐含break,不需要写break.**

# 数组

```go
var arr [3]int32//默认赋值0
var MySlice
```

# slice类型

```go
MySlice :=  []int32 {1,2,3}

MySlice = append(MySlice,4)//slice类型 创建一个新数组容纳数组追加上4

fmt.Println(MySlice)//1,2,3,4

Slice2 := make([]int32,0)//定义slice长度
```

# map类型

```go
var MyMap = map[string/*key类型*/]uint8{"Adam":23,"Sarah":45}

fmt.Println(MyMap["Adam"])//23



/**************************/
m := make(map[string]int)
m["a"] = 1
m["b"] = 2
```

# 循环

```go
file := []string{"a","b"}
for idx/*索引*/,readfile := range file{
    fmt.Printf(readfile)//遍历file,输出副本    
}

```

# os包

```go
os.ReadDIr(dir)//返回([]DirEntry, error)
//读取读取指定目录下的所有文件 / 目录项（不递归）
//目录列表(名称)和错误(权限不足等)


file.IsDir()//返回bool值 true->目录;false->文件
//判断目录项是文件还是目录


os.Rename(oldpath,newpath/*字符串类型*/)
//输入前后文件或目录的路径
//返回一个error
//重命名文件或目录

```

# filepath包

```go
filepath.Ext(file.Name())
//输入文件名或文件路径
//返回包含.的文件后缀(无后缀返回空字符串"")

filepath.Join(dir,file.name())
//拼接文件路径避免手动拼接出错
//遇到绝对路径会覆盖
```

  


