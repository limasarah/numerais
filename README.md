Desafios de Código – DIO (Go)

Este repositório contém a resolução de dois desafios de lógica de programação propostos pela plataforma DIO (Digital Innovation One), utilizando a linguagem Go.

Os desafios têm como objetivo praticar estruturas de repetição, condicionais e o uso do operador de módulo (%).

📌 Desafio 1 – Números divisíveis por 3
🎯 Objetivo

Criar um programa que exiba todos os números de 1 a 100 que sejam divisíveis por 3.

🧠 Lógica utilizada

Utilização de um for para iterar de 1 até 100

Uso do operador de módulo % para verificar se o número é divisível por 3

Impressão apenas dos números cujo resto da divisão por 3 seja igual a 0

💻 Código
package main

import "fmt"

func main() {
    // Loop de 1 a 100
    for i := 1; i <= 100; i++ {
        // Verifica se o número é divisível por 3
        if i%3 == 0 {
            fmt.Println(i)
        }
    }
}
📤 Saída esperada
3, 6, 9, 12, ..., 96, 99
📌 Desafio 2 – Múltiplos de 3 e 5 (Pin e Pan)
🎯 Objetivo

Criar um programa que percorra os números de 0 a 100, seguindo as regras:

Se o número for múltiplo de 3, imprimir "Pin"

Se o número for múltiplo de 5, imprimir "Pan"

Caso não seja múltiplo de nenhum dos dois, imprimir o próprio número

🔎 Observação: Caso o número seja múltiplo de 3 e 5 ao mesmo tempo, pode-se adaptar a lógica para tratar essa condição específica (ex: "PinPan").

🧠 Lógica utilizada

Loop for de 0 a 100

Operador % para verificar múltiplos

Estruturas condicionais if / else if / else para controlar a saída

💻 Código
package main

import "fmt"

func main() {
    for i := 0; i <= 100; i++ {

        if i%3 == 0 && i%5 == 0 {
            fmt.Println("PinPan")
        } else if i%3 == 0 {
            fmt.Println("Pin")
        } else if i%5 == 0 {
            fmt.Println("Pan")
        } else {
            fmt.Println(i)
        }
    }
}
🛠️ Tecnologias utilizadas

Go (Golang)

Estruturas de repetição (for)

Estruturas condicionais (if, else if, else)

Operador de módulo (%)

📚 Aprendizados

Com esses desafios foi possível praticar:

Lógica de programação

Controle de fluxo

Estruturas básicas da linguagem Go

Organização de código simples e legível
