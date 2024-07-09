package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
    "math/rand"
    "sync"
    "time"
)

type Level13 struct{}

func StartLevel13() {
    game.StartLevel(Level13{})
}

func (l Level13) Start() {
    // Definir cores
    //red := "\033[31m"
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 13: Concorrência Avançada" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Crie goroutines, use canais para comunicação e um grupo de espera para sincronização" + string(reset))
    fmt.Println()

    // Criar um canal para comunicação entre goroutines
    results := make(chan int, 5)

    // Usar WaitGroup para sincronizar goroutines
    var wg sync.WaitGroup

    // Função que será executada concorrentemente
    worker := func(id int) {
        defer wg.Done()
        sleepTime := rand.Intn(5) + 1
        time.Sleep(time.Duration(sleepTime) * time.Second)
        fmt.Println(string(green) + fmt.Sprintf("Worker %d completou sua tarefa em %d segundos", id, sleepTime) + string(reset))
        results <- id
    }

    // Iniciar 5 goroutines
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i)
    }

    // Aguardar todas as goroutines terminarem
    go func() {
        wg.Wait()
        close(results)
    }()

    // Processar resultados
    for result := range results {
        fmt.Println(string(green) + fmt.Sprintf("Resultado recebido do worker %d", result) + string(reset))
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
    fmt.Println("    \"math/rand\"")
    fmt.Println("    \"sync\"")
    fmt.Println("    \"time\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    // Criar um canal para comunicação entre goroutines")
    fmt.Println("    results := make(chan int, 5)")
    fmt.Println()
    fmt.Println("    // Usar WaitGroup para sincronizar goroutines")
    fmt.Println("    var wg sync.WaitGroup")
    fmt.Println()
    fmt.Println("    // Função que será executada concorrentemente")
    fmt.Println("    worker := func(id int) {")
    fmt.Println("        defer wg.Done()")
    fmt.Println("        sleepTime := rand.Intn(5) + 1")
    fmt.Println("        time.Sleep(time.Duration(sleepTime) * time.Second)")
    fmt.Println("        fmt.Printf(\"Worker %d completou sua tarefa em %d segundos\\n\", id, sleepTime)")
    fmt.Println("        results <- id")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    // Iniciar 5 goroutines")
    fmt.Println("    for i := 1; i <= 5; i++ {")
    fmt.Println("        wg.Add(1)")
    fmt.Println("        go worker(i)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    // Aguardar todas as goroutines terminarem")
    fmt.Println("    go func() {")
    fmt.Println("        wg.Wait()")
    fmt.Println("        close(results)")
    fmt.Println("    }()")
    fmt.Println()
    fmt.Println("    // Processar resultados")
    fmt.Println("    for result := range results {")
    fmt.Println("        fmt.Printf(\"Resultado recebido do worker %d\\n\", result)")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a criar goroutines, usar canais para comunicação e sincronizar tarefas concorrentes usando WaitGroup. Essas técnicas permitem a execução eficiente de tarefas paralelas em Go." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use `sync.WaitGroup` para aguardar a conclusão de múltiplas goroutines.",
        "Curiosidade: WaitGroup é uma maneira eficiente de sincronizar goroutines.",
        "Dica: Use `chan` para comunicação segura entre goroutines.",
        "Curiosidade: Canais e goroutines facilitam a construção de sistemas concorrentes em Go.",
        "Dica: Use `defer wg.Done()` no início da goroutine para garantir que o contador do WaitGroup seja decrementado.",
        "Curiosidade: Go foi projetado com concorrência em mente, tornando-a fácil e eficiente de implementar.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o décimo terceiro nível!" + string(reset))
}
