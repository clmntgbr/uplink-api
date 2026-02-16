package service

import (
	"errors"
	"time"
	"uplink-api/config"
	"uplink-api/domain"
	"uplink-api/repository"
	"uplink-api/dto"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticateService struct {
	userRepo    *repository.UserRepository
	projectRepo *repository.ProjectRepository
	config      *config.Config
}

type JWTClaims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	jwt.RegisteredClaims
}

func NewAuthenticateService(userRepo *repository.UserRepository, projectRepo *repository.ProjectRepository, cfg *config.Config) *AuthenticateService {
	return &AuthenticateService{
		userRepo: userRepo,
		projectRepo: projectRepo,
		config:   cfg,
	}
}

func (s *AuthenticateService) Login(loginInput *dto.LoginInput) (*dto.LoginOutput, error) {
	user, err := s.userRepo.FindByEmail(loginInput.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	token, err := s.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &dto.LoginOutput{
		Token: token,
		User:  *user,
	}, nil
}


func (s *AuthenticateService) GenerateToken(user *domain.User) (string, error) {
	claims := JWTClaims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.config.JWTExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "uplink-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.JWTSecret))
}