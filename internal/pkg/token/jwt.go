package token

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JWTClaims struct {
	UserRole string `json:"user_role"`
	jwt.StandardClaims
}

func CreateJWT(userID uuid.UUID, userRole string) (token string, err error) {
	op := "internal.pkg.token.CreateJWT()"

	signingKey := []byte(os.Getenv("SIGNING_KEY"))

	sub := userID.String()
	exp := time.Now().Unix() + 10800 // 3 hour

	claims := JWTClaims{
		userRole,

		jwt.StandardClaims{
			Subject:   sub,
			ExpiresAt: exp,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err = tokenClaims.SignedString(signingKey)
	if err != nil {
		return "", fmt.Errorf("%s:%w", op, err)
	}

	return
}

func CheckJWT(tokenString string) (userID string, userRole string, err error) {
	op := "internal.pkg.token.CheckJWT()"
	signingKey := []byte(os.Getenv("SIGNING_KEY"))
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%s:unexpected signing method: %v", op, token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil || !token.Valid {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return "", "", fmt.Errorf("%s:%s", op, "that's not even a token")
			} else if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
				return "", "", fmt.Errorf("%s:%s", op, "timing is everything")
			}
			return "", "", fmt.Errorf("%s,%s", op, "invalid token")
		}
	}

	if claims, ok := token.Claims.(*JWTClaims); ok {
		return claims.Subject, claims.UserRole, nil
	} else {
		return "", "", fmt.Errorf("%s:%s", op, "failed conversion of jwt claims")
	}
}
