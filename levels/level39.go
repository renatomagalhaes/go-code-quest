package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
    "math/rand"
    "time"
)

type Level39 struct{}

func StartLevel39() {
    game.StartLevel(Level39{})
}

func (l Level39) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 39: Utilizando Channels e Select para Comunicação Concorrente em Go" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Aprenda a utilizar channels e a instrução select para comunicação concorrente em Go" + string(reset))
    fmt.Println()

    // Criar channels
    ch1 := make(chan int)
    ch2 := make(chan int)

    // Função para enviar dados para o channel
    go func() {
        for {
            ch1 <- rand.Intn(100)
            time.Sleep(time.Millisecond * 500)
        }
    }()

    // Função para enviar dados para o outro channel
    go func() {
        for {
            ch2 <- rand.Intn(100)
            time.Sleep(time.Millisecond * 700)
        }
    }()

    // Receber dados usando select
    go func() {
        for {
            select {
            case msg1 := <-ch1:
                fmt.Println(string(green), "Recebido do ch1:", msg1, string(reset))
            case msg2 := <-ch2:
                fmt.Println(string(cyan), "Recebido do ch2:", msg2, string(reset))
            }
        }
    }()

    // Simulação de tempo de execução
    time.Sleep(5 * time.Second)

    showChannelsSummary39()
}

func showChannelsSummary39() {
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
    fmt.Println("import (")
    fmt.Println("    \"fmt\"")
    fmt.Println("    \"math/rand\"")
    fmt.Println("    \"time\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    // Criar channels")
    fmt.Println("    ch1 := make(chan int)")
    fmt.Println("    ch2 := make(chan int)")
    fmt.Println()
    fmt.Println("    // Função para enviar dados para o channel")
    fmt.Println("    go func() {")
    fmt.Println("        for {")
    fmt.Println("            ch1 <- rand.Intn(100)")
    fmt.Println("            time.Sleep(time.Millisecond * 500)")
    fmt.Println("        }")
    fmt.Println("    }()")
    fmt.Println()
    fmt.Println("    // Função para enviar dados para o outro channel")
    fmt.Println("    go func() {")
    fmt.Println("        for {")
    fmt.Println("            ch2 <- rand.Intn(100)")
    fmt.Println("            time.Sleep(time.Millisecond * 700)")
    fmt.Println("        }")
    fmt.Println("    }()")
    fmt.Println()
    fmt.Println("    // Receber dados usando select")
    fmt.Println("    go func() {")
    fmt.Println("        for {")
    fmt.Println("            select {")
    fmt.Println("            case msg1 := <-ch1:")
    fmt.Println("                fmt.Println(\"Recebido do ch1:\", msg1)")
    fmt.Println("            case msg2 := <-ch2:")
    fmt.Println("                fmt.Println(\"Recebido do ch2:\", msg2)")
    fmt.Println("            }")
    fmt.Println("        }")
    fmt.Println("    }()")
    fmt.Println()
    fmt.Println("    // Simulação de tempo de execução")
    fmt.Println("    time.Sleep(5 * time.Second)")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a utilizar channels e a instrução select para comunicação concorrente em Go, permitindo a troca de dados entre goroutines de maneira eficiente e segura." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Channels permitem a comunicação segura entre goroutines.",
        "Curiosidade: A instrução select em Go permite aguardar em múltiplos channels simultaneamente.",
        "Dica: Use `make(chan tipo)` para criar um channel.",
        "Curiosidade: Channels podem ser unidirecionais (enviar ou receber) ou bidirecionais.",
        "Dica: A instrução select é útil para evitar bloqueios em operações de comunicação concorrente.",
        "Curiosidade: Channels podem ser usados para implementar padrões como produtor-consumidor e pipelines.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o trigésimo nono nível!" + string(reset))
}
