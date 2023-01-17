package main

import "fmt"

// https://github.com/binance/binance-public-data/
// https://www.binance.com/en/support/faq/how-to-download-historical-market-data-on-binance-5810ae42176b4770b880ce1f14932262
func main() {
	daysEachMonth := map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
	startYear := 2020
	startMonth := 1
	startDay := 1

	for i := startYear; i <= 2022; i++ {
		for j := startMonth; j <= 12; j++ {
			for k := startDay; k <= daysEachMonth[j]; k++ {
				//fmt.Printf("https://data.binance.vision/data/spot/daily/klines/ETHUSDT/1s/ETHUSDT-1s-2019-01.zip")
				fmt.Printf("ETHUSDT-1s-%02d-%02d-%02d.zip\n", i, j, k)
			}
		}
	}
}
