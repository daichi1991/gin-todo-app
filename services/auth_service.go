package services

import (
	"fmt"
	"gin-todo-app/models"
	"gin-todo-app/repositories"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Signup(email string, password string) error
	Login(email string, password string) (*string, error)
	GetUserFromToken(tokenString string) (*models.User, error)
	UpdateUser(user models.User) (*models.User, error)
}

type AuthService struct {
	repository    repositories.IAuthRepository
	statusService IStatusService
}

func NewAuthService(repository repositories.IAuthRepository, statusService IStatusService) IAuthService {
	return &AuthService{
		repository:    repository,
		statusService: statusService,
	}
}

func (s *AuthService) Signup(email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	newUser, err := s.repository.CreateUser(user)
	if err != nil {
		return err
	}

	_, err = s.statusService.CreateDefaultStatus(newUser.ID)
	if err != nil {
		return err
	}

	return nil
}

// ログイン
func (s *AuthService) Login(email string, password string) (*string, error) {
	foundUser, err := s.repository.FindUser(email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	token, err := CreateToken(foundUser.ID, foundUser.Email)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func CreateToken(userId uint, email string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userId,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func (s *AuthService) GetUserFromToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header[""])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	var user *models.User
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, jwt.ErrTokenExpired
		}

		user, err = s.repository.FindUser(claims["email"].(string))
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}

func (s *AuthService) UpdateUser(user models.User) (*models.User, error) {
	targetUser, err := s.repository.FindUserByID(user.ID)
	if err != nil {
		return nil, err
	}
	result, err := s.repository.UpdateUser(*targetUser)
	if err != nil {
		return nil, err
	}
	return result, nil
}
