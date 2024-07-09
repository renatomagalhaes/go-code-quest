package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
		"math/rand"
    "time"
    "log"
    "plugin"
)

type Level38 struct{}

func StartLevel38() {
    game.StartLevel(Level38{})
}

func (l Level38) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 38: Utilizando Plugins Dinâmicos em Go" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Aprenda a criar e utilizar plugins dinâmicos em Go" + string(reset))
    fmt.Println()

    // Carregar o plugin
    plug, err := plugin.Open("plugin38.so")
    if err != nil {
        log.Fatalf("Erro ao carregar plugin: %v", err)
    }

    // Procurar o símbolo (uma função neste caso)
    symFunc, err := plug.Lookup("PluginFunction")
    if err != nil {
        log.Fatalf("Erro ao procurar função no plugin: %v", err)
    }

    // Assegurar que o símbolo é do tipo correto
    pluginFunc, ok := symFunc.(func() string)
    if !ok {
        log.Fatalf("Símbolo tem tipo inesperado: %T", symFunc)
    }

    // Usar a função do plugin
    result := pluginFunc()
    fmt.Println(string(green) + "Resultado da função do plugin: " + result + string(reset))

    showPluginSummary()
}

func showPluginSummary() {
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
    fmt.Println("    \"plugin\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    plug, err := plugin.Open(\"plugin38.so\")")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"Erro ao carregar plugin: %v\", err)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    symFunc, err := plug.Lookup(\"PluginFunction\")")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"Erro ao procurar função no plugin: %v\", err)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    pluginFunc, ok := symFunc.(func() string)")
    fmt.Println("    if !ok {")
    fmt.Println("        log.Fatalf(\"Símbolo tem tipo inesperado: %T\", symFunc)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    result := pluginFunc()")
    fmt.Println("    fmt.Println(\"Resultado da função do plugin: \" + result)")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a criar e utilizar plugins dinâmicos em Go, permitindo a extensão de funcionalidades em tempo de execução sem recompilar a aplicação principal." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use plugins para adicionar funcionalidades de forma modular.",
        "Curiosidade: Plugins são carregados em tempo de execução, permitindo a extensão da aplicação sem recompilação.",
        "Dica: Use `plugin.Open` para carregar plugins dinâmicos em Go.",
        "Curiosidade: Plugins são específicos do sistema operacional e arquitetura para os quais são compilados.",
        "Dica: Use o comando `go build -buildmode=plugin -o nome_do_plugin.so` para compilar plugins.",
        "Curiosidade: Plugins podem conter funções, variáveis e tipos que podem ser usados na aplicação principal.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o trigésimo oitavo nível!" + string(reset))
}
