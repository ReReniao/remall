package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("gaoqian")

type Claims struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

// GenerateTaken 签发 token
func GenerateTaken(id uint, userName string, authority int) (string, error) {
	nowTime := time.Now()
	expiredTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		ID:        id,
		UserName:  userName,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
			Issuer:    "reniao",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken 验证 token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

type EmailClaims struct {
	UserID        uint   `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
	jwt.StandardClaims
}

// GenerateEmailTaken 签发email token
func GenerateEmailTaken(userId, operationType uint, email, password string) (string, error) {
	nowTime := time.Now()
	expiredTime := nowTime.Add(24 * time.Hour)
	emailClaims := EmailClaims{
		UserID:        userId,
		Email:         email,
		Password:      password,
		OperationType: operationType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
			Issuer:    "reniao",
		},
	}
	tokenEmailClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, emailClaims)
	token, err := tokenEmailClaims.SignedString(jwtSecret)
	return token, err
}

// ParseEmailToken 验证email token
func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if emailClaims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			return emailClaims, nil
		}
	}
	return nil, err
}
