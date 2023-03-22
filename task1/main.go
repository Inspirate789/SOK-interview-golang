package main

import "fmt"

func reorder(sl []int) {
	k := 0
	for i := 0; i < len(sl); i++ {
		if sl[i] != 0 {
			sl[k] = sl[i]
			k++
		}
	}

	for i := k; i < len(sl); i++ {
		sl[i] = 0
	}
}

func main() {
	sl1 := []int{1, 2, 3, 4, 5}
	sl2 := []int{1, 0, 3, 0, 5}
	sl3 := []int{0, 0, 0, 0, 0}

	reorder(sl1)
	reorder(sl2)
	reorder(sl3)

	fmt.Printf("%v\n%v\n%v\n", sl1, sl2, sl3)
}
