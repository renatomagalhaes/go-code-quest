package levels

import (
    "context"
    "fmt"
    "go-code-quest/pkg/game"
    "math/rand"
    "sync"
    "time"
)

type Level24 struct{}

func StartLevel24() {
    game.StartLevel(Level24{})
}

func (l Level24) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 24: Uso de Contextos para Controle de Execução e Cancelamento" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Use o pacote context para controlar a execução e cancelamento de goroutines" + string(reset))
    fmt.Println()

    // Exemplo avançado de uso de contextos
    fmt.Println(string(green) + "Exemplo avançado de uso de contextos:" + string(reset))

    // Função que simula um trabalho com goroutines usando contextos
    work := func(ctx context.Context, wg *sync.WaitGroup, id int) {
        defer wg.Done()
        for {
            select {
            case <-ctx.Done():
                fmt.Printf("Goroutine %d cancelada\n", id)
                return
            default:
                fmt.Printf("Goroutine %d trabalhando...\n", id)
                time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
            }
        }
    }

    // Criar um contexto com cancelamento
    ctx, cancel := context.WithCancel(context.Background())

    var wg sync.WaitGroup
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go work(ctx, &wg, i)
    }

    // Esperar um pouco e depois cancelar o contexto
    time.Sleep(2 * time.Second)
    fmt.Println("Cancelando todas as goroutines...")
    cancel()

    // Esperar que todas as goroutines terminem
    wg.Wait()

    // Criar um contexto com timeout
    ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancelTimeout()

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go work(ctxTimeout, &wg, i)
    }

    // Esperar que todas as goroutines terminem ou o timeout ocorra
    wg.Wait()

    showContextSummary()
}

func showContextSummary() {
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
    fmt.Println("    \"context\"")
    fmt.Println("    \"fmt\"")
    fmt.Println("    \"math/rand\"")
    fmt.Println("    \"sync\"")
    fmt.Println("    \"time\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    work := func(ctx context.Context, wg *sync.WaitGroup, id int) {")
    fmt.Println("        defer wg.Done()")
    fmt.Println("        for {")
    fmt.Println("            select {")
    fmt.Println("            case <-ctx.Done():")
    fmt.Println("                fmt.Printf(\"Goroutine %d cancelada\\n\", id)")
    fmt.Println("                return")
    fmt.Println("            default:")
    fmt.Println("                fmt.Printf(\"Goroutine %d trabalhando...\\n\", id)")
    fmt.Println("                time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)")
    fmt.Println("            }")
    fmt.Println("        }")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    ctx, cancel := context.WithCancel(context.Background())")
    fmt.Println()
    fmt.Println("    var wg sync.WaitGroup")
    fmt.Println("    for i := 1; i <= 5; i++ {")
    fmt.Println("        wg.Add(1)")
    fmt.Println("        go work(ctx, &wg, i)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    time.Sleep(2 * time.Second)")
    fmt.Println("    fmt.Println(\"Cancelando todas as goroutines...\")")
    fmt.Println("    cancel()")
    fmt.Println()
    fmt.Println("    wg.Wait()")
    fmt.Println()
    fmt.Println("    ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 2*time.Second)")
    fmt.Println("    defer cancelTimeout()")
    fmt.Println()
    fmt.Println("    for i := 1; i <= 5; i++ {")
    fmt.Println("        wg.Add(1)")
    fmt.Println("        go work(ctxTimeout, &wg, i)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    wg.Wait()")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a usar o pacote context para controlar a execução e cancelamento de goroutines. Isso é útil para garantir que goroutines não continuem a execução desnecessariamente e para controlar o tempo de execução." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use `context.Background()` para criar um contexto vazio.",
        "Curiosidade: `context.WithCancel()` retorna um contexto que pode ser cancelado manualmente.",
        "Dica: Use `context.WithTimeout()` para definir um tempo limite após o qual o contexto será cancelado automaticamente.",
        "Curiosidade: O pacote `context` é amplamente utilizado em servidores web e em programas concorrentes.",
        "Dica: O método `Done()` de um contexto retorna um canal que é fechado quando o contexto é cancelado ou expira.",
        "Curiosidade: O pacote `context` foi introduzido no Go 1.7 e se tornou um padrão para cancelamento e controle de tempo.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o vigésimo quarto nível!" + string(reset))
}
