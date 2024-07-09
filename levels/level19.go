package levels

import (
    "bufio"
    "context"
    "fmt"
    "go-code-quest/pkg/game"
    "log"
    "math/rand"
    "os"
    "strings"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Level19 struct{}

func StartLevel19() {
    game.StartLevel(Level19{})
}

func (l Level19) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 19: Integração com MongoDB" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Conecte-se ao MongoDB, insira, consulte e manipule documentos" + string(reset))
    fmt.Println()

    // Obter informações de conexão ao banco de dados do usuário
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Digite o URI do MongoDB (ex: mongodb://localhost:27017): ")
    uri, _ := reader.ReadString('\n')
    uri = strings.TrimSpace(uri)

    // Configurações de conexão
    clientOptions := options.Client().ApplyURI(uri)

    // Tentar conectar ao MongoDB
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Println("Erro ao conectar ao MongoDB:", err)
        showMongoSummary()
        return
    }
    defer client.Disconnect(context.TODO())

    // Verificar a conexão
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Println("Erro ao conectar ao MongoDB:", err)
        showMongoSummary()
        return
    }
    fmt.Println(string(green) + "Conectado ao MongoDB com sucesso!" + string(reset))

    // Obter a coleção
    collection := client.Database("testdb").Collection("users")

    // Inserir documentos
    insertDocuments(collection)

    // Consultar documentos
    queryDocuments(collection)

    // Atualizar documentos
    updateDocuments(collection)

    // Excluir documentos
    deleteDocuments(collection)

    showMongoSummary()
}

func insertDocuments(collection *mongo.Collection) {
    documents := []interface{}{
        bson.D{{"name", "Alice"}, {"age", 30}},
        bson.D{{"name", "Bob"}, {"age", 25}},
    }

    insertResult, err := collection.InsertMany(context.TODO(), documents)
    if err != nil {
        log.Println("Erro ao inserir documentos:", err)
        return
    }
    fmt.Println("Documentos inseridos:", insertResult.InsertedIDs)
}

func queryDocuments(collection *mongo.Collection) {
    filter := bson.D{{"age", bson.D{{"$gt", 20}}}}

    cur, err := collection.Find(context.TODO(), filter)
    if err != nil {
        log.Println("Erro ao consultar documentos:", err)
        return
    }
    defer cur.Close(context.TODO())

    fmt.Println("Documentos encontrados:")
    for cur.Next(context.TODO()) {
        var result bson.D
        err := cur.Decode(&result)
        if err != nil {
            log.Println("Erro ao decodificar documento:", err)
            return
        }
        fmt.Println(result)
    }

    if err := cur.Err(); err != nil {
        log.Println("Erro durante a iteração dos documentos:", err)
    }
}

func updateDocuments(collection *mongo.Collection) {
    filter := bson.D{{"name", "Alice"}}
    update := bson.D{{"$set", bson.D{{"age", 32}}}}

    updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        log.Println("Erro ao atualizar documentos:", err)
        return
    }
    fmt.Println("Documentos atualizados:", updateResult.ModifiedCount)
}

func deleteDocuments(collection *mongo.Collection) {
    filter := bson.D{{"name", "Bob"}}

    deleteResult, err := collection.DeleteOne(context.TODO(), filter)
    if err != nil {
        log.Println("Erro ao excluir documentos:", err)
        return
    }
    fmt.Println("Documentos excluídos:", deleteResult.DeletedCount)
}

func showMongoSummary() {
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
    fmt.Println("    \"go.mongodb.org/mongo-driver/bson\"")
    fmt.Println("    \"go.mongodb.org/mongo-driver/mongo\"")
    fmt.Println("    \"go.mongodb.org/mongo-driver/mongo/options\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    // Configurações de conexão")
    fmt.Println("    uri := \"mongodb://localhost:27017\"")
    fmt.Println("    clientOptions := options.Client().ApplyURI(uri)")
    fmt.Println()
    fmt.Println("    // Tentar conectar ao MongoDB")
    fmt.Println("    client, err := mongo.Connect(context.TODO(), clientOptions)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao conectar ao MongoDB:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    defer client.Disconnect(context.TODO())")
    fmt.Println()
    fmt.Println("    // Verificar a conexão")
    fmt.Println("    err = client.Ping(context.TODO(), nil)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao conectar ao MongoDB:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"Conectado ao MongoDB com sucesso!\")")
    fmt.Println()
    fmt.Println("    // Obter a coleção")
    fmt.Println("    collection := client.Database(\"testdb\").Collection(\"users\")")
    fmt.Println()
    fmt.Println("    // Inserir documentos")
    fmt.Println("    insertDocuments(collection)")
    fmt.Println()
    fmt.Println("    // Consultar documentos")
    fmt.Println("    queryDocuments(collection)")
    fmt.Println()
    fmt.Println("    // Atualizar documentos")
    fmt.Println("    updateDocuments(collection)")
    fmt.Println()
    fmt.Println("    // Excluir documentos")
    fmt.Println("    deleteDocuments(collection)")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func insertDocuments(collection *mongo.Collection) {")
    fmt.Println("    documents := []interface{}{")
    fmt.Println("        bson.D{{\"name\", \"Alice\"}, {\"age\", 30}},")
    fmt.Println("        bson.D{{\"name\", \"Bob\"}, {\"age\", 25}},")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    insertResult, err := collection.InsertMany(context.TODO(), documents)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao inserir documentos:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"Documentos inseridos:\", insertResult.InsertedIDs)")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func queryDocuments(collection *mongo.Collection) {")
    fmt.Println("    filter := bson.D{{\"age\", bson.D{{\"$gt\", 20}}}}")
    fmt.Println()
    fmt.Println("    cur, err := collection.Find(context.TODO(), filter)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao consultar documentos:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    defer cur.Close(context.TODO())")
    fmt.Println()
    fmt.Println("    fmt.Println(\"Documentos encontrados:\")")
    fmt.Println("    for cur.Next(context.TODO()) {")
    fmt.Println("        var result bson.D")
    fmt.Println("        err := cur.Decode(&result)")
    fmt.Println("        if err != nil {")
    fmt.Println("            log.Println(\"Erro ao decodificar documento:\", err)")
    fmt.Println("            return")
    fmt.Println("        }")
    fmt.Println("        fmt.Println(result)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    if err := cur.Err(); err != nil {")
    fmt.Println("        log.Println(\"Erro durante a iteração dos documentos:\", err)")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func updateDocuments(collection *mongo.Collection) {")
    fmt.Println("    filter := bson.D{{\"name\", \"Alice\"}}")
    fmt.Println("    update := bson.D{{\"$set\", bson.D{{\"age\", 32}}}}")
    fmt.Println()
    fmt.Println("    updateResult, err := collection.UpdateOne(context.TODO(), filter, update)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao atualizar documentos:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"Documentos atualizados:\", updateResult.ModifiedCount)")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func deleteDocuments(collection *mongo.Collection) {")
    fmt.Println("    filter := bson.D{{\"name\", \"Bob\"}}")
    fmt.Println()
    fmt.Println("    deleteResult, err := collection.DeleteOne(context.TODO(), filter)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao excluir documentos:\", err)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"Documentos excluídos:\", deleteResult.DeletedCount)")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a integrar Golang com MongoDB. Aprendemos a conectar ao MongoDB, inserir, consultar, atualizar e excluir documentos. A integração com MongoDB permite manipular grandes volumes de dados não estruturados de forma eficiente." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use o pacote `go.mongodb.org/mongo-driver/mongo` para integrar com MongoDB.",
        "Curiosidade: MongoDB é um banco de dados NoSQL amplamente utilizado para armazenamento de dados não estruturados.",
        "Dica: Use `context.TODO()` quando não tiver um contexto específico.",
        "Curiosidade: MongoDB armazena dados em formato BSON (Binary JSON).",
        "Dica: Sempre verifique erros ao interagir com MongoDB.",
        "Curiosidade: MongoDB suporta consultas avançadas e índices para melhorar o desempenho.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o décimo nono nível!" + string(reset))
}
