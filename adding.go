package main

func Gcd(a *fraction,
	b *fraction) {
	print("this is where we'll calc gcd, though I am not exactly sure how, and in my my mind Im thinking in c#")
	// coment
	aa := a.copy()
	//bb:=b
	aa.denom = a.denom * b.denom
	aa.num = a.num * b.denom

	bb := b.copy()
	bb.denom = b.denom * a.denom
	bb.num = b.num * a.denom

	a.num = aa.num
	a.denom = aa.denom

	b.num = bb.num
	b.denom = bb.denom

}

func Add(a fraction, b fraction) fraction {
	//first of all, are denominators the same, if not then we gotta make them the same
	// first reduce and then find gcd
	if a.denom != b.denom {
		a.reduceMe()
		b.reduceMe()

		Gcd(&a, &b)

	}

	c := fraction{-1, a.denom} //temporary -1 for now
	c.num = a.num + b.num
	return c
}
