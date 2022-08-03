package main

import (
	"fmt"
)

func main() {
	var stockCode = 123
	var endDate = "2022-08-03"
	var url = "Code=%d&endDate=%s"
	// fmt.Sprintf格式化字符串
	var targetUrl = fmt.Sprintf(url, stockCode, endDate)
	fmt.Println(targetUrl)
}