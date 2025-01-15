package variadicFunctions

import "fmt"

//This functions is accpeting individual integer parameters and convert it into slice and store onto numbers variable (That's how the function is variadic)
func sum(firstParameter int, numbers ...int) int {
  sum := 0;
  for _, val := range numbers {
    sum += val;
  }
  return sum + firstParameter;
}

func VariadicFunctionsMainFn(){
  numbers := []int{1, 2, 3, 4};
  sumOfNumbers := sum(5, numbers...);
  fmt.Println("Sum of the numbers is:", sumOfNumbers);
}
