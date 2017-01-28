package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
)

//类型在变量名之后
func add(x int, y int) int {
	return x + y
}

//当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略
func add2(x, y int) int {
	return x + y
}

//函数可以返回任意数量的返回值
func swap(x, y string) (string, string) {
	return y, x
}

//命名返回值
func div(x, y int) (result, rest int) {
	result = int(x / y)
	rest = x % y
	return
}

//变量定义
func variableDef() {
	//方式一： var 关键字定义
	var a, b int
	var c, d, e int = 1, 2, 3
	//方式二: := 运算符定义
	f, g, h := true, false, "testString"
	//错误，会报未定义
	//i = 4
	//基本数据类型
	var ( //变量打包定义
		v1  bool
		v2  string
		v3  int
		v4  int8
		v5  int16
		v6  int32
		v7  int64
		v8  uint
		v9  uint8
		v10 uint16
		v11 uint32
		v12 uint64
		v13 uintptr
		v14 byte // uint8 的别名
		v15 rune // int32 的别名
		// 代表一个Unicode码
		v16 float32
		v17 float64
		v18 complex64
		v19 complex128
	)
	//常量定义
	const World = "世界"
	fmt.Println(a, b, c, d, e, f, g, h)
	fmt.Println(v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19)
}

func loop() {
	//for 是c中不要括号
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	//for 是c 中 while 不要括号
	total := 100
	for total < 1000 {
		total += total
	}
	fmt.Println(total)
	// if 是c中不要括号
	if total < 1000 {

	} else if total == 1600 && sum == 45 {
		fmt.Println("test if success!")
	} else {
	}
	//switch 1 初始化 2 没有括号 3 反向break
	switch os := runtime.GOOS;os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s",os)
	}
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0 ; i < 100000; i ++{
		z = z - ((z*z-x)/2*z)
 	}
	return z
}

func main() {
	fmt.Println("hello Go")
	fmt.Println("rand number:", rand.Intn(100))
	fmt.Println("Pi is value:", math.Pi)
	fmt.Println(add(4, 6))
	fmt.Println(swap("abc", "edf"))
	fmt.Println(div(21, 5))
	variableDef()
	loop()
	fmt.Println(Sqrt(2))
}
