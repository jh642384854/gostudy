package main

import "fmt"

const (
	x = iota
	y = iota
	z = iota
	w
)

const (
	a = iota
	b
	c
	d
	e
)
const v = iota

func main() {
	fmt.Printf("x=%d\n y=%d\n z=%d\n w=%d\n v=%d\n",x,y,z,w,v)

	fmt.Printf("a=%d\n b=%d\n c=%d\n d=%d\n e=%d\n",a,b,c,d,e)
}