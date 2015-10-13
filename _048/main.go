package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := complex128(2)
	for i := 1; i < 11; i++ {
		z = z - (cmplx.Pow(z, 3)-x)/(3*(z*z))
	}
	return z
}

func main() {
	x0 := Cbrt(2)
	x1 := cmplx.Pow(2, 1.0/3)
	fmt.Println(x0)
	fmt.Println(x1)
	fmt.Println(cmplx.Abs(x0 - x1))
}
