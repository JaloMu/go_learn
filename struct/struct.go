package main

import (
	"fmt"
)

type Rect struct {
	x, y          float64
	Width, Height float64
}
//
//func (r Rect) Area(i float64) float64 {
//	r.x = i
//	return r.Width * r.Height
//}

type slice []int

func (s slice)Append(i int) {
	s = append(s, i)
	fmt.Printf("%p-----%v\n",s, s)
}

func main() {
	var s slice = []int{}
	fmt.Printf("%p-----%v\n",s, s)
	var i = 2
	s.Append(i)
	fmt.Printf("%p-----%v\n",s, s)
	m := make(map[int]string)
	fmt.Printf("%p--------%v\n", m,m)
	m2 := m
	fmt.Printf("%p--------%v\n", m2,m2)
	m2[1] = "A"
	fmt.Printf("%p--------%v\n", m,m)
	fmt.Printf("%p--------%v\n", m2,m2)
	r := Rect{1,2,3,4}
	fmt.Printf("%p ---Rect--- %v\n", &r, r)
	r1 := r
	fmt.Printf("%p ---Rect--- %v\n", &r1, r1)
	r1.x = 3
	fmt.Printf("%p ---Rect--- %v\n", &r, r)
	fmt.Printf("%p ---Rect--- %v\n", &r1, r1)
	//var r *Rect
	//s1 := make([]int , 2, 2)
	//s2 := s1
	//fmt.Println(s1, s2)
	//s2[0]= 1
	//fmt.Println(s1, s2)
	//fmt.Printf("%p ---> %v\n", &r, r)
	//r = new(Rect)
	//fmt.Printf("%p ---> %v\n", &r, r)
	//r.x = 2
	//(*r).Width = 3
	//(*r).Height = 4
	//var i float64 = 3
	//fmt.Printf("%p --Rect-> %v\n", &r, r)
	//
	//fmt.Println(r.Area(i))
	//fmt.Printf("%p --Rect-> %v\n", &r, r)
	///*  New and make distinction*/
	//p := new([]int)
	//fmt.Printf("%p  --new([]int)--  %v\n", p, *p)
	//fmt.Println(p)
	//(*p) = make([]int,2)
	//fmt.Printf("%p  --new([]int) make--  %v\n", p, *p)
	//
	//v:= make([]int , 2,3)
	//fmt.Printf("%p  ----  %v\n", v, v)
	//fmt.Println(v)
	//(*p)[0]=18
	//v[1] = 18
	//fmt.Printf("%p  ----  %v\n", v, v)
	//fmt.Printf("%p  ----  %v\n", p, *p)

}
