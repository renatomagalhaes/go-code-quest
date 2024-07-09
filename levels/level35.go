package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
		"math/rand"
		"time"
    "log"
    //"os"

    "github.com/spf13/viper"
)

type Level35 struct{}

func StartLevel35() {
    game.StartLevel(Level35{})
}

func (l Level35) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 35: Configurações de Aplicações com Viper" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Aprenda a gerenciar configurações em suas aplicações Go usando o pacote Viper" + string(reset))
    fmt.Println()

    // Configurar Viper
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")

    // Ler arquivo de configuração
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Erro ao ler arquivo de configuração: %v", err)
    }

    // Configurar variáveis de ambiente
    viper.SetEnvPrefix("app")
    viper.AutomaticEnv()

    // Acessar valores de configuração
    serverPort := viper.GetInt("server.port")
    dbHost := viper.GetString("database.host")
    dbPort := viper.GetInt("database.port")

    fmt.Printf(string(green)+"Servidor rodando na porta: %d\n"+string(reset), serverPort)
    fmt.Printf(string(green)+"Banco de dados rodando em: %s:%d\n"+string(reset), dbHost, dbPort)

    showViperSummary()
}

func showViperSummary() {
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
    fmt.Println()
    fmt.Println("    \"github.com/spf13/viper\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    // Configurar Viper")
    fmt.Println("    viper.SetConfigName(\"config\")")
    fmt.Println("    viper.SetConfigType(\"yaml\")")
    fmt.Println("    viper.AddConfigPath(\".\")")
    fmt.Println()
    fmt.Println("    // Ler arquivo de configuração")
    fmt.Println("    if err := viper.ReadInConfig(); err != nil {")
    fmt.Println("        log.Fatalf(\"Erro ao ler arquivo de configuração: %v\", err)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    // Configurar variáveis de ambiente")
    fmt.Println("    viper.SetEnvPrefix(\"app\")")
    fmt.Println("    viper.AutomaticEnv()")
    fmt.Println()
    fmt.Println("    // Acessar valores de configuração")
    fmt.Println("    serverPort := viper.GetInt(\"server.port\")")
    fmt.Println("    dbHost := viper.GetString(\"database.host\")")
    fmt.Println("    dbPort := viper.GetInt(\"database.port\")")
    fmt.Println()
    fmt.Println("    fmt.Printf(\"Servidor rodando na porta: %d\\n\", serverPort)")
    fmt.Println("    fmt.Printf(\"Banco de dados rodando em: %s:%d\\n\", dbHost, dbPort)")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a usar o pacote Viper para gerenciar configurações em suas aplicações Go, carregando valores de arquivos de configuração e variáveis de ambiente." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use Viper para centralizar o gerenciamento de configurações.",
        "Curiosidade: Viper suporta vários formatos de arquivo, como JSON, TOML, YAML, HCL e ENV.",
        "Dica: Use `viper.AutomaticEnv` para carregar automaticamente variáveis de ambiente.",
        "Curiosidade: Viper permite a sobrecarga de configurações a partir de múltiplas fontes.",
        "Dica: Use `viper.SetDefault` para definir valores padrão para as configurações.",
        "Curiosidade: Viper pode ser usado junto com o pacote Cobra para gerenciar configurações em aplicações CLI.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o trigésimo quinto nível!" + string(reset))
}
