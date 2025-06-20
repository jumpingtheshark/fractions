package main

import "fmt"

func main() {
	a := fraction{num: 10, denom: 8}
	b := fraction{num: 6, denom: 8}
	b = fraction{num: 3, denom: 4}
	fmt.Println(a)
	// fmt.Println("")
	//a = reduceVal(a)
	a.reduceMe()

	print("a: ", a.num, " ", a.denom)
	println("")
	//reduceRef(&b)
	//	print("b: ", b.Num, " ", b.Denom)
	Gcd(&a, &b)
}

func reduceRef(f *fraction) {
	n := f.num
	d := f.denom
	divisor := f.num
	for true {
		//fmt.Println("trying ", divisor)
		if n%divisor == 0 && d%divisor == 0 {
			n = n / divisor
			d = d / divisor
			fmt.Println(divisor, " bingo")
			break
		} else {
			divisor--

		}
	}
	f.num = n
	f.denom = d
	println(f.num, " ", f.denom)

}

// pass by value, turn in into pass by reference
func reduceVal(f fraction) fraction {
	// plan of action - figure out which one is smaller, and then test if a mod b == 0 and then just keep going down
	// until the divisor is 1, and then say that you are done.

	n := f.num
	d := f.denom
	divisor := f.num
	for true {
		fmt.Println("trying ", divisor)
		if n%divisor == 0 && d%divisor == 0 {
			n = n / divisor
			d = d / divisor
			fmt.Println(divisor, " bingo")
			break
		} else {
			divisor--

		}
	}
	f.num = n
	f.denom = d
	println(f.num, " ", f.denom)
	return f

}

// made a change
