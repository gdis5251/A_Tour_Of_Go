# Go 语言--变量逃逸

堆和栈各有优缺点，该怎么在编程中处理这个问题呢？在 C/C++语言中，需要开发者自己学习如何进行内存分配，选用怎样的内存分配方式来适应不同的算法需求。比如，函数局部变量尽量使用栈；全局变量、结构体成员使用堆分配等。程序员不得不花费很多年的时间在不同的项目中学习、记忆这些概念并加以实践和使用。

Go语言将这个过程整合到编译器中，命名为“变量逃逸分析”。这个技术由编译器分析代码的特征和代码生命期，决定应该如何堆还是栈进行内存分配，即使程序员使用Go语言完成了整个工程后也不会感受到这个过程。  

这样的概念可能会让你一知半解，来先看一段代码

## 逃逸分析

```go
package main

import "fmt"

func dummy(a int) int {
	var b int = a
	return b
}

func void() {

}

func main() {
	// 声明变量 c 并打印
	var c int

	void()

	fmt.Println(c, dummy(0))
}
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190917195745407.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MjY3ODUwNw==,size_16,color_FFFFFF,t_70)

使用 go run 运行程序时，-gcflags 参数是编译参数。其中 -m 表示进行内存分配分析，-l 表示避免程序内联，也就是避免进行程序优化。

程序运行结果分析如下：

- 输出第 2 行告知“main 的变量 c 逃逸到堆”。
- 第 3 行告知“dummy(0)调用逃逸到堆”。由于 dummy() 函数会返回一个整型值，这个值被 fmt.Println 使用后还是会在其声明后继续在 main() 函数中存在。
- 第 4 行，这句提示是默认的，可以忽略。

## 取地址发生逃逸

```go
package main

import "fmt"

// 声明空结构体测试结构体逃逸情况
type Data struct {
}

func dummy() *Data {

    // 实例化c为Data类型
    var c Data

    //返回函数局部变量地址
    return &c
}

func main() {

    fmt.Println(dummy())
}
```

代码说明如下：

- 第 6 行，声明一个空的结构体做结构体逃逸分析。
- 第 9 行，将 dummy() 函数的返回值修改为 *Data 指针类型。
- 第 12 行，将 c 变量声明为 Data 类型，此时 c 的结构体为值类型。
- 第 15 行，取函数局部变量 c 的地址并返回。Go语言的特性允许这样做。
- 第 20 行，打印 dummy() 函数的返回值。

执行逃逸分析：

$ go run -gcflags "-m -l" main.go

\# command-line-arguments

./main.go:15:9: &c escapes to heap

./main.go:12:6: moved to heap: c

./main.go:20:19: dummy() escapes to heap

./main.go:20:13: main ... argument does not escape

&{}

注意第 4 行出现了新的提示：将 c 移到堆中。这句话表示，Go 编译器已经确认如果将 c 变量分配在栈上是无法保证程序最终结果的。如果坚持这样做，dummy() 的返回值将是 Data 结构的一个不可预知的内存地址。这种情况一般是 C/C++ 语言中容易犯错的地方：引用了一个函数局部变量的地址。

Go语言最终选择将 c 的 Data 结构分配在堆上。然后由垃圾回收器去回收 c 的内存。

## 总结

在使用Go语言进行编程时，Go语言的设计者不希望开发者将精力放在内存应该分配在栈还是堆上的问题。编译器会自动帮助开发者完成这个纠结的选择。但变量逃逸分析也是需要了解的一个编译器技术，这个技术不仅用于Go语言，在 Java等语言的编译器优化上也使用了类似的技术。

编译器觉得变量应该分配在堆和栈上的原则是：

- 变量是否被取地址。
- 变量是否发生逃逸。