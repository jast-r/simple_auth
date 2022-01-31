package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	simpleauth "github.com/jast-r/simple-auth"
	"github.com/jast-r/simple-auth/pkg/repository"
	"github.com/sirupsen/logrus"
)

const (
	salt      = "bhu32yg45hjthg-f9_8drty*hnjkl*_y324jn"
	tokenTTL  = 12 * time.Hour
	signedKey = "ugy3364ny3#___ASDNAQ3I2"
)

type AuthService struct {
	repos repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

func (s *AuthService) CreateUser(user simpleauth.User) error {
	user.Password = genPasswordHash(user.Password)
	fmt.Println(user.Password)
	return s.repos.CreateUser(user)
}

func genPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenJWT(username, password string) (string, error) {
	user, err := s.repos.GetUser(username, genPasswordHash(password))
	if err != nil {
		logrus.Errorf("get user failed: %v", err)
		return "", err
	}
	fmt.Println(user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Username,
	})
	return token.SignedString([]byte(signedKey))
}
