package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
    "time"
    "math/rand"
)

type Level7 struct{}

func StartLevel7() {
    game.StartLevel(Level7{})
}

func (l Level7) Start() {
    // Definir cores
    //red := "\033[31m"
    green := "\033[32m"
    yellow := "\033[33m"
    //blue := "\033[34m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 7: Concurrency" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Use goroutines e canais para implementar execução concorrente" + string(reset))
    fmt.Println()

    // Criar um canal para comunicação entre goroutines
    messageChannel := make(chan string)

    // Definir uma função para ser executada como goroutine
    go func() {
        time.Sleep(2 * time.Second) // Simular trabalho
        messageChannel <- "Trabalho concluído na goroutine 1"
    }()

    go func() {
        time.Sleep(1 * time.Second) // Simular trabalho
        messageChannel <- "Trabalho concluído na goroutine 2"
    }()

    // Receber mensagens das goroutines
    for i := 0; i < 2; i++ {
        fmt.Println(string(green) + <-messageChannel + string(reset))
    }

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Exemplo de solução em Go" + string(reset))
    fmt.Println(separator)
    fmt.Println("```go")
    fmt.Println("package main")
    fmt.Println()
    fmt.Println("import (")
    fmt.Println("    \"fmt\"")
    fmt.Println("    \"time\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    // Criar um canal para comunicação entre goroutines")
    fmt.Println("    messageChannel := make(chan string)")
    fmt.Println()
    fmt.Println("    // Definir uma função para ser executada como goroutine")
    fmt.Println("    go func() {")
    fmt.Println("        time.Sleep(2 * time.Second) // Simular trabalho")
    fmt.Println("        messageChannel <- \"Trabalho concluído na goroutine 1\"")
    fmt.Println("    }()")
    fmt.Println()
    fmt.Println("    go func() {")
    fmt.Println("        time.Sleep(1 * time.Second) // Simular trabalho")
    fmt.Println("        messageChannel <- \"Trabalho concluído na goroutine 2\"")
    fmt.Println("    }()")
    fmt.Println()
    fmt.Println("    // Receber mensagens das goroutines")
    fmt.Println("    for i := 0; i < 2; i++ {")
    fmt.Println("        fmt.Println(<-messageChannel)")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a usar goroutines e canais em Go. Goroutines permitem execução concorrente de funções, e canais permitem comunicação segura entre goroutines." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use goroutines para realizar tarefas que podem ser executadas em paralelo.",
        "Curiosidade: Goroutines são leves, o que permite a criação de milhares de goroutines em um único programa.",
        "Dica: Use canais para comunicar e sincronizar goroutines de forma segura.",
        "Curiosidade: A palavra-chave `go` é usada para iniciar uma goroutine em Go.",
        "Dica: Combine goroutines e canais para implementar padrões de concorrência como fan-out e fan-in.",
        "Curiosidade: Go foi projetado com suporte robusto para concorrência, facilitando a criação de programas concorrentes.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o sétimo nível!" + string(reset))
}
