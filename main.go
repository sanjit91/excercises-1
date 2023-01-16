package main

import "fmt"

func main() {

	// first question output function

	// fmt.Println(evenNumbers([]int{1, 2, 3, 4, 5})) // output: [2, 4]

	// second question output function

	// frequencies := characterFrequency("hello")
	// fmt.Println(frequencies)

	// third question output function

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sumSlice(nums))

}

// first question solution

// func evenNumbers(nums []int) []int {
// 	evens := []int{}
// 	for _, num := range nums {
// 		if num%2 == 0 {
// 			evens = append(evens, num)
// 		}
// 	}
// 	return evens
// }

// second question solution

// func characterFrequency(s string) map[string]int {
// 	m := make(map[string]int)
// 	for _, char := range s {
// 		if _, ok := m[string(char)]; ok {
// 			m[string(char)]++
// 		} else {
// 			m[string(char)] = 1
// 		}
// 	}
// 	return m
// }

// third question solution

func sumSlice(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
