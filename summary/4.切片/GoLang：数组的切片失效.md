# GoLang：数组的切片失效

学习切片的时候想到一个问题：如果一个切片是引用一个完整的数组，那么给这个切片 append 一个元素后，这个 append 的元素会以什么样的形式存在？

先看代码：

```go
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := arr[:]
	s[8] = 666
	fmt.Println(arr)
	fmt.Println(s)
	fmt.Printf("&arr[8] = %p, &s[8] = %p\n", &arr[8], &s[8])


	// a := 10
	// fmt.Println(&a)

	s = append(s, 10)
	fmt.Println(s)

	fmt.Printf("&arr[8] = %p, &s[8] = %p, &s[9] = %p\n", &arr[8], &s[8], &s[9])

	s[8] = 921
	fmt.Println(arr)
	fmt.Println(s)
```

当切片 s 去引用整个数组 arr 时，改变切片就是改变数组，这个都清楚。

可以如果我现在给这个切片 s ，append 一个元素 10，那么这个 10 将会存在哪里呢？切片是否还是数组的引用呢？

所以我做个这个小实验

结论：

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191008211518392.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MjY3ODUwNw==,size_16,color_FFFFFF,t_70)

当切片引用整个数组时，它的长度和容量就是数组的长度。所以当我 append 的时候，会发现容量不够，会发生扩容操作。

所以我推断，当发生扩容时，会在堆上开辟一段连续的空间将数组完整的拷贝过来（因为此时的数组中的元素存在栈上），再加上追加的元素。所以我们看到后来 arr[8] 和 s[8] 的地址都已经不一样了。之后再操作切片就跟数组没关系了。

叮~:bell: