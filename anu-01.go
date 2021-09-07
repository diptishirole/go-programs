package main

import "fmt"

func main() {
   
   p:= 34
   q:= 20

  //Addition
  result1:= p + q
  fmt.Printf("Result of p + q = %d", result1)

  //Subtraction 
  result2:= p - q
  fmt.Printf("\nResult of p - q = %d", result2)

  //Multiplication
  result3:= p * q
  fmt.Printf("\nResult of p * q = %d", result3)

  //Division
  result4:= p / q
  fmt.Printf("\nResult of p / q = %d", result4)

  //Modulus
  result5:= p % q
  fmt.Printf("\nResult of p % q = %%d", result5)

  num1 := 100
  num2 := 200
  operator := "+"
  result := mathFunc(num1, num2, operator)
 
  fmt.Printf("\nResult of p %s q = %d", operator, result5)  

}
package main

import "fmt"

func main() {
   a:= 2

   //Using address og operator(&) and
   //pointer indirection(*) operator
   b:= &a
   fmt.Println(*b)
   *b = 9
   fmt.Println(a)

}