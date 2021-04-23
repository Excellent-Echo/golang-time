package time

import (
	"fmt"
	"time"
)

func timeDuration() {
	// time.Duration
	// bechmarking, mengecek durasi eksekusi , dll

	var start1 = time.Date(2021, 04, 23, 15, 30, 0, 0, time.Local)

	// time.Sleep(2 * time.Second)

	var start2 = time.Now()
	// ----

	var duration = start2.Sub(start1)

	fmt.Println(duration.Minutes())

	var result uint64
	var i uint64

	var start = time.Now() // 15.56.00

	// kecepatan = optimasi code, clean code, tidak redundan, clean arsitektur
	// benchmarking semakin cepat semakin baik

	for i = 0; i < 99_000_000; i++ { // anggap eksekusi 2 detik
		result += i
	}

	// time.Sleep(1 * time.Second)

	fmt.Println(result)

	// ----

	var finish = time.Since(start) // 15.56.02 // memakan waktu 2 detik

	fmt.Println(finish)
}
