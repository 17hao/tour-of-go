package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

func parseNums(nums string) ([]decimal.Decimal, error) {
	strings.TrimSpace(nums)
	strs := strings.Split(nums, " ")
	res := make([]decimal.Decimal, 0)
	for _, str := range strs {
		if str == "" {
			continue
		}
		d, err := decimal.NewFromString(str)
		if err != nil {
			return nil, err
		}
		res = append(res, d)
	}
	return res, nil
}

func parseTarget(target string) (decimal.Decimal, error) {
	strings.TrimSpace(target)
	d, err := decimal.NewFromString(target)
	if err != nil {
		return decimal.NewFromFloat(0), err
	}
	return d, nil
}

func parseCount(count string) (int, error) {
	strings.TrimSpace(count)
	i, err := strconv.ParseInt(count, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

var combinationLists [][]decimal.Decimal

func combinations(nums []decimal.Decimal, k int, list []decimal.Decimal, idx int) {
	if len(list) == k {
		newList := make([]decimal.Decimal, len(list))
		copy(newList, list)
		combinationLists = append(combinationLists, newList)
		return
	}
	for i := idx; i < len(nums); i++ {
		list = append(list, nums[i])
		combinations(nums, k, list, i+1)
		list = list[:len(list)-1]
	}
}

func f1(nums []decimal.Decimal, k int, target decimal.Decimal) []decimal.Decimal {
	combinations(nums, k, []decimal.Decimal{}, 0)
	for _, list := range combinationLists {
		sum := decimal.NewFromFloat(0)
		for _, n := range list {
			sum = sum.Add(n)
		}
		if sum.Equal(target) {
			return list
		}
	}
	return nil
}

// GOOS=windows GOARCH=amd64 go build -o lyj.exe
func main() {
	fmt.Print("候选数字: ")
	scanner := bufio.NewScanner(os.Stdin)
	var (
		numsStr   string
		countStr  string
		targetStr string
	)

	if scanner.Scan() {
		numsStr = scanner.Text()
	}
	nums, err := parseNums(numsStr)
	if err != nil {
		panic(err)
	}
	//fmt.Println(nums)

	fmt.Print("挑选个数: ")
	if scanner.Scan() {
		countStr = scanner.Text()
	}
	count, err := parseCount(countStr)
	if err != nil {
		panic(err)
	}
	//fmt.Println(count)

	fmt.Print("目标值: ")
	if scanner.Scan() {
		targetStr = scanner.Text()
	}
	target, err := parseTarget(targetStr)
	if err != nil {
		panic(err)
	}
	//fmt.Println(target)

	res := f1(nums, count, target)
	fmt.Print("答案: ")
	if len(res) == 0 {
		fmt.Println("无解")
		return
	}
	for _, n := range res {
		fmt.Printf("%s ", n.String())
	}
	fmt.Println("\n按任意键退出")
	if scanner.Scan() {
	}
}
