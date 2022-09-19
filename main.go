// demo project main.go
package main

import (
	"bufio"
	"bytes"
	"container/list"
	"demo/model"
	"errors"
	"flag"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"regexp"
	"sort"
	"sync"
	"time"
)

var a int = 100
var b = 200

func main() {
	fmt.Println("Hello World!")
	a, b = b, a
	fmt.Println(a, b)
	fmt.Printf("%d", a)

	x := 100
	fmt.Println(x)

	fmt.Println(float32(a - b))

	/* fmt.Println(math.MaxFloat32)
	   fmt.Println(math.MaxFloat64) */

	complex_test()

	data_type_test()

	pointer_test()

	const_test()

	array_test()

	slice_test()

	range_test()

	map_test()

	sync_map_test()

	list_test()

	for_test()

	lambda_test()

	accumulate_test()

	defer_test()

	since_test()

	struct_test()

	link_list_test()

	io_test()

	interface_test()

	assert_test()

	sort_test()

	interface_more_test()

	error_test()

	package_test()

	regexp_test()

	time_test()

	lock_test()

	chan_test()

	chan_for_test()

	play_ball_test()

	chan_buffer_test()

	// http.Handle("/", http.FileServer(http.Dir(".")))
	// http.ListenAndServe(":8090", nil)
}

// 复数
func complex_test() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
}

// 数据类型转换
func data_type_test() {
	fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)

	var a int32 = 1047483647
	fmt.Printf("int32: 0x%x %d\n", a, a)

	b := int16(a)
	fmt.Printf("int32: 0x%x %d\n", b, b)

	var c float32 = math.Pi
	fmt.Printf("pi: %0.8f\n", c)
	fmt.Println(int(c))
}

//指针
func pointer_test() {
	//变量地址
	var cat int = 1
	var str string = "banana"
	fmt.Printf("%p %p\n", &cat, &str)

	//取值
	var house = "golang pointer"
	ptr := &house
	fmt.Printf("ptr type: %T\n", ptr)
	fmt.Printf("ptr : %p\n", ptr)

	value := *ptr
	fmt.Printf("value type: %T\n", value)
	fmt.Printf("value : %s\n", value)

	//flag 命令行输入
	/* var mode = flag.String("a", "", "process mode")
	flag.Parse()
	fmt.Println(*mode) */
}

//常量
func const_test() {
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d)

	//iota 常量生成器(第一个声明默认为0)
	type weekday int
	const (
		sunday weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}

//数组
func array_test() {
	// var q [3]int = [3]int{1, 2, 3}
	var q [3]int = [...]int{1, 2, 3}
	for _, v := range q {
		fmt.Println("array_test", v)
	}

	//多维数组
	var array [4][2]int = [4][2]int{1: {2, 3}, 3: {1, 2}}
	fmt.Printf("type %T, %d\n", array, array)
}

//切片
func slice_test() {
	a := [3]int{1, 2, 3}
	fmt.Println(a, a[1:2])

	var house [30]int
	for i := 0; i < 30; i++ {
		house[i] = i
	}

	fmt.Println(house[10:15])
	fmt.Println(house[20:])
	fmt.Println(house[:2])
	fmt.Println(house[:])
	fmt.Println(house[0:0]) //重置切片

	//声明切片
	var strList []string
	var numList []int
	var nunListEmpty = []int{} //空切片(已经分配地址)
	fmt.Println(strList == nil, numList == nil, nunListEmpty == nil)

	//make 函数构造切片
	c := make([]int, 2)
	d := make([]int, 2, 10)
	fmt.Println(c, d)

	// append
	numList = append(numList, 1)
	numList = append(numList, 2, 3)
	numList = append(numList, []int{4, 5, 6}...)
	numList = append([]int{-1, -2, -3}, numList...)
	fmt.Println(numList)

	//切片复制 copy (按小的切片元素个数复制)
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8}
	// copy(slice2, slice1)
	// fmt.Println(slice2)
	copy(slice1, slice2)
	fmt.Println(slice1)

	// 切片中删除
	e := []int{1, 2, 3}
	e = e[:2]
	fmt.Println(e)

	e = []int{1, 2, 3}
	e = append(e[:1], e[2:]...)
	fmt.Println(e)
}

// range
func range_test() {
	slice := []int{10, 20, 30, 40}
	for index, value := range slice {
		fmt.Printf("Vaule:%d Vaule-Addr:%X ElemAddr:%X\n", value, &value, &slice[index])
	}
}

//map
func map_test() {
	var mapList map[string]int
	mapList = map[string]int{"one": 1, "two": 2}
	fmt.Println(mapList)

	mapCreate := make(map[string]float64) // map[string]float64{}
	mapCreate["a"] = 1
	mapCreate["b"] = 2.1
	fmt.Println(mapCreate)

	var mapAssigned map[string]float64
	mapAssigned = mapCreate //引用
	mapAssigned["c"] = 3.3
	fmt.Println(mapCreate)

	mp1 := make(map[int][]int)
	mp2 := make(map[int]*[]int)
	fmt.Println(mp1, mp2)

	//遍历map
	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	//排序
	var sceneList []string
	for k, v := range scene {
		fmt.Println(k, v)
		sceneList = append(sceneList, k)
	}
	sort.Strings(sceneList)
	fmt.Println(sceneList)

	//delete 删除键
	delete(scene, "brazil")
	fmt.Println(scene)

}

//并发中使用map 只读是线程安全的，同时读写是线程不安全的
// sync.Map
/* sync.Map 有以下特性：
1、无须初始化，直接声明即可。
2、sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，Store 表示存储，Load 表示获取，Delete 表示删除。
3、使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false。 */

func sync_map_test() {
	var scene sync.Map //直接声明 不能用make创建

	//存
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	//取
	fmt.Println(scene.Load("london"))

	//删除
	scene.Delete("london")

	//遍历
	scene.Range(func(key, value interface{}) bool {
		fmt.Println("interate:", key, value)
		return true //继续
	})

}

//列表
//container/list 是双链表
func list_test() {
	list_a := list.New()
	// var list_b list.List

	list_a.PushFront("aaa")
	list_a.PushBack(99)
	// fmt.Println(list_a)

	element := list_a.PushBack(101)

	list_a.InsertAfter("qer", element)
	list_a.InsertBefore("ccc", element)

	list_a.Remove(element)

	//遍历
	for i := list_a.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}

//for
func for_test() {
	var i int
	for {
		if i > 10 {
			break
		}
		i++
	}
	fmt.Println(i)

	//只有一个循环条件的循环
	for i <= 20 {
		i++
	}
	fmt.Println(i)

	//99乘法表
	for y := 1; y <= 9; y++ {
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%d ", x, y, x*y)
		}
		fmt.Println()
	}
}

// 匿名函数&map
func lambda_test() {
	var skillParam = flag.String("skill", "", "skill to perform")
	flag.Parse()
	// fmt.Println(*skillParam)

	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
	}

	if f, ok := skill[*skillParam]; ok {
		fmt.Println(ok)
		f()
	} else {
		fmt.Println("skill not found")
	}
}

// 累加
func lambda_test_accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}

func accumulate_test(args ...interface{}) {
	accumulate := lambda_test_accumulate(1)
	fmt.Println(accumulate())

	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println("type is int")
		}
	}
}

// defer延迟
func defer_test() {
	fmt.Println("defer-no")
	defer fmt.Println("defer-1")
	defer fmt.Println("defer-2")
}

func fileSize(filename string) int64 {
	f, err := os.Open(filename)

	if err != nil {
		return 0
	}

	defer f.Close()

	info, err := f.Stat()

	if err != nil {
		return 0
	}

	size := info.Size()
	return size
}

// time since
func since_test() {
	start := time.Now()
	sum := 0
	for i := 0; i < 10000000; i++ {
		sum++
	}
	elapsed := time.Since(start)
	// elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

//初始化结构体
func struct_test() {

	//1、键值对初始化结构体
	type People struct {
		name  string
		child *People
	}

	relation := People{
		name: "grand father",
		child: &People{
			name: "father",
			child: &People{
				name: "me",
			},
		},
	}
	fmt.Println(relation)
	fmt.Println(relation.name)
	fmt.Println(relation.child.name)

	//2、使用多个值的列表初始化结构体(无需键)
	type Address struct {
		Province    string
		City        string
		ZipCode     int
		PhoneNumber string
	}

	addr := Address{
		"Fu Jian",
		"Fu Zhou",
		350000,
		"13000000000",
	}
	fmt.Println(addr)
	fmt.Println(addr.ZipCode)

	// 3、初始化匿名结构体
	msg := &struct {
		id   int
		data string
	}{
		1024,
		"hello",
	}
	fmt.Println(msg)
	fmt.Println(msg.id)

	// 4、结构体内嵌
	type innerS struct {
		int1, int2 int
	}

	type outerS struct {
		b int
		c float32
		int
		innerS
	}

	outer := new(outerS)
	outer.int1 = 1
	outer.int2 = 2
	fmt.Println(outer)

	//5、初始化内嵌结构体
	type Wheel struct {
		Size int
	}

	type Car struct {
		Wheel
		Engine struct {
			Power int
			Type  string
		}
	}

	c := Car{
		Wheel: Wheel{
			Size: 18,
		},

		Engine: struct {
			Power int
			Type  string
		}{
			1,
			"engine",
		},
	}

	fmt.Println(c)
}

// 链表、用struct定义链表
func link_list_test() {
	//1、单链表
	type Node struct {
		data int
		next *Node
	}

	head := new(Node)
	head.data = 1

	node1 := new(Node)
	node1.data = 2
	head.next = node1

	node2 := new(Node)
	node2.data = 3
	node1.next = node2

	for head != nil {
		fmt.Println(*head)
		head = head.next
	}

	//2、单链表插入
	var head2 = new(Node)
	head2.data = 0
	var tail *Node
	tail = head2
	for i := 1; i < 10; i++ {
		var node = Node{data: i}
		// node.next = tail //头插法
		(*tail).next = &node //尾插法
		tail = &node
	}
	for tail != nil {
		fmt.Println(*tail)
		tail = tail.next
	}

}

// io测试
func io_test() {
	data := []byte("golang io test")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [128]byte
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)
}

//interface
//1、接口的方法与实现接口的类型方法格式一致
//2、接口中所有方法均被实现
type DataWriter interface {
	WriteData(data interface{}) error
	CanWrite() bool
}

type file struct {
}

func (d *file) WriteData(data interface{}) error {
	fmt.Println("WriteData:", data)
	return nil
}

func (d *file) CanWrite() bool {
	return false
}

func interface_test() {
	f := new(file)
	var writer DataWriter
	writer = f
	writer.WriteData("test data data1")
}

//assert断言
func assert_test() {
	var x interface{}
	x = 10
	value, ok := x.(int)
	fmt.Println(value, ok)
}

//sort
type MyStringList []string

func (m MyStringList) Len() int {
	return len(m)
}

func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func sort_test() {
	names := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}

	sort.Sort(names)

	for _, v := range names {
		fmt.Printf("%s\n", v)
	}

	//或者去实现方法
	names1 := sort.StringSlice{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sort.Sort(names1)
	for _, v := range names1 {
		fmt.Printf("%s\n", v)
	}

	//或者用已封装方法
	names2 := []string{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sort.Strings(names2)
	for _, v := range names1 {
		fmt.Printf("%s\n", v)
	}

}

//接口嵌套
type device struct{}

func (d *device) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (d *device) Close() error {
	return nil
}

func interface_more_test() {
	var wc io.WriteCloser = new(device)
	wc.Write(nil)
	wc.Close()

	var writeOnly io.Writer = new(device)
	writeOnly.Write(nil)

	/* var aaa interface{} = "1"
	b, isInt := aaa.(int)
	fmt.Println(b, isInt)
	errors.New()
	*/
}

//error
func error_msg() (int, error) {
	return 1, errors.New("no known error......")
}

type dualError struct {
	Num int
	msg string
}

func (d dualError) Error() string {
	return fmt.Sprintf("Wrong, because \"%d\"", d.Num)
}

func error_msg1() (int, error) {
	return 1, dualError{Num: 100}
}

func error_test() {
	// 1、errors.New
	result, err := error_msg()
	fmt.Println(result, err)

	// 2、自定义报错
	ret, err1 := error_msg1()
	fmt.Println(ret, err1)
}

//package person
func package_test() {
	person := model.NewPerson("guarder")
	person.SetAge(30)
	person.SetSal(100000000.00)
	fmt.Println(person, person.Name, person.GetAge(), person.GetSal())

	animal := model.NewAnimal("dog")
	fmt.Println(animal.Name)
	animal.SetAge(2)
	fmt.Println(animal.GetAge())
}

//regexp
func regexp_test() {
	buf := "abc azc a7c aac 888 a9c  tac"
	reg1 := regexp.MustCompile(`a.c`)
	result1 := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println(result1)
}

//time
func time_test() {
	now := time.Now()
	fmt.Printf("current time:%v\n", now)
	fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(now.Month(), int(now.Month()))

	//时间戳
	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano() //纳秒
	fmt.Println(timestamp1, timestamp2)

	//时间戳格式化
	timeObj := time.Unix(timestamp1, 0)
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", timeObj.Year(), timeObj.Month(), timeObj.Day(), timeObj.Hour(), timeObj.Minute(), timeObj.Second())

	//周几
	fmt.Println(now.Weekday())

	//add函数
	one_hour_later := now.Add(time.Hour) //1小时后
	fmt.Println(one_hour_later)

	//定时器
	/* ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i)
	} */

	// 格式化
	//需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S 而是使用Go语言的诞生时间 2006 年 1 月 2 号 15 点 04 分 05 秒
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	//字符串转time
	var layout string = "2006-01-02 15:04:05"
	var timeStr string = "2022-10-01 10:01:00"
	timeObj1, _ := time.Parse(layout, timeStr)
	timeObj2, _ := time.ParseInLocation(layout, timeStr, time.Local)
	fmt.Println(timeObj1, timeObj1.Unix(), timeObj2, timeObj2.Unix())

	/* var input string
	fmt.Scanln(&input) */
}

//goroutine
func lock_test() {
	model.Lock_test()
}

//chan
func chan_test() {
	ch := make(chan int)
	go func() {
		fmt.Println("start goroutine")
		ch <- 0
		fmt.Println("exit goroutine")
	}()
	fmt.Println("wait goroutine")
	<-ch
	fmt.Println("all done")
}

//chan foreach
func chan_for_test() {
	ch := make(chan int)
	go func() {
		for i := 3; i >= 0; i-- {
			ch <- i
			// time.Sleep(time.Second)
		}
	}()

	for data := range ch {
		fmt.Println(data)
		if data == 0 {
			break
		}
	}

	close(ch) //关闭通道

	data1, ok := <-ch
	fmt.Println(data1, ok) //可判断通道是否成功关闭，false为关闭了
}

//模拟网球
var wg sync.WaitGroup

func play_ball_test() {
	court := make(chan int)
	wg.Add(2)

	go player("Mike", court)
	go player("guarder", court)

	//发球
	court <- 1

	wg.Wait() //等待结束
}

func player(name string, court chan int) {
	defer wg.Done() // 在函数退出时调用Done 来通知main 函数工作已经完成

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		n := rand.Intn(100) //选随机数判断是否丢球
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		court <- ball
	}
}

//缓冲通道 make(chan int, 3) ,缓冲元素3个
func chan_buffer_test() {
	ch := make(chan int, 3)
	fmt.Println(len(ch))

	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(len(ch))
}
