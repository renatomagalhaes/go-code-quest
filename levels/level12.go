package levels

import (
    "encoding/json"
    "fmt"
    "go-code-quest/pkg/game"
    "log"
    "math/rand"
    "time"
)

type Level12 struct{}

func StartLevel12() {
    game.StartLevel(Level12{})
}

func (l Level12) Start() {
    // Definir cores
    red := "\033[31m"
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 12: Manipulação de JSON" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Codifique e decodifique dados JSON" + string(reset))
    fmt.Println()

    // Estrutura de exemplo para codificação JSON
    type User struct {
        Name  string `json:"name"`
        Age   int    `json:"age"`
        Email string `json:"email"`
    }

    // Criar instância da estrutura User
    user := User{
        Name:  "Alice",
        Age:   30,
        Email: "alice@example.com",
    }

    // Codificar a estrutura em JSON
    userJSON, err := json.Marshal(user)
    if err != nil {
        log.Fatal(string(red) + "Erro ao codificar JSON: " + string(reset), err)
    }
    fmt.Println(string(green) + "JSON codificado: " + string(reset) + string(userJSON))

    // Decodificar JSON para a estrutura
    var decodedUser User
    err = json.Unmarshal(userJSON, &decodedUser)
    if err != nil {
        log.Fatal(string(red) + "Erro ao decodificar JSON: " + string(reset), err)
    }
    fmt.Printf(string(green)+"JSON decodificado: %+v\n"+string(reset), decodedUser)

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Exemplo de solução em Go" + string(reset))
    fmt.Println(separator)
    fmt.Println("```go")
    fmt.Println("package main")
    fmt.Println()
    fmt.Println("import (")
    fmt.Println("    \"encoding/json\"")
    fmt.Println("    \"fmt\"")
    fmt.Println("    \"log\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("// Estrutura de exemplo para codificação JSON")
    fmt.Println("type User struct {")
    fmt.Println("    Name  string `json:\"name\"`")
    fmt.Println("    Age   int    `json:\"age\"`")
    fmt.Println("    Email string `json:\"email\"`")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    // Criar instância da estrutura User")
    fmt.Println("    user := User{")
    fmt.Println("        Name:  \"Alice\",")
    fmt.Println("        Age:   30,")
    fmt.Println("        Email: \"alice@example.com\",")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    // Codificar a estrutura em JSON")
    fmt.Println("    userJSON, err := json.Marshal(user)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatal(\"Erro ao codificar JSON: \", err)")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"JSON codificado: \", string(userJSON))")
    fmt.Println()
    fmt.Println("    // Decodificar JSON para a estrutura")
    fmt.Println("    var decodedUser User")
    fmt.Println("    err = json.Unmarshal(userJSON, &decodedUser)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatal(\"Erro ao decodificar JSON: \", err)")
    fmt.Println("    }")
    fmt.Println("    fmt.Printf(\"JSON decodificado: %+v\\n\", decodedUser)")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a codificar e decodificar dados JSON em Go. JSON é um formato leve de troca de dados, amplamente utilizado em APIs e armazenamento de dados." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use `json.Marshal` para codificar dados em JSON.",
        "Curiosidade: JSON (JavaScript Object Notation) é um formato leve de troca de dados.",
        "Dica: Use `json.Unmarshal` para decodificar dados JSON em estruturas Go.",
        "Curiosidade: JSON é amplamente utilizado em APIs web devido à sua simplicidade e legibilidade.",
        "Dica: Adicione tags JSON às suas estruturas Go para controlar a codificação e decodificação.",
        "Curiosidade: Go fornece suporte nativo para JSON através do pacote `encoding/json`.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o décimo segundo nível!" + string(reset))
}
