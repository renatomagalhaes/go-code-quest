package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
    "math/rand"
	"time"
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)

type Level25 struct{}

func StartLevel25() {
    game.StartLevel(Level25{})
}

func (l Level25) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 25: Uso de WebSockets para Comunicação em Tempo Real" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Configure um servidor WebSocket e implemente comunicação em tempo real" + string(reset))
    fmt.Println()

    // Iniciar o servidor WebSocket
    http.HandleFunc("/ws", handleConnections)
    go handleMessages()

    fmt.Println(string(green) + "Servidor WebSocket iniciado em ws://localhost:8080/ws" + string(reset))
    fmt.Println("Abra um navegador e conecte-se ao WebSocket para enviar e receber mensagens.")
    log.Fatal(http.ListenAndServe(":8080", nil))

    showWebSocketSummary()
}

// Definir o atualizador de WebSocket
var upgrader = websocket.Upgrader{}

// Canal para mensagens
var broadcast = make(chan Message)

// Cliente conectado
var clients = make(map[*websocket.Conn]bool)

// Estrutura para mensagem
type Message struct {
    Username string `json:"username"`
    Message  string `json:"message"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer ws.Close()

    clients[ws] = true

    for {
        var msg Message
        err := ws.ReadJSON(&msg)
        if err != nil {
            log.Printf("erro: %v", err)
            delete(clients, ws)
            break
        }
        broadcast <- msg
    }
}

func handleMessages() {
    for {
        msg := <-broadcast
        for client := range clients {
            err := client.WriteJSON(msg)
            if err != nil {
                log.Printf("erro: %v", err)
                client.Close()
                delete(clients, client)
            }
        }
    }
}

func showWebSocketSummary() {
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
    fmt.Println()
    fmt.Println("    \"github.com/gorilla/websocket\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("var upgrader = websocket.Upgrader{}")
    fmt.Println("var broadcast = make(chan Message)")
    fmt.Println("var clients = make(map[*websocket.Conn]bool)")
    fmt.Println()
    fmt.Println("type Message struct {")
    fmt.Println("    Username string `json:\"username\"`")
    fmt.Println("    Message  string `json:\"message\"`")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    http.HandleFunc(\"/ws\", handleConnections)")
    fmt.Println("    go handleMessages()")
    fmt.Println()
    fmt.Println("    fmt.Println(\"Servidor WebSocket iniciado em ws://localhost:8080/ws\")")
    fmt.Println("    log.Fatal(http.ListenAndServe(\":8080\", nil))")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func handleConnections(w http.ResponseWriter, r *http.Request) {")
    fmt.Println("    ws, err := upgrader.Upgrade(w, r, nil)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatal(err)")
    fmt.Println("    }")
    fmt.Println("    defer ws.Close()")
    fmt.Println()
    fmt.Println("    clients[ws] = true")
    fmt.Println()
    fmt.Println("    for {")
    fmt.Println("        var msg Message")
    fmt.Println("        err := ws.ReadJSON(&msg)")
    fmt.Println("        if err != nil {")
    fmt.Println("            log.Printf(\"erro: %v\", err)")
    fmt.Println("            delete(clients, ws)")
    fmt.Println("            break")
    fmt.Println("        }")
    fmt.Println("        broadcast <- msg")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func handleMessages() {")
    fmt.Println("    for {")
    fmt.Println("        msg := <-broadcast")
    fmt.Println("        for client := range clients {")
    fmt.Println("            err := client.WriteJSON(msg)")
    fmt.Println("            if err != nil {")
    fmt.Println("                log.Printf(\"erro: %v\", err)")
    fmt.Println("                client.Close()")
    fmt.Println("                delete(clients, client)")
    fmt.Println("            }")
    fmt.Println("        }")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a usar WebSockets em Go para permitir comunicação em tempo real entre um servidor e clientes. Você configurou um servidor WebSocket, enviou e recebeu mensagens em tempo real." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use o pacote `github.com/gorilla/websocket` para adicionar suporte a WebSockets no seu servidor Go.",
        "Curiosidade: WebSockets permitem comunicação bidirecional entre o cliente e o servidor.",
        "Dica: Use `upgrader.Upgrade` para atualizar uma conexão HTTP para uma conexão WebSocket.",
        "Curiosidade: WebSockets são amplamente usados em aplicações de tempo real, como chat, jogos e notificações ao vivo.",
        "Dica: Sempre feche a conexão WebSocket com `defer ws.Close()` para liberar recursos.",
        "Curiosidade: WebSockets usam um protocolo diferente do HTTP, mas a conexão inicial é feita através de um handshake HTTP.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o vigésimo quinto nível!" + string(reset))
}
