package levels

import (
    "bufio"
    "fmt"
    "go-code-quest/pkg/game"
    "os"
    "strconv"
    "strings"
    "math/rand"
    "time"
)

type Level3 struct{}

func StartLevel3() {
    game.StartLevel(Level3{})
}

func (l Level3) Start() {
    // Definir cores
    red := "\033[31m"
    green := "\033[32m"
    yellow := "\033[33m"
    blue := "\033[34m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 3: Funções" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Crie funções para calcular a soma e a média de uma série de números" + string(reset))
    fmt.Println()

    reader := bufio.NewReader(os.Stdin)

    // Solicitar uma série de números do jogador
    fmt.Print(string(blue) + "Digite uma série de números separados por espaço: " + string(reset))
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    numStrings := strings.Split(input, " ")
    numbers := make([]int, len(numStrings))

    for i, ns := range numStrings {
        n, err := strconv.Atoi(ns)
        if err != nil {
            fmt.Println(string(red) + "Entrada inválida, usando valor padrão: 0" + string(reset))
            n = 0
        }
        numbers[i] = n
    }

    // Calcular soma e média usando funções
    sum := calculateSum(numbers)
    avg := calculateAverage(numbers)

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resultados" + string(reset))
    fmt.Println(separator)
    fmt.Printf(string(green)+"Soma dos números: %d\n"+string(reset), sum)
    fmt.Printf(string(green)+"Média dos números: %.2f\n"+string(reset), avg)

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Exemplo de solução em Go" + string(reset))
    fmt.Println(separator)
    fmt.Println("```go")
    fmt.Println("package main")
    fmt.Println()
    fmt.Println("import (")
    fmt.Println("    \"fmt\"")
    fmt.Println("    \"strings\"")
    fmt.Println("    \"strconv\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    // Exemplo de entrada como string")
    fmt.Println("    input := \"1 2 3 4 5\"")
    fmt.Println("    // Dividir a entrada em substrings usando espaço como delimitador")
    fmt.Println("    numStrings := strings.Split(input, \" \")")
    fmt.Println("    // Criar um slice de inteiros com o mesmo tamanho das substrings")
    fmt.Println("    numbers := make([]int, len(numStrings))")
    fmt.Println()
    fmt.Println("    // Converter cada substring para inteiro e armazenar no slice")
    fmt.Println("    for i, ns := range numStrings {")
    fmt.Println("        n, err := strconv.Atoi(ns)")
    fmt.Println("        if err != nil {")
    fmt.Println("            n = 0 // Usar 0 se a conversão falhar")
    fmt.Println("        }")
    fmt.Println("        numbers[i] = n")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    // Calcular soma e média usando funções")
    fmt.Println("    sum := calculateSum(numbers)")
    fmt.Println("    avg := calculateAverage(numbers)")
    fmt.Println()
    fmt.Println("    // Exibir resultados")
    fmt.Println("    fmt.Printf(\"Soma dos números: %d\\n\", sum)")
    fmt.Println("    fmt.Printf(\"Média dos números: %.2f\\n\", avg)")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("// Função para calcular a soma de um slice de inteiros")
    fmt.Println("func calculateSum(numbers []int) int {")
    fmt.Println("    sum := 0")
    fmt.Println("    for _, num := range numbers {")
    fmt.Println("        sum += num")
    fmt.Println("    }")
    fmt.Println("    return sum")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("// Função para calcular a média de um slice de inteiros")
    fmt.Println("func calculateAverage(numbers []int) float64 {")
    fmt.Println("    sum := calculateSum(numbers)")
    fmt.Println("    return float64(sum) / float64(len(numbers))")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a criar e usar funções em Go. Funções permitem encapsular blocos de código reutilizáveis e tornar o código mais modular e legível. Você também aprendeu a passar argumentos e retornar valores de funções." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Funções em Go podem retornar múltiplos valores, o que é útil para retornar resultados e erros.",
        "Curiosidade: Funções em Go são cidadãs de primeira classe, o que significa que podem ser atribuídas a variáveis e passadas como argumentos.",
        "Dica: Use funções para dividir seu código em partes menores e mais gerenciáveis.",
        "Curiosidade: O nome 'Golang' vem do nome do domínio 'golang.org', mas a linguagem é oficialmente chamada de 'Go'.",
        "Dica: Funções anônimas são funções sem nome e podem ser definidas dentro de outras funções.",
        "Curiosidade: O design simples e eficiente de Go permite a criação de funções de maneira intuitiva.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o terceiro nível!" + string(reset))
}

func calculateSum(numbers []int) int {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return sum
}

func calculateAverage(numbers []int) float64 {
    sum := calculateSum(numbers)
    return float64(sum) / float64(len(numbers))
}
