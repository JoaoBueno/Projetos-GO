package main

import (
	"time"

	dataframe "github.com/rocketlaunchr/dataframe-go"

	"cloud.google.com/go/civil"
)

func main() {
	sg := dataframe.NewSeriesGeneric("date", civil.Date{}, nil, civil.Date{2018, time.May, 01}, civil.Date{2018, time.May, 02}, civil.Date{2018, time.May, 03})
	s2 := dataframe.NewSeriesFloat64("sales", nil, 50.3, 23.4, 56.2)

	df := dataframe.NewDataFrame(sg, s2)

}
