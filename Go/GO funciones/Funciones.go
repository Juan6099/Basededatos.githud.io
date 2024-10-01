package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Introduce una cadena: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    if input == "" {
        fmt.Println("La cadena está vacía.")
    } else {
      
        length := len(input)
        fmt.Printf("La longitud de la cadena es: %d\n", length)

        upper := strings.ToUpper(input)
        fmt.Printf("En mayúsculas: %s\n", upper)

    
        lower := strings.ToLower(input)
        fmt.Printf("En minúsculas: %s\n", lower)

      
        palabra := "golang"
        if strings.Contains(input, palabra) {
            fmt.Printf("La cadena contiene la palabra: %s\n", palabra)
        } else {
            fmt.Printf("La cadena NO contiene la palabra: %s\n", palabra)
        }

        
        nuevaCadena := strings.ReplaceAll(input, " ", "_")
        fmt.Printf("Cadena con espacios reemplazados: %s\n", nuevaCadena)
    }
}