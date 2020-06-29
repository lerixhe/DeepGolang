package main

import (
	"log"
	"runtime"
	"time"
)

func readMemStatus() {
	ms := runtime.MemStats{}
	runtime.ReadMemStats(&ms)
	log.Printf("alloc:%d,heapidle:%d,heapreleased:%d", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
	// 分配的内存字节数、堆内存中空闲的字节数（包括申请未分配的堆内存、未回收的堆内存）、返还给OS的堆内存
}
func main() {
	readMemStatus()
	log.Print("Start ...")
	test()
	readMemStatus()
	log.Print("Force GC after 10s...")
	time.Sleep(10 * time.Second)
	readMemStatus()
	runtime.GC()
	readMemStatus()
	log.Print("Done ...")
	go func() {
		for {
			readMemStatus()
			time.Sleep(10 * time.Second)
		}
	}()
	select {}
}
func test() {
	mySlice := make([]int, 8)
	log.Print("-->loop begion...")
	for i := 0; i < 32*1000*1000; i++ {
		mySlice = append(mySlice, i)
		if i == 16*1000*1000 {
			readMemStatus()
		}
	}
	log.Print("-->loop end...")
}
