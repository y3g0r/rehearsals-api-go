package service

import (
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type CryptoService struct {
	secretKey []byte
}

func (s *CryptoService) ComparePasswordAndHash(password string, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err
	}
	return nil
}

func NewCryptoService(secretKey []byte) *CryptoService {
	return &CryptoService{
		secretKey: secretKey,
	}
}

func (s *CryptoService) CreateAccessToken(subject string, expiresDelta time.Duration) (string, error) {
	expire := time.Now().UTC().Add(expiresDelta)
	claims := jwt.MapClaims{
		"exp": expire.Unix(),
		"sub": subject,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	encodedJwt, err := token.SignedString(s.secretKey)
	if err != nil {
		return "", err
	}
	return encodedJwt, nil
}

func (s *CryptoService) HashPassword(password string) (string, error) {
	// hash password with bcrypt and return it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
