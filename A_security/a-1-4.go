package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 乱数の種を設定
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		// 浮動小数点数（float64）の乱数を生成
		fmt.Println(rand.Float64())
	}
}
