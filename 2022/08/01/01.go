package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {
	//strs := []string{"浙江@省", "浙江1省", "浙江z省", "浙江a省", "浙江省", "中文", "英文", "汉字", "拼音"}
	//strs := []string{"浙江@省", "浙江1省", "浙江z省", "浙江a省", "浙江省", "中文", "英文", "汉字", "拼音", "abc", "ad", "zx"}
	strs := []string{
		"1234",
		"-1234",
		"中文排序",
		"汉字a排序",
		"字符串z排序",
		"标点符号@排序",
		"中文序a列",
		"拼音词z汇",
		"姓+名",
		"abc",
		"ABC",
		"ad",
		"xyz",
		"abc！@#",
		"a！@#bc",
		"，。/；'【】《》？：「」、｜",
		"！@#¥%……&*（）——+｜～",
		"焱犇骉麤顕誩甦",
		"汉字",
		"杭晗",
		"憨憨",
		"喝喝",
		"和和",
		"呵呵",
	}
	sort.SliceStable(strs, func(i, j int) bool {
		return strs[i] < strs[j]
	})
	for _, s := range strs {
		fmt.Println(s)
	}

	fmt.Println("====")

	sort.SliceStable(strs, func(i, j int) bool {
		a, _ := UTF82GBK(strs[i])
		b, _ := UTF82GBK(strs[j])
		return string(a) < string(b)
	})
	for _, s := range strs {
		fmt.Println(s)
	}
}

//UTF82GBK : transform UTF8 rune into GBK byte array
func UTF82GBK(src string) ([]byte, error) {
	GB18030 := simplifiedchinese.All[0]
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), GB18030.NewEncoder()))
}

//GBK2UTF8 : transform  GBK byte array into UTF8 string
func GBK2UTF8(src []byte) (string, error) {
	GB18030 := simplifiedchinese.All[0]
	b, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(src), GB18030.NewDecoder()))
	return string(b), err
}
