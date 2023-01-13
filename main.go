package main

import "fmt"

func main() {
	fmt.Println(evenNumbers([]int{1, 2, 3, 4, 5})) // output: [2, 4]

}

func evenNumbers(nums []int) []int {
	evens := []int{}
	for _, num := range nums {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	return evens
}
