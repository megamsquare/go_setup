package main_old

import "fmt"

func main() {
	fmt.Println("Start Golang")
	A := []int{2,5,2,7,4}
	fmt.Println(distAbsDiff(A))
}

func distAbsDiff(A []int) int {
	result := make(map[int]bool)
	for i := 1; i < len(A); i++ {
		diff := abs(A[i] - A[i - 1])
		result[diff] = true
	}

	return len(result)
}

func abs(x int) int {
	// fmt.Println(x)
	if x < 0 {
		return -x
	}
	return x
}