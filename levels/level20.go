package levels

import (
    "context"
    "fmt"
    "go-code-quest/pkg/game"
    "log"
    "os"
		"math/rand"
    "time"

    "github.com/joho/godotenv"
    "github.com/streadway/amqp"
)

type Level20 struct{}

func StartLevel20() {
    game.StartLevel(Level20{})
}

func (l Level20) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 20: Publicação e Consumo de Mensagens com RabbitMQ" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Conecte-se ao RabbitMQ, publique e consuma mensagens usando variáveis de ambiente" + string(reset))
    fmt.Println()

    // Carregar variáveis de ambiente do arquivo .env
    err := godotenv.Load()
    if err != nil {
        log.Println("Erro ao carregar o arquivo .env:", err)
        showRabbitMQSummary()
        return
    }

    rabbitmqURI := os.Getenv("RABBITMQ_URI")
    queueName := os.Getenv("QUEUE_NAME")

    // Tentar conectar ao RabbitMQ
    conn, err := amqp.Dial(rabbitmqURI)
    if err != nil {
        log.Println("Erro ao conectar ao RabbitMQ:", err)
        showRabbitMQSummary()
        return
    }
    defer conn.Close()

    // Criar um canal
    ch, err := conn.Channel()
    if err != nil {
        log.Println("Erro ao criar canal RabbitMQ:", err)
        showRabbitMQSummary()
        return
    }
    defer ch.Close()

    // Declarar uma fila
    q, err := ch.QueueDeclare(
        queueName, // Nome da fila
        false,     // Durável
        false,     // Deletar automaticamente
        false,     // Exclusiva
        false,     // Sem espera
        nil,       // Argumentos adicionais
    )
    if err != nil {
        log.Println("Erro ao declarar fila RabbitMQ:", err)
        showRabbitMQSummary()
        return
    }

    // Publicar uma mensagem
    body := "Hello, RabbitMQ!"
    err = ch.Publish(
        "",     // Exchange
        q.Name, // Rota de chave (Queue)
        false,  // Obrigatório
        false,  // Imediato
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        })
    if err != nil {
        log.Println("Erro ao publicar mensagem:", err)
        showRabbitMQSummary()
        return
    }
    fmt.Println(string(green) + "Mensagem publicada com sucesso!" + string(reset))

    // Consumir mensagens
    msgs, err := ch.Consume(
        q.Name, // Nome da fila
        "",     // Nome do consumidor
        true,   // Auto-ack
        false,  // Exclusivo
        false,  // Sem espera
        false,  // Não local
        nil,    // Argumentos adicionais
    )
    if err != nil {
        log.Println("Erro ao consumir mensagens:", err)
        showRabbitMQSummary()
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    go func() {
        for d := range msgs {
            fmt.Printf("Mensagem recebida: %s\n", d.Body)
        }
    }()

    // Esperar um pouco para receber mensagens
    <-ctx.Done()

    showRabbitMQSummary()
}

func showRabbitMQSummary() {
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
    fmt.Println("    \"context\"")
    fmt.Println("    \"fmt\"")
    fmt.Println("    \"log\"")
    fmt.Println("    \"os\"")
    fmt.Println("    \"time\"")
    fmt.Println()
    fmt.Println("    \"github.com/joho/godotenv\"")
    fmt.Println("    \"github.com/streadway/amqp\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    // Carregar variáveis de ambiente do arquivo .env")
    fmt.Println("    err := godotenv.Load()")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao carregar o arquivo .env:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    rabbitmqURI := os.Getenv(\"RABBITMQ_URI\")")
    fmt.Println("    queueName := os.Getenv(\"QUEUE_NAME\")")
    fmt.Println()
    fmt.Println("    // Tentar conectar ao RabbitMQ")
    fmt.Println("    conn, err := amqp.Dial(rabbitmqURI)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao conectar ao RabbitMQ:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    defer conn.Close()")
    fmt.Println()
    fmt.Println("    // Criar um canal")
    fmt.Println("    ch, err := conn.Channel()")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao criar canal RabbitMQ:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    defer ch.Close()")
    fmt.Println()
    fmt.Println("    // Declarar uma fila")
    fmt.Println("    q, err := ch.QueueDeclare(")
    fmt.Println("        queueName, // Nome da fila")
    fmt.Println("        false,     // Durável")
    fmt.Println("        false,     // Deletar automaticamente")
    fmt.Println("        false,     // Exclusiva")
    fmt.Println("        false,     // Sem espera")
    fmt.Println("        nil,       // Argumentos adicionais")
    fmt.Println("    )")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao declarar fila RabbitMQ:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    // Publicar uma mensagem")
    fmt.Println("    body := \"Hello, RabbitMQ!\"")
    fmt.Println("    err = ch.Publish(")
    fmt.Println("        \"\",     // Exchange")
    fmt.Println("        q.Name, // Rota de chave (Queue)")
    fmt.Println("        false,  // Obrigatório")
    fmt.Println("        false,  // Imediato")
    fmt.Println("        amqp.Publishing{")
    fmt.Println("            ContentType: \"text/plain\",")
    fmt.Println("            Body:        []byte(body),")
    fmt.Println("        })")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao publicar mensagem:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"Mensagem publicada com sucesso!\")")
    fmt.Println()
    fmt.Println("    // Consumir mensagens")
    fmt.Println("    msgs, err := ch.Consume(")
    fmt.Println("        q.Name, // Nome da fila")
    fmt.Println("        \"\",     // Nome do consumidor")
    fmt.Println("        true,   // Auto-ack")
    fmt.Println("        false,  // Exclusivo")
    fmt.Println("        false,  // Sem espera")
    fmt.Println("        false,  // Não local")
    fmt.Println("        nil,    // Argumentos adicionais")
    fmt.Println("    )")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao consumir mensagens:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)")
    fmt.Println("    defer cancel()")
    fmt.Println()
    fmt.Println("    go func() {")
    fmt.Println("        for d := range msgs {")
    fmt.Println("            fmt.Printf(\"Mensagem recebida: %s\\n\", d.Body)")
    fmt.Println("        }")
    fmt.Println("    }()")
    fmt.Println()
    fmt.Println("    // Esperar um pouco para receber mensagens")
    fmt.Println("    <-ctx.Done()")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a integrar Golang com RabbitMQ usando variáveis de ambiente. Aprendemos a publicar e consumir mensagens de uma fila RabbitMQ." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use o pacote `github.com/joho/godotenv` para carregar variáveis de ambiente a partir de um arquivo .env.",
        "Curiosidade: RabbitMQ é um sistema de mensageria amplamente utilizado para comunicação assíncrona entre serviços.",
        "Dica: Use `os.Getenv` para obter valores das variáveis de ambiente em Go.",
        "Curiosidade: Mensagens em RabbitMQ são entregues de forma confiável através de filas duráveis.",
        "Dica: Sempre verifique erros ao interagir com RabbitMQ.",
        "Curiosidade: RabbitMQ suporta diferentes tipos de trocas como direto, tópico e fanout.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o vigésimo nível!" + string(reset))
}
