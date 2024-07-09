package levels

import (
    "bufio"
    "fmt"
    "go-code-quest/pkg/game"
    "math/rand"
    "os"
    "strconv"
    "time"
)

type Level1 struct{}

func StartLevel1() {
    game.StartLevel(Level1{})
}

func (l Level1) Start() {
    // Definir cores
    red := "\033[31m"
    green := "\033[32m"
    yellow := "\033[33m"
    blue := "\033[34m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "=============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 1: Variáveis e Tipos de Dados" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Declarar e imprimir variáveis de diferentes tipos" + string(reset))
    fmt.Println()

    reader := bufio.NewReader(os.Stdin)

    // Solicitar e ler um inteiro do jogador
    fmt.Print(string(blue) + "Digite um número inteiro: " + string(reset))
    intInput, _ := reader.ReadString('\n')
    intInput = intInput[:len(intInput)-1]
    i, err := strconv.Atoi(intInput)
    if err != nil {
        fmt.Println(string(red) + "Entrada inválida para inteiro. Usando valor padrão: 10" + string(reset))
        i = 10
    }

    // Solicitar e ler uma string do jogador
    fmt.Print(string(blue) + "Digite uma string: " + string(reset))
    s, _ := reader.ReadString('\n')
    s = s[:len(s)-1]

    // Solicitar e ler um booleano do jogador
    fmt.Print(string(blue) + "Digite um valor booleano (true/false): " + string(reset))
    boolInput, _ := reader.ReadString('\n')
    boolInput = boolInput[:len(boolInput)-1]
    b, err := strconv.ParseBool(boolInput)
    if err != nil {
        fmt.Println(string(red) + "Entrada inválida para booleano. Usando valor padrão: true" + string(reset))
        b = true
    }

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resultados" + string(reset))
    fmt.Println(separator)

    // Exibir os valores inseridos
    fmt.Printf(string(green)+"Inteiro: %d\n"+string(reset), i)
    fmt.Printf(string(green)+"String: %s\n"+string(reset), s)
    fmt.Printf(string(green)+"Booleano: %t\n"+string(reset), b)

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Exemplo de solução em Go" + string(reset))
    fmt.Println(separator)
    fmt.Println("```go")
    fmt.Printf("var i int = %d\n", i)
    fmt.Printf("var s string = \"%s\"\n", s)
    fmt.Printf("var b bool = %t\n", b)
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu sobre variáveis em Go. O tipo int armazena números inteiros, string armazena sequências de caracteres, e bool armazena valores booleanos (true/false). Golang facilita a manipulação dessas variáveis com uma sintaxe clara e eficiente." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Em Go, as variáveis não utilizadas causam um erro de compilação, o que ajuda a manter o código limpo e eficiente.",
        "Curiosidade: Go foi desenvolvido pelo Google para melhorar a produtividade na programação.",
        "Dica: Use `:=` para declarar e inicializar variáveis curtas, uma característica única e conveniente do Go.",
        "Curiosidade: Go tem coleta de lixo automática, o que facilita o gerenciamento de memória.",
        "Dica: As variáveis em Go têm valores zero padrão, como 0 para inteiros e \"\" para strings.",
        "Curiosidade: A mascote do Go é um gopher, um pequeno roedor, que se tornou um símbolo popular entre os desenvolvedores de Go.",
        "Dica: Use `var` para declarar variáveis de escopo mais amplo, como no nível de pacote.",
        "Curiosidade: Go compila o código extremamente rápido, tornando o desenvolvimento mais eficiente.",
        "Dica: Use `const` para declarar constantes, que são imutáveis e podem melhorar a legibilidade do código.",
        "Curiosidade: Go foi criado por Robert Griesemer, Rob Pike, e Ken Thompson, figuras influentes no mundo da computação.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o primeiro nível!" + string(reset))
}
