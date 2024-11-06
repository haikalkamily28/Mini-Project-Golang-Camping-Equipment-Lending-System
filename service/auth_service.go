package service

import (
    "errors"
    "mini/entity"
    "mini/repository"
    "golang.org/x/crypto/bcrypt"
)

type UserService interface {
    Register(user *entity.User) error
}

type userService struct {
    userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
    return &userService{userRepo: repo}
}

// Fungsi untuk meng-hash password
func hashPassword(password string) (string, error) {
    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedBytes), nil
}

func (s *userService) Register(user *entity.User) error {
    // Periksa apakah email sudah terdaftar
    existingUser, _ := s.userRepo.GetUserByEmail(user.Email)
    if existingUser != nil {
        return errors.New("email already registered")
    }

    // Hash password
    hashedPassword, err := hashPassword(user.Password)
    if err != nil {
        return err
    }
    user.Password = hashedPassword

    // Simpan user
    return s.userRepo.CreateUser(user)
}
