package main

import (
	"fmt"
	"github.com/uniplaces/carbon"
)

func main() {
	today, _ := carbon.NowInLocation("Asia/Shanghai")
	today.SubDay()
	fmt.Printf("Right now in Vancouver is %s\n", today)
}
