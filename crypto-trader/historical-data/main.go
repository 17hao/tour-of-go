package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/alexmullins/zip"
	"github.com/pkg/errors"
)

// https://github.com/binance/binance-public-data/
// https://www.binance.com/en/support/faq/how-to-download-historical-market-data-on-binance-5810ae42176b4770b880ce1f14932262
func main() {
	//daysEachMonth := map[int]int{
	//	1:  31,
	//	2:  28,
	//	3:  31,
	//	4:  30,
	//	5:  31,
	//	6:  30,
	//	7:  31,
	//	8:  31,
	//	9:  30,
	//	10: 31,
	//	11: 30,
	//	12: 31,
	//}
	//startYear := 2020
	//startMonth := 1
	//startDay := 1
	//
	//for i := startYear; i <= 2022; i++ {
	//	for j := startMonth; j <= 12; j++ {
	//		for k := startDay; k <= daysEachMonth[j]; k++ {
	//			//fmt.Printf("https://data.binance.vision/data/spot/daily/klines/ETHUSDT/1s/ETHUSDT-1s-2019-01.zip")
	//			fmt.Printf("ETHUSDT-1s-%02d-%02d-%02d.zip\n", i, j, k)
	//		}
	//	}
	//}

	downloadZip("file:///Users/bytedance/Downloads/organizations-10000.zip")
	unzip("/tmp/organizations-10000.zip")
}

func unzip(path string) {
	zipFile, err := zip.OpenReader(path)
	if err != nil {
		log.Fatal(errors.WithStack(err))
		return
	}
	defer func(zipFile *zip.ReadCloser) {
		err := zipFile.Close()
		if err != nil {

		}
	}(zipFile)

	if len(zipFile.File) == 0 {
		return
	}

	f := zipFile.File[0]
	rc, err := f.Open()
	if err != nil {
		log.Fatal(errors.WithStack(err))
		return
	}

	reader := bufio.NewReader(rc)
	for i := 0; i < 10; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(errors.WithStack(err))
		}
		fmt.Print(line)
	}

	bs, err := io.ReadAll(rc)
	if err != nil {
		log.Fatal(errors.WithStack(err))
		return
	}
	fmt.Println(string(bs))
}

func addFileProtocol(c *http.Client) {
	tr := &http.Transport{}
	tr.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
	c.Transport = tr
}

// download zip file and write to /tmp
func downloadZip(url string) {
	c := http.Client{}
	addFileProtocol(&c)
	resp, err := c.Get(url)
	if err != nil {
		log.Fatal(errors.WithStack(err))
		return
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		log.Fatal("download zip file failed.")
		return
	}

	zipFile, err := os.Create("/tmp/organizations-10000.zip")
	if err != nil {
		log.Fatal(errors.WithStack(err))
		return
	}
	defer func(zipFile *os.File) {
		err := zipFile.Close()
		if err != nil {

		}
	}(zipFile)

	_, err = io.Copy(zipFile, resp.Body)
	if err != nil {
		log.Fatal(errors.WithStack(err))
		return
	}
}
