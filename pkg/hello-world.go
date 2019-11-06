//package main
//
//import (
//	"fmt"
//	"os"
//)
//
////func main() {
////	fmt.Printf("xuanz")
////	var s, sep string
////	for i := 1; i < len(os.Args); i++ {
////		s += sep + os.Args[i]
////		sep = " +"
////	}
////	fmt.Printf(s + "cdcd")
////}
//
//func main() {
//	fmt.Printf("zhendong")
//	var args string
//	s, sep := "", ""
//	for args := range os.Args[1:] {
//		s += sep + args
//		sep = " "
//	}
//	fmt.Printf(s + "cdcd")
//}
// 2.
//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func main() {
//	fmt.Println(math.SqrtPi)
//}
//
////3.
//package main
//
//import "fmt"
//
//func add(x, y, z int) int {
//	return (x + y) * z
//}
//
//func main() {
//	fmt.Println(add(42, 13, 66))
//}

//4
//package main
//
//import "fmt"
//
//func swap(x, y string) (string, string, string) {
//	return y, x, "zhendong"
//}
//
//func main() {
//	a, b, c := swap("hello", "world")
//	fmt.Println(a, b, c)
//}

//5
//package main
//
//import "fmt"
//
////var z string
//var c, i int
//
//func split(sum int) (x, y int, z string) {
//	x = sum * 4 / 9
//	y = sum - x
//	z = "zhendong"
//	//return x, y, z
//	return
//}
//
//func main() {
//	fmt.Println(split(10))
//}

//6
//package main
//
//import "fmt"
//
//var i, j int = 1, 2
//
//func main() {
//	var c, python, java = true, false, "no!"
//	fmt.Println(i, j, c, python, java)
//}

//7
//package main
//
//import "fmt"
//
//func main() {
//	var i, j int = 1, 2
//	k := 3
//	c, python, java := true, false, "no!"
//
//	fmt.Println(i, j, k, c, python, java)
//}

//8
//package main
//
//import (
//	"fmt"
//	"math/cmplx"
//)
//
//var (
//	ToBe   bool       = false
//	MaxInt uint64     = 1<<64 - 1
//	z      complex128 = cmplx.Sqrt(-5 + 12i)
//)
//
//func main() {
//	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
//	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
//	fmt.Printf("Type: %T Value: %v\n", z, z)
//}

//9
//package main
//
//import "fmt"
//
//func main() {
//	var i int
//	var f float64
//	var b bool
//	var s string
//	fmt.Printf("%v %v %v %q\n", i, f, b, s)
//	fmt.Printf("%T %T %T %T\n", i, f, b, s)
//}

//10
//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func main() {
//	var x, y int = 3, 4
//	var f float64 = math.Sqrt(float64(x*x + y*y))
//	//var z uint = uint(f)
//	fmt.Println(x, y, f)
//	fmt.Printf(string("%T %T %T\n"), x, y, f)
//}

//////////////////////////////////////////////////////////////////////////////////////////////

//类型推导
//
//在声明一个变量而不指定其类型时（即使用不带类型的 := 语法或 var = 表达式语法），变量的类型由右值推导得出。
//
//当右值声明了类型时，新变量的类型与其相同：
//
//var i int
//j := i // j 也是一个 int
//
//不过当右边包含未指明类型的数值常量时，新变量的类型就可能是 int, float64 或 complex128 了，这取决于常量的精度：
//
//i := 42           // int
//f := 3.142        // float64
//g := 0.867 + 0.5i // complex128
//
//尝试修改示例代码中 v 的初始值，并观察它是如何影响类型的。

//package main
//
//import "fmt"
//
//func main() {
//	//var k int
//	//v := k
//	v := 42.241651 + 0.3i // 修改这里！
//	fmt.Printf("%v is of type %T\n", v, v)
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func main() {
//	fmt.Println(string("Π："), math.Pi)
//}

//package main
//
//import "fmt"
//
//func swap(x, y string) (string, string, string) {
//	var z string = "zhendong"
//	return y, x, z
//}
//
//func main() {
//	a, b, f := swap("hello,", "world")
//	fmt.Println(a, b, f)
//}

//package main
//
//import "fmt"
//
//func split(sum int) (x, y int, z string) {
//	x = sum * 4 / 9
//	y = sum - x
//	z = "zhendong"
//	return
//}
//
//func main() {
//	fmt.Println(split(17))
//}

//package main
//
//import "fmt"
//
//var java = true
//var c, python bool
//
//func main() {
//	var i int
//	fmt.Println(i, c, python, java)
//}

//package main
//
//import "fmt"
//
//var i, j = 1, "zhendong"
//var k int = 3
//
//func main() {
//	var c, python, java = true, false, "no!"
//	fmt.Println(i, j, c, python, java, k)
//}

//package main
//
//import "fmt"
//
//func main() {
//	var i, j int = 1, 2
//	k := 3
//	c, python, java := true, false, "no!"
//
//	fmt.Println(i, j, k, c, python, java)
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func main() {
//	var x, y int = 3, 4
//	var f float64 = math.Sqrt(float64(x*x + y*y))
//	var z uint = uint(f)
//	fmt.Println(x, y, z)
//	fmt.Printf(string("%T %T %T %T\n"), x, y, f, z)
//}

//package main
//
//import "fmt"
//
//func main() {
//	v := 42           // 修改这里！
//	f := 3.142        // float64
//	g := 0.867 + 0.5i // complex128
//	fmt.Printf("%v is of type %T\n", v, v)
//	fmt.Printf("%v is of type %T\n", f, f)
//	fmt.Printf("%v is of type %T\n", g, g)
//}

//package main
//
//import "fmt"
//
//const Pi = 3.14
//
//func main() {
//	const World = "世界"
//	fmt.Println("Hello", World)
//	fmt.Println("Happy", Pi, "Day")
//
//	const Truth = true
//	fmt.Println("Go rules?", Truth)
//}

//package main
//
//import "fmt"
//
//const (
//	// 将 1 左移 100 位来创建一个非常大的数字
//	// 即这个数的二进制是 1 后面跟着 100 个 0
//	Big = 3 << 100
//	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
//	Small = Big >> 98
//)
//
//func needInt(x int) int { return x*10 + 1 }
//func needFloat(x float64) float64 {
//	return x * 0.1
//}
//
//func main() {
//	fmt.Println(needInt(Small))
//	fmt.Println(needFloat(Small))
//	fmt.Println(needFloat(Big))
//}

//package main
//
//import "fmt"
//
//func main() {
//	sum := 0
//	for i := 0; i < 10; i++ {
//		sum += i
//		fmt.Println(sum)
//	}
//	fmt.Println(sum)
//}

//package main
//
//import "fmt"
//
//func main() {
//	sum := 1
//	for sum < 1000 {
//		sum += sum
//	}
//	fmt.Println(sum)
//}

//package main
//
//import "fmt"
//
//func main() {
//	for {
//		fmt.Println("zhendong.....")
//	}
//}
//
//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func sqrt(x float64) string {
//	if x < 0 {
//		return sqrt(-x)
//	}
//	return fmt.Sprint(math.Sqrt(x))
//}
//
//func main() {
//	fmt.Println(sqrt(9), sqrt(-4))
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func pow(x, n, lim float64) float64 {
//	if v := math.Pow(x, n); v < lim {
//		return v
//	}
//	return lim
//}
//
//func main() {
//	fmt.Println(
//		pow(3, 2, 10),
//		pow(3, 3, 20),
//	)
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func pow(x, n, lim float64) float64 {
//	if v := math.Pow(x, n); v < lim {
//		return v
//	} else {
//		fmt.Printf("%g >= %g\n", v, lim)
//	}
//	// 这里开始就不能使用 v 了
//	return lim
//}
//
//func main() {
//	fmt.Println(
//		pow(3, 2, 10),
//		pow(3, 3, 20),
//	)
//}
//package main
//
//import (
//	"fmt"
//)
//
//func Sqrt(x float64) float64 {
//	z := 1.0
//	//z := float64(1)
//	z -= (z*z - x) / (2 * z)
//	return z
//}
//
//func main() {
//	fmt.Println(Sqrt(2))
//}

//package main
//
//import (
//	"fmt"
//	"runtime"
//)
//
//func main() {
//	fmt.Print("Go runs on ")
//	switch os := runtime.GOOS; os {
//	case "darwin":
//		fmt.Println("OS X.")
//	case "linux":
//		fmt.Println("Linux.")
//	default:
//		// freebsd, openbsd,
//		// plan9, windows...
//		fmt.Printf("%s.\n", os)
//	}
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	fmt.Println("When's Saturday?")
//	fmt.Println("time.Now().Weekday(): ", time.Now().Weekday())
//	fmt.Println("time.Saturday: ", time.Saturday)
//	today := time.Now().Weekday()
//	switch time.Saturday {
//	case today + 0:
//		fmt.Println("Today.")
//	case today + 1:
//		fmt.Println("Tomorrow.")
//	case today + 2:
//		fmt.Println("In two days.")
//	default:
//		fmt.Println("Too far away.")
//	}
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	t := time.Now()
//	switch {
//	case t.Hour() < 12:
//		fmt.Println("Good morning!")
//	case t.Hour() < 17:
//		fmt.Println("Good afternoon.")
//	default:
//		fmt.Println("Good evening.")
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	defer fmt.Println("world")
//
//	fmt.Println("hello")
//}

//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("counting")
//
//	for i := 0; i < 10; i++ {
//		defer fmt.Println(i)
//	}
//
//	fmt.Println("done")
//}

//package main
//
//import "fmt"
//
//func main() {
//	i, j := 42, 2701
//
//	fmt.Println("i 内存地址：", &i)
//	// 指向 i
//	p := &i
//	// 通过指针读取 i 的值
//	fmt.Println(*p)
//	// 通过指针读取 i 的值
//	fmt.Println("P的值，再次进行内存地址映射：", &*p)
//	// 通过指针设置 i 的值
//	*p = 21
//	// 查看 i 的值
//	fmt.Println(i)
//
//	p = &j         // 指向 j
//	*p = *p / 37   // 通过指针对 j 进行除法运算
//	fmt.Println(j) // 查看 j 的值
//}

//package main
//
//import (
//	"fmt"
//)
//
//type Vertex struct {
//	X int
//	Y int
//	z string
//}
//
//func main() {
//	fmt.Println(Vertex{1, 2, "zhendong"})
//}

//package main
//
//import "fmt"
//
//type Vertex struct {
//	X int
//	Y int
//	z string
//}
//
//func main() {
//	fmt.Println(Vertex{1, 2, "zhendong"})
//	v := Vertex{1, 2, "zhendong"}
//	v.X = 7
//	fmt.Println(v.X)
//	fmt.Println(v.z)
//	v.z = "zhendongran"
//	fmt.Println(v.z)
//}

//package main
//
//import (
//	"fmt"
//)
//
//type Vertex struct {
//	X int
//	Y int
//}
//
//func main() {
//	v := Vertex{1, 2}
//	p := &v
//	p.X = len("zhendongran")
//	fmt.Println(v)
//}

//package main
//
//import "fmt"
//
//type Vertex struct {
//	X, Y int
//}
//
//var (
//	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
//	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
//	v3 = Vertex{}      // X:0 Y:0
//	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
//)
//
//func main() {
//	p.Y = 4
//	v1.Y = 4
//	fmt.Println(v1, p, v2, v3)
//}

//package main
//
//import "fmt"
//
//func main() {
//
//	var a [2]string
//	a[0] = "Hello"
//	a[1] = "World"
//	fmt.Println(a[0], a[1])
//	fmt.Println(a)
//
//	primes := [6]int{2, 3, 5, 7, 11, 13}
//	fmt.Println(primes)
//}

//package main
//
//import "fmt"
//
//func main() {
//	names := [4]string{
//		"John",
//		"Paul",
//		"George",
//		"Ringo",
//	}
//	fmt.Println(names)
//
//	a := names[0:2]
//	b := names[1:3]
//	fmt.Println(a, b)
//
//	b[0] = "ZhenDongXXXXXX"
//	fmt.Println(a, b)
//	fmt.Println(names)
//}

//package main
//
//import "fmt"
//
//func main() {
//	q := []int{2, 3, 5, 7, 11, 13}
//	fmt.Println(q)
//
//	r := []bool{true, false, true, true, false, true}
//	fmt.Println(r)
//
//	s := []struct {
//		i int
//		b bool
//	}{
//		{2, true},
//		{3, false},
//		{5, true},
//		{7, true},
//		{11, false},
//		{13, true},
//	}
//	fmt.Println(s)
//}

//package main
//
//import "fmt"
//
//func main() {
//	s := []int{2, 3, 5, 7, 11, 13}
//
//	s = s[1:4]
//	fmt.Println("[1:4]: ", s)
//
//	s = s[:2]
//	fmt.Println("[:2]: ", s)
//
//	s = s[1:]
//	fmt.Println("[1:]: ", s)
//
//	s = s[:]
//	fmt.Println("[:]: ", s)
//}

//package main
//
//import "fmt"
//
//func main() {
//	var a, b, d, e int
//	var c, f []int
//	s := []int{2, 3, 5, 7, 11, 13}
//	a, b, c = printSlice(s)
//	d, e, f = printSlices(s)
//	fmt.Println("NORMAL RETURN: ", a, b, c)
//	fmt.Println("SPEICAL RETURN: ", d, e, f, "\n")
//
//	// 截取切片使其长度为 0
//	s = s[:0]
//	printSlice(s)
//	printSlices(s)
//
//	// 拓展其长度
//	s = s[:4]
//	printSlice(s)
//	printSlices(s)
//
//	// 舍弃前两个值
//	s = s[2:]
//	printSlice(s)
//	printSlices(s)
//}
//
//func printSlice(s []int) (int, int, []int) {
//	fmt.Printf("NORMAL RETURN: len=%d cap=%d %v\n", len(s), cap(s), s)
//	return len(s), cap(s), s
//}
//
//func printSlices(s []int) (lens, caps int, sa []int) {
//	fmt.Printf("SPEICAL RETURN: len=%d cap=%d %v\n\n", len(s), cap(s), s)
//	lens = len(s)
//	caps = cap(s)
//	sa = s
//	return
//}

//package main
//
//import "fmt"
//
//func main() {
//	var s []int
//	fmt.Println(s, len(s), cap(s))
//	if s == nil {
//		fmt.Println("nil!")
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	a := make([]int, 5)
//	printSlice("a", a)
//
//	b := make([]int, 0, 5)
//	printSlice("b", b)
//
//	c := b[:2]
//	printSlice("c", c)
//
//	d := c[2:5]
//	printSlice("d", d)
//}
//
//func printSlice(s string, x []int) {
//	fmt.Printf("%s len=%d cap=%d %v\n",
//		s, len(x), cap(x), x)
//}

//package main
//
//import (
//	"fmt"
//	"strings"
//)
//
//func main() {
//	//sArr := []string{"a", "b", "c", "d"}
//	////fmt.Println(strings.Join(sArr, ""))
//	//var b strings.Builder
//	//for _, str := range sArr {
//	//	b.WriteString(str) //也可以用fmt.FPrintf(&b,"%s",str)写入其它类型
//	//}
//	//fmt.Println(b)
//	//fmt.Println(b.String())
//	// 创建一个井字板（经典游戏）
//	a := []strings.Builder{}
//	fmt.Println(a)
//	s := [][][]int{
//		{
//			[]int{1},
//			[]int{1, 1},
//		},
//		{
//			[]int{2},
//			[]int{2, 2},
//		},
//	}
//	fmt.Println(s)
//	board := [][]string{
//		[]string{"_", "_", "_"},
//		[]string{"_", "_", "_"},
//		[]string{"_", "_", "_"},
//		[]string{"_", "_", "_"},
//	}
//
//	// 两个玩家轮流打上 X 和 O
//	board[0][0] = "X"
//	board[2][2] = "O"
//	board[1][2] = "X"
//	board[1][0] = "O"
//	board[0][2] = "X"
//
//	fmt.Println("board长度:", len(board))
//	fmt.Println("board嵌套数组:", board)
//	fmt.Println("board第一个数组:", board[0])
//	fmt.Println("board嵌套第一个数组的第一个值:", board[1][1])
//	fmt.Println("board嵌套第一个数组进行字符***拼接:", strings.Join(board[1], "***"))
//
//	for i := 0; i < len(board); i++ {
//		fmt.Printf("%s\n", strings.Join(board[i], " "))
//	}
//}
//
//package main
//
//import "fmt"
//
//func main() {
//	var s []int
//	printSlice(s)
//
//	// 添加一个空切片
//	s = append(s, 0)
//	printSlice(s)
//
//	// 这个切片会按需增长
//	s = append(s, 1)
//	printSlice(s)
//
//	// 可以一次性添加多个元素
//	s = append(s, 2, 3, 4)
//	printSlice(s)
//	s = append(s, 5, 6)
//	printSlice(s)
//	//s = append(s, 7, 8, 9, 10, 11, 11, 11, 11, 11)
//	//printSlice(s)
//}
//
//func printSlice(s []int) {
//	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
//}

//package main
//
//import "fmt"
//
////var pow []int{1, 2, 4, 8, 16, 32, 64, 128}
//
//func main() {
//	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
//	for i, v := range pow {
//		fmt.Printf("角标:%d ；角标值:%d\n", i, v)
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	pow := make([]int, 10)
//	for i := range pow {
//		// 将1的二级制进行左移i位，因为i是
//		fmt.Printf("%T", i)
//		//i := uint(i)
//		//pow[i] = 1 << i // == 2**i
//		pow[i] = 1 << uint(i)
//		// == 2**i
//		//fmt.Printf("%T", i)
//		fmt.Print(i, "\n")
//	}
//	for _, value := range pow {
//		fmt.Printf("%d\n", value)
//	}
//}

//package main
//
//import "fmt"
//
//type Vertex struct {
//	Lat, Long float64
//}
//
//var m map[string]Vertex
//
//func main() {
//	m = make(map[string]Vertex)
//	m["Bell Labs"] = Vertex{
//		40.68433, -74.39967,
//	}
//	fmt.Println(m["Bell Labs"])
//}

//package main
//
//import (
//	"fmt"
//)
//
//type info struct {
//	name string
//	age  int
//}
//
//type Vertex struct {
//	Lat, Long float64
//}
//
//var m = map[string]Vertex{
//	"Bell Labs": Vertex{
//		40.68433, -74.39967,
//	},
//	//"Google": Vertex{
//	//	37.42202, -122.08408,
//	//},
//}
//
//func main() {
//	//m["Zhendong"] = Vertex{
//	//	1994, -9,
//	//}
//	fmt.Println(m)
//}

//package main
//
//import "fmt"
//
//type Vertex struct {
//	Lat, Long float64
//	info      struct {
//		name string
//	}
//}
//
//var m = map[string]Vertex{
//	"Bell Labs": {40.68433, -74.39967, {"sss"}},
//	"Google":    {37.42202, -122.08408},
//}
//
//func main() {
//	fmt.Println(m)
//}

//package main
//
//import "fmt"
//
//var m map[string]int
//
//func main() {
//	m = make(map[string]int)
//
//	m["Answer"] = 42
//	fmt.Println("The value:", m["Answer"])
//
//	m["Answer"] = 48
//	fmt.Println("The value:", m["Answer"])
//
//	delete(m, "Answer")
//	fmt.Println("The value:", m["Answer"])
//
//	v, ok := m["Answer"]
//	fmt.Println("The value:", v, "Present?", ok)
//}

//package main
//
//import (
//	"fmt"
//	"math"
//	"strconv"
//)
//
//func compute(fn func(float64, float64, string) string) string {
//	return fn(3, 4, "zhendong")
//}
//
//func main() {
//	hypot := func(x, y float64, z string) string {
//		//fmt.Println(z)
//		//return math.Sqrt(x*x+y*y)
//		b := math.Sqrt(x*x + y*y)
//		a := z + strconv.Itoa(int(b))
//		fmt.Println(x, y, a, b)
//		fmt.Println(strconv.Itoa(int(b)))
//		//return a
//		return z + strconv.Itoa(int(math.Sqrt(x*x+y*y)))
//	}
//
//	fmt.Println(compute(hypot))
//
//}

//package main
//
//import (
//	"fmt"
//)
//
//func adder() func(int) int {
//	sum := 0
//	return func(x int) int {
//		sum += x
//		return sum
//	}
//}
//
//func main() {
//	adder()
//	pos, neg := adder(), adder()
//	fmt.Println(pos)
//	for i := 0; i < 10; i++ {
//		fmt.Println(
//			pos(i),
//			neg(-2*i),
//		)
//	}
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//type Vertex struct {
//	X, Y float64
//}
//
//func (v Vertex) Abs() float64 {
//	fmt.Println("函数内", v.Y, v.X)
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func main() {
//	v := Vertex{3, 4}
//	fmt.Println("函数外", v.Y, v.X)
//	fmt.Println(v.Abs())
//}

//package main
//
//import (
//	"fmt"
//	"math"
//	"strconv"
//)
//
//type Vertex struct {
//	X, Y float64
//	Z    string
//}
//
//func Abs(v Vertex) string {
//	return strconv.Itoa(int(math.Sqrt(v.X*v.X+v.Y*v.Y))) + v.Z
//}
//
//func main() {
//	v := Vertex{3, 4, "zhendong"}
//	fmt.Println(Abs(v))
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//type MyFloat float64
//
//func (f MyFloat) Abs() float64 {
//	if f < 0 {
//		return float64(-f)
//	}
//	return float64(f)
//}
//
//func main() {
//	f := MyFloat(-math.Sqrt2)
//	fmt.Println(f.Abs())
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//type Vertex struct {
//	X, Y float64
//}
//
//func (v Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func (v *Vertex) Scale(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
//func main() {
//	v := Vertex{3, 4}
//	fmt.Println(v.Abs())
//	fmt.Println(v)
//	v.Scale(10)
//	fmt.Println(v.Abs())
//	fmt.Println(v)
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//type Vertex struct {
//	X, Y float64
//}
//
//func Abs(v Vertex) float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func Scale(v Vertex, f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
//func main() {
//	v := Vertex{3, 4}
//	fmt.Println(Abs(v))
//	fmt.Println(v)
//	Scale(v, 10)
//	fmt.Println(Abs(v))
//	fmt.Println(v)
//}

//package main
//
//import "fmt"
//
//type Vertex struct {
//	X, Y float64
//}
//
//func (v *Vertex) Scale(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
//func ScaleFunc(v *Vertex, f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
//func main() {
//	v := Vertex{3, 4}
//	fmt.Println(v)
//	v.Scale(2)
//	ScaleFunc(&v, 10)
//	fmt.Println(v)
//
//	p := &Vertex{4, 3}
//	fmt.Println(p)
//	p.Scale(3)
//	ScaleFunc(p, 8)
//	fmt.Println(p)
//
//	fmt.Println(v, p)
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//type Vertex struct {
//	X, Y float64
//}
//
//func (v Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func AbsFunc(v Vertex) float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func main() {
//	v := Vertex{3, 4}
//	fmt.Println(v.Abs())
//	fmt.Println(AbsFunc(v))
//
//	p := &Vertex{4, 3}
//	fmt.Println(p.Abs())
//	fmt.Println(AbsFunc(*p))
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//type Vertex struct {
//	X, Y float64
//}
//
//func (v *Vertex) Scale(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
//func (v Vertex) Scales(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
//func (v *Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func main() {
//	v := &Vertex{3, 4}
//	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
//	v.Scale(5)
//	fmt.Printf("After scaling: %v, Abs: %v\n", v, v.Abs())
//	v.Scales(5)
//	fmt.Printf("After scaling: %v, Abs: %v\n", v, v.Abs())
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//type Abser interface {
//	Abs() float64
//	//Abs2() float64
//}
//
//type MyFloat float64
//
//type Vertex struct {
//	X, Y float64
//}
//
//func (f MyFloat) Abs() float64 {
//	fmt.Println("MyFloat")
//	if f < 0 {
//		return float64(-f)
//	}
//	return float64(f)
//}
//
//func (v *Vertex) Abs() float64 {
//	fmt.Println("Vertex")
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func main() {
//	var a Abser
//	//f := MyFloat(-math.Sqrt2)
//	//v := Vertex{3, 4}
//
//	//a = f // a MyFloat 实现了 Abser
//	//fmt.Println(a.Abs())
//	//a = &v // a *Vertex 实现了 Abser
//
//	x := Vertex{3, 4}
//	a = new(Vertex)
//	a.Abs()
//	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
//	// 所以没有实现 Abser。
//	//a = v
//
//	//fmt.Println(a.Abs())
//}

//package main
//
//import (
//	"fmt"
//)
//
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
//
//type IPhone struct {
//}
//
//func (iPhone IPhone) call() {
//	fmt.Println("I am iPhone, I can call you!")
//}
//
//func main() {
//	var phone Phone
//
//	phone = new(NokiaPhone)
//	phone.call()
//
//	phone = new(IPhone)
//	phone.call()
//
//}

//package main
//
//import "fmt"
//
////定义接口
//type Skills interface {
//	Running()
//	Getname() string
//}
//
//type Student struct {
//	Name string
//	Age  int
//}
//
//// 实现接口
//func (p Student) Getname() string { //实现Getname方法
//	fmt.Println(p.Name)
//	return p.Name
//}
//
//func (p Student) Running() { // 实现 Running方法
//	fmt.Printf("%s running", p.Name)
//}
//func main() {
//	var skill Skills
//	var stu1 Student
//	stu1.Name = "wd"
//	stu1.Age = 22
//	skill = stu1
//	skill.Running() //调用接口
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//type Abser interface {
//	Abs() float64
//}
//
//func main() {
//	var a Abser
//	f := MyFloat(-math.Sqrt2)
//	v := Vertex{3, 4}
//
//	a = f  // a MyFloat 实现了 Abser
//	a = &v // a *Vertex 实现了 Abser
//
//	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
//	// 所以没有实现 Abser。
//	// a = v
//
//	fmt.Println(a.Abs())
//}
//
//type MyFloat float64
//
//func (f MyFloat) Abs() float64 {
//	if f < 0 {
//		return float64(-f)
//	}
//	return float64(f)
//}
//
//type Vertex struct {
//	X, Y float64
//}
//
//func (v *Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}

//package main
//
//import "fmt"
//
//type I interface {
//	M()
//}
//
//type T struct {
//	S string
//	N string
//}
//
//type C struct {
//	NAME string
//	AGE  string
//}
//
//// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
//func (t T) M() {
//	fmt.Println("M函数")
//	fmt.Println(t.S, t.N)
//}
//
//func (t C) M() {
//	fmt.Println("K函数")
//	fmt.Println(t.NAME, t.AGE)
//}
//
//func main() {
//	a := C{"zhendong", "25"}
//	var im I = a
//	var i I = T{"hello", "world"}
//	i.M()
//	im.M()
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//type I interface {
//	M()
//}
//
//type Ts struct {
//	S string
//}
//
//func (t *Ts) M() {
//	fmt.Println(t.S)
//}
//
//type F float64
//
//func (f F) M() {
//	fmt.Println(f)
//}
//
//func main() {
//	var i I
//
//	i = &Ts{"Hello"}
//	describe(i)
//	i.M()
//
//	i = F(math.Pi)
//	describe(i)
//	i.M()
//}
//
//func describe(i I) {
//	fmt.Printf("(VAUEL：%v ; TYPE：%T)\n", i, i)
//}

//package main
//
//import "fmt"
//
//type I interface {
//	M()
//}
//
//type T struct {
//	S string
//}
//
//func (m *T) M() {
//	if m == nil {
//		fmt.Println("<nil>")
//		return
//	}
//	fmt.Println(m.S)
//}
//
//func main() {
//	var i I
//
//	var t *T
//	i = t
//	describe(i)
//	i.M()
//
//	i = &T{"hello"}
//	describe(i)
//	i.M()
//}
//
//func describe(i I) {
//	fmt.Printf("(%v, %T)\n", i, i)
//}

//package main
//
//import "fmt"
//
//type I interface {
//	M()
//}
//
//func main() {
//	var i I
//	describe(i)
//	i.M()
//}
//
//func describe(i I) {
//	fmt.Printf("(%v, %T)\n", i, i)
//}

//package main
//
//import "fmt"
//
//func main() {
//	var i interface{}
//	describe(i)
//
//	i = 42
//	describe(i)
//
//	i = "hello"
//	describe(i)
//}
//
//func describe(i interface{}) {
//	fmt.Printf("(%v, %T)\n", i, i)
//}

//package main
//
//import "fmt"
//
//func main() {
//	var i interface{} = "hello"
//
//	fmt.Printf("%v,%T \n", i, i)
//	s := i.(string)
//	fmt.Println(s)
//
//	s, ok := i.(string)
//	fmt.Println(s, ok)
//
//	f, ok := i.(float64)
//	fmt.Println(f, ok)
//
//	f = i.(float64) // 报错(panic)
//	fmt.Println(f)
//}

//package main
//
//import "fmt"
//
//func do(i interface{}) {
//	switch v := i.(type) {
//	case int:
//		fmt.Printf("Twice %v is %v\n", v, v*2)
//	case string:
//		fmt.Printf("%q is %v bytes long\n", v, len(v))
//	default:
//		fmt.Printf("I don't know about type %T!\n", v)
//	}
//}
//
//func main() {
//	do(21)
//	do("hello")
//	do(true)
//}

//package main
//
//import (
//	"fmt"
//	"strconv"
//)
//
//type Person struct {
//	Name string
//	Age  int
//}
//
//func (p Person) String() string {
//	//return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
//	k := p.Name + strconv.Itoa(p.Age)
//	return k
//}
//
//func main() {
//	a := Person{"Arthur Dent", 42}
//	z := Person{"Zaphod Beeblebrox", 9001}
//	fmt.Println(a, z)
//}

//package main
//
//import "fmt"
//
//type IPAddr [4]byte
//
//// TODO: 给 IPAddr 添加一个 "String() string" 方法
//
//func main() {
//	hosts := map[string]IPAddr{
//		"loopback":  {127, 0, 0, 1},
//		"googleDNS": {8, 8, 8, 8},
//	}
//	for name, ip := range hosts {
//		fmt.Printf("%v: %v\n", name, ip)
//	}
//}

//package main
//
//import (
//	"fmt"
//	"strconv"
//	"time"
//)
//
//type MyError struct {
//	When time.Time
//	What string
//}
//
//func (e *MyError) Error() string {
//	return fmt.Sprintf("at %v, %s",
//		e.When, e.What)
//}
//
//func run() error {
//	return &MyError{
//		time.Now(),
//		"it didn't work",
//	}
//}
//
//func main() {
//	if err := run(); err != nil {
//		fmt.Println(err)
//	}
//
//	// second exp
//	// 字符串转换数字
//	i, err := strconv.Atoi("81")
//	if err != nil {
//		fmt.Printf("couldn't convert number: %v\n", err)
//		return
//	}
//	fmt.Println("Converted integer:", i)
//
//	k, err := strconv.Atoi("zhendong")
//	if err != nil {
//		fmt.Printf("couldn't convert number: %v\n", err)
//		return
//	}
//	fmt.Println("Converted integer:", k)
//}
//
//package main
//
//import (
//	"bytes"
//	"fmt"
//	"io"
//	"strings"
//)
//
//func main() {
//	k := "Hello, Readers, zhendong!!"
//	fmt.Println("[]byte(k):", []byte(k))
//
//	r := strings.NewReader("Hello, Readers, zhendong!!")
//	buf := new(bytes.Buffer)
//	buf.ReadFrom(r)
//	fmt.Println("strings.NewReader:", buf.Bytes())
//	fmt.Println(buf)
//	s := strings.NewReader("Hello, Readers, zhendong!!")
//	b := make([]byte, 8)
//	for {
//		n, err := s.Read(b)
//		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
//		fmt.Printf("b[:n] = %q\n\n", b[:n])
//		if err == io.EOF {
//			break
//		}
//	}
//	//a := "Hello, playground"
//	//fmt.Println([]byte(a))
//	//
//	//buf := new(bytes.Buffer)
//	//n, err := buf.ReadFrom(strings.NewReader(a))
//	//fmt.Printf("n = %v err = %v\n", n, err)
//	//fmt.Println(buf.Bytes())
//	//fmt.Println(buf)
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func say(s string) {
//	for i := 0; i < 5; i++ {
//		time.Sleep(100 * time.Millisecond)
//		fmt.Println(s)
//	}
//}
//
//func main() {
//	say("world")
//	say("hello")
//}

//package main
//
//import "fmt"
//
//func sum(s []int, c chan int) {
//	fmt.Println("ZHEN")
//	sum := 0
//	for _, v := range s {
//		sum += v
//	}
//	c <- sum // 将和送入 c
//	//k := <-c
//	fmt.Println("ZHEN")
//}
//
//func main() {
//	s := []int{7, 2, 8, -9, 4, 0}
//
//	c := make(chan int)
//	go sum(s[:len(s)/2], c)
//	go sum(s[len(s)/2:], c)
//	fmt.Println("ready for x, y := <-c, <-c")
//	x, y := <-c, <-c // 从 c 中接收
//	fmt.Println("END for x, y := <-c, <-c")
//
//	fmt.Println(x, y, x+y)
//}

//package main
//
//import "fmt"
//
//func main() {
//	ch := make(chan int, 2)
//	ch <- 1
//	ch <- 2
//	fmt.Println(<-ch)
//	fmt.Println(<-ch)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func fibonacci(n int, c chan int) {
//	x, y := 0, 1
//	for i := 0; i < n; i++ {
//		c <- x
//		x, y = y, x+y
//	}
//	close(c)
//}
//
//func main() {
//	c := make(chan int, 10)
//	go fibonacci(cap(c), c)
//	fmt.Println(<-c)
//	for i := range c {
//		//fmt.Println("---")
//		fmt.Println(i)
//	}
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func fibonacci(c, quit, down chan int) {
//	x, y := 0, 1
//	for {
//		select {
//		case c <- x:
//			x, y = y, x+y
//		case <-quit:
//			fmt.Println("quit")
//			//return
//			//default:
//			//	fmt.Println("my default")
//		case <-down:
//			fmt.Println("down")
//			return
//		}
//	}
//}
//
//func main() {
//	c := make(chan int)
//	quit := make(chan int)
//	down := make(chan int)
//	go func() {
//		time.Sleep(2020 * time.Millisecond)
//		for i := 0; i < 10; i++ {
//			fmt.Println(<-c)
//		}
//	}()
//	go func() {
//		quit <- 0
//	}()
//	go func() {
//		time.Sleep(2050 * time.Millisecond)
//		down <- 1
//	}()
//	fibonacci(c, quit, down)
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	tick := time.Tick(100 * time.Millisecond)
//	boom := time.After(500 * time.Millisecond)
//	//tick := make(chan int)
//	//boom := make(chan int)
//	for {
//		select {
//		case <-tick:
//			fmt.Println("tick.")
//		case <-boom:
//			fmt.Println("BOOM!")
//			return
//		default:
//			fmt.Println("    .")
//			time.Sleep(50 * time.Millisecond)
//		}
//	}
//}

//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//// SafeCounter 的并发使用是安全的。
//type SafeCounter struct {
//	v   map[string]int
//	mux sync.Mutex
//}
//
//// Inc 增加给定 key 的计数器的值。
//func (c *SafeCounter) Inc(key string) {
//	//c.mux.Lock()
//	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
//	c.v[key]++
//	fmt.Println(c.v[key])
//	//c.mux.Unlock()
//}
//
//// Value 返回给定 key 的计数器的当前值。
//func (c *SafeCounter) Value(key string) int {
//	c.mux.Lock()
//	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
//	c.mux.Unlock()
//	//defer c.mux.Unlock()
//	return c.v[key]
//}
//
//func main() {
//	//k := make(map[string]int)
//	//c := SafeCounter{v: k}
//	//计时器
//	bT := time.Now()
//	c := SafeCounter{v: make(map[string]int)}
//	for i := 0; i < 1000; i++ {
//		fmt.Println("WOSHI I ", i)
//		c.Inc("zhendong")
//	}
//	eT := time.Since(bT)
//	fmt.Println("Inc Run time: ", eT)
//
//	//time.Sleep(3 * time.Second)
//	fmt.Println("delay value:", c.Value("zhendong"))
//}

//package main
//
//import "fmt"
//
//func main() {
//	//	fmt.Println("函数:", 1)
//	//	goto End
//	//	fmt.Println("函数:", 2)
//	//End:
//	//	fmt.Println("函数:", 3)
//
//	//for i := 0; i < 10; i++ {
//	//	//OuterLoop:
//	//	for j := 0; j < 10; j++ {
//	//		fmt.Printf("i=%v, j=%v\n", i, j)
//	//		//break OuterLoop
//	//	}
//	//}
//
//SwitchStatement:
//	switch 1 {
//	case 1:
//		fmt.Println(1)
//		for i := 0; i < 10; i++ {
//			break SwitchStatement
//		}
//		fmt.Println(2)
//	}
//	fmt.Println(3)
//}

//package main
//
//import "fmt"
//
//func trace(s string) string {
//	fmt.Println("entering:", s)
//	return s
//}
//
//func un(s string) {
//	fmt.Println("leaving:", s)
//}
//
//func a() {
//	defer un(trace("a"))
//	fmt.Println("in a")
//}
//
//func b() {
//	defer un(trace("b"))
//	fmt.Println("in b")
//	a()
//}
//
//func main() {
//	b()
//}

//package main
//
//import "fmt"
//
//type Allergen int
//
//const (
//	IgEggs         Allergen = 1 << iota // 1 << 0 which is 00000001
//	IgChocolate                         // 1 << 1 which is 00000010
//	IgNuts                              // 1 << 2 which is 00000100
//	IgStrawberries                      // 1 << 3 which is 00001000
//	IgShellfish                         // 1 << 4 which is 00010000
//)
//
//func main() {
//	fmt.Println(IgEggs)
//	fmt.Println(IgChocolate)
//	fmt.Println(IgNuts)
//	fmt.Println(IgStrawberries)
//}

//package main
//
//import (
//	"fmt"
//	"sort"
//)
//
//type Sequence []int
//
//// Methods required by sort.Interface.
//// sort.Interface 所需的方法。
//func (s Sequence) Len() int {
//	return len(s)
//}
//func (s Sequence) Less(i, j int) bool {
//	return s[i] < s[j]
//}
//func (s Sequence) Swap(i, j int) {
//	s[i], s[j] = s[j], s[i]
//}
//
//// Method for printing - sorts the elements before printing.
//// 用于打印的方法 - 在打印前对元素进行排序。
//func (s Sequence) String() string {
//	sort.Sort(s)
//	return fmt.Sprint([]int(s))
//}
//
//func main() {
//	k := Sequence{2, 4, 4}
//	fmt.Println(k.String())
//}

//package main
//
//import (
//	"fmt"
//	"io"
//	"log"
//	"os"
//)
//
//type Job struct {
//	Command string
//	*log.Logger
//}
//
//var _ = fmt.Printf // For debugging; delete when done. // 用于调试，结束时删除。
//var _ io.Reader    // For debugging; delete when done. // 用于调试，结束时删除。
//cs := make(chan *os.File, 100)
//
//func NewJob(command string, logger *log.Logger) *Job {
//	return &Job{command, logger}
//}
//
//func main() {
//	fd, err := os.Open("test.go")
//	if err != nil {
//		log.Fatal(err)
//	}
//	// TODO: use fd.
//	_ = fd
//
//}

//package main
//
//import (
//	"flag"
//	"html/template"
//	"log"
//	"net/http"
//)
//
//var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18
//
//var templ = template.Must(template.New("qr").Parse(templateStr))
//
//func main() {
//	flag.Parse()
//	http.Handle("/", http.HandlerFunc(QR))
//	err := http.ListenAndServe(*addr, nil)
//	if err != nil {
//		log.Fatal("ListenAndServe:", err)
//	}
//}
//
//func QR(w http.ResponseWriter, req *http.Request) {
//	templ.Execute(w, req.FormValue("s"))
//}
//
//const templateStr = `
//<html>
//<head>
//<title>QR Link Generator</title>
//</head>
//<body>
//{{if .}}
//<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
//<br>
//{{.}}
//<br>
//<br>
//{{end}}
//<form action="/" name=f method="GET"><input maxLength=1024 size=70
//name=s value="" title="Text to QR Encode"><input type=submit
//value="Show QR" name=qr>
//</form>
//</body>
//</html>
//`

package pkg

import "fmt"

// PathError 记录一个错误以及产生该错误的路径和操作。
type PathError struct {
	Op   string // "open"、"unlink" 等等。
	Path string // 相关联的文件。
	Err  error  // 由系统调用返回。
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

func main() {
	//a := PathError{"open", "sssss"}
	fmt.Println("aaaa")
}
