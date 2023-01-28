package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/alexmullins/zip"
	"github.com/pkg/errors"
)

var historicalDataDir = "/tmp/crypto-historical"

// https://github.com/binance/binance-public-data/
// https://www.binance.com/en/support/faq/how-to-download-historical-market-data-on-binance-5810ae42176b4770b880ce1f14932262
func main() {
	if _, err := os.Stat(historicalDataDir); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(historicalDataDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

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
	startYear := 2023
	startMonth := 1
	startDay := 27
	dailyETHBaseURL := "https://data.binance.vision/data/spot/daily/klines/ETHUSDT/1d"
	for i := startYear; i <= 2023; i++ {
		for j := startMonth; j <= 12; j++ {
			for k := startDay; k <= daysEachMonth[j]; k++ {
				dailyETH := fmt.Sprintf("ETHUSDT-1d-%02d-%02d-%02d.zip", i, j, k)
				url := fmt.Sprintf("%s/%s", dailyETHBaseURL, dailyETH)
				path := downloadZip(url)
				unzip(path)
			}
		}
	}
}

// download zip file and write to /tmp/crypto-historical
func downloadZip(url string) string {
	c := http.Client{}
	addFileProtocol(&c)
	resp, err := c.Get(url)
	if err != nil {
		log.Fatal(errors.WithStack(err))
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		log.Fatalf("download zip file failed, url=%s\n", url)
	}

	tmp := strings.Split(url, "/")
	fileName := tmp[len(tmp)-1]
	path := fmt.Sprintf("%s/%s", historicalDataDir, fileName)
	if _, err := os.Stat(path); errors.Is(err, os.ErrExist) {
		return path
	}
	zipFile, err := os.Create(path)
	if err != nil {
		log.Fatal(errors.WithStack(err))
	}
	defer func(zipFile *os.File) {
		err := zipFile.Close()
		if err != nil {

		}
	}(zipFile)

	_, err = io.Copy(zipFile, resp.Body)
	if err != nil {
		log.Fatal(errors.WithStack(err))
	}

	return path
}

func addFileProtocol(c *http.Client) {
	tr := &http.Transport{}
	tr.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
	c.Transport = tr
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

	// buffered io
	scanner := bufio.NewScanner(rc)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
