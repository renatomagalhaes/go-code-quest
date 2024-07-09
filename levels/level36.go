package levels

import (
    "encoding/json"
    "fmt"
    "go-code-quest/pkg/game"
		"math/rand"
		"time"
    "log"
)

type Level36 struct{}

func StartLevel36() {
    game.StartLevel(Level36{})
}

func (l Level36) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 36: Manipulação de JSON em Go" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Aprenda a trabalhar com JSON em Go, incluindo serialização, desserialização e uso de tags de struct" + string(reset))
    fmt.Println()

    // Exemplo de JSON simples
    person := Person{
        FirstName: "John",
        LastName:  "Doe",
        Age:       30,
        Address: Address{
            Street: "123 Main St",
            City:   "Springfield",
            State:  "IL",
        },
    }

    jsonData, err := json.Marshal(person)
    if err != nil {
        log.Fatalf("Erro ao serializar JSON: %v", err)
    }

    fmt.Println(string(green) + "JSON Serializado: " + string(jsonData) + string(reset))

    var person2 Person
    err = json.Unmarshal(jsonData, &person2)
    if err != nil {
        log.Fatalf("Erro ao desserializar JSON: %v", err)
    }

    fmt.Println(string(green) + "JSON Desserializado: " + fmt.Sprintf("%+v", person2) + string(reset))

    showJSONSummary()
}

type Person struct {
    FirstName string  `json:"first_name"`
    LastName  string  `json:"last_name"`
    Age       int     `json:"age"`
    Address   Address `json:"address"`
}

type Address struct {
    Street string `json:"street"`
    City   string `json:"city"`
    State  string `json:"state"`
}

func showJSONSummary() {
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
    fmt.Println("    \"encoding/json\"")
    fmt.Println("    \"fmt\"")
    fmt.Println("    \"log\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("type Person struct {")
    fmt.Println("    FirstName string  `json:\"first_name\"`")
    fmt.Println("    LastName  string  `json:\"last_name\"`")
    fmt.Println("    Age       int     `json:\"age\"`")
    fmt.Println("    Address   Address `json:\"address\"`")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("type Address struct {")
    fmt.Println("    Street string `json:\"street\"`")
    fmt.Println("    City   string `json:\"city\"`")
    fmt.Println("    State  string `json:\"state\"`")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    person := Person{")
    fmt.Println("        FirstName: \"John\",")
    fmt.Println("        LastName:  \"Doe\",")
    fmt.Println("        Age:       30,")
    fmt.Println("        Address: Address{")
    fmt.Println("            Street: \"123 Main St\",")
    fmt.Println("            City:   \"Springfield\",")
    fmt.Println("            State:  \"IL\",")
    fmt.Println("        },")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    jsonData, err := json.Marshal(person)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"Erro ao serializar JSON: %v\", err)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    fmt.Println(\"JSON Serializado: \" + string(jsonData))")
    fmt.Println()
    fmt.Println("    var person2 Person")
    fmt.Println("    err = json.Unmarshal(jsonData, &person2)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"Erro ao desserializar JSON: %v\", err)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    fmt.Println(\"JSON Desserializado: \" + fmt.Sprintf(\"%+v\", person2))")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a trabalhar com JSON em Go, incluindo serialização e desserialização de dados e uso de tags de struct para customizar a conversão." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use `json.Marshal` para serializar dados em JSON.",
        "Curiosidade: Use `json.Unmarshal` para desserializar dados JSON em structs Go.",
        "Dica: Use tags de struct para definir nomes de campo JSON personalizados.",
        "Curiosidade: O pacote `encoding/json` é parte da biblioteca padrão do Go.",
        "Dica: As tags de struct permitem definir opções como `omitempty` para ignorar campos vazios na serialização.",
        "Curiosidade: JSON (JavaScript Object Notation) é um formato de intercâmbio de dados leve e amplamente utilizado.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o trigésimo sexto nível!" + string(reset))
}
