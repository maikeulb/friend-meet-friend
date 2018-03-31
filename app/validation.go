package app

import (
	"net/http"
	"strings"
	"context"

	"github.com/maikeulb/friend-meet-friend/app/auth"
)

func (a *App) ValidationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := auth.ParseToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}
		userID := int(claims["userId"].(float64))
		context := context.WithValue(r.Context(), "userId", userID)
		next.ServeHTTP(w, r.WithContext(context))
	})
}
