//A program that shows how to converts types in Go at command line

package main

import (
	"fmt"
	"strconv"
)

func main(){
	edad := 22
	edad_str := strconv.Itoa(edad)

	numero := "30"
	numero_int,_ := strconv.Atoi(numero)

	fmt.Println("Mi edad es "+edad_str)
	fmt.Println(numero_int+10)
}