package main

import "fmt"

func main() {
	// Loop de 1 a 100
	for i := 1; i <= 100; i++ {
		// Verifica se o número é divisível por 3
		// (o resto da divisão inteira por 3 é 0)
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
