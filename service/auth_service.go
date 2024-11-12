package service

import (
    "errors"
    "mini/entity"
    "mini/repository"
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
    "time"
	"log"
)

type UserService interface {
    Register(user *entity.User) error
    Login(email, password string) (string, error)
}

type userService struct {
    userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
    return &userService{userRepo: repo}
}

func hashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Println("Error hashing password:", err)
        return "", err
    }
    return string(hashedPassword), nil
}

func (s *userService) Register(user *entity.User) error {
    existingUser, _ := s.userRepo.GetUserByEmail(user.Email)
    if existingUser != nil {
        return errors.New("email already registered")
    }

    hashedPassword, err := hashPassword(user.Password)
    if err != nil {
        return err
    }
    user.Password = hashedPassword

    return s.userRepo.CreateUser(user)
}

func checkPassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *userService) Login(email, password string) (string, error) {
    user, err := s.userRepo.GetUserByEmail(email)
    if err != nil {
        return "", errors.New("invalid email or password")
    }

    if err := checkPassword(user.Password, password); err != nil {
        return "", errors.New("invalid email or password")
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString([]byte("aji"))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

