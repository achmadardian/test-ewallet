package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

var SecretKey = []byte(os.Getenv("SECRET_KEY"))

type TokenType string

const (
	TypeAccessToken  TokenType = "access_token"
	TypeRefreshToken TokenType = "refresh_token"
)

type Claim struct {
	TokenType TokenType
	jwt.RegisteredClaims
}

// returning token
func (a *AuthService) GenerateToken(userId uuid.UUID, tokenType TokenType, d time.Duration) (string, error) {
	now := time.Now()
	exp := now.Add(d)

	claims := &Claim{
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userId.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}

	return tokenString, nil
}

// validate token
func (a *AuthService) ValidateToken(token string) (*Claim, error) {
	validate, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, fmt.Errorf("validate token: %w", err)
	}

	claims, ok := validate.Claims.(*Claim)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
