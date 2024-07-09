package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
	"math/rand"
	"time"
    "log"
    "net/http"
    "os"
    "runtime/pprof"
)

type Level27 struct{}

func StartLevel27() {
    game.StartLevel(Level27{})
}

func (l Level27) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 27: Profiling e Otimização de Código" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Use ferramentas de profiling para identificar e otimizar gargalos de desempenho" + string(reset))
    fmt.Println()

    // Iniciar servidor HTTP para profiling
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()

    // Exemplo de código para profiling
    fmt.Println(string(green) + "Executando código de exemplo para profiling:" + string(reset))

    f, err := os.Create("cpu.prof")
    if err != nil {
        log.Fatal(err)
    }
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()

    for i := 0; i < 1000000; i++ {
        _ = fibonacci(30)
    }

    showProfilingSummary()
}

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func showProfilingSummary() {
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
    fmt.Println("    \"log\"")
    fmt.Println("    \"net/http\"")
    fmt.Println("    \"os\"")
    fmt.Println("    \"runtime/pprof\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    go func() {")
    fmt.Println("        log.Println(http.ListenAndServe(\"localhost:6060\", nil))")
    fmt.Println("    }()")
    fmt.Println()
    fmt.Println("    f, err := os.Create(\"cpu.prof\")")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatal(err)")
    fmt.Println("    }")
    fmt.Println("    pprof.StartCPUProfile(f)")
    fmt.Println("    defer pprof.StopCPUProfile()")
    fmt.Println()
    fmt.Println("    for i := 0; i < 1000000; i++ {")
    fmt.Println("        _ = fibonacci(30)")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func fibonacci(n int) int {")
    fmt.Println("    if n <= 1 {")
    fmt.Println("        return n")
    fmt.Println("    }")
    fmt.Println("    return fibonacci(n-1) + fibonacci(n-2)")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a usar o pacote pprof para fazer profiling de CPU e memória em Go. Isso ajuda a identificar gargalos de desempenho e a otimizar seu código." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use o pacote `runtime/pprof` para fazer profiling de CPU e memória em Go.",
        "Curiosidade: O profiling ajuda a identificar funções que consomem muito tempo ou memória.",
        "Dica: Use ferramentas como `go tool pprof` para analisar os resultados do profiling.",
        "Curiosidade: Profiling pode revelar gargalos de desempenho que não são óbvios apenas olhando o código.",
        "Dica: Sempre pare o profiling com `pprof.StopCPUProfile` para garantir que os dados sejam salvos corretamente.",
        "Curiosidade: Você pode fazer profiling de memória com `pprof.WriteHeapProfile`.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o vigésimo sétimo nível!" + string(reset))
}
