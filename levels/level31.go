package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
		"math/rand"
		"time"
    "io/ioutil"
    "log"
    "os"
)

type Level31 struct{}

func StartLevel31() {
    game.StartLevel(Level31{})
}

func (l Level31) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 31: Manipulação de Arquivos e Diretórios" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Aprenda a ler e escrever arquivos, criar e remover diretórios" + string(reset))
    fmt.Println()

    // Criar diretório
    dirName := "example_dir"
    err := os.Mkdir(dirName, 0755)
    if err != nil {
        log.Fatalf("Erro ao criar diretório: %v", err)
    }
    fmt.Println(string(green) + "Diretório criado: " + dirName + string(reset))

    // Criar arquivo
    fileName := dirName + "/example_file.txt"
    content := []byte("Olá, Go Code Quest!")
    err = ioutil.WriteFile(fileName, content, 0644)
    if err != nil {
        log.Fatalf("Erro ao criar arquivo: %v", err)
    }
    fmt.Println(string(green) + "Arquivo criado: " + fileName + string(reset))

    // Ler arquivo
    data, err := ioutil.ReadFile(fileName)
    if err != nil {
        log.Fatalf("Erro ao ler arquivo: %v", err)
    }
    fmt.Println(string(green) + "Conteúdo do arquivo: " + string(data) + string(reset))

    // Remover arquivo
    err = os.Remove(fileName)
    if err != nil {
        log.Fatalf("Erro ao remover arquivo: %v", err)
    }
    fmt.Println(string(green) + "Arquivo removido: " + fileName + string(reset))

    // Remover diretório
    err = os.Remove(dirName)
    if err != nil {
        log.Fatalf("Erro ao remover diretório: %v", err)
    }
    fmt.Println(string(green) + "Diretório removido: " + dirName + string(reset))

    showFileManipulationSummary()
}

func showFileManipulationSummary() {
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
    fmt.Println("    \"io/ioutil\"")
    fmt.Println("    \"log\"")
    fmt.Println("    \"os\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    dirName := \"example_dir\"")
    fmt.Println("    err := os.Mkdir(dirName, 0755)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"Erro ao criar diretório: %v\", err)")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"Diretório criado: \" + dirName)")
    fmt.Println()
    fmt.Println("    fileName := dirName + \"/example_file.txt\"")
    fmt.Println("    content := []byte(\"Olá, Go Code Quest!\")")
    fmt.Println("    err = ioutil.WriteFile(fileName, content, 0644)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"Erro ao criar arquivo: %v\", err)")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"Arquivo criado: \" + fileName)")
    fmt.Println()
    fmt.Println("    data, err := ioutil.ReadFile(fileName)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"Erro ao ler arquivo: %v\", err)")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"Conteúdo do arquivo: \" + string(data))")
    fmt.Println()
    fmt.Println("    err = os.Remove(fileName)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"Erro ao remover arquivo: %v\", err)")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"Arquivo removido: \" + fileName)")
    fmt.Println()
    fmt.Println("    err = os.Remove(dirName)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"Erro ao remover diretório: %v\", err)")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"Diretório removido: \" + dirName)")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a ler e escrever arquivos, criar e remover diretórios usando os pacotes `os` e `io/ioutil` em Go." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use `os.Mkdir` para criar diretórios.",
        "Curiosidade: O pacote `io/ioutil` oferece funções convenientes para ler e escrever arquivos.",
        "Dica: Use `os.Remove` para remover arquivos e diretórios.",
        "Curiosidade: A manipulação de arquivos é uma habilidade essencial para muitas aplicações, como processamento de dados e geração de relatórios.",
        "Dica: Use `ioutil.ReadFile` e `ioutil.WriteFile` para operações de leitura e escrita simples de arquivos.",
        "Curiosidade: O pacote `os` permite manipular permissões de arquivos e diretórios.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o trigésimo primeiro nível!" + string(reset))
}
