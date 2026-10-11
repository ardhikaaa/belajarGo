package main

import "fmt"

/*func main() {
	for i := 1; i <= 5; i++ {
		for j :=1; j <= i; j++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
output:
* 
* * 
* * * 
* * * * 
* * * * * 	
*/

/*func main() {
	for i := 1; i <= 4; i++ {
		for s := 1; s <= 4 - i; s++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i - 1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	
}
output:
   *
  ***
 *****
*******	
*/

/*func main() {
	for i := 4; i >= 1; i-- {
		for s := 1; s <= 4 - i; s++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i - 1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	
}
output:
*******
 *****
  ***
   *
*/

/*func main() {
	for i := 1; i <= 4; i++ {
		for s := 1; s <= 4 - i; s++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i - 1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 4 - 1; i >= 1; i-- {
		for s := 1; s <= 4 - i; s++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i - 1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
output:
   *
  ***
 *****
*******
 *****
  ***
   *
*/

func main() {
	for i := 1; i <= 3; i++ {
		for s := 1; s <= 3 - i; s++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i - 1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}