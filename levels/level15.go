package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
		"math/rand"
    "os"
    "os/exec"
		"time"
)

type Level15 struct{}

func StartLevel15() {
    game.StartLevel(Level15{})
}

func (l Level15) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 15: Build e Execução de Programas" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Construa e execute um programa Go" + string(reset))
    fmt.Println()

    // Exibir instruções para o jogador
    fmt.Println(string(green) + "1. Crie um arquivo chamado `hello.go` com o seguinte conteúdo:" + string(reset))
    fmt.Println("```go")
    fmt.Println("package main")
    fmt.Println()
    fmt.Println("import \"fmt\"")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    fmt.Println(\"Hello, Go!\")")
    fmt.Println("}")
    fmt.Println("```")
    fmt.Println()
    fmt.Println(string(green) + "2. Compile o programa usando o comando:" + string(reset))
    fmt.Println("```bash")
    fmt.Println("go build hello.go")
    fmt.Println("```")
    fmt.Println()
    fmt.Println(string(green) + "3. Execute o programa compilado usando o comando:" + string(reset))
    fmt.Println("```bash")
    fmt.Println("./hello")
    fmt.Println("```")
    fmt.Println()
    fmt.Println(string(green) + "4. Ou execute o programa diretamente usando o comando:" + string(reset))
    fmt.Println("```bash")
    fmt.Println("go run hello.go")
    fmt.Println("```")
    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Exemplo de solução em Go" + string(reset))
    fmt.Println(separator)
    fmt.Println("```go")
    fmt.Println("package main")
    fmt.Println()
    fmt.Println("import \"fmt\"")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    fmt.Println(\"Hello, Go!\")")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a compilar e executar programas Go. A compilação transforma o código fonte em um executável, enquanto `go run` compila e executa o código em um único passo." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use `go build` para compilar seu programa Go em um executável.",
        "Curiosidade: `go run` compila e executa o código fonte em um único passo.",
        "Dica: Use `go build -o <nome_do_arquivo>` para especificar o nome do arquivo executável.",
        "Curiosidade: Go gera executáveis estáticos por padrão, que não dependem de bibliotecas externas.",
        "Dica: O comando `go install` compila e instala o pacote especificado em `$GOPATH/bin`.",
        "Curiosidade: Go foi projetado para ser simples e eficiente na construção de programas compilados.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Criar o arquivo `hello.go` para demonstrar
    createDemoFile()
}

func createDemoFile() {
    demoCode := `package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}`

    file, err := os.Create("hello.go")
    if err != nil {
        fmt.Println("Erro ao criar arquivo de demonstração:", err)
        return
    }
    defer file.Close()

    _, err = file.WriteString(demoCode)
    if err != nil {
        fmt.Println("Erro ao escrever no arquivo de demonstração:", err)
    } else {
        fmt.Println("Arquivo de demonstração `hello.go` criado com sucesso.")
    }

    // Compilar o arquivo `hello.go`
    cmd := exec.Command("go", "build", "hello.go")
    if err := cmd.Run(); err != nil {
        fmt.Println("Erro ao compilar o arquivo de demonstração:", err)
    } else {
        fmt.Println("Arquivo `hello.go` compilado com sucesso.")
    }

    // Executar o arquivo compilado
    fmt.Println("Executando o arquivo compilado:")
    cmd = exec.Command("./hello")
    output, err := cmd.Output()
    if err != nil {
        fmt.Println("Erro ao executar o arquivo compilado:", err)
    } else {
        fmt.Println(string(output))
    }
}
