package main

import "fmt"

type Node struct {
	array     []int
	response bool
}

func main() {

	casos := []Node{
		{[]int{1, 3, 2, 1}, false},
		{[]int{1, 3, 2}, true},
		{[]int{3, 6, 5, 8, 10, 20, 15}, false},
		{[]int{1, 2, 1, 2}, false},
		{[]int{10, 1, 2, 3, 4, 5}, true},
		{[]int{1, 2, 3, 4, 99, 5, 6}, true},
		{[]int{10, 1, 2, 3, 4, 5, 6, 1}, false},
		{[]int{3, 5, 67, 98, 3}, true},
		{[]int{123, -17, -5, 1, 2, 3, 12, 43, 45}, true},
	}

	for _, caso := range casos {

		fmt.Println(fmt.Sprintf("Caso: %v, Esperado: %v, Obtenido: %v", caso.array, caso.response, solution(caso.array)))

	}

}

func solution(sequence []int) bool {
	for i := 1; i < len(sequence); i++ {
		if sequence[i] <= sequence[i-1] {
			withoutPrev := make([]int, len(sequence)-1)
			withoutCurr := make([]int, len(sequence)-1)
			copy(withoutPrev, sequence[:i])
			copy(withoutPrev[i:], sequence[i+1:])
			copy(withoutCurr, sequence[:i-1])
			copy(withoutCurr[i-1:], sequence[i:])
			if isAscending(withoutPrev) || isAscending(withoutCurr) {
				return true
			}
			return false
		}
	}
	return true
}

func isAscending(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] {
			return false
		}
	}
	return true
}
