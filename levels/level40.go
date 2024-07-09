package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
		"math/rand"
    "time"
    "sync"
)

type Level40 struct{}

func StartLevel40() {
    game.StartLevel(Level40{})
}

func (l Level40) Start() {
    // Definir cores
		red := "\033[31m"
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 40: Padrão de Projeto Singleton em Go" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Aprenda a implementar o padrão de projeto Singleton em Go" + string(reset))
    fmt.Println()

    // Testar o Singleton
    instance1 := GetInstance()
    instance2 := GetInstance()

    fmt.Println(string(green), "Instância 1:", instance1, string(reset))
    fmt.Println(string(green), "Instância 2:", instance2, string(reset))

    if instance1 == instance2 {
        fmt.Println(string(green), "Instância única garantida!", string(reset))
    } else {
        fmt.Println(string(red), "Falha ao garantir instância única.", string(reset))
    }

    showSingletonSummary()
}

var instance *Singleton
var once sync.Once

// Singleton struct
type Singleton struct{}

// GetInstance retorna a única instância do Singleton
func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{}
    })
    return instance
}

func showSingletonSummary() {
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
    fmt.Println(")")
    fmt.Println()
    fmt.Println("var instance *Singleton")
    fmt.Println("var once sync.Once")
    fmt.Println()
    fmt.Println("// Singleton struct")
    fmt.Println("type Singleton struct{}")
    fmt.Println()
    fmt.Println("// GetInstance retorna a única instância do Singleton")
    fmt.Println("func GetInstance() *Singleton {")
    fmt.Println("    once.Do(func() {")
    fmt.Println("        instance = &Singleton{}")
    fmt.Println("    })")
    fmt.Println("    return instance")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    instance1 := GetInstance()")
    fmt.Println("    instance2 := GetInstance()")
    fmt.Println()
    fmt.Println("    fmt.Println(\"Instância 1:\", instance1)")
    fmt.Println("    fmt.Println(\"Instância 2:\", instance2)")
    fmt.Println()
    fmt.Println("    if instance1 == instance2 {")
    fmt.Println("        fmt.Println(\"Instância única garantida!\")")
    fmt.Println("    } else {")
    fmt.Println("        fmt.Println(\"Falha ao garantir instância única.\")")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a implementar o padrão de projeto Singleton em Go, garantindo que uma classe tenha apenas uma instância e fornecendo um ponto global de acesso a essa instância." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use o padrão Singleton para garantir que apenas uma instância de uma classe seja criada.",
        "Curiosidade: O Singleton é útil em cenários onde é necessário um ponto global de acesso, como gerenciadores de configuração ou conexão de banco de dados.",
        "Dica: Em Go, use o pacote `sync` com `sync.Once` para garantir a criação de uma única instância de forma thread-safe.",
        "Curiosidade: O padrão Singleton é amplamente utilizado em linguagens orientadas a objetos, como Java e C++.",
        "Dica: Evite usar Singletons em excesso, pois podem introduzir dependências globais e dificultar o teste do código.",
        "Curiosidade: O Singleton é um dos padrões de projeto do livro \"Design Patterns: Elements of Reusable Object-Oriented Software\" dos autores Erich Gamma, Richard Helm, Ralph Johnson e John Vlissides.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o quadragésimo nível!" + string(reset))
}
