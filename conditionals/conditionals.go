// A program that shows different forms of use if and else

package main

import "fmt"

func main(){

	x := 10
	y:= 5

	if x > y {
		fmt.Printf("%d es mayor que %d\n",x,y)
	} else if x < y{
		fmt.Printf("%d es menor que %d\n",x,y)
	} else {
		fmt.Printf("Son iguales \n")
	}
}