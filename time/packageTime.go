package time

import (
	"fmt"
	"time"
)

type Buy struct {
	ID                int
	Name, Description string
	Date              time.Time
}

func packageTime() {
	// time
	// tipe data = time.Time

	// var time1 time.Time

	// time.Now() = tanggal sekarang / time.Date() sesuai tanggal yg diinginkan

	// time1 = time.Now()

	// time2 := time.Date(2021, 02, 23, 13, 57, 10, 0, time.Local)

	// fmt.Println(time2.Year(), time2.Month(), time2.Day(), time2.UnixNano())
	// fmt.Println(time2)

	// var layoutFormat
	// var value string
	var date time.Time
	var err error

	date, err = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")

	format1 := date.Format(time.RFC850)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(format1)
	//default format = YYYY/MM/DD

	// layoutFormat = "2006-01-02 15:04:05"
	// value = "2015-09-02 08:04:00"
	// date, err = time.Parse(layoutFormat, value)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(value, "\t->", date.String())
	// // 2015-09-02 08:04:00 +0000 UTC

	// layoutFormat = "02/01/2006 MST"
	// value = "02/09/2015 WIB"
	// date, err = time.Parse(layoutFormat, value)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(value, "\t\t->", date.String())
	// 2015-09-02 00:00:00 +0700 WIB

}
