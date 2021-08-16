// A program that shows different forms of use printf

package main

import "fmt"

func main(){

	edad := 22
	nombre := "Juanma"
	precio := 14.3

	fmt.Printf("Mi edad es %d\n", edad)
	fmt.Printf("Mi edad es %v\n", edad)
	fmt.Printf("Mi nombre es %s\n", nombre)
	fmt.Printf("Mi nombre es %v\n", nombre)
	fmt.Printf("El precio es %f\n", precio)
}