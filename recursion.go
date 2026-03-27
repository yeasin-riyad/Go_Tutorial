//A function is recursive if it calls itself and reaches a stop condition.
// package main
// import ("fmt")

// func testcount(x int) int {
//   if x == 11 {
//     return 0
//   }
//   fmt.Println(x)
//   return testcount(x + 1)
// }

// func main(){
//   testcount(1)
// }


package main
import ("fmt")

func factorial_recursion(x float64) (y float64) {
  if x > 0 {
     y = x * factorial_recursion(x-1)
  } else {
     y = 1
  }
  return
}

func main() {
  fmt.Println(factorial_recursion(4))
}