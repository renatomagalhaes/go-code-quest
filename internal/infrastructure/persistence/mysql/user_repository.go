package mysql

import (
    "database/sql"
    "go-code-quest/internal/domain/entity"
    "go-code-quest/internal/domain/repository"
)

type UserRepository struct {
    DB *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
    return &UserRepository{DB: db}
}

func (repo *UserRepository) GetByID(id int) (*entity.User, error) {
    user := &entity.User{}
    err := repo.DB.QueryRow("SELECT id, name, age FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Age)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (repo *UserRepository) Create(user *entity.User) error {
    _, err := repo.DB.Exec("INSERT INTO users (name, age) VALUES (?, ?)", user.Name, user.Age)
    return err
}
