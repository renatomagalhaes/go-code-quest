package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
    "math/rand"
    "sync"
    "time"
)

type Level33 struct{}

func StartLevel33() {
    game.StartLevel(Level33{})
}

func (l Level33) Start() {
    // Definir cores
    //green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 33: Trabalhando com Canais Avançados" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Explore o uso avançado de canais para sincronização e comunicação entre goroutines" + string(reset))
    fmt.Println()

    // Exemplos de canais avançados

    // Exemplo 1: Canal com buffer
    bufferedChannel()

    // Exemplo 2: Canal com timeout
    channelWithTimeout()

    // Exemplo 3: Fan-in e Fan-out
    fanInFanOut()

    showChannelsSummary()
}

func bufferedChannel() {
    green := "\033[32m"
    reset := "\033[0m"

    fmt.Println(string(green) + "Exemplo de Canal com Buffer" + string(reset))
    ch := make(chan int, 3)
    ch <- 1
    ch <- 2
    ch <- 3
    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}

func channelWithTimeout() {
    green := "\033[32m"
    reset := "\033[0m"

    fmt.Println(string(green) + "Exemplo de Canal com Timeout" + string(reset))
    ch := make(chan int)

    go func() {
        time.Sleep(2 * time.Second)
        ch <- 42
    }()

    select {
    case res := <-ch:
        fmt.Println("Recebido:", res)
    case <-time.After(1 * time.Second):
        fmt.Println("Timeout")
    }
}

func fanInFanOut() {
    green := "\033[32m"
    reset := "\033[0m"

    fmt.Println(string(green) + "Exemplo de Fan-In e Fan-Out" + string(reset))
    in := make(chan int)
    out1 := make(chan int)
    out2 := make(chan int)

    var wg sync.WaitGroup

    go func() {
        for i := 0; i < 10; i++ {
            in <- i
        }
        close(in)
    }()

    for i := 0; i < 2; i++ {
        wg.Add(1)
        go func() {
            for n := range in {
                if rand.Intn(2) == 0 {
                    out1 <- n
                } else {
                    out2 <- n
                }
            }
            wg.Done()
        }()
    }

    go func() {
        wg.Wait()
        close(out1)
        close(out2)
    }()

    go func() {
        for n := range out1 {
            fmt.Println("out1:", n)
        }
    }()

    for n := range out2 {
        fmt.Println("out2:", n)
    }
}

func showChannelsSummary() {
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
    fmt.Println("    \"sync\"")
    fmt.Println("    \"time\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    // Exemplo de Canal com Buffer")
    fmt.Println("    ch := make(chan int, 3)")
    fmt.Println("    ch <- 1")
    fmt.Println("    ch <- 2")
    fmt.Println("    ch <- 3")
    fmt.Println("    fmt.Println(<-ch)")
    fmt.Println("    fmt.Println(<-ch)")
    fmt.Println("    fmt.Println(<-ch)")
    fmt.Println()
    fmt.Println("    // Exemplo de Canal com Timeout")
    fmt.Println("    ch2 := make(chan int)")
    fmt.Println()
    fmt.Println("    go func() {")
    fmt.Println("        time.Sleep(2 * time.Second)")
    fmt.Println("        ch2 <- 42")
    fmt.Println("    }()")
    fmt.Println()
    fmt.Println("    select {")
    fmt.Println("    case res := <-ch2:")
    fmt.Println("        fmt.Println(\"Recebido:\", res)")
    fmt.Println("    case <-time.After(1 * time.Second):")
    fmt.Println("        fmt.Println(\"Timeout\")")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    // Exemplo de Fan-In e Fan-Out")
    fmt.Println("    in := make(chan int)")
    fmt.Println("    out1 := make(chan int)")
    fmt.Println("    out2 := make(chan int)")
    fmt.Println()
    fmt.Println("    var wg sync.WaitGroup")
    fmt.Println()
    fmt.Println("    go func() {")
    fmt.Println("        for i := 0; i < 10; i++ {")
    fmt.Println("            in <- i")
    fmt.Println("        }")
    fmt.Println("        close(in)")
    fmt.Println("    }()")
    fmt.Println()
    fmt.Println("    for i := 0; i < 2; i++ {")
    fmt.Println("        wg.Add(1)")
    fmt.Println("        go func() {")
    fmt.Println("            for n := range in {")
    fmt.Println("                if rand.Intn(2) == 0 {")
    fmt.Println("                    out1 <- n")
    fmt.Println("                } else {")
    fmt.Println("                    out2 <- n")
    fmt.Println("                }")
    fmt.Println("            }")
    fmt.Println("            wg.Done()")
    fmt.Println("        }()")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    go func() {")
    fmt.Println("        wg.Wait()")
    fmt.Println("        close(out1)")
    fmt.Println("        close(out2)")
    fmt.Println("    }()")
    fmt.Println()
    fmt.Println("    go func() {")
    fmt.Println("        for n := range out1 {")
    fmt.Println("            fmt.Println(\"out1:\", n)")
    fmt.Println("        }")
    fmt.Println("    }()")
    fmt.Println()
    fmt.Println("    for n := range out2 {")
    fmt.Println("        fmt.Println(\"out2:\", n)")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu sobre o uso avançado de canais para comunicação entre goroutines, incluindo canais com buffer, canais com timeout, e padrões de Fan-In e Fan-Out." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use canais com buffer para evitar bloqueios quando uma goroutine envia dados.",
        "Curiosidade: O padrão Fan-In combina dados de vários canais em um único canal.",
        "Dica: Use `select` para implementar timeouts em operações de canal.",
        "Curiosidade: O padrão Fan-Out distribui dados de um único canal para várias goroutines.",
        "Dica: Sincronize goroutines usando o pacote `sync` para garantir que todas as operações sejam concluídas.",
        "Curiosidade: O uso correto de canais pode simplificar a sincronização de goroutines e evitar condições de corrida.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o trigésimo terceiro nível!" + string(reset))
}
