package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
    "sync"
		"math/rand"
    "time"
)

type Level41 struct{}

func StartLevel41() {
    game.StartLevel(Level41{})
}

func (l Level41) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 41: Trabalhando com Tarefas Assíncronas e Go Routines" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Aprenda a gerenciar tarefas assíncronas usando Go Routines e WaitGroups" + string(reset))
    fmt.Println()

    // Criar WaitGroup
    var wg sync.WaitGroup

    // Iniciar múltiplas Go Routines
    wg.Add(3)
    go task("Tarefa 1", 1, &wg)
    go task("Tarefa 2", 2, &wg)
    go task("Tarefa 3", 3, &wg)

    // Aguardar todas as Go Routines terminarem
    wg.Wait()

    fmt.Println(string(green), "Todas as tarefas foram concluídas!", string(reset))

    showAsyncSummary()
}

func task(name string, delay time.Duration, wg *sync.WaitGroup) {
    defer wg.Done()
    time.Sleep(delay * time.Second)
    fmt.Printf("Tarefa %s concluída após %d segundos\n", name, delay)
}

func showAsyncSummary() {
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
    fmt.Println("    \"sync\"")
    fmt.Println("    \"time\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    var wg sync.WaitGroup")
    fmt.Println()
    fmt.Println("    wg.Add(3)")
    fmt.Println("    go task(\"Tarefa 1\", 1, &wg)")
    fmt.Println("    go task(\"Tarefa 2\", 2, &wg)")
    fmt.Println("    go task(\"Tarefa 3\", 3, &wg)")
    fmt.Println()
    fmt.Println("    wg.Wait()")
    fmt.Println("    fmt.Println(\"Todas as tarefas foram concluídas!\")")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func task(name string, delay time.Duration, wg *sync.WaitGroup) {")
    fmt.Println("    defer wg.Done()")
    fmt.Println("    time.Sleep(delay * time.Second)")
    fmt.Println("    fmt.Printf(\"Tarefa %s concluída após %d segundos\\n\", name, delay)")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a gerenciar tarefas assíncronas usando Go Routines e WaitGroups para sincronizar a execução de múltiplas tarefas de forma eficiente." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use Go Routines para executar tarefas de forma concorrente.",
        "Curiosidade: Go Routines são mais leves que threads do sistema operacional.",
        "Dica: Use `sync.WaitGroup` para aguardar a conclusão de múltiplas Go Routines.",
        "Curiosidade: Go Routines permitem um alto grau de paralelismo em aplicações Go.",
        "Dica: Sempre use `defer wg.Done()` para garantir que o contador do WaitGroup seja decrementado.",
        "Curiosidade: O uso eficiente de Go Routines pode melhorar a performance e a capacidade de resposta de suas aplicações.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o quadragésimo primeiro nível!" + string(reset))
}
