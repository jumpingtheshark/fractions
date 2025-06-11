package main

import "fmt"

type fraction struct {
	Num   int
	Denom int
}

func (f *fraction) reduceMe() {
	n := f.Num
	d := f.Denom
	divisor := f.Num
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
	f.Num = n
	f.Denom = d
	println(f.Num, " ", f.Denom)

}
