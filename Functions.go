// package main
// import ("fmt")

// func familyName(fname string, age int) {
//   fmt.Println("Hello", age, "year old", fname, "Refsnes")
// }

// func main() {
//   familyName("Liam", 3)
//   familyName("Jenny", 14)
//   familyName("Anja", 30)
// }


// package main
// import ("fmt")

// func myFunction(x int, y int) int {
//   return x + y
// }

// func main() {
//   fmt.Println(myFunction(1, 2))
// }

// package main
// import ("fmt")

// func myFunction(x int, y int) (result int) {
//   result = x + y
//   return
// }

// func main() {
//   fmt.Println(myFunction(1, 2))
// }



// multiple return values
// package main
// import ("fmt")

// func myFunction(x int, y string) (result int, txt1 string) {
//   result = x + x
//   txt1 = y + " World!"
//   return
// }

// func main() {
//   fmt.Println(myFunction(5, "Hello"))
// }

// Here, we store the two return values into two variables (a and b):
// package main
// import ("fmt")

// func myFunction(x int, y string) (result int, txt1 string) {
//   result = x + x
//   txt1 = y + " World!"
//   return
// }

// func main() {
//   a, b := myFunction(5, "Hello")
//   fmt.Println(a, b)
// }

//Here, we want to omit the first returned value and only store the second returned value into a variable (b):
package main
import ("fmt")

func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func main() {
   _, b := myFunction(5, "Hello")
  fmt.Println(b)
}