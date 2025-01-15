package main

import (
	//"fmt"
	//"example.com/functionsDeepDive/recursion"
	"example.com/functionsDeepDive/variadicFunctions"
)

func transformNumbers(numbers *[]int, transform func(int) int ) []int {
	transformedNumbers := []int{}
	for _, value := range *numbers {
		transformedNumbers = append(transformedNumbers, transform(value))
	}
	return transformedNumbers
}

func createTransformFunction(factor int) func(int) int {
  return func (number int) int {
    return number * factor
  }
}

func main() {
	/*numbers := []int{1, 2, 3, 4, 5}
	doubledNumbers := transformNumbers(&numbers, func(number int) int {
    return number * 2; 
  });
  tripleTransformFn := createTransformFunction(3);
  tripledNumbers := transformNumbers(&numbers, tripleTransformFn)

	fmt.Println("Doubled numbers is:", doubledNumbers)
	fmt.Println("Tripled numbers is:", tripledNumbers)*/
  
  //Recursion
  //recursion.RecursionMainFn()
  variadicFunctions.VariadicFunctionsMainFn();
}
