package levels

import (
    "bufio"
    "database/sql"
    "fmt"
    "go-code-quest/pkg/game"
    "html/template"
    "log"
    "os"
    "strings"
	"math/rand"
	"time"

    _ "github.com/go-sql-driver/mysql"
)

type Level17 struct{}

func StartLevel17() {
    game.StartLevel(Level17{})
}

func (l Level17) Start() {
    // Definir cores
		red := "\033[31m"
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 17: Uso de Templates SQL" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Use templates SQL para criar consultas dinâmicas e manipular dados" + string(reset))
    fmt.Println()

    // Obter informações de conexão ao banco de dados do usuário
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Digite o usuário do banco de dados: ")
    user, _ := reader.ReadString('\n')
    user = strings.TrimSpace(user)

    fmt.Print("Digite a senha do banco de dados: ")
    password, _ := reader.ReadString('\n')
    password = strings.TrimSpace(password)

    fmt.Print("Digite o host do banco de dados: ")
    host, _ := reader.ReadString('\n')
    host = strings.TrimSpace(host)

    fmt.Print("Digite a porta do banco de dados: ")
    port, _ := reader.ReadString('\n')
    port = strings.TrimSpace(port)

    fmt.Print("Digite o nome do banco de dados: ")
    dbname, _ := reader.ReadString('\n')
    dbname = strings.TrimSpace(dbname)

    // Montar a string de conexão
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

    // Tentar conectar ao banco de dados
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Println(string(red) + "Erro ao conectar ao banco de dados:" + string(reset), err)
        showSummary()
        return
    }
    defer db.Close()

    // Verificar a conexão
    err = db.Ping()
    if err != nil {
        log.Println(string(red) + "Erro ao conectar ao banco de dados:" + string(reset), err)
        showSummary()
        return
    }
    fmt.Println(string(green) + "Conectado ao banco de dados com sucesso!" + string(reset))

    // Criar tabelas de exemplo
    createTables(db)

    // Inserir dados de exemplo
    insertData(db)

    // Criar templates SQL
    simpleQueryTemplate, err := template.New("simpleQuery").Parse(`SELECT id, name, age FROM users WHERE age > {{.Age}};`)
    if err != nil {
        log.Fatal(err)
    }

    joinQueryTemplate, err := template.New("joinQuery").Parse(`SELECT u.name, u.age, o.order_id, o.amount 
                                                               FROM users u 
                                                               INNER JOIN orders o ON u.id = o.user_id 
                                                               WHERE o.amount > {{.Amount}};`)
    if err != nil {
        log.Fatal(err)
    }

    // Dados para os templates
    simpleQueryData := struct {
        Age int
    }{
        Age: 25,
    }

    joinQueryData := struct {
        Amount float64
    }{
        Amount: 50.0,
    }

    // Executar consultas usando os templates
    executeTemplateQuery(db, simpleQueryTemplate, simpleQueryData)
    executeTemplateQuery(db, joinQueryTemplate, joinQueryData)

    showSummary()
}

func showSummary() {
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
    fmt.Println("    \"bufio\"")
    fmt.Println("    \"database/sql\"")
    fmt.Println("    \"fmt\"")
    fmt.Println("    \"html/template\"")
    fmt.Println("    \"log\"")
    fmt.Println("    \"os\"")
    fmt.Println("    \"strings\"")
    fmt.Println()
    fmt.Println("    _ \"github.com/go-sql-driver/mysql\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    // Obter informações de conexão ao banco de dados do usuário")
    fmt.Println("    reader := bufio.NewReader(os.Stdin)")
    fmt.Println("    fmt.Print(\"Digite o usuário do banco de dados: \")")
    fmt.Println("    user, _ := reader.ReadString('\\n')")
    fmt.Println("    user = strings.TrimSpace(user)")
    fmt.Println()
    fmt.Println("    fmt.Print(\"Digite a senha do banco de dados: \")")
    fmt.Println("    password, _ := reader.ReadString('\\n')")
    fmt.Println("    password = strings.TrimSpace(password)")
    fmt.Println()
    fmt.Println("    fmt.Print(\"Digite o host do banco de dados: \")")
    fmt.Println("    host, _ := reader.ReadString('\\n')")
    fmt.Println("    host = strings.TrimSpace(host)")
    fmt.Println()
    fmt.Println("    fmt.Print(\"Digite a porta do banco de dados: \")")
    fmt.Println("    port, _ := reader.ReadString('\\n')")
    fmt.Println("    port = strings.TrimSpace(port)")
    fmt.Println()
    fmt.Println("    fmt.Print(\"Digite o nome do banco de dados: \")")
    fmt.Println("    dbname, _ := reader.ReadString('\\n')")
    fmt.Println("    dbname = strings.TrimSpace(dbname)")
    fmt.Println()
    fmt.Println("    // Montar a string de conexão")
    fmt.Println("    dsn := fmt.Sprintf(\"%s:%s@tcp(%s:%s)/%s\", user, password, host, port, dbname)")
    fmt.Println()
    fmt.Println("    // Tentar conectar ao banco de dados")
    fmt.Println("    db, err := sql.Open(\"mysql\", dsn)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao conectar ao banco de dados:\", err)")
    fmt.Println("        showSummary()")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    defer db.Close()")
    fmt.Println()
    fmt.Println("    // Verificar a conexão")
    fmt.Println("    err = db.Ping()")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Println(\"Erro ao conectar ao banco de dados:\", err)")
    fmt.Println("        showSummary()")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"Conectado ao banco de dados com sucesso!\")")
    fmt.Println()
    fmt.Println("    // Criar tabelas de exemplo")
    fmt.Println("    createTables(db)")
    fmt.Println()
    fmt.Println("    // Inserir dados de exemplo")
    fmt.Println("    insertData(db)")
    fmt.Println()
    fmt.Println("    // Criar templates SQL")
    fmt.Println("    simpleQueryTemplate, err := template.New(\"simpleQuery\").Parse(`SELECT id, name, age FROM users WHERE age > {{.Age}};`)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatal(err)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    joinQueryTemplate, err := template.New(\"joinQuery\").Parse(`SELECT u.name, u.age, o.order_id, o.amount ")
    fmt.Println("                                                               FROM users u ")
    fmt.Println("                                                               INNER JOIN orders o ON u.id = o.user_id ")
    fmt.Println("                                                               WHERE o.amount > {{.Amount}};`)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatal(err)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    // Dados para os templates")
    fmt.Println("    simpleQueryData := struct {")
    fmt.Println("        Age int")
    fmt.Println("    }{")
    fmt.Println("        Age: 25,")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    joinQueryData := struct {")
    fmt.Println("        Amount float64")
    fmt.Println("    }{")
    fmt.Println("        Amount: 50.0,")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    // Executar consultas usando os templates")
    fmt.Println("    executeTemplateQuery(db, simpleQueryTemplate, simpleQueryData)")
    fmt.Println("    executeTemplateQuery(db, joinQueryTemplate, joinQueryData)")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func executeTemplateQuery(db *sql.DB, tmpl *template.Template, data interface{}) {")
    fmt.Println("    var sqlQuery strings.Builder")
    fmt.Println("    err := tmpl.Execute(&sqlQuery, data)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatal(err)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    rows, err := db.Query(sqlQuery.String())")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatal(err)")
    fmt.Println("    }")
    fmt.Println("    defer rows.Close()")
    fmt.Println()
    fmt.Println("    columns, err := rows.Columns()")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatal(err)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    values := make([]sql.RawBytes, len(columns))")
    fmt.Println("    scanArgs := make([]interface{}, len(values))")
    fmt.Println("    for i := range values {")
    fmt.Println("        scanArgs[i] = &values[i]")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    fmt.Printf(\"Resultado da consulta %s:\\n\", tmpl.Name())")
    fmt.Println("    for rows.Next() {")
    fmt.Println("        err = rows.Scan(scanArgs...)")
    fmt.Println("        if err != nil {")
    fmt.Println("            log.Fatal(err)")
    fmt.Println("        }")
    fmt.Println()
    fmt.Println("        for i, col := range values {")
    fmt.Println("            fmt.Printf(\"%s: %s\\n\", columns[i], string(col))")
    fmt.Println("        }")
    fmt.Println("        fmt.Println()")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a usar templates SQL em Go para criar consultas dinâmicas. Templates SQL permitem criar consultas reutilizáveis e dinâmicas, facilitando a manipulação de dados." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use o pacote `html/template` para criar templates SQL.",
        "Curiosidade: Templates SQL permitem criar consultas dinâmicas de forma segura e eficiente.",
        "Dica: Use `template.New` para criar um novo template.",
        "Curiosidade: Você pode passar qualquer tipo de dados para um template, desde que ele implemente a interface `fmt.Stringer`.",
        "Dica: Sempre verifique erros retornados por operações com templates.",
        "Curiosidade: Templates em Go são seguros contra injeções SQL, pois os dados são escapados automaticamente.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o décimo sétimo nível!" + string(reset))
}

func createTables(db *sql.DB) {
    createUsersTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(50),
        age INT
    );`
    _, err := db.Exec(createUsersTableSQL)
    if err != nil {
        log.Fatal(err)
    }

    createOrdersTableSQL := `CREATE TABLE IF NOT EXISTS orders (
        order_id INT AUTO_INCREMENT PRIMARY KEY,
        user_id INT,
        amount DECIMAL(10, 2),
        FOREIGN KEY (user_id) REFERENCES users(id)
    );`
    _, err = db.Exec(createOrdersTableSQL)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Tabelas 'users' e 'orders' criadas ou já existentes.")
}

func insertData(db *sql.DB) {
    insertUsersSQL := `INSERT INTO users (name, age) VALUES ('Alice', 30), ('Bob', 25);`
    _, err := db.Exec(insertUsersSQL)
    if err != nil {
        log.Fatal(err)
    }

    insertOrdersSQL := `INSERT INTO orders (user_id, amount) VALUES (1, 100.00), (1, 75.50), (2, 50.00);`
    _, err = db.Exec(insertOrdersSQL)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Dados de exemplo inseridos nas tabelas 'users' e 'orders'.")
}

func executeTemplateQuery(db *sql.DB, tmpl *template.Template, data interface{}) {
    var sqlQuery strings.Builder
    err := tmpl.Execute(&sqlQuery, data)
    if err != nil {
        log.Fatal(err)
    }

    rows, err := db.Query(sqlQuery.String())
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    columns, err := rows.Columns()
    if err != nil {
        log.Fatal(err)
    }

    values := make([]sql.RawBytes, len(columns))
    scanArgs := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }

    fmt.Printf("Resultado da consulta %s:\n", tmpl.Name())
    for rows.Next() {
        err = rows.Scan(scanArgs...)
        if err != nil {
            log.Fatal(err)
        }

        for i, col := range values {
            fmt.Printf("%s: %s\n", columns[i], string(col))
        }
        fmt.Println()
    }
}
