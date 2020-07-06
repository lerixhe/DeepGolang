package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	log.Print("Start ...")
	test()
	log.Print("Force GC after 10s...")
	time.Sleep(10 * time.Second)
	runtime.GC()
	log.Print("Done ...")
	time.Sleep(3600 * time.Second)
}
func test() {
	mySlice := make([]int, 8)
	log.Print("-->loop begion...")
	for i := 0; i < 32*1000*1000; i++ {
		mySlice = append(mySlice, i)
	}
	log.Print("-->loop end...")
}
