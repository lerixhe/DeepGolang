package main

import (
	"log"
	"time"
)

func main() {
	log.Print("Start ...")
	test()
	log.Print("Done ...")
	time.Sleep(3600 * time.Second)
}
func test() {
	mySlice := make([]int, 8)
	log.Print("-->loop begion...")
	for i := 0; i < 32*1000*100; i++ {
		mySlice = append(mySlice, i)
	}
	log.Print("-->loop end...")
}
