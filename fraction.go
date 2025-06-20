package main

import "fmt"

type fraction struct {
	num   int
	denom int
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
