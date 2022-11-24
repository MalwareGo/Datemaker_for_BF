package main

import (
	"fmt"
)

func main() {
	k := "01"
	l := "1952"
	for i := 0; i <= 31; i++ {
		if i == 1 {
			for j := 0; j <= 30; j++ {
				fmt.Printf("%v%v%v\n", i+j, k, l)
			}
		}
	}
	c := "02"
	for a := 0; a <= 31; a++ {
		if a == 1 {
			for m := 0; m <= 27; m++ {
				fmt.Printf("%v%v%v\n", a+m, c, l)
			}
		}
	}

	d := "03"
	for e := 0; e <= 31; e++ {
		if e == 1 {
			for f := 0; f <= 30; f++ {
				fmt.Printf("%v%v%v\n", e+f, d, l)
			}
		}
	}

	g := "04"
	for h := 0; h <= 31; h++ {
		if h == 1 {
			for n := 0; n <= 29; n++ {
				fmt.Printf("%v%v%v\n", h+n, g, l)
			}
		}
	}

	o := "05"
	for p := 0; p <= 31; p++ {
		if p == 1 {
			for q := 0; q <= 30; q++ {
				fmt.Printf("%v%v%v\n", p+q, o, l)
			}
		}
	}

	r := "06"
	for s := 0; s <= 31; s++ {
		if s == 1 {
			for t := 0; t <= 29; t++ {
				fmt.Printf("%v%v%v\n", s+t, r, l)
			}
		}
	}

	u := "07"
	for v := 0; v <= 31; v++ {
		if v == 1 {
			for w := 0; w <= 30; w++ {
				fmt.Printf("%v%v%v\n", v+w, u, l)
			}
		}
	}

	x := "08"
	for y := 0; y <= 31; y++ {
		if y == 1 {
			for z := 0; z <= 30; z++ {
				fmt.Printf("%v%v%v\n", y+z, x, l)
			}
		}
	}

	c_ := "09"
	for a_ := 0; a_ <= 31; a_++ {
		if a_ == 1 {
			for m_ := 0; m_ <= 29; m_++ {
				fmt.Printf("%v%v%v\n", a_+m_, c_, l)
			}
		}
	}

	d_ := "10"
	for e_ := 0; e_ <= 31; e_++ {
		if e_ == 1 {
			for f_ := 0; f_ <= 30; f_++ {
				fmt.Printf("%v%v%v\n", e_+f_, d_, l)
			}
		}
	}

	g_ := "11"
	for h_ := 0; h_ <= 31; h_++ {
		if h_ == 1 {
			for n_ := 0; n_ <= 29; n_++ {
				fmt.Printf("%v%v%v\n", h_+n_, g_, l)
			}
		}
	}

	o_ := "12"
	for p_ := 0; p_ <= 31; p_++ {
		if p_ == 1 {
			for q_ := 0; q_ <= 30; q_++ {
				fmt.Printf("%v%v%v\n", p_+q_, o_, l)
			}
		}
	}
}
