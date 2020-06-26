package main

import "log"

func main() {
	log.Print("Start ...")
	test()
	log.Print("Done ...")
}
func test() {
	mySlice := make([]int, 8)
	log.Print("-->loop begion...")
	for i := 0; i < 32*1000*1000; i++ {
		mySlice = append(mySlice, i)
	}
	log.Print("-->loop end...")
}
