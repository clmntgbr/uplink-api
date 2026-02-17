package service

import (
	"context"
	"time"
	"uplink-api/config"
	"uplink-api/domain"
	"uplink-api/dto"
	"uplink-api/errors"
	"uplink-api/repository"

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
		userRepo:    userRepo,
		projectRepo: projectRepo,
		config:      cfg,
	}
}

func (s *AuthenticateService) Login(loginInput dto.LoginInput) (*dto.LoginOutput, error) {
	user, err := s.userRepo.FindByEmail(loginInput.Email)
	if err != nil {
		return nil, errors.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password)); err != nil {
		return nil, errors.ErrInvalidCredentials
	}

	token, err := s.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &dto.LoginOutput{
		Token: token,
		User:  dto.NewUserOutput(*user),
	}, nil
}

func (s *AuthenticateService) Register(ctx context.Context, registerInput dto.RegisterInput) (*dto.RegisterOutput, error) {
	existingUser, _ := s.userRepo.FindByEmail(registerInput.Email)
	if existingUser != nil {
		return nil, errors.ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	project := &domain.Project{
		Name: "Default Project",
	}

	if err := s.projectRepo.Create(ctx, project); err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:           registerInput.Email,
		Password:        string(hashedPassword),
		FirstName:       registerInput.FirstName,
		LastName:        registerInput.LastName,
		Avatar:          "https://avatar-placeholder.iran.liara.run/avatars/male",
		Projects:        []domain.Project{*project},
		ActiveProjectID: project.ID,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	token, err := s.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &dto.RegisterOutput{
		Token: token,
		User:  dto.NewUserOutput(*user),
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

func (s *AuthenticateService) ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrUnexpectedSigningMethod
		}
		return []byte(s.config.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.ErrInvalidToken
}
