package recursion

import "fmt"

//Iterative approach to calculate a factorial of a number;
func factorialCalculationFnWithIterative(number int) int {
  result := 1;
  for i := 1; i<=number; i++ {
    result = result * i
  }
  return result
}

// Recursive approach to calculate a factorial of a number;
func factorialCalculationFnWithRecursive(number int) int {
  if number > 1 {
    number = number * factorialCalculationFnWithRecursive(number - 1)
  }
  return number;
}

func RecursionMainFn(){
  factorialOf5WithIterative := factorialCalculationFnWithIterative(5);
  factorialOf5WithRecursive := factorialCalculationFnWithRecursive(5);

  fmt.Println("Factorial of 5 with iterative approach is", factorialOf5WithIterative);
  fmt.Println("Factorial of 5 with recursive approach is", factorialOf5WithRecursive);
}
