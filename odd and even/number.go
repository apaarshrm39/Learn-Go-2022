package main

import "fmt"

type numbers []int

func newNumbers(howMany int) numbers {
	list := []int{}
	for i := 0; i <= howMany; i++ {
		list = append(list, i)
	}

	return list
}

func evenOrOdd(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

func (n numbers) printEvenAndOdd() {
	for _, i := range n {
		if evenOrOdd(i) {
			fmt.Println(i, " is Even")
		} else {
			fmt.Println(i, " is Odd")
		}
	}
}
