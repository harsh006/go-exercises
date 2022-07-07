package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)
func add(wg *sync.WaitGroup,m *sync.Mutex,points *int64){
	defer wg.Done()
	n := rand.Intn(10)
	time.Sleep(time.Duration(n)*time.Millisecond) //simulating random time taken by students :>
	atomic.AddInt64(points, int64(rand.Intn(10)))
}
func main() {
	rand.Seed(time.Now().UnixNano())
	students := 200
	var AvgPoints int64 = 0
	wg := new(sync.WaitGroup)
	m:= new(sync.Mutex)
	for i:=0;i<students;i++{
		wg.Add(1)
		go add(wg,m,&AvgPoints)
	}
	wg.Wait()
	fmt.Println(float64(AvgPoints)/float64(students))
}
