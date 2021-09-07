package main
import "fmt"
func main() {
   p:= 30
   q:= 10
   
   //& (bitwise AND)
   result1:= p & q
   fmt.Printf("Result of p & q = %d", result1)

   //| (bitwise OR)
   result2:= p | q
   fmt.Printf("\nResult of p | q = %d", result2)

   //^ (bitwise XOR)
   result3:= p ^ q
   fmt.Printf("\nResult of p ^ q = %d", result3)

   //<< (left shift)
   result4:= p << q 
   fmt.Printf("\nResult of p << q = %d", result4)

   //>> (right shift)
   result5:= p >> q
   fmt.Printf("\nResult of p >> q = %d", result5)

   //&^ (AND NOT)
   result6:= p &^ q
   fmt.Printf("\nResult of p &^ q = %d", result6)


}
package main
import "fmt"

func main() {
   var p int = 30
   var q int = 40
 
   //"="(Simple Assignment)
   p = q
   fmt.Println(p)

   //"+="(Add Assignment)
   p += q
   fmt.Println(p)

   //"-="(Subtract Assignment)
   p -= q
   fmt.Println(p)

   //"*="(Multiply Assignment)
   p *= q
   fmt.Println(p)

   //"/="(Division Assignment)
   p /= q
   fmt.Println(p)
   
   //"%="(Modulus Assignment)
   p %= q
   fmt.Println(p)