package main

func Gcd(a *fraction, b *fraction) {
	print("this is where we'll calc gcd, though I am not exactly sure how, and in my my mind Im thinking in c#")
	// coment
	aa := a
	//bb:=b
	aa.denom = a.denom * b.denom
	aa.num = a.num * b.denom

}

func Add(a fraction, b fraction) fraction {
	c := fraction{-1, a.denom} //temporary -1 for now
	c.num = a.num + b.num
	return c
}
