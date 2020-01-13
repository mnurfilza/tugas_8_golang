package main

import (
	"belajargolang"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	var messages = make(chan int)

	go belajargolang.Senddata(messages)
	belajargolang.RetreiveData(messages)
}
