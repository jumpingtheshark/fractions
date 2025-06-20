package main

func Gcd(a *fraction, b *fraction) {
	print("this is where we'll calc gcd, though I am not exactly sure how, and in my my mind Im thinking in c#")
	// coment
	aa := a
	//bb:=b
	aa.denom = a.denom * b.denom
	aa.num = a.num * b.denom
	print(aa)

}
