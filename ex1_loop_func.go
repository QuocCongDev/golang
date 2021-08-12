package main

import (
	"fmt"
)
const num =0.0000001
func Sqrt(x float64) float64 {
	var z float64
	for z=x;(z*z - x) / (2*z)>=num;z -= (z*z - x)/(2*z){
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}