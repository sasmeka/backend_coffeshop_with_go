package pkg

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type claims struct {
	Id    string `json:"id"`
	Role  string `json:"role"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func NewToken(uid, role, email string) *claims {
	minutes, _ := strconv.Atoi(viper.GetString("jwt.expireminutes"))
	return &claims{
		Id:    uid,
		Role:  role,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "backend coffee shop",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(minutes))),
		},
	}
}

func (c *claims) Generate() (string, error) {
	secrets := viper.GetString("jwt.secrets")
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return tokens.SignedString([]byte(secrets))
}

func VerifyToken(token string) (*claims, error) {
	secrets := viper.GetString("jwt.secrets")
	data, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secrets), nil
	})

	if err != nil {
		return nil, err
	}

	claimData := data.Claims.(*claims)
	return claimData, nil

}
