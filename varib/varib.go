package main

import (
	"fmt"
)

type slice []int

func (s slice) Apend(i int) {
	s = append(s, i)
}
func (s *slice) Aspend(i int) {
	*s = append(*s, i)
}

func main() {
	var i slice = make([]int, 2)
	var a = 1
	i.Apend(a)
	fmt.Println(i)
	i.Aspend(a)
	fmt.Println(i)
}
