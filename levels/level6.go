package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
    "math/rand"
    "time"
)

// Definir a interface Shape
type Shape interface {
    Area() float64
}

// Definir a estrutura Circle
type Circle struct {
    Radius float64
}

// Implementar o método Area para Circle
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// Definir a estrutura Rectangle
type Rectangle struct {
    Width, Height float64
}

// Implementar o método Area para Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

type Level6 struct{}

func StartLevel6() {
    game.StartLevel(Level6{})
}

func (l Level6) Start() {
    // Definir cores
    //red := "\033[31m"
    green := "\033[32m"
    yellow := "\033[33m"
    //blue := "\033[34m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 6: Interfaces" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Crie interfaces para calcular a área de diferentes formas geométricas" + string(reset))
    fmt.Println()

    // Criar instâncias das estruturas Circle e Rectangle
    circle := Circle{Radius: 5}
    rectangle := Rectangle{Width: 4, Height: 3}

    // Criar um slice de Shape e adicionar as formas
    shapes := []Shape{circle, rectangle}

    // Calcular e exibir as áreas das formas
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resultados" + string(reset))
    fmt.Println(separator)
    for _, shape := range shapes {
        fmt.Printf(string(green)+"Área: %.2f\n"+string(reset), shape.Area())
    }

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Exemplo de solução em Go" + string(reset))
    fmt.Println(separator)
    fmt.Println("```go")
    fmt.Println("package main")
    fmt.Println()
    fmt.Println("import \"fmt\"")
    fmt.Println()
    fmt.Println("// Definir a interface Shape")
    fmt.Println("type Shape interface {")
    fmt.Println("    Area() float64")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("// Definir a estrutura Circle")
    fmt.Println("type Circle struct {")
    fmt.Println("    Radius float64")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("// Implementar o método Area para Circle")
    fmt.Println("func (c Circle) Area() float64 {")
    fmt.Println("    return 3.14 * c.Radius * c.Radius")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("// Definir a estrutura Rectangle")
    fmt.Println("type Rectangle struct {")
    fmt.Println("    Width, Height float64")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("// Implementar o método Area para Rectangle")
    fmt.Println("func (r Rectangle) Area() float64 {")
    fmt.Println("    return r.Width * r.Height")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    // Criar instâncias das estruturas Circle e Rectangle")
    fmt.Println("    circle := Circle{Radius: 5}")
    fmt.Println("    rectangle := Rectangle{Width: 4, Height: 3}")
    fmt.Println()
    fmt.Println("    // Criar um slice de Shape e adicionar as formas")
    fmt.Println("    shapes := []Shape{circle, rectangle}")
    fmt.Println()
    fmt.Println("    // Calcular e exibir as áreas das formas")
    fmt.Println("    for _, shape := range shapes {")
    fmt.Println("        fmt.Printf(\"Área: %.2f\\n\", shape.Area())")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a criar e usar interfaces em Go. Interfaces permitem definir métodos que devem ser implementados por tipos que satisfaçam a interface, facilitando a generalização de comportamentos." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use interfaces para definir comportamentos comuns que diferentes tipos podem implementar.",
        "Curiosidade: Interfaces em Go são satisfeitas implicitamente, o que simplifica a implementação.",
        "Dica: Combine interfaces menores em interfaces maiores para maior flexibilidade e modularidade.",
        "Curiosidade: Em Go, qualquer tipo que implemente os métodos de uma interface satisfaz a interface, sem necessidade de declaração explícita.",
        "Dica: Use a palavra-chave `type` para definir novas interfaces e structs.",
        "Curiosidade: Interfaces em Go ajudam a promover a composição sobre herança, um princípio importante de design de software.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o sexto nível!" + string(reset))
}
