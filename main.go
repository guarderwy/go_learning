// demo project main.go
package main

import (
	"container/list"
	"flag"
	"fmt"
	"math"
	"sort"
	"sync"
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
	var mode = flag.String("a", "", "process mode")
	flag.Parse()
	fmt.Println(*mode)
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
