package main

import "fmt"

func main() {
	// Loop de 0 até 100
	for i := 0; i <= 100; i++ {
		// Verifica primeiro se é múltiplo de 3 E de 5 (opcional, mas comum em lógica "FizzBuzz")
		// Se o número for 0, ele entra em ambas as condições.
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("pinpan")
		} else if i%3 == 0 {
			// Verifica se é múltiplo de 3
			fmt.Println("pin")
		} else if i%5 == 0 {
			// Verifica se é múltiplo de 5
			fmt.Println("pan")
		} else {
			// Caso contrário, imprime o número
			fmt.Println(i)
		}
	}
}
