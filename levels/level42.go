package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
    "log"
    "os"
    "sync"
		"math/rand"
    "time"

    "github.com/streadway/amqp"
)

type Level42 struct{}

func StartLevel42() {
    game.StartLevel(Level42{})
}

func (l Level42) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 42: Conexão Avançada com RabbitMQ e Consumidores" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Aprenda a implementar uma conexão avançada com RabbitMQ, incluindo criação de consumidores, leitura de uma fila de alto movimento, e uso de ack e nack." + string(reset))
    fmt.Println()

    // Configurar RabbitMQ
    rabbitMQURL := os.Getenv("RABBITMQ_URL")
    if rabbitMQURL == "" {
        rabbitMQURL = "amqp://guest:guest@localhost:5672/"
    }

    conn, err := amqp.Dial(rabbitMQURL)
    if err != nil {
        log.Fatalf("Falha ao conectar ao RabbitMQ: %v", err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Falha ao abrir um canal: %v", err)
    }
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "task_queue",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        log.Fatalf("Falha ao declarar uma fila: %v", err)
    }

    // Quantidade de consumidores
    var consumers int
    fmt.Print("Digite a quantidade de consumidores: ")
    fmt.Scan(&consumers)

    var wg sync.WaitGroup
    wg.Add(consumers)

    for i := 0; i < consumers; i++ {
        go func(id int) {
            defer wg.Done()
            consumeMessages(ch, q.Name, id)
        }(i)
    }

    wg.Wait()

    fmt.Println(string(green), "Todos os consumidores concluíram o processamento das mensagens!", string(reset))

    showRabbitMQSummary42()
}

func consumeMessages(ch *amqp.Channel, queueName string, id int) {
    msgs, err := ch.Consume(
        queueName,
        "",
        false,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        log.Fatalf("Falha ao registrar um consumidor: %v", err)
    }

    for msg := range msgs {
        log.Printf("Consumidor %d: Mensagem recebida: %s", id, msg.Body)

        // Processar a mensagem
        if err := processMessage(msg.Body); err != nil {
            log.Printf("Consumidor %d: Falha ao processar mensagem: %v", id, err)
            msg.Nack(false, true) // Rejeitar a mensagem e reencaminhá-la para a fila
        } else {
            msg.Ack(false) // Confirmar processamento da mensagem
        }
    }
}

func processMessage(body []byte) error {
    // Simular processamento
    time.Sleep(time.Second * 2)
    if string(body) == "error" {
        return fmt.Errorf("erro simulado")
    }
    return nil
}

func showRabbitMQSummary42() {
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
    fmt.Println("    \"os\"")
    fmt.Println("    \"sync\"")
    fmt.Println("    \"time\"")
    fmt.Println()
    fmt.Println("    \"github.com/streadway/amqp\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    rabbitMQURL := os.Getenv(\"RABBITMQ_URL\")")
    fmt.Println("    if rabbitMQURL == \"\" {")
    fmt.Println("        rabbitMQURL = \"amqp://guest:guest@localhost:5672/\"")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    conn, err := amqp.Dial(rabbitMQURL)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"Falha ao conectar ao RabbitMQ: %v\", err)")
    fmt.Println("    }")
    fmt.Println("    defer conn.Close()")
    fmt.Println()
    fmt.Println("    ch, err := conn.Channel()")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"Falha ao abrir um canal: %v\", err)")
    fmt.Println("    }")
    fmt.Println("    defer ch.Close()")
    fmt.Println()
    fmt.Println("    q, err := ch.QueueDeclare(")
    fmt.Println("        \"task_queue\",")
    fmt.Println("        true,")
    fmt.Println("        false,")
    fmt.Println("        false,")
    fmt.Println("        false,")
    fmt.Println("        nil,")
    fmt.Println("    )")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"Falha ao declarar uma fila: %v\", err)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    var consumers int")
    fmt.Println("    fmt.Print(\"Digite a quantidade de consumidores: \")")
    fmt.Println("    fmt.Scan(&consumers)")
    fmt.Println()
    fmt.Println("    var wg sync.WaitGroup")
    fmt.Println("    wg.Add(consumers)")
    fmt.Println()
    fmt.Println("    for i := 0; i < consumers; i++ {")
    fmt.Println("        go func(id int) {")
    fmt.Println("            defer wg.Done()")
    fmt.Println("            consumeMessages(ch, q.Name, id)")
    fmt.Println("        }(i)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    wg.Wait()")
    fmt.Println("    fmt.Println(\"Todos os consumidores concluíram o processamento das mensagens!\")")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func consumeMessages(ch *amqp.Channel, queueName string, id int) {")
    fmt.Println("    msgs, err := ch.Consume(")
    fmt.Println("        queueName,")
    fmt.Println("        \"\",")
    fmt.Println("        false,")
    fmt.Println("        false,")
    fmt.Println("        false,")
    fmt.Println("        false,")
    fmt.Println("        nil,")
    fmt.Println("    )")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"Falha ao registrar um consumidor: %v\", err)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    for msg := range msgs {")
    fmt.Println("        log.Printf(\"Consumidor %d: Mensagem recebida: %s\", id, msg.Body)")
    fmt.Println()
    fmt.Println("        if err := processMessage(msg.Body); err != nil {")
    fmt.Println("            log.Printf(\"Consumidor %d: Falha ao processar mensagem: %v\", id, err)")
    fmt.Println("            msg.Nack(false, true) // Rejeitar a mensagem e reencaminhá-la para a fila")
    fmt.Println("        } else {")
    fmt.Println("            msg.Ack(false) // Confirmar processamento da mensagem")
    fmt.Println("        }")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func processMessage(body []byte) error {")
    fmt.Println("    time.Sleep(time.Second * 2)")
    fmt.Println("    if string(body) == \"error\" {")
    fmt.Println("        return fmt.Errorf(\"erro simulado\")")
    fmt.Println("    }")
    fmt.Println("    return nil")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a implementar uma conexão avançada com RabbitMQ, incluindo a criação de consumidores, leitura de uma fila de alto movimento, e uso de ack e nack para gerenciamento de mensagens." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use `ack` para confirmar o processamento de uma mensagem e removê-la da fila.",
        "Curiosidade: `nack` pode ser usado para rejeitar uma mensagem e reencaminhá-la para a fila.",
        "Dica: Configure o RabbitMQ para lidar com filas de alto movimento e garantir alta disponibilidade.",
        "Curiosidade: O RabbitMQ é um sistema de mensagens amplamente utilizado em arquiteturas de microserviços.",
        "Dica: Use múltiplos consumidores para processar mensagens em paralelo e aumentar a eficiência.",
        "Curiosidade: O RabbitMQ suporta vários padrões de mensagens, incluindo filas, tópicos e RPC.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o quadragésimo segundo nível!" + string(reset))
}
