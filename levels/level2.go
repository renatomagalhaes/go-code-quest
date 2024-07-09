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

type Level2 struct{}

func StartLevel2() {
    game.StartLevel(Level2{})
}

func (l Level2) Start() {
    // Definir cores
    red := "\033[31m"
    green := "\033[32m"
    yellow := "\033[33m"
    blue := "\033[34m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 2: Estruturas de Controle" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Use if-else, switch e loops para resolver o problema abaixo" + string(reset))
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

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resultados" + string(reset))
    fmt.Println(separator)

    // Processar os números e exibir resultados usando estruturas de controle
    sum := 0
    evenCount := 0
    oddCount := 0

    for _, num := range numbers {
        sum += num
        if num%2 == 0 {
            evenCount++
        } else {
            oddCount++
        }
    }

    fmt.Printf(string(green)+"Soma dos números: %d\n"+string(reset), sum)
    fmt.Printf(string(green)+"Quantidade de números pares: %d\n"+string(reset), evenCount)
    fmt.Printf(string(green)+"Quantidade de números ímpares: %d\n"+string(reset), oddCount)

    // Exemplo de uso de switch
    fmt.Println()
    fmt.Print(string(blue) + "Escolha um número da lista para verificar sua paridade: " + string(reset))
    choice, _ := reader.ReadString('\n')
    choice = strings.TrimSpace(choice)
    chosenNumber, err := strconv.Atoi(choice)
    if err != nil {
        fmt.Println(string(red) + "Entrada inválida, usando valor padrão: 0" + string(reset))
        chosenNumber = 0
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
    fmt.Println("    // Inicializar soma e contadores de números pares e ímpares")
    fmt.Println("    sum := 0")
    fmt.Println("    evenCount := 0")
    fmt.Println("    oddCount := 0")
    fmt.Println("    // Calcular a soma e contar números pares e ímpares")
    fmt.Println("    for _, num := range numbers {")
    fmt.Println("        sum += num")
    fmt.Println("        if num%2 == 0 {")
    fmt.Println("            evenCount++")
    fmt.Println("        } else {")
    fmt.Println("            oddCount++")
    fmt.Println("        }")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    // Exibir resultados")
    fmt.Println("    fmt.Printf(\"Soma dos números: %d\\n\", sum)")
    fmt.Println("    fmt.Printf(\"Quantidade de números pares: %d\\n\", evenCount)")
    fmt.Println("    fmt.Printf(\"Quantidade de números ímpares: %d\\n\", oddCount)")
    fmt.Println()
    fmt.Println("    // Exemplo de uso de switch para verificar a paridade de um número escolhido")
    fmt.Println("    choice := 3 // Exemplo de escolha")
    fmt.Println("    switch choice {")
    fmt.Println("    case 0:")
    fmt.Println("        fmt.Println(\"O número é zero.\")")
    fmt.Println("    case 1, 3, 5, 7, 9:")
    fmt.Println("        fmt.Println(\"O número é ímpar.\")")
    fmt.Println("    case 2, 4, 6, 8, 10:")
    fmt.Println("        fmt.Println(\"O número é par.\")")
    fmt.Println("    default:")
    fmt.Println("        fmt.Println(\"Número desconhecido.\")")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println("```")

    switch chosenNumber {
    case 0:
        fmt.Println(string(green) + "O número é zero." + string(reset))
    case 1, 3, 5, 7, 9:
        fmt.Println(string(green) + "O número é ímpar." + string(reset))
    case 2, 4, 6, 8, 10:
        fmt.Println(string(green) + "O número é par." + string(reset))
    default:
        fmt.Println(string(green) + "Número desconhecido." + string(reset))
    }

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a usar estruturas de controle em Go. O if-else permite executar blocos de código com base em condições. O switch é uma forma mais limpa de escrever múltiplas condições. O for é usado para criar loops." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: O loop for em Go é bastante flexível e pode ser usado como um while.",
        "Curiosidade: Go não tem um loop while, apenas o for que pode ser usado de várias maneiras.",
        "Dica: Use switch em vez de múltiplos if-else para tornar o código mais legível.",
        "Curiosidade: Go foi projetado para ser simples e eficiente, evitando a complexidade desnecessária.",
        "Dica: Combine condições em um switch usando vírgulas para testar múltiplos casos.",
        "Curiosidade: O design de Go incentiva a escrita de código claro e legível.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o segundo nível!" + string(reset))
}
