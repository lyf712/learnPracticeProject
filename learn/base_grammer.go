package main

/**
 * stater
 */
//func main() {
//	fmt.Printf("start your first go app\n")
//	//dataType()
//	//advanceType()
//	//baseSentence()
//	testRange()
//}

/*
 * 语法基础知识
 */
func dataType() {
	// ==================变量&常量基础====================
	var var1 int
	var1 = 1
	println(var1)
	// int 可省略，可自动判断语言类型（该种语言应当效率较低？需要运行时检查？
	// TODO Think it
	var var2 int = 2
	println(var2)
	var3 := 3
	println(var3)

	var var4, var5 int = 1, 2
	println(var4 + var5)

	const constVal = 1
	//Cannot assign to constVal
	//constVal = 2
	/**
	iota，特殊常量，可以认为是一个可以被编译器修改的常量。
	在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。
	iota 可以被用作枚举值
	*/
	const name1 = iota
	// 枚举
	const (
		a = iota
		b = iota
	)
	println(a, b)

}

/**
 * 数组、指针、结构体
 */
func advanceType() {
	// 对比类 --- 未有封装的概念？
	//type struct_variable_type struct {
	//	member definition
	//	member definition
	//	...
	//	member definition
	//}
	type Person struct {
		age  int
		name string
	}
	type Compeny struct {
		id     int
		leader Person
	}
	var Person1 Person
	Person1.age = 1
	Person1.name = "hello"

	// Slice 切片--动态数组
	var arr1 = make([]int, 2)
	arr1[0] = 1
	arr1 = append(arr1, 2)
	arr1 = append(arr1, 3, 4)
	println(arr1)
	for i := 0; i < len(arr1); i++ {
		println(arr1[i])
	}
}

/**
 * 基础语句
 */
func baseSentence() {
	// if else;switch;
	// 基本小括号可省去，但大括号必须有
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			println(i)
		}
	}

	//for i, i2 := range collection {
	//
	//}

	//var count int = 1
	//for {
	//	println("count:%d", count)
	//	time.Sleep(2000)
	//}

	// select!!
	/**
	A channel provides a mechanism for concurrently executing functions
	to communicate by sending  and receiving  values of a specified element type.
	The value of an uninitialized channel is nil.
	*/
	//var c1, c2, c3 chan int
	//var i1, i2 int

	//select {
	//case i1 = <-c1:
	//	fmt.Printf("received ", i1, " from c1\n")
	//case c2 <- i2:
	//	fmt.Printf("sent ", i2, " to c2\n")
	//case i3, ok := (<-c3): // same as: i3, ok := <-c3
	//	if ok {
	//		fmt.Printf("received ", i3, " from c3\n")
	//	} else {
	//		fmt.Printf("c3 is closed\n")
	//	}
	//default:
	//	fmt.Printf("no communication\n")
	//}

}

/**
 * 语言函数
 * 1.基本的写法
 * 2.闭包的概念及使用（额外掌握、同scala）
 * 3.defer语句
 */
func swap(var1 int, var2 int) {
	var tmp = var1
	var2 = tmp
}
func max(var1, var2 int) int {
	if var1 > var2 {
		return var1
	} else {
		return var2
	}
}

func testRange() {
	//Go 语言中 range 关键字用于for循环中迭代数组(array)、
	//切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引值，在集合中返回 key-value 对的 key 值。

	//
	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	// index,val
	for i, val := range arr {
		println(i, val)
	}

	kvs := map[string]string{"a": "1", "b": "2"}
	for key, value := range kvs {
		println(key, value)
	}
}

type IPhone interface {
	call()
}
type Phone struct {
}

// 结构体类似Java中类的概念
func (phone Phone) call() {

}

//type Phone interface {
//	call()
//}
//
//type NokiaPhone struct {
//}
//
//func (nokiaPhone NokiaPhone) call() {
//	fmt.Println("I am Nokia, I can call you!")
//}
