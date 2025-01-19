package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager interface {
	AccessToken(userID int, ttl time.Duration) (string, error)
	RefreshToken(userID int) (string, error)
}

type Manager struct {
	secretKey []byte
}

func NewManager(secretKey string) (*Manager, error) {
	if secretKey == "" {
		return nil, fmt.Errorf("error: empty secret key")
	}

	return &Manager{secretKey: []byte(secretKey)}, nil
}

func (m *Manager) AccessToken(userID int, ttl time.Duration) (string, error) {
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(ttl).Unix(),
	}).SignedString(m.secretKey)

	if err != nil {
		return "", err
	}

	return accessToken, nil
}

// todo. Сделать подпись токена изменяемой
func (m *Manager) RefreshToken(userID int) (string, error) {
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
	}).SignedString(m.secretKey)

	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
