// A program that shows how to get text form command line with bufio

package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Escribe tu nombre: ")
	nombre,err := reader.ReadString('\n')

	if err!=nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Tu nombre es %s\n", nombre)
	}
}