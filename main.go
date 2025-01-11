package main

import "fmt"
type transformFn func(int) int

func transformNumbers(numbers *[]int, transform transformFn) []int{
  transformedNumbers := []int{};
  for _, value := range *numbers{
    transformedNumbers = append(transformedNumbers, transform(value));
  }
  return transformedNumbers;
} 

func double(number int) int {
  return number * 2;
}

func triple(number int) int {
  return number * 3;
}

func main(){
  numbers := []int{1, 2, 3, 4, 5};
  doubledNumbers := transformNumbers(&numbers, double);
  tripledNumbers := transformNumbers(&numbers, triple);
  fmt.Println("Doubled numbers is:", doubledNumbers);
  fmt.Println("Tripled numbers is:", tripledNumbers);
}
