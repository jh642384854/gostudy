package storage

import "log"

func SaveItem() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemcount := 0
		for {
			item := <-out
			log.Printf("Got item,item index :#%d，value:%s \n", itemcount, item)
			itemcount ++
		}
	}()
	return out
}
