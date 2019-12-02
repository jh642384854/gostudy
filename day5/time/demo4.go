package main

import (
	"fmt"
	"github.com/chenhg5/collection"
	"time"
)

func main() {
	var monthMaps = map[string]interface{}{"January":"01","February":"02","March":"03","April":"04","May":"05","June":"06","July":"07","August":"08","September":"09","October":"10","November":"11","December":"12"}
	fmt.Println(monthMaps["April"])
	fmt.Println(time.Now().Day())
	//var keys = collection.Collect(monthMaps).Keys().ToStringArray()
	fmt.Println(collection.Collect(monthMaps).Has("March22"))
	var timeNow = time.Now()
	fmt.Println(fmt.Sprintf("%d%s%d",timeNow.Year(),collection.Collect(monthMaps).Get(timeNow.Month().String()),timeNow.Day()))
}
