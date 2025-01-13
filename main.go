package main

import "fmt"

func transformNumbers(numbers *[]int, transform func(int) int ) []int {
	transformedNumbers := []int{}
	for _, value := range *numbers {
		transformedNumbers = append(transformedNumbers, transform(value))
	}
	return transformedNumbers
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubledNumbers := transformNumbers(&numbers, func(number int) int {
    return number * 2; 
  })

	fmt.Println("Doubled numbers is:", doubledNumbers)
}
