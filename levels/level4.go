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

type Level4 struct{}

func StartLevel4() {
    game.StartLevel(Level4{})
}

func (l Level4) Start() {
    // Definir cores
    red := "\033[31m"
    green := "\033[32m"
    yellow := "\033[33m"
    blue := "\033[34m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 4: Estruturas de Dados" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Manipule arrays, slices e mapas para resolver o problema abaixo" + string(reset))
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

    // Calcular a soma e criar um mapa de contagem de ocorrências
    sum := 0
    occurrences := make(map[int]int)
    for _, num := range numbers {
        sum += num
        occurrences[num]++
    }

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resultados" + string(reset))
    fmt.Println(separator)
    fmt.Printf(string(green)+"Soma dos números: %d\n"+string(reset), sum)
    fmt.Println(string(green) + "Contagem de ocorrências:" + string(reset))
    for num, count := range occurrences {
        fmt.Printf(string(green)+"Número %d: %d vezes\n"+string(reset), num, count)
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
    fmt.Println("    // Inicializar soma e mapa de contagem de ocorrências")
    fmt.Println("    sum := 0")
    fmt.Println("    occurrences := make(map[int]int)")
    fmt.Println("    // Calcular a soma e a contagem de ocorrências")
    fmt.Println("    for _, num := range numbers {")
    fmt.Println("        sum += num")
    fmt.Println("        occurrences[num]++")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    // Exibir resultados")
    fmt.Println("    fmt.Printf(\"Soma dos números: %d\\n\", sum)")
    fmt.Println("    fmt.Println(\"Contagem de ocorrências:\")")
    fmt.Println("    for num, count := range occurrences {")
    fmt.Println("        fmt.Printf(\"Número %d: %d vezes\\n\", num, count)")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a manipular arrays, slices e mapas em Go. Arrays têm tamanho fixo, slices são mais flexíveis e podem crescer dinamicamente, e mapas são usados para associar chaves a valores." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use slices em vez de arrays para maior flexibilidade ao trabalhar com listas de dados.",
        "Curiosidade: Slices em Go são construídos sobre arrays, proporcionando uma visão mais poderosa e flexível dos dados subjacentes.",
        "Dica: Use o operador make para criar slices e mapas de maneira eficiente.",
        "Curiosidade: Mapas em Go são implementados como tabelas hash, oferecendo operações de leitura e escrita rápidas.",
        "Dica: Aproveite a função append para adicionar elementos a um slice dinamicamente.",
        "Curiosidade: O design simples de Go torna a manipulação de estruturas de dados intuitiva e direta.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o quarto nível!" + string(reset))
}
