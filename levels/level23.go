package levels

import (
    "bytes"
    "encoding/json"
    "fmt"
    "go-code-quest/pkg/game"
		"math/rand"
		"time"
    "io/ioutil"
    "log"
    "net/http"
)

type Level23 struct{}

func StartLevel23() {
    game.StartLevel(Level23{})
}

func (l Level23) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 23: Integração com APIs REST usando HTTP" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Integre-se com uma API REST usando as operações HTTP básicas (GET, POST, PUT, DELETE)" + string(reset))
    fmt.Println()

    // URL base da API
    apiUrl := "https://jsonplaceholder.typicode.com/posts"

    // Fazer uma requisição GET
    getResponse, err := http.Get(apiUrl + "/1")
    if err != nil {
        log.Println("Erro ao fazer requisição GET:", err)
        showApiSummary()
        return
    }
    defer getResponse.Body.Close()

    var getResult map[string]interface{}
    if err := json.NewDecoder(getResponse.Body).Decode(&getResult); err != nil {
        log.Println("Erro ao processar resposta GET:", err)
        showApiSummary()
        return
    }
    fmt.Println(string(green) + "GET Resultado:" + string(reset), getResult)

    // Fazer uma requisição POST
    postData := map[string]string{"title": "foo", "body": "bar", "userId": "1"}
    postBody, _ := json.Marshal(postData)
    postResponse, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(postBody))
    if err != nil {
        log.Println("Erro ao fazer requisição POST:", err)
        showApiSummary()
        return
    }
    defer postResponse.Body.Close()

    var postResult map[string]interface{}
    if err := json.NewDecoder(postResponse.Body).Decode(&postResult); err != nil {
        log.Println("Erro ao processar resposta POST:", err)
        showApiSummary()
        return
    }
    fmt.Println(string(green) + "POST Resultado:" + string(reset), postResult)

    // Fazer uma requisição PUT
    putData := map[string]string{"id": "1", "title": "foo", "body": "bar", "userId": "1"}
    putBody, _ := json.Marshal(putData)
    client := &http.Client{}
    putRequest, err := http.NewRequest(http.MethodPut, apiUrl+"/1", bytes.NewBuffer(putBody))
    if err != nil {
        log.Println("Erro ao criar requisição PUT:", err)
        showApiSummary()
        return
    }
    putRequest.Header.Set("Content-Type", "application/json")
    putResponse, err := client.Do(putRequest)
    if err != nil {
        log.Println("Erro ao fazer requisição PUT:", err)
        showApiSummary()
        return
    }
    defer putResponse.Body.Close()

    var putResult map[string]interface{}
    if err := json.NewDecoder(putResponse.Body).Decode(&putResult); err != nil {
        log.Println("Erro ao processar resposta PUT:", err)
        showApiSummary()
        return
    }
    fmt.Println(string(green) + "PUT Resultado:" + string(reset), putResult)

    // Fazer uma requisição DELETE
    deleteRequest, err := http.NewRequest(http.MethodDelete, apiUrl+"/1", nil)
    if err != nil {
        log.Println("Erro ao criar requisição DELETE:", err)
        showApiSummary()
        return
    }
    deleteResponse, err := client.Do(deleteRequest)
    if err != nil {
        log.Println("Erro ao fazer requisição DELETE:", err)
        showApiSummary()
        return
    }
    defer deleteResponse.Body.Close()

    deleteBody, err := ioutil.ReadAll(deleteResponse.Body)
    if err != nil {
        log.Println("Erro ao ler resposta DELETE:", err)
        showApiSummary()
        return
    }
    fmt.Println(string(green) + "DELETE Resultado:" + string(reset), string(deleteBody))

    showApiSummary()
}

func showApiSummary() {
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
    fmt.Println("    \"bytes\"")
    fmt.Println("    \"encoding/json\"")
    fmt.Println("    \"fmt\"")
    fmt.Println("    \"io/ioutil\"")
    fmt.Println("    \"log\"")
    fmt.Println("    \"net/http\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    // URL base da API")
    fmt.Println("    apiUrl := \"https://jsonplaceholder.typicode.com/posts\"")
    fmt.Println()
    fmt.Println("    // Fazer uma requisição GET")
    fmt.Println("    getResponse, err := http.Get(apiUrl + \"/1\")")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao fazer requisição GET:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    defer getResponse.Body.Close()")
    fmt.Println()
    fmt.Println("    var getResult map[string]interface{}")
    fmt.Println("    if err := json.NewDecoder(getResponse.Body).Decode(&getResult); err != nil {")
    fmt.Println("        log.Println(\"Erro ao processar resposta GET:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"GET Resultado:\", getResult)")
    fmt.Println()
    fmt.Println("    // Fazer uma requisição POST")
    fmt.Println("    postData := map[string]string{\"title\": \"foo\", \"body\": \"bar\", \"userId\": \"1\"}")
    fmt.Println("    postBody, _ := json.Marshal(postData)")
    fmt.Println("    postResponse, err := http.Post(apiUrl, \"application/json\", bytes.NewBuffer(postBody))")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao fazer requisição POST:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    defer postResponse.Body.Close()")
    fmt.Println()
    fmt.Println("    var postResult map[string]interface{}")
    fmt.Println("    if err := json.NewDecoder(postResponse.Body).Decode(&postResult); err != nil {")
    fmt.Println("        log.Println(\"Erro ao processar resposta POST:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"POST Resultado:\", postResult)")
    fmt.Println()
    fmt.Println("    // Fazer uma requisição PUT")
    fmt.Println("    putData := map[string]string{\"id\": \"1\", \"title\": \"foo\", \"body\": \"bar\", \"userId\": \"1\"}")
    fmt.Println("    putBody, _ := json.Marshal(putData)")
    fmt.Println("    client := &http.Client{}")
    fmt.Println("    putRequest, err := http.NewRequest(http.MethodPut, apiUrl+\"/1\", bytes.NewBuffer(putBody))")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao criar requisição PUT:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    putRequest.Header.Set(\"Content-Type\", \"application/json\")")
    fmt.Println("    putResponse, err := client.Do(putRequest)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao fazer requisição PUT:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    defer putResponse.Body.Close()")
    fmt.Println()
    fmt.Println("    var putResult map[string]interface{}")
    fmt.Println("    if err := json.NewDecoder(putResponse.Body).Decode(&putResult); err != nil {")
    fmt.Println("        log.Println(\"Erro ao processar resposta PUT:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"PUT Resultado:\", putResult)")
    fmt.Println()
    fmt.Println("    // Fazer uma requisição DELETE")
    fmt.Println("    deleteRequest, err := http.NewRequest(http.MethodDelete, apiUrl+\"/1\", nil)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao criar requisição DELETE:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    deleteResponse, err := client.Do(deleteRequest)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao fazer requisição DELETE:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    defer deleteResponse.Body.Close()")
    fmt.Println()
    fmt.Println("    deleteBody, err := ioutil.ReadAll(deleteResponse.Body)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao ler resposta DELETE:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"DELETE Resultado:\", string(deleteBody))")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a integrar Golang com APIs REST usando HTTP. Aprendemos a fazer requisições GET, POST, PUT e DELETE, e a processar respostas JSON." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use o pacote `net/http` para fazer requisições HTTP em Go.",
        "Curiosidade: APIs REST são amplamente usadas para comunicação entre serviços web.",
        "Dica: Use `json.NewDecoder` para decodificar respostas JSON em Go.",
        "Curiosidade: `ioutil.ReadAll` lê o corpo da resposta HTTP em um único byte slice.",
        "Dica: Sempre feche o corpo da resposta HTTP (`Body`) para evitar vazamentos de memória.",
        "Curiosidade: A função `http.NewRequest` permite criar requisições HTTP customizadas.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o vigésimo terceiro nível!" + string(reset))
}
