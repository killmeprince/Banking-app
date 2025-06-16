package jwt

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"banking-app/config"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

type ctxKey string

const UserIDKey ctxKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("Authorization")
		if !strings.HasPrefix(h, "Bearer ") {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(h, "Bearer ")

		claims := &jwtlib.RegisteredClaims{}
		tok, err := jwtlib.ParseWithClaims(tokenStr, claims, func(t *jwtlib.Token) (interface{}, error) {
			return config.JWTSecret, nil
		})
		if err != nil || !tok.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, claims.Subject)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserIDFromRequest(r *http.Request) (int64, error) {
	val := r.Context().Value(UserIDKey)
	if val == nil {
		return 0, errors.New("no user id in context")
	}
	userIDStr, ok := val.(string)
	if !ok {
		return 0, errors.New("user id in context is not string")
	}
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return 0, errors.New("user id not int")
	}
	return userID, nil
}
