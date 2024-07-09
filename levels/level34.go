package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
		"math/rand"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

type Level34 struct{}

func StartLevel34() {
    game.StartLevel(Level34{})
}

func (l Level34) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 34: Monitoramento e Logging em Go" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Aprenda a implementar monitoramento e logging eficaz em suas aplicações Go" + string(reset))
    fmt.Println()

    // Logging básico
    logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Erro ao abrir arquivo de log: %v", err)
    }
    log.SetOutput(logFile)
    log.Println("Aplicação iniciada")

    // Configurar métricas Prometheus
    requests := prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Número total de requisições HTTP",
        },
        []string{"path"},
    )
    prometheus.MustRegister(requests)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println("Requisição recebida")
        requests.WithLabelValues(r.URL.Path).Inc()
        fmt.Fprintln(w, "Olá, Go Code Quest!")
    })

    // Expor métricas via HTTP
    http.Handle("/metrics", promhttp.Handler())

    go func() {
        fmt.Println(string(green) + "Servidor HTTP iniciado em http://localhost:8080" + string(reset))
        if err := http.ListenAndServe(":8080", nil); err != nil {
            log.Fatalf("Erro ao iniciar servidor HTTP: %v", err)
        }
    }()

    time.Sleep(2 * time.Second) // Simulação de tempo de execução

    log.Println("Aplicação finalizada")

    showMonitoringSummary()
}

func showMonitoringSummary() {
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
    fmt.Println("    \"os\"")
    fmt.Println("    \"time\"")
    fmt.Println()
    fmt.Println("    \"github.com/prometheus/client_golang/prometheus\"")
    fmt.Println("    \"github.com/prometheus/client_golang/prometheus/promhttp\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    logFile, err := os.OpenFile(\"app.log\", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"Erro ao abrir arquivo de log: %v\", err)")
    fmt.Println("    }")
    fmt.Println("    log.SetOutput(logFile)")
    fmt.Println("    log.Println(\"Aplicação iniciada\")")
    fmt.Println()
    fmt.Println("    requests := prometheus.NewCounterVec(")
    fmt.Println("        prometheus.CounterOpts{")
    fmt.Println("            Name: \"http_requests_total\",")
    fmt.Println("            Help: \"Número total de requisições HTTP\",")
    fmt.Println("        },")
    fmt.Println("        []string{\"path\"},")
    fmt.Println("    )")
    fmt.Println("    prometheus.MustRegister(requests)")
    fmt.Println()
    fmt.Println("    http.HandleFunc(\"/\", func(w http.ResponseWriter, r *http.Request) {")
    fmt.Println("        log.Println(\"Requisição recebida\")")
    fmt.Println("        requests.WithLabelValues(r.URL.Path).Inc()")
    fmt.Println("        fmt.Fprintln(w, \"Olá, Go Code Quest!\")")
    fmt.Println("    })")
    fmt.Println()
    fmt.Println("    http.Handle(\"/metrics\", promhttp.Handler())")
    fmt.Println()
    fmt.Println("    go func() {")
    fmt.Println("        fmt.Println(\"Servidor HTTP iniciado em http://localhost:8080\")")
    fmt.Println("        if err := http.ListenAndServe(\":8080\", nil); err != nil {")
    fmt.Println("            log.Fatalf(\"Erro ao iniciar servidor HTTP: %v\", err)")
    fmt.Println("        }")
    fmt.Println("    }()")
    fmt.Println()
    fmt.Println("    time.Sleep(2 * time.Second)")
    fmt.Println()
    fmt.Println("    log.Println(\"Aplicação finalizada\")")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a implementar monitoramento e logging em aplicações Go usando Prometheus para métricas e o pacote `log` para logging." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use arquivos de log para armazenar logs de maneira persistente.",
        "Curiosidade: Prometheus é uma ferramenta de monitoramento e alerta de código aberto.",
        "Dica: Use `log.SetOutput` para redirecionar logs para um arquivo.",
        "Curiosidade: O Prometheus foi originalmente desenvolvido pela SoundCloud.",
        "Dica: Use métricas personalizadas para monitorar aspectos específicos da sua aplicação.",
        "Curiosidade: O formato de logs estruturados facilita a análise e o monitoramento.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o trigésimo quarto nível!" + string(reset))
}
