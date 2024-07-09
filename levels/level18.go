package levels

import (
    "context"
    "fmt"
    "go-code-quest/pkg/game"
    "math/rand"
    "sync"
    "time"
)

type Level18 struct{}

func StartLevel18() {
    game.StartLevel(Level18{})
}

func (l Level18) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 18: Concorrência Avançada" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Use WaitGroup, canais e contexto para gerenciar goroutines" + string(reset))
    fmt.Println()

    // Exemplo de uso avançado de concorrência
    fmt.Println(string(green) + "Exemplo de uso avançado de concorrência:" + string(reset))

    // Canal para comunicação entre goroutines
    results := make(chan int, 5)

    // WaitGroup para sincronizar goroutines
    var wg sync.WaitGroup

    // Contexto para cancelar goroutines
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    worker := func(id int, ctx context.Context, results chan<- int, wg *sync.WaitGroup) {
        defer wg.Done()
        for {
            select {
            case <-ctx.Done():
                fmt.Printf("Worker %d parou\n", id)
                return
            default:
                result := rand.Intn(100)
                time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
                results <- result
                fmt.Printf("Worker %d enviou: %d\n", id, result)
            }
        }
    }

    // Iniciar várias goroutines
    numWorkers := 5
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go worker(i, ctx, results, &wg)
    }

    // Goroutine para coletar resultados
    go func() {
        wg.Wait()
        close(results)
    }()

    // Processar resultados
    for result := range results {
        fmt.Printf("Resultado recebido: %d\n", result)
    }

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
    fmt.Println("    // Canal para comunicação entre goroutines")
    fmt.Println("    results := make(chan int, 5)")
    fmt.Println()
    fmt.Println("    // WaitGroup para sincronizar goroutines")
    fmt.Println("    var wg sync.WaitGroup")
    fmt.Println()
    fmt.Println("    // Contexto para cancelar goroutines")
    fmt.Println("    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)")
    fmt.Println("    defer cancel()")
    fmt.Println()
    fmt.Println("    worker := func(id int, ctx context.Context, results chan<- int, wg *sync.WaitGroup) {")
    fmt.Println("        defer wg.Done()")
    fmt.Println("        for {")
    fmt.Println("            select {")
    fmt.Println("            case <-ctx.Done():")
    fmt.Println("                fmt.Printf(\"Worker %d parou\\n\", id)")
    fmt.Println("                return")
    fmt.Println("            default:")
    fmt.Println("                result := rand.Intn(100)")
    fmt.Println("                time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))")
    fmt.Println("                results <- result")
    fmt.Println("                fmt.Printf(\"Worker %d enviou: %d\\n\", id, result)")
    fmt.Println("            }")
    fmt.Println("        }")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    // Iniciar várias goroutines")
    fmt.Println("    numWorkers := 5")
    fmt.Println("    for i := 1; i <= numWorkers; i++ {")
    fmt.Println("        wg.Add(1)")
    fmt.Println("        go worker(i, ctx, results, &wg)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    // Goroutine para coletar resultados")
    fmt.Println("    go func() {")
    fmt.Println("        wg.Wait()")
    fmt.Println("        close(results)")
    fmt.Println("    }()")
    fmt.Println()
    fmt.Println("    // Processar resultados")
    fmt.Println("    for result := range results {")
    fmt.Println("        fmt.Printf(\"Resultado recebido: %d\\n\", result)")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a usar WaitGroup, canais e contexto para gerenciar goroutines em Go. Essas ferramentas permitem uma sincronização e comunicação eficaz entre goroutines, facilitando o desenvolvimento de programas concorrentes robustos." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use `sync.WaitGroup` para aguardar a conclusão de várias goroutines.",
        "Curiosidade: Canais (`chan`) são usados para comunicação segura entre goroutines.",
        "Dica: Use o pacote `context` para controlar o ciclo de vida de goroutines.",
        "Curiosidade: `context.WithTimeout` permite definir um limite de tempo para operações concorrentes.",
        "Dica: Feche canais para sinalizar que nenhuma nova mensagem será enviada.",
        "Curiosidade: Concorrência em Go é inspirada pelo modelo de comunicação sequencial de processos (CSP).",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o décimo oitavo nível!" + string(reset))
}
