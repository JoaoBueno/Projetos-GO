package main

import (
	"fmt"
	"time"

	"github.com/snabb/isoweek"
)

func main() {
	data := "10/Jun/2019 - 08:30:45"
	fmt.Println(data)
	// t2 := time.Now()
	t2, err := time.Parse("02-01-2006", "30-12-2012")
	fmt.Println(t2.Format(("02/Jan/2006 - 15:04:05")))
	fmt.Println(t2.Format(("02/01/2006")))

	// t1, err := time.Parse("02-01-2006", "07-06-1965")
	t1, err := time.Parse("02-01-2006", "01-02-2022")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t1.Format(("02/01/2006")))
	fmt.Println(t2.Sub(t1))

	fmt.Println("-------------------------------")
	fmt.Println(t2)
	md := time.Date(t2.Year(), t2.Month(), t2.Day()+1, 0, 0, 0, 0, t2.Location())
	fmt.Println(md)
	md = time.Date(t2.Year(), t2.Month(), t2.Day()-1, 0, 0, 0, 0, t2.Location())
	fmt.Println(md)
	mm := time.Date(t2.Year(), t2.Month()+1, t2.Day(), 0, 0, 0, 0, t2.Location())
	fmt.Println(mm)
	yy := time.Date(t2.Year()+1, t2.Month(), t2.Day(), 0, 0, 0, 0, t2.Location())
	fmt.Println(yy)

	// Calling YearDay method
	yrday := t1.YearDay()

	// Prints the day
	// of the year as specified
	fmt.Printf("The day of the year "+
		"in the 't' specified is: %v\n", yrday)

	fmt.Println("-------------------------------")
	tn := time.Now().UTC()
	fmt.Println(tn)
	year, week := tn.ISOWeek()
	fmt.Println(year, week)

	ts := time.Now().UTC().Unix()
	tn = time.Unix(ts, 0)
	fmt.Println(tn)
	year, week = tn.ISOWeek()
	fmt.Println(year, week)

	fmt.Println("-------------------------------")
	fmt.Println(isoweek.DateToJulian(2012, 12, 30))
	fmt.Println(isoweek.JulianToDate(2456292))
	fmt.Println(isoweek.JulianToDate(2456293))
}
