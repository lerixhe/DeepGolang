package main

import (
	"bytes"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

// 不断生成并打印字符串

func printBytes() {
	for i := 0; i < 1000; i++ {
		log.Println(genBytes)
	}
}
func genBytes() *bytes.Buffer {
	// 生成一块buffer,用来写入字符
	var buff bytes.Buffer
	for i := 0; i < 2000; i++ {
		buff.Write([]byte{'0' + byte(rand.Intn(10))})
	}
	return &buff
}
func main() {
	go func() {
		http.ListenAndServe(":6060", nil)
	}()
	for {
		printBytes()
		time.Sleep(10*time.Microsecond)
	}

}
