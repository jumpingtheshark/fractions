package main

import "fmt"

func main() {

	a := fraction{num: 10, denom: 8}
	b := fraction{num: 6, denom: 8}
	b = b
	// fmt.Println("")
	//a = reduceVal(a)
	a.reduceMe()

	print("a: ", a.num, " ", a.denom)
	println("")

	//reduceRef(&b)
	//	print("b: ", b.Num, " ", b.Denom)
	//Gcd(&a, &b)

	c := Add(a, b)
	fmt.Println(a.show(), "+", b.show(), "=", c.show())

}
