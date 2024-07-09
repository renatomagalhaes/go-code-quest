package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
    "math/rand"
    "runtime"
    "sync"
    "sync/atomic"
    "time"
)

type Level14 struct{}

func StartLevel14() {
    game.StartLevel(Level14{})
}

func (l Level14) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 14: Paralelismo" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Use goroutines em conjunto com operações atômicas para executar tarefas paralelas" + string(reset))
    fmt.Println()

    // Definir o número de threads do sistema
    numCPU := runtime.NumCPU()
    runtime.GOMAXPROCS(numCPU)
    fmt.Println(string(green) + fmt.Sprintf("Usando %d CPUs", numCPU) + string(reset))

    // Variável atômica
    var counter int64

    // WaitGroup para sincronizar goroutines
    var wg sync.WaitGroup

    // Função que incrementa o contador atômico
    worker := func(id int) {
        defer wg.Done()
        for i := 0; i < 1000; i++ {
            atomic.AddInt64(&counter, 1)
            time.Sleep(time.Millisecond)
        }
        fmt.Println(string(green) + fmt.Sprintf("Worker %d completou", id) + string(reset))
    }

    // Iniciar várias goroutines
    numGoroutines := 10
    for i := 1; i <= numGoroutines; i++ {
        wg.Add(1)
        go worker(i)
    }

    // Aguardar todas as goroutines terminarem
    wg.Wait()

    // Exibir o valor final do contador
    fmt.Println(string(green) + fmt.Sprintf("Valor final do contador: %d", counter) + string(reset))

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Exemplo de solução em Go" + string(reset))
    fmt.Println(separator)
    fmt.Println("```go")
    fmt.Println("package main")
    fmt.Println()
    fmt.Println("import (")
    fmt.Println("    \"fmt\"")
    fmt.Println("    \"runtime\"")
    fmt.Println("    \"sync\"")
    fmt.Println("    \"sync/atomic\"")
    fmt.Println("    \"time\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    // Definir o número de threads do sistema")
    fmt.Println("    numCPU := runtime.NumCPU()")
    fmt.Println("    runtime.GOMAXPROCS(numCPU)")
    fmt.Println("    fmt.Printf(\"Usando %d CPUs\\n\", numCPU)")
    fmt.Println()
    fmt.Println("    // Variável atômica")
    fmt.Println("    var counter int64")
    fmt.Println()
    fmt.Println("    // WaitGroup para sincronizar goroutines")
    fmt.Println("    var wg sync.WaitGroup")
    fmt.Println()
    fmt.Println("    // Função que incrementa o contador atômico")
    fmt.Println("    worker := func(id int) {")
    fmt.Println("        defer wg.Done()")
    fmt.Println("        for i := 0; i < 1000; i++ {")
    fmt.Println("            atomic.AddInt64(&counter, 1)")
    fmt.Println("            time.Sleep(time.Millisecond)")
    fmt.Println("        }")
    fmt.Println("        fmt.Printf(\"Worker %d completou\\n\", id)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    // Iniciar várias goroutines")
    fmt.Println("    numGoroutines := 10")
    fmt.Println("    for i := 1; i <= numGoroutines; i++ {")
    fmt.Println("        wg.Add(1)")
    fmt.Println("        go worker(i)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    // Aguardar todas as goroutines terminarem")
    fmt.Println("    wg.Wait()")
    fmt.Println()
    fmt.Println("    // Exibir o valor final do contador")
    fmt.Println("    fmt.Printf(\"Valor final do contador: %d\\n\", counter)")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a usar goroutines em conjunto com operações atômicas para executar tarefas paralelas. Essas técnicas permitem a execução eficiente de tarefas em sistemas multi-core." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use `runtime.GOMAXPROCS` para definir o número de threads do sistema.",
        "Curiosidade: Go permite a execução eficiente de tarefas paralelas em sistemas multi-core.",
        "Dica: Use `sync/atomic` para operações atômicas seguras entre goroutines.",
        "Curiosidade: Operações atômicas são essenciais para evitar condições de corrida em programas concorrentes.",
        "Dica: Use `sync.WaitGroup` para aguardar a conclusão de múltiplas goroutines.",
        "Curiosidade: O pacote `sync` fornece várias ferramentas úteis para sincronização em Go.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o décimo quarto nível!" + string(reset))
}
