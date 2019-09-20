package main

import (
	"fmt"
	"io"
	"strings"
)

func demoseek()  {
	reader := strings.NewReader("GO语言中文网")
	reader.Seek(-6,io.SeekEnd)
	r,_,_ := reader.ReadRune()
	fmt.Printf("%c\n",r)
}



func main() {
	demoseek()

}