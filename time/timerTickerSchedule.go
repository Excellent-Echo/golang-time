package time

import (
	"fmt"
	"time"
)

func timer() {
	// ticker (scheduler)
	done := make(chan bool)
	ticker := time.NewTicker(500 * time.Millisecond)

	go func() {
		time.Sleep(20 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case t := <-ticker.C:
			fmt.Println("hello ", t)
		}
	}
	// AfterFunc
	// var ch = make(chan bool)

	// var timer = time.AfterFunc(2*time.Second, func() {
	// 	fmt.Println("process")
	// 	fmt.Println("process 2")
	// 	// ch <- true
	// })
	// <-ch
	// <-timer.C

	// fmt.Println("start")

	// <-time.After(2 * time.Second)
	// fmt.Println("finish")

	// scheduler
	// for true {
	// 	fmt.Println("hello !!")
	// 	time.Sleep(1 * time.Second)
	// }

	// NewTimer()
	// var timer = time.NewTimer(2 * time.Second)

	// fmt.Println("start")
	// fmt.Println("start2")
	// <-timer.C
	// fmt.Println("finish")
}
