package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
		"math/rand"
		"time"
)

type Level16 struct{}

func StartLevel16() {
    game.StartLevel(Level16{})
}

func (l Level16) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 16: Ponteiros" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Trabalhe com ponteiros em Golang" + string(reset))
    fmt.Println()

    // Exemplo de uso de ponteiros
    fmt.Println(string(green) + "Exemplo de uso de ponteiros:" + string(reset))
    var x int = 10
    var p *int = &x // Declaração de um ponteiro para x
    fmt.Printf("Valor de x: %d\n", x)
    fmt.Printf("Endereço de x: %p\n", p)
    fmt.Printf("Valor apontado por p: %d\n", *p)

    // Função que aceita um ponteiro como argumento
    increment := func(n *int) {
        *n++
    }
    increment(&x)
    fmt.Printf("Valor de x após incremento: %d\n", x)

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Exemplo de solução em Go" + string(reset))
    fmt.Println(separator)
    fmt.Println("```go")
    fmt.Println("package main")
    fmt.Println()
    fmt.Println("import \"fmt\"")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    var x int = 10")
    fmt.Println("    var p *int = &x // Declaração de um ponteiro para x")
    fmt.Println("    fmt.Printf(\"Valor de x: %d\\n\", x)")
    fmt.Println("    fmt.Printf(\"Endereço de x: %p\\n\", p)")
    fmt.Println("    fmt.Printf(\"Valor apontado por p: %d\\n\", *p)")
    fmt.Println()
    fmt.Println("    // Função que aceita um ponteiro como argumento")
    fmt.Println("    increment := func(n *int) {")
    fmt.Println("        *n++")
    fmt.Println("    }")
    fmt.Println("    increment(&x)")
    fmt.Println("    fmt.Printf(\"Valor de x após incremento: %d\\n\", x)")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a usar ponteiros em Go. Ponteiros permitem passar referências de variáveis para funções, evitando cópias desnecessárias e permitindo modificações diretas nos valores originais." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use `&` para obter o endereço de uma variável.",
        "Curiosidade: Ponteiros permitem manipular diretamente a memória, oferecendo mais controle sobre o comportamento do programa.",
        "Dica: Use `*` para desreferenciar um ponteiro e acessar o valor apontado.",
        "Curiosidade: Em Go, os ponteiros são seguros e evitam muitos problemas comuns em linguagens como C.",
        "Dica: Passe ponteiros para funções quando precisar modificar o valor original da variável.",
        "Curiosidade: Go não possui aritmética de ponteiros, o que torna o uso de ponteiros mais seguro e menos propenso a erros.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o décimo sexto nível!" + string(reset))
}
