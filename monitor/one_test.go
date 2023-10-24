package monitor

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	num := 23102402
	binStr := fmt.Sprintf("%b", num)
	fmt.Println(binStr)

}
