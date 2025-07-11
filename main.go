package main

import "fmt"

func main() {

	a := fraction{num: 1, denom: 7}
	b := fraction{num: 1, denom: 8}
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

	//now for no good reason let's add these things to a list

	var fractions []fraction
	fractions = append(fractions, a)
	fractions = append(fractions, b)
	fractions = append(fractions, c)
	fmt.Println("fractions: ", fractions)

	for _, v := range fractions {

		fmt.Println(v.show())
	}

}
