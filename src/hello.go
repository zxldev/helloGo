package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"strings"

	"image"
	"bytes"
	"image/png"
	"encoding/base64"
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
	// if 可以像for一样先执行一个语句
	if x := 100 ; total < 1000 {
		fmt.Print(x)
	} else if total == 1600 && sum == 45 {
		fmt.Println("test if success!")
	} else {
	}
	//switch 1 初始化 2 没有括号 3 反向break,关键字 fallthrough
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

func pointer()  {
	//指针和C中指针类似，但是没有指针的运算 不能执行 p = p + 1
	j :=100
	p := &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

type Point struct {
	x int
	y int
}



type Circle struct {
	radius int
	middle Point
}

func (c *Circle)getArea() float32{
	return math.Pi*float32(c.radius)*float32(c.radius);
}

func printStruct()  {
	circle := Circle{10,Point{1,1}}
	fmt.Println(circle)
	fmt.Println(circle.radius)
	fmt.Println(circle.middle.x)
}

func printArray(){
	var x [2]int
	x[0] = 1
	x[1] = 2
	fmt.Println(x)

	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	//切片
	fmt.Println("s[:2] ==",s[:2])

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	table := [][]string{
		[]string{"A","B","C"},
		[]string{"A","B","C"},
		[]string{"A","B","C"},
	}
	fmt.Println("二维数组",table[1][1],"长度",len(table),"链接",strings.Join(table[2]," "))

	test := []int{1,2,3,4,5,6,7,8}
	test2 := test[3:]
	test3 := test[:3]
	//len当前元素大小 cap 当前空间大小
	fmt.Println(len(test2),cap(test2))
	fmt.Println(len(test3),cap(test3))

	s = append(s,3)
	printSlice("x",s)

	// range 返回切片key,val键值对 想单于 in
	for i,v := range s {
		fmt.Printf("%d => %d\n", i, v)
	}

	for i:= range s {
		fmt.Printf("%d => \n", i)
	}

	for _,v:= range s {
		fmt.Printf(" => %d\n", v)
	}

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func Pic(dx, dy int) [][]uint8 {
	table := [][]uint8{}

	for i := 0; i < dy; i++ {
		line := []uint8{}
		for j := 0; j < dx; j++ {
			line = append(line, uint8(float64(j)*math.Log(float64(i))))
		}
		table = append(table, line)
	}
	return table
}


func WordCount(s string) map[string]int {

	data := strings.Split(s," ");
	mapData := map[string]int{}
	for _,val := range data {
		mapData[val]++
	}
	for key,val := range mapData{
		fmt.Println(key,val)
	}
	return mapData
}


func main() {
	for x:=1 ;x <5 ; x++ {
		//先进后出
		defer fmt.Println("这句话写在前面，但是会等待main(上层函数)调用后执行",x)
	}
	fmt.Println("hello Go")
	fmt.Println("rand number:", rand.Intn(100))
	fmt.Println("Pi is value:", math.Pi)
	fmt.Println(add(4, 6))
	fmt.Println(swap("abc", "edf"))
	fmt.Println(div(21, 5))
	variableDef()
	loop()
	fmt.Println(Sqrt(2))
	pointer()
	printStruct()
	printArray()
	Show(Pic)
	WordCount("ssss ss s s ss s s ");

	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
	//删除map元素
	delete(addrs,"loopback")
	//检测是否存在
	elem,isExist := addrs["loopback"]
	fmt.Println(elem,isExist)


}

func Show(f func(int, int) [][]uint8) {
	const (
		dx = 2
		dy = 2
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	ShowImage(m)
}


type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ad IPAddr) String() string{
	return fmt.Sprintf("%v.%v.%v.%v", ad[0],ad[1],ad[2],ad[3])
}


func ShowImage(m image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println("IMAGE:" + enc)
}
