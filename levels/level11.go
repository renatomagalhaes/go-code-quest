package levels

import (
    "bufio"
    "fmt"
    "go-code-quest/pkg/game"
    //"log"
    "os"
    "math/rand"
    "time"

    "github.com/streadway/amqp"
)

type Level11 struct{}

func StartLevel11() {
    game.StartLevel(Level11{})
}

func (l Level11) Start() {
    // Definir cores
    red := "\033[31m"
    green := "\033[32m"
    yellow := "\033[33m"
    blue := "\033[34m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 11: Integração com RabbitMQ" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Conecte-se a um broker RabbitMQ, envie mensagens a uma fila e consuma mensagens da fila" + string(reset))
    fmt.Println()

    reader := bufio.NewReader(os.Stdin)

    // Solicitar parâmetros de conexão ao usuário
    fmt.Print(string(blue) + "Digite o usuário do RabbitMQ: " + string(reset))
    rabbitUser, _ := reader.ReadString('\n')
    rabbitUser = rabbitUser[:len(rabbitUser)-1]

    fmt.Print(string(blue) + "Digite a senha do RabbitMQ: " + string(reset))
    rabbitPassword, _ := reader.ReadString('\n')
    rabbitPassword = rabbitPassword[:len(rabbitPassword)-1]

    fmt.Print(string(blue) + "Digite o endereço do RabbitMQ (ex: localhost): " + string(reset))
    rabbitAddress, _ := reader.ReadString('\n')
    rabbitAddress = rabbitAddress[:len(rabbitAddress)-1]

    fmt.Print(string(blue) + "Digite a porta do RabbitMQ (ex: 5672): " + string(reset))
    rabbitPort, _ := reader.ReadString('\n')
    rabbitPort = rabbitPort[:len(rabbitPort)-1]

    rabbitURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitUser, rabbitPassword, rabbitAddress, rabbitPort)

    // Tentar conectar ao RabbitMQ
    conn, err := amqp.Dial(rabbitURL)
    if err != nil {
        fmt.Println(string(red) + "Erro ao conectar ao RabbitMQ: " + string(reset), err)
    } else {
        defer conn.Close() // Fechar a conexão quando a função terminar

        // Abrir um canal de comunicação
        ch, err := conn.Channel()
        if err != nil {
            fmt.Println(string(red) + "Erro ao abrir canal: " + string(reset), err)
        } else {
            defer ch.Close() // Fechar o canal quando a função terminar

            // Declarar fila
            q, err := ch.QueueDeclare(
                "hello", // Nome da fila
                false,   // Durável
                false,   // Auto exclusão
                false,   // Exclusivo
                false,   // Sem espera
                nil,     // Argumentos adicionais
            )
            if err != nil {
                fmt.Println(string(red) + "Erro ao declarar fila: " + string(reset), err)
            }

            // Enviar mensagem
            body := "Hello, RabbitMQ!"
            err = ch.Publish(
                "",     // Troca
                q.Name, // Roteamento da chave (routing key)
                false,  // Obrigatório
                false,  // Imediato
                amqp.Publishing{
                    ContentType: "text/plain", // Tipo de conteúdo
                    Body:        []byte(body), // Mensagem
                })
            if err != nil {
                fmt.Println(string(red) + "Erro ao enviar mensagem: " + string(reset), err)
            }

            // Consumir mensagens
            msgs, err := ch.Consume(
                q.Name, // Nome da fila
                "",     // Consumidor
                false,  // Auto ack
                false,  // Exclusivo
                false,  // Sem espera
                false,  // Sem argumentos adicionais
                nil,
            )
            if err != nil {
                fmt.Println(string(red) + "Erro ao consumir mensagens: " + string(reset), err)
            }

            go func() {
                for d := range msgs {
                    fmt.Println(string(green) + "Recebido: " + string(d.Body) + string(reset))
                    d.Ack(false) // Acknowledge (ack) a mensagem
                    // Comentado: Exemplo de uso de nack para simular falha e retornar a mensagem para a fila
                    // d.Nack(false, true) // Nack (negativa) a mensagem
                }
            }()

            fmt.Println(string(green) + "Pressione Enter para sair..." + string(reset))
            reader.ReadString('\n')
        }
    }

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
    fmt.Println("    \"github.com/streadway/amqp\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    // Conectar ao RabbitMQ")
    fmt.Println("    conn, err := amqp.Dial(\"amqp://guest:guest@localhost:5672/\")")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatal(\"Erro ao conectar ao RabbitMQ: \", err)")
    fmt.Println("    }")
    fmt.Println("    defer conn.Close() // Fechar a conexão quando a função terminar")
    fmt.Println()
    fmt.Println("    // Abrir um canal de comunicação")
    fmt.Println("    ch, err := conn.Channel()")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatal(\"Erro ao abrir canal: \", err)")
    fmt.Println("    }")
    fmt.Println("    defer ch.Close() // Fechar o canal quando a função terminar")
    fmt.Println()
    fmt.Println("    // Declarar fila")
    fmt.Println("    q, err := ch.QueueDeclare(")
    fmt.Println("        \"hello\", // Nome da fila")
    fmt.Println("        false,    // Durável")
    fmt.Println("        false,    // Auto exclusão")
    fmt.Println("        false,    // Exclusivo")
    fmt.Println("        false,    // Sem espera")
    fmt.Println("        nil,      // Argumentos adicionais")
    fmt.Println("    )")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatal(\"Erro ao declarar fila: \", err)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    // Enviar mensagem")
    fmt.Println("    body := \"Hello, RabbitMQ!\"")
    fmt.Println("    err = ch.Publish(")
    fmt.Println("        \"\",     // Troca")
    fmt.Println("        q.Name, // Roteamento da chave (routing key)")
    fmt.Println("        false,  // Obrigatório")
    fmt.Println("        false,  // Imediato")
    fmt.Println("        amqp.Publishing{")
    fmt.Println("            ContentType: \"text/plain\", // Tipo de conteúdo")
    fmt.Println("            Body:        []byte(body), // Mensagem")
    fmt.Println("        })")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatal(\"Erro ao enviar mensagem: \", err)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    // Consumir mensagens")
    fmt.Println("    msgs, err := ch.Consume(")
    fmt.Println("        q.Name, // Nome da fila")
    fmt.Println("        \"\",     // Consumidor")
    fmt.Println("        false,  // Auto ack")
    fmt.Println("        false,  // Exclusivo")
    fmt.Println("        false,  // Sem espera")
    fmt.Println("        false,  // Sem argumentos adicionais")
    fmt.Println("        nil,")
    fmt.Println("    )")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatal(\"Erro ao consumir mensagens: \", err)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    go func() {")
    fmt.Println("        for d := range msgs {")
    fmt.Println("            fmt.Println(\"Recebido: \" + string(d.Body))")
    fmt.Println("            d.Ack(false) // Acknowledge (ack) a mensagem")
    fmt.Println("            // Comentado: Exemplo de uso de nack para simular falha e retornar a mensagem para a fila")
    fmt.Println("            // d.Nack(false, true) // Nack (negativa) a mensagem")
    fmt.Println("        }")
    fmt.Println("    }()")
    fmt.Println()
    fmt.Println("    fmt.Println(\"Pressione Enter para sair...\")")
    fmt.Println("    fmt.Scanln()")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a integrar-se com RabbitMQ em Go. A comunicação assíncrona usando filas de mensagens é uma técnica poderosa para construir sistemas distribuídos e desacoplados." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use o pacote `github.com/streadway/amqp` para integrar-se com RabbitMQ em Go.",
        "Curiosidade: RabbitMQ é um broker de mensagens amplamente utilizado para comunicação assíncrona.",
        "Dica: Sempre feche a conexão e o canal com RabbitMQ usando `defer conn.Close()` e `defer ch.Close()`.",
        "Curiosidade: Mensageria é uma técnica essencial para construção de sistemas escaláveis e resilientes.",
        "Dica: Use `ch.QueueDeclare` para garantir que a fila exista antes de publicar ou consumir mensagens.",
        "Curiosidade: RabbitMQ suporta vários padrões de troca de mensagens, incluindo filas, tópicos e roteamento direto.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o décimo primeiro nível!" + string(reset))
}
