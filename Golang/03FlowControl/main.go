package main

import (
	"fmt"
	"strconv"
)

func main() {

	if num, err := strconv.Atoi("100"); err == nil {
		fmt.Println("转换成功", num)
	} else {
		fmt.Println("转换失败", err)
	}
}
