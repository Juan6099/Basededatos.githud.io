package main

import (
    "fmt"
    "math"
)

func main() {
    var peso, altura float64

    fmt.Print("Ingrese su peso en kilogramos: ")
    fmt.Scanln(&peso)

    fmt.Print("Ingrese su altura en metros: ")
    fmt.Scanln(&altura)

    imc := calcularIMC(peso, altura)

    fmt.Printf("Su IMC es: %.2f\n", imc)

    interpretarIMC(imc)
}

func calcularIMC(peso, altura float64) float64 {
    return peso / math.Pow(altura, 2)
}

func interpretarIMC(imc float64) {
    if imc < 18.5 {
        fmt.Println("Bajo peso")
    } else if imc < 25 {
        fmt.Println("Peso normal")
    } else if imc < 30 {
        fmt.Println("Sobrepeso")
    } else {
        fmt.Println("Obesidad")
    }
}