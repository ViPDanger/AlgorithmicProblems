package tinkoff

import fmt "github.com/ViPDanger/AlgoritmicProblems/algoritms/output"

func Test5() {
	var l, r, x uint64
	fmt.Scan(&l, &r)
	var count uint
	for i := 0; i < 10; i++ {
		x = 0
		for j := 0; j < 18; j++ {
			x = x*10 + uint64(i)
			if l <= x && x <= r {
				count++
			}
		}
	}
	fmt.Println(count)
}
