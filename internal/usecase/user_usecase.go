package usecase

import (
    "go-code-quest/internal/domain/entity"
    "go-code-quest/internal/domain/repository"
)

type UserUseCase struct {
    UserRepo repository.UserRepository
}

func (uc *UserUseCase) GetUserByID(id int) (*entity.User, error) {
    return uc.UserRepo.GetByID(id)
}

func (uc *UserUseCase) CreateUser(user *entity.User) error {
    return uc.UserRepo.Create(user)
}
