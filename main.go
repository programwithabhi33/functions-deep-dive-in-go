package main

import "fmt"
type transformFn func(int) int

func getTransformerFn(numbers *[]int) func(int) int {
  if (*numbers)[0] == 1 {
    return double;
  }else{
    return triple;
  }
}

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
  numbers2 := []int{5, 6, 11, 30, 101};

  doubledNumbers := transformNumbers(&numbers, double);
  tripledNumbers := transformNumbers(&numbers, triple);
  fmt.Println("Doubled numbers is:", doubledNumbers);
  fmt.Println("Tripled numbers is:", tripledNumbers);

  transformFn := getTransformerFn(&numbers)
  transformFn2 := getTransformerFn(&numbers2)

  transformedNumbers := transformNumbers(&numbers, transformFn);
  transformedNumbers2 := transformNumbers(&numbers2, transformFn2);

  fmt.Println(transformedNumbers)
  fmt.Println(transformedNumbers2)
}
