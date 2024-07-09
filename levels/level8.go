package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
    "io/ioutil"
    "math/rand"
    "time"
)

type Level8 struct{}

func StartLevel8() {
    game.StartLevel(Level8{})
}

func (l Level8) Start() {
    // Definir cores
    red := "\033[31m"
    green := "\033[32m"
    yellow := "\033[33m"
    //blue := "\033[34m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 8: Manipulação de Arquivos" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Leia e escreva em arquivos" + string(reset))
    fmt.Println()

    // Definir nome do arquivo
    filename := "example.txt"

    // Criar e escrever no arquivo
    err := ioutil.WriteFile(filename, []byte("Hello, Go!"), 0644)
    if err != nil {
        fmt.Println(string(red) + "Erro ao escrever no arquivo:" + string(reset), err)
        return
    }
    fmt.Println(string(green) + "Arquivo escrito com sucesso." + string(reset))

    // Ler o conteúdo do arquivo
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(string(red) + "Erro ao ler o arquivo:" + string(reset), err)
        return
    }
    fmt.Println(string(green) + "Conteúdo do arquivo:" + string(reset))
    fmt.Println(string(cyan) + string(data) + string(reset))

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Exemplo de solução em Go" + string(reset))
    fmt.Println(separator)
    fmt.Println("```go")
    fmt.Println("package main")
    fmt.Println()
    fmt.Println("import (")
    fmt.Println("    \"fmt\"")
    fmt.Println("    \"io/ioutil\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    // Definir nome do arquivo")
    fmt.Println("    filename := \"example.txt\"")
    fmt.Println()
    fmt.Println("    // Criar e escrever no arquivo")
    fmt.Println("    err := ioutil.WriteFile(filename, []byte(\"Hello, Go!\"), 0644)")
    fmt.Println("    if err != nil {")
    fmt.Println("        fmt.Println(\"Erro ao escrever no arquivo:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"Arquivo escrito com sucesso.\")")
    fmt.Println()
    fmt.Println("    // Ler o conteúdo do arquivo")
    fmt.Println("    data, err := ioutil.ReadFile(filename)")
    fmt.Println("    if err != nil {")
    fmt.Println("        fmt.Println(\"Erro ao ler o arquivo:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"Conteúdo do arquivo:\")")
    fmt.Println("    fmt.Println(string(data))")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a ler e escrever arquivos em Go. A leitura e escrita de arquivos são operações comuns em muitas aplicações, e Go fornece um suporte robusto para essas operações." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use `ioutil.WriteFile` para escrever dados rapidamente em um arquivo.",
        "Curiosidade: A leitura e escrita de arquivos são operações fundamentais em muitos programas, e Go simplifica essas operações com sua biblioteca padrão.",
        "Dica: Sempre verifique e trate erros ao trabalhar com arquivos para evitar falhas inesperadas.",
        "Curiosidade: Go usa permissões Unix para definir quem pode ler ou escrever em um arquivo.",
        "Dica: Use `ioutil.ReadFile` para ler o conteúdo de um arquivo de forma simples e eficiente.",
        "Curiosidade: Manipulação de arquivos é essencial para operações de entrada e saída em sistemas de arquivos.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o oitavo nível!" + string(reset))
}
