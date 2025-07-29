package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret")

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := jwt.MapClaims{
		"username": username,
		"exp":      expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenStr string) (*jwt.MapClaims, error) {
	claims := &jwt.MapClaims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !tkn.Valid {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	return claims, nil
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		parts := strings.Split(authHeader, "Bearer ")
		if len(parts) != 2 {
			http.Error(w, "invalid Authorization header", http.StatusUnauthorized)
			return
		}

		_, err := ValidateJWT(parts[1])
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
