package main

import "fmt"

type Matrix struct {
	name string
	m    int
	n    int
}

type Aux struct {
	cost int
	name string
	row  int
	col  int
}

type P struct {
	row int
	col int
}

const BitsPerWord = 32 << (^uint(0) >> 63) // either 32 or 64
const MaxInt int = 1<<(BitsPerWord-1) - 1

func matrixMult(m []Matrix) Aux {
	n := len(m)
	aux := make(map[P]Aux)
	for i := 0; i < len(m); i++ {
		aux[P{i, i}] = Aux{0, m[i].name, m[i].m, m[i].n}
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n-i; j++ {
			best := MaxInt

			for k := j; k < j+i; k++ {
				a := aux[P{j, k}]
				lcost, lname, lrow, lcol := a.cost, a.name, a.row, a.col
				b := aux[P{k + 1, j + i}]
				rcost, rname, rcol := b.cost, b.name, b.col
				cost := lcost + rcost + lrow*lcol*rcol
				name := fmt.Sprintf("(%s%s)", lname, rname)
				if cost < best {
					best = cost
					aux[P{j, j + i}] = Aux{cost, name, lrow, rcol}
				}
			}
		}
	}
	return aux[P{0, n - 1}]
}

func main() {
	m := []Matrix{Matrix{"A", 10, 20}, Matrix{"B", 20, 30}, Matrix{"C", 30, 40}}
	res := matrixMult(m)
	fmt.Println("cost:", res.cost, "order:", res.name, "rows:", res.row, "cols:", res.col)
}
