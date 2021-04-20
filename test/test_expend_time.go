package main

import (
	"fmt"
	"time"
)

func sum(n int) {
	startT := time.Now() //计算当前时间

	//total := 0
	//for i := 1; i <= n; i++ {
	//	total += i
	//}

	tc := time.Since(startT) //计算耗时
	fmt.Printf("time cost = %v\n", tc)
	//return total
}

func main() {
	startT := time.Now() //计算当前时间

	//total := 0
	//for i := 1; i <= n; i++ {
	//	total += i
	//}
	time.Sleep(100 * time.Millisecond)
	tc := time.Since(startT) //计算耗时
	tc.Seconds()
	fmt.Printf("time cost = %v\n", tc.Seconds())
}
