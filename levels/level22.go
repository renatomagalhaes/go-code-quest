package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
		"math/rand"
		"time"
)

type Level22 struct{}

func StartLevel22() {
    game.StartLevel(Level22{})
}

func (l Level22) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 22: Uso Avançado de Ponteiros" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Entenda e use ponteiros para manipular dados de forma eficiente" + string(reset))
    fmt.Println()

    // Exemplo avançado de uso de ponteiros
    fmt.Println(string(green) + "Exemplo avançado de uso de ponteiros:" + string(reset))

    // Função que altera um valor usando um ponteiro
    changeValue := func(val *int) {
        *val = 42
    }

    var number int = 10
    fmt.Println("Valor original:", number)

    changeValue(&number)
    fmt.Println("Valor alterado:", number)

    // Analogias
    fmt.Println("\n" + string(cyan) + "Analogias:" + string(reset))
    fmt.Println("Usar ponteiros é como enviar o endereço de uma casa em vez de enviar a casa inteira.")
    fmt.Println("Você pode acessar e modificar o conteúdo da casa (valor) conhecendo apenas seu endereço (ponteiro).")

    showPointersSummary()
}

func showPointersSummary() {
    green := "\033[32m"
    cyan := "\033[36m"
    yellow := "\033[33m"
    reset := "\033[0m"
    separator := "==============================================="

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Exemplo de solução em Go" + string(reset))
    fmt.Println(separator)
    fmt.Println("```go")
    fmt.Println("package main")
    fmt.Println()
    fmt.Println("import \"fmt\"")
    fmt.Println()
    fmt.Println("func changeValue(val *int) {")
    fmt.Println("    *val = 42")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    var number int = 10")
    fmt.Println("    fmt.Println(\"Valor original:\", number)")
    fmt.Println()
    fmt.Println("    changeValue(&number)")
    fmt.Println("    fmt.Println(\"Valor alterado:\", number)")
    fmt.Println()
    fmt.Println("    // Analogias")
    fmt.Println("    fmt.Println(\"Usar ponteiros é como enviar o endereço de uma casa em vez de enviar a casa inteira.\")")
    fmt.Println("    fmt.Println(\"Você pode acessar e modificar o conteúdo da casa (valor) conhecendo apenas seu endereço (ponteiro).\")")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a usar ponteiros em Go para manipular dados de forma eficiente. Ao passar o endereço de uma variável, você pode acessar e modificar seu valor diretamente, sem precisar copiar a variável inteira." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use o operador `&` para obter o endereço de uma variável.",
        "Curiosidade: Ponteiros são amplamente utilizados em linguagens de baixo nível como C e C++.",
        "Dica: Use o operador `*` para acessar o valor armazenado no endereço de um ponteiro.",
        "Curiosidade: Em Go, todos os tipos possuem um ponteiro correspondente, como `*int` para `int`.",
        "Dica: Ponteiros podem ajudar a melhorar o desempenho evitando cópias de grandes estruturas de dados.",
        "Curiosidade: A linguagem Go gerencia automaticamente a memória dos ponteiros, evitando muitos dos problemas de gerenciamento manual de memória encontrados em C e C++.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o vigésimo segundo nível!" + string(reset))
}
