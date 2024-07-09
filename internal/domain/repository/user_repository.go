package repository

import "go-code-quest/internal/domain/entity"

type UserRepository interface {
    GetByID(id int) (*entity.User, error)
    Create(user *entity.User) error
}
