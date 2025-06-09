package main

import "fmt"

type Fraction struct {
	Num   int
	Denom int
}

func main() {
	a := Fraction{Num: 6, Denom: 8}
	b := Fraction{Num: 6, Denom: 8}
	//b := Fraction{Num=3, Denom=4}
	fmt.Println(a)
	// fmt.Println("")
	a = reduceVal(a)

	print("a: ", a.Num, " ", a.Denom)
	println("")
	reduceRef(&b)
	print("b: ", b.Num, " ", b.Denom)
}

func reduceRef(f *Fraction) {
	n := f.Num
	d := f.Denom
	divisor := f.Num
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
	f.Num = n
	f.Denom = d
	println(f.Num, " ", f.Denom)

}

//pass by value, turn in into pass by reference
func reduceVal(f Fraction) Fraction {
	// plan of action - figure out which one is smaller, and then test if a mod b == 0 and then just keep going down
	// until the divisor is 1, and then say that you are done.

	n := f.Num
	d := f.Denom
	divisor := f.Num
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
	f.Num = n
	f.Denom = d
	println(f.Num, " ", f.Denom)
	return f

}

// made a change
