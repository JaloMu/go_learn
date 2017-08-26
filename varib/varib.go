package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type slice []int
type sli []*int
type HEAD map[string][]string

// 定义一个通用的类型切片
type sl []interface{}

func (s slice) Apend(i int) {
	s = append(s, i)
}
func (s *slice) Aspend(i int) {
	*s = append(*s, i)
}
func (s *sli) Aspend(i *int) {
	*s = append(*s, i)
}

func (s *sl) Apend(i interface{}) {
	*s = append(*s, i)
}

func (h HEAD) Add(key, value string) {
	h[key] = append(h[key], value)
}

func main() {
	var w, r int32 = 2, 2
	fmt.Println(w, r, unsafe.Sizeof(w), unsafe.Sizeof(r))
	var i slice
	fmt.Println(i)
	var ko []int
	fmt.Println(len(ko), unsafe.Sizeof(ko))
	fmt.Println("---------------")
	fmt.Println(len(i), unsafe.Sizeof(i))
	var p sli = make([]*int, 3)
	var a int = 1
	fmt.Println("a 占用空间大小：", unsafe.Sizeof(a))
	var g int64 = 3
	fmt.Println("g 占用空间大小：", unsafe.Sizeof(g))
	var c = "12"
	var d sl
	f := []interface{}{}
	var n []int
	m := []int{}
	//////////////
	fmt.Printf("%p---------------i: %v\n", i, i)
	i.Apend(a)
	fmt.Printf("%p---------------i: %v\n", i, i)
	i.Aspend(a)
	fmt.Printf("%p---------------i: %v\n", i, i)
	p.Aspend(&a)
	fmt.Println(p)

	// 输出d的值
	fmt.Println(d)
	fmt.Println(n)
	fmt.Println(m)
	// 变量类型
	v := reflect.TypeOf(d)
	fmt.Println(v)
	v = reflect.TypeOf(f)
	fmt.Println(v)
	v = reflect.TypeOf(n)
	fmt.Println(v)
	v = reflect.TypeOf(m)
	fmt.Println(v)
	// END----------
	d.Apend(a)
	fmt.Println(d)
	d.Apend(c)
	fmt.Println(d)

	for _, v := range d {
		is_type(v)
		fmt.Println("v 占用空间大小：", unsafe.Sizeof(v))
	}
	var h HEAD = make(map[string][]string, 1)
	fmt.Printf("%p, %v\n", h, h)
	h.Add("1", "A")
	h.Add("2", "A")
	h.Add("3", "A")
	h.Add("3", "A")
	h.Add("3", "A")
	h.Add("3", "A")
	h.Add("3", "A")
	h.Add("3", "A")
	h.Add("3", "A")
	h.Add("3", "A")
	h.Add("3", "A")
	h.Add("3", "A")
	h.Add("3", "A")
	h.Add("3", "A")
	h.Add("3", "A")
	h.Add("3", "A")
	h.Add("3", "A")
	h.Add("3", "A")
	h.Add("3", "A")
	h.Add("3", "A")
	h.Add("3", "A")
	fmt.Printf("%p, %v\n", h, h)
}

func is_type(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case int64:
		fmt.Println("int64")
	default:
		fmt.Println("Unknow")
	}
}
