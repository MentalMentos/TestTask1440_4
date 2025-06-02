package main

import (
	"fmt"
)

func f(n int) int {
	res := 0
	for n > 0 {
		res *= 10
		res += n % 10
		n /= 10
	}
	return res
}

func main() {
	mp := make(map[float64]struct{})
	for i := 2; i < 10^30; i++ {
		fn := f(i)
		ffn := f(fn)
		mp[float64(ffn)/float64(i)] = struct{}{}
	}
	fmt.Println("Ответ:", len(mp))
}
