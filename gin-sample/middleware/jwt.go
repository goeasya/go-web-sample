package middleware

import (
	"errors"
	"net/http"
	"time"

	"gin-sample/global"
	"gin-sample/model/response"
	"gin-sample/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var ErrInvalidToken = errors.New("invalid token")

type Claims struct {
	UserID   string `json:"userId,omitempty"`
	UserName string `json:"username,omitempty"`
	NickName string `json:"NickName,omitempty"`
	Email    string `json:"email,omitempty"`
	jwt.StandardClaims
}

// JWT only parse token, not validate user
func JWT() gin.HandlerFunc {
	secretKey := global.GetSecretKey()
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		claim, err := parseToken(secretKey, token)
		if err != nil {
			util.ResponseError(c, http.StatusUnauthorized, err.Error())
			c.Abort()
		}
		if claim != nil {
			global.Logger.Debugf("jwt pass: %v",
				claim.UserName, claim.UserID, c.Request.Method, c.Request.URL)
		}
		c.Next()
	}
}

// parseToken
func parseToken(secretKey []byte, token string) (*Claims, error) {
	if token == "" {
		return nil, ErrInvalidToken
	}
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil || tokenClaims == nil {
		return nil, ErrInvalidToken
	}
	claims, ok := tokenClaims.Claims.(*Claims)
	if !ok || !tokenClaims.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

func GenerateToken(args *response.UserBase) (string, error) {
	claim := Claims{
		args.UserId,
		args.UserName,
		args.NickName,
		args.Email,
		jwt.StandardClaims{
			ExpiresAt: int64(time.Hour * 24 * 7),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(global.GetSecretKey())
}
