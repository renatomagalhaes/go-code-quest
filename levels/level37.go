package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
    "log"
    "net/http"
		"math/rand"
    "time"

    "github.com/gorilla/websocket"
)

type Level37 struct{}

func StartLevel37() {
    game.StartLevel(Level37{})
}

var wsUpgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func (l Level37) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 37: Utilizando WebSockets em Go" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Aprenda a implementar WebSockets em suas aplicações Go usando a biblioteca `gorilla/websocket`" + string(reset))
    fmt.Println()

    // Iniciar servidor WebSocket
    http.HandleFunc("/ws", wsHandler37)
    go func() {
        fmt.Println(string(green) + "Servidor WebSocket iniciado em ws://localhost:8080/ws" + string(reset))
        if err := http.ListenAndServe(":8080", nil); err != nil {
            log.Fatalf("Erro ao iniciar servidor WebSocket: %v", err)
        }
    }()

    // Simulação de tempo de execução
    time.Sleep(5 * time.Second) // Manter o servidor rodando por 5 segundos

    showWebSocketSummary37()
}

func wsHandler37(w http.ResponseWriter, r *http.Request) {
    conn, err := wsUpgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Erro ao atualizar para WebSocket:", err)
        return
    }
    defer conn.Close()

    // Goroutine para enviar mensagens
    go func() {
        for {
            message := "Olá do servidor!"
            err := conn.WriteMessage(websocket.TextMessage, []byte(message))
            if err != nil {
                log.Println("Erro ao enviar mensagem:", err)
                return
            }
            time.Sleep(2 * time.Second) // Esperar 2 segundos antes de enviar a próxima mensagem
        }
    }()

    // Ler mensagens do cliente
    for {
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            log.Println("Erro ao ler mensagem:", err)
            break
        }
        log.Printf("Mensagem recebida do cliente: %s", message)
        if err := conn.WriteMessage(messageType, message); err != nil {
            log.Println("Erro ao escrever mensagem:", err)
            break
        }
    }
}

func showWebSocketSummary37() {
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
    fmt.Println("    \"log\"")
    fmt.Println("    \"net/http\"")
    fmt.Println("    \"time\"")
    fmt.Println()
    fmt.Println("    \"github.com/gorilla/websocket\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("var wsUpgrader = websocket.Upgrader{")
    fmt.Println("    ReadBufferSize:  1024,")
    fmt.Println("    WriteBufferSize: 1024,")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    http.HandleFunc(\"/ws\", wsHandler37)")
    fmt.Println("    go func() {")
    fmt.Println("        fmt.Println(\"Servidor WebSocket iniciado em ws://localhost:8080/ws\")")
    fmt.Println("        if err := http.ListenAndServe(\":8080\", nil); err != nil {")
    fmt.Println("            log.Fatalf(\"Erro ao iniciar servidor WebSocket: %v\", err)")
    fmt.Println("        }")
    fmt.Println("    }()")
    fmt.Println()
    fmt.Println("    time.Sleep(5 * time.Second) // Manter o servidor rodando por 5 segundos")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func wsHandler37(w http.ResponseWriter, r *http.Request) {")
    fmt.Println("    conn, err := wsUpgrader.Upgrade(w, r, nil)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao atualizar para WebSocket:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    defer conn.Close()")
    fmt.Println()
    fmt.Println("    // Goroutine para enviar mensagens")
    fmt.Println("    go func() {")
    fmt.Println("        for {")
    fmt.Println("            message := \"Olá do servidor!\"")
    fmt.Println("            err := conn.WriteMessage(websocket.TextMessage, []byte(message))")
    fmt.Println("            if err != nil {")
    fmt.Println("                log.Println(\"Erro ao enviar mensagem:\", err)")
    fmt.Println("                return")
    fmt.Println("            }")
    fmt.Println("            time.Sleep(2 * time.Second) // Esperar 2 segundos antes de enviar a próxima mensagem")
    fmt.Println("        }")
    fmt.Println("    }()")
    fmt.Println()
    fmt.Println("    // Ler mensagens do cliente")
    fmt.Println("    for {")
    fmt.Println("        messageType, message, err := conn.ReadMessage()")
    fmt.Println("        if err != nil {")
    fmt.Println("            log.Println(\"Erro ao ler mensagem:\", err)")
    fmt.Println("            break")
    fmt.Println("        }")
    fmt.Println("        log.Printf(\"Mensagem recebida do cliente: %s\", message)")
    fmt.Println("        if err := conn.WriteMessage(messageType, message); err != nil {")
    fmt.Println("            log.Println(\"Erro ao escrever mensagem:\", err)")
    fmt.Println("            break")
    fmt.Println("        }")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a implementar WebSockets em suas aplicações Go usando a biblioteca `gorilla/websocket`, incluindo a configuração do servidor WebSocket e o manejo de conexões e mensagens bidirecionais." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use WebSockets para comunicação bidirecional em tempo real entre cliente e servidor.",
        "Curiosidade: O protocolo WebSocket permite uma comunicação eficiente e de baixa latência.",
        "Dica: Use a biblioteca `gorilla/websocket` para gerenciar conexões WebSocket em Go.",
        "Curiosidade: WebSockets são amplamente utilizados em aplicações de chat, jogos online e atualizações em tempo real.",
        "Dica: Lembre-se de lidar com erros de conexão e mensagens ao usar WebSockets.",
        "Curiosidade: O WebSocket é uma tecnologia padrão definida pelo IETF como RFC 6455.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o trigésimo sétimo nível!" + string(reset))
}
