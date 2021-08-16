// A program that shows how to get text form command line

package main

import "fmt"

func main(){

	var edad int
	
	fmt.Printf("Inserta tu edad: ")
	fmt.Scanf("%d\n",&edad)
	fmt.Printf("Tu edad es %d años\n",edad)

	var nombre string

	fmt.Printf("Inserta tu nombre: ")
	fmt.Scanf("%s\n",&nombre)
	fmt.Printf("Tu nombre es %s\n",nombre)
}