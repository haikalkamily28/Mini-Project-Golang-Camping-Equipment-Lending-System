package repository

import (
    "mini/entity"

    "gorm.io/gorm"
)

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *entity.User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) GetUserByEmail(email string) (*entity.User, error) {
    var user entity.User
    if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
