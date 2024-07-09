package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
		"math/rand"
		"time"
    //"testing"
)

type Level26 struct{}

func StartLevel26() {
    game.StartLevel(Level26{})
}

func (l Level26) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 26: Testes Unitários e de Integração com Go" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Escreva testes unitários e de integração para verificar o comportamento de funções e componentes" + string(reset))
    fmt.Println()

    // Exemplo de função a ser testada
    sum := func(a, b int) int {
        return a + b
    }

    // Exemplo de função a ser testada
    multiply := func(a, b int) int {
        return a * b
    }

    // Exemplo de testes unitários
    fmt.Println(string(green) + "Executando testes unitários:" + string(reset))

    var tests = []struct {
        name string
        a, b int
        want int
    }{
        {"TestSum1", 1, 2, 3},
        {"TestSum2", 2, 3, 5},
        {"TestMultiply1", 2, 3, 6},
        {"TestMultiply2", 3, 4, 12},
    }

    for _, tt := range tests {
        if tt.name == "TestSum1" || tt.name == "TestSum2" {
            if got := sum(tt.a, tt.b); got != tt.want {
                fmt.Printf("sum(%d, %d) = %d; want %d\n", tt.a, tt.b, got, tt.want)
            } else {
                fmt.Printf("%s passed\n", tt.name)
            }
        }
        if tt.name == "TestMultiply1" || tt.name == "TestMultiply2" {
            if got := multiply(tt.a, tt.b); got != tt.want {
                fmt.Printf("multiply(%d, %d) = %d; want %d\n", tt.a, tt.b, got, tt.want)
            } else {
                fmt.Printf("%s passed\n", tt.name)
            }
        }
    }

    // Exemplo de teste de integração
    fmt.Println(string(green) + "Executando teste de integração:" + string(reset))

    integrationTest := func() bool {
        a, b := 2, 3
        sumResult := sum(a, b)
        multiplyResult := multiply(a, b)
        return sumResult == 5 && multiplyResult == 6
    }

    if integrationTest() {
        fmt.Println("Teste de integração passou")
    } else {
        fmt.Println("Teste de integração falhou")
    }

    showTestSummary()
}

func showTestSummary() {
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
    fmt.Println("    \"testing\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func sum(a, b int) int {")
    fmt.Println("    return a + b")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func multiply(a, b int) int {")
    fmt.Println("    return a * b")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func TestSum(t *testing.T) {")
    fmt.Println("    var tests = []struct {")
    fmt.Println("        a, b int")
    fmt.Println("        want int")
    fmt.Println("    }{")
    fmt.Println("        {1, 2, 3},")
    fmt.Println("        {2, 3, 5},")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    for _, tt := range tests {")
    fmt.Println("        if got := sum(tt.a, tt.b); got != tt.want {")
    fmt.Println("            t.Errorf(\"sum(%d, %d) = %d; want %d\", tt.a, tt.b, got, tt.want)")
    fmt.Println("        }")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func TestMultiply(t *testing.T) {")
    fmt.Println("    var tests = []struct {")
    fmt.Println("        a, b int")
    fmt.Println("        want int")
    fmt.Println("    }{")
    fmt.Println("        {2, 3, 6},")
    fmt.Println("        {3, 4, 12},")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    for _, tt := range tests {")
    fmt.Println("        if got := multiply(tt.a, tt.b); got != tt.want {")
    fmt.Println("            t.Errorf(\"multiply(%d, %d) = %d; want %d\", tt.a, tt.b, got, tt.want)")
    fmt.Println("        }")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func TestIntegration(t *testing.T) {")
    fmt.Println("    a, b := 2, 3")
    fmt.Println("    sumResult := sum(a, b)")
    fmt.Println("    multiplyResult := multiply(a, b)")
    fmt.Println("    if sumResult != 5 || multiplyResult != 6 {")
    fmt.Println("        t.Errorf(\"integration test failed: sum(%d, %d) = %d; multiply(%d, %d) = %d\", a, b, sumResult, a, b, multiplyResult)")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    fmt.Println(\"Executando testes unitários e de integração\")")
    fmt.Println("    testing.M{}")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a escrever testes unitários e de integração em Go. Testes unitários verificam o comportamento de funções individuais, enquanto testes de integração verificam o comportamento de múltiplos componentes juntos." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use o pacote `testing` para escrever testes em Go.",
        "Curiosidade: Testes ajudam a garantir que seu código funcione corretamente e facilitam a manutenção.",
        "Dica: Use `t.Errorf` para relatar falhas em testes.",
        "Curiosidade: Testes de integração verificam se diferentes partes do seu sistema funcionam bem juntas.",
        "Dica: Execute seus testes com `go test` no terminal.",
        "Curiosidade: Escrever testes pode ajudar a identificar problemas de design no seu código.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o vigésimo sexto nível!" + string(reset))
}
