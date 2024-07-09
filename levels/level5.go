package levels

import (
    "bufio"
    "fmt"
    "go-code-quest/pkg/game"
    "os"
    "strconv"
    "strings"
    "math/rand"
    "time"
)

type Level5 struct{}

func StartLevel5() {
    game.StartLevel(Level5{})
}

func (l Level5) Start() {
    // Definir cores
    red := "\033[31m"
    green := "\033[32m"
    yellow := "\033[33m"
    blue := "\033[34m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 5: Estruturas e Métodos" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Crie uma estrutura para representar um aluno e métodos para calcular a média das notas e determinar se o aluno está aprovado" + string(reset))
    fmt.Println()

    reader := bufio.NewReader(os.Stdin)

    // Solicitar nome do aluno
    fmt.Print(string(blue) + "Digite o nome do aluno: " + string(reset))
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    // Solicitar notas do aluno
    fmt.Print(string(blue) + "Digite as notas do aluno separadas por espaço: " + string(reset))
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    gradeStrings := strings.Split(input, " ")
    grades := make([]float64, len(gradeStrings))

    for i, gs := range gradeStrings {
        g, err := strconv.ParseFloat(gs, 64)
        if err != nil {
            fmt.Println(string(red) + "Entrada inválida, usando nota padrão: 0.0" + string(reset))
            g = 0.0
        }
        grades[i] = g
    }

    // Criar instância da estrutura Student
    student := Student{name: name, grades: grades}

    // Calcular média e verificar aprovação
    avg := student.CalculateAverage()
    passed := student.IsPassed()

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resultados" + string(reset))
    fmt.Println(separator)
    fmt.Printf(string(green)+"Nome do aluno: %s\n"+string(reset), student.name)
    fmt.Printf(string(green)+"Média das notas: %.2f\n"+string(reset), avg)
    if passed {
        fmt.Println(string(green) + "O aluno está aprovado!" + string(reset))
    } else {
        fmt.Println(string(red) + "O aluno está reprovado." + string(reset))
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
    fmt.Println("    \"strings\"")
    fmt.Println("    \"strconv\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("type Student struct {")
    fmt.Println("    name   string")
    fmt.Println("    grades []float64")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("// Método para calcular a média das notas do aluno")
    fmt.Println("func (s Student) CalculateAverage() float64 {")
    fmt.Println("    sum := 0.0")
    fmt.Println("    for _, grade := range s.grades {")
    fmt.Println("        sum += grade")
    fmt.Println("    }")
    fmt.Println("    return sum / float64(len(s.grades))")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("// Método para verificar se o aluno está aprovado")
    fmt.Println("func (s Student) IsPassed() bool {")
    fmt.Println("    return s.CalculateAverage() >= 6.0")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    // Exemplo de entrada")
    fmt.Println("    name := \"John Doe\"")
    fmt.Println("    grades := []float64{7.0, 8.5, 6.0, 9.0}")
    fmt.Println()
    fmt.Println("    // Criar instância da estrutura Student")
    fmt.Println("    student := Student{name: name, grades: grades}")
    fmt.Println()
    fmt.Println("    // Calcular média e verificar aprovação")
    fmt.Println("    avg := student.CalculateAverage()")
    fmt.Println("    passed := student.IsPassed()")
    fmt.Println()
    fmt.Println("    // Exibir resultados")
    fmt.Println("    fmt.Printf(\"Nome do aluno: %s\\n\", student.name)")
    fmt.Println("    fmt.Printf(\"Média das notas: %.2f\\n\", avg)")
    fmt.Println("    if passed {")
    fmt.Println("        fmt.Println(\"O aluno está aprovado!\")")
    fmt.Println("    } else {")
    fmt.Println("        fmt.Println(\"O aluno está reprovado.\")")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a criar e usar estruturas (structs) e métodos em Go. Structs permitem agrupar dados relacionados, e métodos permitem associar funções a esses dados para manipulação e lógica." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use structs para agrupar dados relacionados em uma única entidade.",
        "Curiosidade: Métodos em Go são funções que têm um receptor, permitindo associar comportamento a structs.",
        "Dica: Use métodos para encapsular a lógica de negócios relacionada a uma struct.",
        "Curiosidade: Em Go, métodos podem ser definidos tanto para tipos de valor quanto para tipos de ponteiro.",
        "Dica: Use funções para manipular dados independentes e métodos para manipular dados associados a structs.",
        "Curiosidade: Go não tem classes, mas structs e métodos proporcionam uma maneira eficiente de modelar dados e comportamento.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o quinto nível!" + string(reset))
}

// Definir a estrutura Student
type Student struct {
    name   string
    grades []float64
}

// Método para calcular a média das notas do aluno
func (s Student) CalculateAverage() float64 {
    sum := 0.0
    for _, grade := range s.grades {
        sum += grade
    }
    return sum / float64(len(s.grades))
}

// Método para verificar se o aluno está aprovado
func (s Student) IsPassed() bool {
    return s.CalculateAverage() >= 6.0
}
