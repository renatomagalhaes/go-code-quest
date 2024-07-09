package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
		"math/rand"
    "log"
    "net/http"
    "time"
)

type Level28 struct{}

func StartLevel28() {
    game.StartLevel(Level28{})
}

func (l Level28) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 28: Implementação de Middlewares em Go" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Crie e use middlewares para gerenciar requisições HTTP" + string(reset))
    fmt.Println()

    // Configurar servidor HTTP com middlewares
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Bem-vindo ao Level 28!")
    })

    // Aplicar middlewares
    loggedMux := loggingMiddleware(mux)
    authMux := authenticationMiddleware(loggedMux)

    // Iniciar servidor HTTP
    fmt.Println(string(green) + "Servidor HTTP iniciado em http://localhost:8080" + string(reset))
    log.Fatal(http.ListenAndServe(":8080", authMux))

    showMiddlewareSummary()
}

// Middleware de logging
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}

// Middleware de autenticação
func authenticationMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Simulação de autenticação
        user := r.Header.Get("X-User")
        if user == "" {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func showMiddlewareSummary() {
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
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    mux := http.NewServeMux()")
    fmt.Println("    mux.HandleFunc(\"/\", func(w http.ResponseWriter, r *http.Request) {")
    fmt.Println("        fmt.Fprintln(w, \"Bem-vindo ao Level 28!\")")
    fmt.Println("    })")
    fmt.Println()
    fmt.Println("    loggedMux := loggingMiddleware(mux)")
    fmt.Println("    authMux := authenticationMiddleware(loggedMux)")
    fmt.Println()
    fmt.Println("    fmt.Println(\"Servidor HTTP iniciado em http://localhost:8080\")")
    fmt.Println("    log.Fatal(http.ListenAndServe(\":8080\", authMux))")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func loggingMiddleware(next http.Handler) http.Handler {")
    fmt.Println("    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {")
    fmt.Println("        start := time.Now()")
    fmt.Println("        next.ServeHTTP(w, r)")
    fmt.Println("        log.Printf(\"%s %s %v\", r.Method, r.URL.Path, time.Since(start))")
    fmt.Println("    })")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func authenticationMiddleware(next http.Handler) http.Handler {")
    fmt.Println("    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {")
    fmt.Println("        user := r.Header.Get(\"X-User\")")
    fmt.Println("        if user == \"\" {")
    fmt.Println("            http.Error(w, \"Forbidden\", http.StatusForbidden)")
    fmt.Println("            return")
    fmt.Println("        }")
    fmt.Println("        next.ServeHTTP(w, r)")
    fmt.Println("    })")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a criar e usar middlewares para gerenciar requisições HTTP em Go. Middlewares podem ser usados para autenticação, logging, e tratamento de erros." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Middlewares são funções que envolvem manipuladores de requisições HTTP para adicionar funcionalidades.",
        "Curiosidade: Middlewares são amplamente usados em frameworks web como Express.js e Django.",
        "Dica: Use o padrão de função `http.HandlerFunc` para criar middlewares em Go.",
        "Curiosidade: Middlewares podem ser compostos, permitindo a criação de pipelines de processamento de requisições.",
        "Dica: Middlewares podem melhorar a modularidade e reutilização de código em aplicações web.",
        "Curiosidade: O pacote `net/http` do Go facilita a criação de middlewares com sua interface `http.Handler`.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o vigésimo oitavo nível!" + string(reset))
}
