package levels

import (
    "fmt"
    //"go-code-quest/internal/domain/entity"
    "go-code-quest/internal/infrastructure/config"
    "go-code-quest/internal/infrastructure/persistence/mysql"
    "go-code-quest/internal/interface/http/handler"
    "go-code-quest/internal/interface/http/router"
    "go-code-quest/internal/usecase"
    "go-code-quest/pkg/game"
	"math/rand"
	"time"
    "log"
    "net/http"
)

type Level30 struct{}

func StartLevel30() {
    game.StartLevel(Level30{})
}

func (l Level30) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 30: Aplicação de Clean Architecture em Go" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Estruture um projeto Go seguindo a Clean Architecture" + string(reset))
    fmt.Println()

    // Configurar a conexão com o banco de dados
    db := config.NewDB()
    defer db.Close()

    // Configurar repositório, caso de uso e manipulador
    userRepo := mysql.NewUserRepository(db)
    userUseCase := &usecase.UserUseCase{UserRepo: userRepo}
    userHandler := &handler.UserHandler{UserUseCase: userUseCase}

    // Configurar e iniciar o servidor HTTP
    r := router.NewRouter(userHandler)
    fmt.Println(string(green) + "Servidor HTTP iniciado em http://localhost:8080" + string(reset))
    log.Fatal(http.ListenAndServe(":8080", r))

    showCleanArchitectureSummary()
}

func showCleanArchitectureSummary() {
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
    fmt.Println("package entity")
    fmt.Println()
    fmt.Println("type User struct {")
    fmt.Println("    ID   int")
    fmt.Println("    Name string")
    fmt.Println("    Age  int")
    fmt.Println("}")
    fmt.Println("```")
    fmt.Println()
    fmt.Println(separator)
    fmt.Println("```go")
    fmt.Println("package repository")
    fmt.Println()
    fmt.Println("import \"go-code-quest/internal/domain/entity\"")
    fmt.Println()
    fmt.Println("type UserRepository interface {")
    fmt.Println("    GetByID(id int) (*entity.User, error)")
    fmt.Println("    Create(user *entity.User) error")
    fmt.Println("}")
    fmt.Println("```")
    fmt.Println()
    fmt.Println(separator)
    fmt.Println("```go")
    fmt.Println("package usecase")
    fmt.Println()
    fmt.Println("import (")
    fmt.Println("    \"go-code-quest/internal/domain/entity\"")
    fmt.Println("    \"go-code-quest/internal/domain/repository\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("type UserUseCase struct {")
    fmt.Println("    UserRepo repository.UserRepository")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func (uc *UserUseCase) GetUserByID(id int) (*entity.User, error) {")
    fmt.Println("    return uc.UserRepo.GetByID(id)")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func (uc *UserUseCase) CreateUser(user *entity.User) error {")
    fmt.Println("    return uc.UserRepo.Create(user)")
    fmt.Println("}")
    fmt.Println("```")
    fmt.Println()
    fmt.Println(separator)
    fmt.Println("```go")
    fmt.Println("package mysql")
    fmt.Println()
    fmt.Println("import (")
    fmt.Println("    \"database/sql\"")
    fmt.Println("    \"go-code-quest/internal/domain/entity\"")
    fmt.Println("    \"go-code-quest/internal/domain/repository\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("type UserRepository struct {")
    fmt.Println("    DB *sql.DB")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func NewUserRepository(db *sql.DB) repository.UserRepository {")
    fmt.Println("    return &UserRepository{DB: db}")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func (repo *UserRepository) GetByID(id int) (*entity.User, error) {")
    fmt.Println("    user := &entity.User{}")
    fmt.Println("    err := repo.DB.QueryRow(\"SELECT id, name, age FROM users WHERE id = ?\", id).Scan(&user.ID, &user.Name, &user.Age)")
    fmt.Println("    if err != nil {")
    fmt.Println("        return nil, err")
    fmt.Println("    }")
    fmt.Println("    return user, nil")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func (repo *UserRepository) Create(user *entity.User) error {")
    fmt.Println("    _, err := repo.DB.Exec(\"INSERT INTO users (name, age) VALUES (?, ?)\", user.Name, user.Age)")
    fmt.Println("    return err")
    fmt.Println("}")
    fmt.Println("```")
    fmt.Println()
    fmt.Println(separator)
    fmt.Println("```go")
    fmt.Println("package handler")
    fmt.Println()
    fmt.Println("import (")
    fmt.Println("    \"encoding/json\"")
    fmt.Println("    \"go-code-quest/internal/domain/entity\"")
    fmt.Println("    \"go-code-quest/internal/usecase\"")
    fmt.Println("    \"net/http\"")
    fmt.Println("    \"strconv\"")
    fmt.Println()
    fmt.Println("    \"github.com/gorilla/mux\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("type UserHandler struct {")
    fmt.Println("    UserUseCase *usecase.UserUseCase")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {")
    fmt.Println("    vars := mux.Vars(r)")
    fmt.Println("    id, err := strconv.Atoi(vars[\"id\"])")
    fmt.Println("    if err != nil {")
    fmt.Println("        http.Error(w, \"Invalid ID\", http.StatusBadRequest)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    user, err := h.UserUseCase.GetUserByID(id)")
    fmt.Println("    if err != nil {")
    fmt.Println("        http.Error(w, \"User not found\", http.StatusNotFound)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    json.NewEncoder(w).Encode(user)")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {")
    fmt.Println("    var user entity.User")
    fmt.Println("    err := json.NewDecoder(r.Body).Decode(&user)")
    fmt.Println("    if err != nil {")
    fmt.Println("        http.Error(w, \"Invalid input\", http.StatusBadRequest)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    err = h.UserUseCase.CreateUser(&user)")
    fmt.Println("    if err != nil {")
    fmt.Println("        http.Error(w, \"Failed to create user\", http.StatusInternalServerError)")
    fmt.Println("        return")
    fmt.Println("    }")
    fmt.Println("    w.WriteHeader(http.StatusCreated)")
    fmt.Println("    json.NewEncoder(w).Encode(user)")
    fmt.Println("}")
    fmt.Println("```")
    fmt.Println()
    fmt.Println(separator)
    fmt.Println("```go")
    fmt.Println("package config")
    fmt.Println()
    fmt.Println("import (")
    fmt.Println("    \"database/sql\"")
    fmt.Println("    \"log\"")
    fmt.Println()
    fmt.Println("    _ \"github.com/go-sql-driver/mysql\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func NewDB() *sql.DB {")
    fmt.Println("    db, err := sql.Open(\"mysql\", \"user:password@tcp(localhost:3306)/dbname\")")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"failed to connect to database: %v\", err)")
    fmt.Println("    }")
    fmt.Println("    return db")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a estruturar um projeto Go seguindo os princípios da Clean Architecture. Isso inclui a separação das camadas de domínio, aplicação e infraestrutura, e a implementação de casos de uso e interfaces." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: A Clean Architecture facilita a manutenção e escalabilidade do código.",
        "Curiosidade: A Clean Architecture foi popularizada por Robert C. Martin (Uncle Bob).",
        "Dica: Separe o código em camadas distintas para melhor organização e teste.",
        "Curiosidade: A Clean Architecture permite trocar detalhes técnicos (como o banco de dados) sem impactar as regras de negócio.",
        "Dica: Use interfaces para definir contratos entre as camadas.",
        "Curiosidade: A Clean Architecture promove a inversão de dependências, onde os detalhes dependem das abstrações e não o contrário.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o trigésimo nível!" + string(reset))
}
