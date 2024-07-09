package levels

import (
    "bufio"
    "database/sql"
    "fmt"
    "go-code-quest/pkg/game"
    //"log"
    "os"
    "math/rand"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

type Level10 struct{}

func StartLevel10() {
    game.StartLevel(Level10{})
}

func (l Level10) Start() {
    // Definir cores
    red := "\033[31m"
    green := "\033[32m"
    yellow := "\033[33m"
		blue := "\033[34m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 10: Integração com Banco de Dados" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Conecte-se a um banco de dados, crie uma tabela, insira dados e execute uma consulta" + string(reset))
    fmt.Println()

    reader := bufio.NewReader(os.Stdin)

    // Solicitar parâmetros de conexão ao usuário
    fmt.Print(string(blue) + "Digite o usuário do banco de dados: " + string(reset))
    dbUser, _ := reader.ReadString('\n')
    dbUser = dbUser[:len(dbUser)-1]

    fmt.Print(string(blue) + "Digite a senha do banco de dados: " + string(reset))
    dbPassword, _ := reader.ReadString('\n')
    dbPassword = dbPassword[:len(dbPassword)-1]

    fmt.Print(string(blue) + "Digite o endereço do banco de dados (ex: 127.0.0.1:3306): " + string(reset))
    dbAddress, _ := reader.ReadString('\n')
    dbAddress = dbAddress[:len(dbAddress)-1]

    fmt.Print(string(blue) + "Digite o nome do banco de dados: " + string(reset))
    dbName, _ := reader.ReadString('\n')
    dbName = dbName[:len(dbName)-1]

    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbAddress, dbName)

    // Tentar conectar ao banco de dados
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println(string(red) + "Erro ao conectar ao banco de dados: " + string(reset), err)
    } else {
        defer db.Close()
        // Criar tabela
        _, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT, name VARCHAR(50), PRIMARY KEY (id))")
        if err != nil {
            fmt.Println(string(red) + "Erro ao criar tabela: " + string(reset), err)
        }

        // Inserir dados
        _, err = db.Exec("INSERT INTO users (name) VALUES ('Alice'), ('Bob'), ('Charlie')")
        if err != nil {
            fmt.Println(string(red) + "Erro ao inserir dados: " + string(reset), err)
        }

        // Consultar dados
        rows, err := db.Query("SELECT id, name FROM users")
        if err != nil {
            fmt.Println(string(red) + "Erro ao consultar dados: " + string(reset), err)
        } else {
            defer rows.Close()
            // Exibir resultados da consulta
            fmt.Println(string(green) + "Dados da tabela users:" + string(reset))
            for rows.Next() {
                var id int
                var name string
                err := rows.Scan(&id, &name)
                if err != nil {
                    fmt.Println(string(red) + "Erro ao escanear linha: " + string(reset), err)
                }
                fmt.Printf(string(green)+"ID: %d, Nome: %s\n"+string(reset), id, name)
            }
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
    fmt.Println("    \"database/sql\"")
    fmt.Println("    \"fmt\"")
    fmt.Println("    \"log\"")
    fmt.Println("    _ \"github.com/go-sql-driver/mysql\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    // Conectar ao banco de dados")
    fmt.Println("    db, err := sql.Open(\"mysql\", \"user:password@tcp(127.0.0.1:3306)/testdb\")")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatal(\"Erro ao conectar ao banco de dados: \", err)")
    fmt.Println("    }")
    fmt.Println("    defer db.Close()")
    fmt.Println()
    fmt.Println("    // Criar tabela")
    fmt.Println("    _, err = db.Exec(\"CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT, name VARCHAR(50), PRIMARY KEY (id))\")")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatal(\"Erro ao criar tabela: \", err)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    // Inserir dados")
    fmt.Println("    _, err = db.Exec(\"INSERT INTO users (name) VALUES ('Alice'), ('Bob'), ('Charlie')\")")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatal(\"Erro ao inserir dados: \", err)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    // Consultar dados")
    fmt.Println("    rows, err := db.Query(\"SELECT id, name FROM users\")")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatal(\"Erro ao consultar dados: \", err)")
    fmt.Println("    }")
    fmt.Println("    defer rows.Close()")
    fmt.Println()
    fmt.Println("    // Exibir resultados da consulta")
    fmt.Println("    for rows.Next() {")
    fmt.Println("        var id int")
    fmt.Println("        var name string")
    fmt.Println("        err := rows.Scan(&id, &name)")
    fmt.Println("        if err != nil {")
    fmt.Println("            log.Fatal(\"Erro ao escanear linha: \", err)")
    fmt.Println("        }")
    fmt.Println("        fmt.Printf(\"ID: %d, Nome: %s\\n\", id, name)")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a conectar-se a um banco de dados, criar tabelas, inserir dados e executar consultas em Go. Essas operações são fundamentais para muitas aplicações que dependem de persistência de dados." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use o pacote `database/sql` para conectar-se a bancos de dados SQL em Go.",
        "Curiosidade: Go suporta várias bibliotecas de drivers para diferentes bancos de dados, incluindo MySQL, PostgreSQL e SQLite.",
        "Dica: Sempre feche a conexão com o banco de dados usando `defer db.Close()`.",
        "Curiosidade: A integração com bancos de dados é uma habilidade essencial para desenvolvedores back-end.",
        "Dica: Use `db.Query` para executar consultas que retornam linhas e `db.Exec` para executar comandos que não retornam linhas.",
        "Curiosidade: Go facilita a manipulação de dados com suporte robusto para SQL e bibliotecas de ORM.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o décimo nível!" + string(reset))
}
