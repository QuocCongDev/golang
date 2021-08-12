package main

import (
	"fmt"
)
const errorNumber = 0.0000001
func Sqrt(x float64) (float64, error) {
	z := 1.0
	if(x<0){
		return x , ErrNegativeSqrt(x)
	}
	for z = x ; (z*z-x)/(2*z) >= errorNumber; z -= (z*z - x) / (2 * z) {
		
	}
	return z,nil
}
type ErrNegativeSqrt float64
func(e ErrNegativeSqrt) Error() string{	
		return fmt.Sprintf("cannot Sqrt negative number %f:",e)	
}
func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}