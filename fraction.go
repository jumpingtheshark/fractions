package main

import (
	"fmt"
	"strconv"
)

type fraction struct {
	num   int
	denom int
}

func (f fraction) show() string {
	a := strconv.Itoa(f.num) + "/" + strconv.Itoa(f.denom)
	return a
}
func (f *fraction) reduceMe() {
	n := f.num
	d := f.denom
	divisor := f.num
	for true {
		//	fmt.Println("trying ", divisor)
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

// useless functions :)
func reduceReference(f *fraction) {
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
