package main

import (
	"fmt"
	"math"
)

func main() {
	var v1, v2 int64
	count, _ := fmt.Scanln(&v1, &v2)
	for (count == 2) {
		var tmp float64 = float64(v1) - float64(v2)
		var res int64 = int64(math.Abs(tmp))
		fmt.Println(res)
		count, _ = fmt.Scanln(&v1, &v2)
	}
}