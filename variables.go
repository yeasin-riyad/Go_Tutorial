package main
import "fmt";

var name string = "Go Tutorial"  //type is string
var age int = 1  //type is int
var isFun bool = true  //type is bool
var name1="Go Tutorial";  //type is inferred as string
var age1=1;  //type is inferred as int
var isFun1=true;  //type is inferred as bool	

  var a, b, c, d int = 1, 3, 5, 7

   


func main() {
	 e, f := 7, "World!"
	fmt.Println("Name:", name);
	fmt.Println("Age:", age);
	fmt.Println("Is Go fun?", isFun);
	fmt.Println(a," ",b," ",c," ",d);
	fmt.Println(c," ",d);
	fmt.Println(e," ",f);
};
