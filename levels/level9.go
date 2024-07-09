package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
    "testing"
    "math/rand"
    "time"
)

type Level9 struct{}

func StartLevel9() {
    game.StartLevel(Level9{})
}

func (l Level9) Start() {
    // Definir cores
    red := "\033[31m"
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 9: Testes" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Escreva testes unitários para uma função simples" + string(reset))
    fmt.Println()

    // Exemplo de função a ser testada
    fmt.Println("Rodando testes...")

    // Chamar a função de teste diretamente
    result := runTests()

    if result {
        fmt.Println(string(green) + "Todos os testes passaram!" + string(reset))
    } else {
        fmt.Println(string(red) + "Alguns testes falharam." + string(reset))
    }

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Exemplo de solução em Go" + string(reset))
    fmt.Println(separator)
    fmt.Println("```go")
    fmt.Println("package main")
    fmt.Println()
    fmt.Println("import \"testing\"")
    fmt.Println()
    fmt.Println("// Função a ser testada")
    fmt.Println("func Add(a, b int) int {")
    fmt.Println("    return a + b")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("// Teste para a função Add")
    fmt.Println("func TestAdd(t *testing.T) {")
    fmt.Println("    result := Add(2, 3)")
    fmt.Println("    if result != 5 {")
    fmt.Println("        t.Errorf(\"Add(2, 3) = %d; want 5\", result)")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("// Função principal para rodar os testes")
    fmt.Println("func main() {")
    fmt.Println("    testing.Main(func(pat, str string) (bool, error) { return true, nil },")
    fmt.Println("        []testing.InternalTest{{\"TestAdd\", TestAdd}},")
    fmt.Println("        nil, nil)")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a escrever testes unitários em Go. Testes são essenciais para garantir que seu código funcione conforme o esperado e para prevenir a introdução de bugs." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use o pacote `testing` para escrever testes unitários em Go.",
        "Curiosidade: Testes ajudam a manter a qualidade do código e facilitam a refatoração.",
        "Dica: Utilize `t.Errorf` para relatar falhas em testes.",
        "Curiosidade: Go possui um suporte robusto para testes, incluindo benchmarks e testes de concorrência.",
        "Dica: Escreva testes pequenos e focados que testem uma única funcionalidade.",
        "Curiosidade: Testes automatizados são uma prática recomendada no desenvolvimento de software para garantir a confiabilidade do código.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o nono nível!" + string(reset))
}

// Função a ser testada
func Add(a, b int) int {
    return a + b
}

// Teste para a função Add
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}

// Função para rodar os testes e retornar um booleano indicando sucesso ou falha
func runTests() bool {
    tests := []struct {
        name string
        test func(*testing.T)
    }{
        {"TestAdd", TestAdd},
    }

    allPassed := true

    for _, tt := range tests {
        t := &testing.T{}
        tt.test(t)
        if t.Failed() {
            allPassed = false
        }
    }

    return allPassed
}
