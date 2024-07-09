package levels

import (
    "bufio"
    "fmt"
    "go-code-quest/pkg/game"
		"math/rand"
		"time"
    "log"
    "os"
    "strings"
    "text/template"
)

type Level21 struct{}

func StartLevel21() {
    game.StartLevel(Level21{})
}

func (l Level21) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 21: Uso de Templates para Geração de Código" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Use templates em Go para gerar código dinamicamente" + string(reset))
    fmt.Println()

    // Obter informações do usuário
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Digite o nome do pacote: ")
    packageName, _ := reader.ReadString('\n')
    packageName = strings.TrimSpace(packageName)

    fmt.Print("Digite o nome da função: ")
    functionName, _ := reader.ReadString('\n')
    functionName = strings.TrimSpace(functionName)

    fmt.Print("Digite a mensagem que a função deve retornar: ")
    message, _ := reader.ReadString('\n')
    message = strings.TrimSpace(message)

    // Template de código Go
    const goTemplate = `package {{.PackageName}}

import "fmt"

// {{.FunctionName}} retorna uma mensagem específica
func {{.FunctionName}}() string {
    return "{{.Message}}"
}

func main() {
    fmt.Println({{.FunctionName}}())
}
`

    // Estrutura de dados para preencher o template
    data := struct {
        PackageName  string
        FunctionName string
        Message      string
    }{
        PackageName:  packageName,
        FunctionName: functionName,
        Message:      message,
    }

    // Criar e preencher o template
    tmpl, err := template.New("goTemplate").Parse(goTemplate)
    if err != nil {
        log.Fatal("Erro ao criar template:", err)
    }

    // Criar arquivo de saída
    outputFile, err := os.Create("generated_code.go")
    if err != nil {
        log.Fatal("Erro ao criar arquivo:", err)
    }
    defer outputFile.Close()

    // Executar o template e escrever no arquivo
    err = tmpl.Execute(outputFile, data)
    if err != nil {
        log.Fatal("Erro ao executar template:", err)
    }

    fmt.Println(string(green) + "Código gerado com sucesso!" + string(reset))
    fmt.Println("Verifique o arquivo 'generated_code.go' no diretório atual.")

    showTemplateSummary()
}

func showTemplateSummary() {
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
    fmt.Println("    \"os\"")
    fmt.Println("    \"text/template\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    const goTemplate = `package {{.PackageName}}")
    fmt.Println()
    fmt.Println("import \"fmt\"")
    fmt.Println()
    fmt.Println("// {{.FunctionName}} retorna uma mensagem específica")
    fmt.Println("func {{.FunctionName}}() string {")
    fmt.Println("    return \"{{.Message}}\"")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    fmt.Println({{.FunctionName}}())")
    fmt.Println("}`")
    fmt.Println()
    fmt.Println("    data := struct {")
    fmt.Println("        PackageName  string")
    fmt.Println("        FunctionName string")
    fmt.Println("        Message      string")
    fmt.Println("    }{")
    fmt.Println("        PackageName:  \"main\",")
    fmt.Println("        FunctionName: \"HelloWorld\",")
    fmt.Println("        Message:      \"Hello, Go!\",")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    tmpl, err := template.New(\"goTemplate\").Parse(goTemplate)")
    fmt.Println("    if err != nil {")
    fmt.Println("        fmt.Println(\"Erro ao criar template:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    outputFile, err := os.Create(\"generated_code.go\")")
    fmt.Println("    if err != nil {")
    fmt.Println("        fmt.Println(\"Erro ao criar arquivo:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    defer outputFile.Close()")
    fmt.Println()
    fmt.Println("    err = tmpl.Execute(outputFile, data)")
    fmt.Println("    if err != nil {")
    fmt.Println("        fmt.Println(\"Erro ao executar template:\", err)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    fmt.Println(\"Código gerado com sucesso! Verifique o arquivo 'generated_code.go' no diretório atual.\")")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a usar templates em Go para gerar código dinamicamente. Usamos o pacote `text/template` para criar um template que gera uma função Go com base em dados fornecidos pelo usuário." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use o pacote `text/template` para gerar qualquer tipo de texto, incluindo código.",
        "Curiosidade: Templates são uma maneira poderosa de automatizar a geração de código repetitivo.",
        "Dica: Variáveis de template são delimitadas por `{{` e `}}`.",
        "Curiosidade: Você pode usar estruturas de dados para passar múltiplos valores para um template.",
        "Dica: Templates podem ser usados para gerar HTML, XML, e-mail, documentação e muito mais.",
        "Curiosidade: O pacote `html/template` é semelhante a `text/template`, mas é seguro para uso com HTML.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o vigésimo primeiro nível!" + string(reset))
}
