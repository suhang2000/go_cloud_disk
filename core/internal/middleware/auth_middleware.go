package middleware

import (
	"go_cloud_disk/core/helper"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		userClaim, err := helper.ParseToken(auth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte(err.Error()))
			if err != nil {
				return
			}
			return
		}
		// token is valid
		r.Header.Set("UserId", string(rune(userClaim.Id)))
		r.Header.Set("UserIdentity", userClaim.Identity)
		r.Header.Set("UserName", userClaim.Name)
		// Passthrough to next handler if need
		next(w, r)
	}
}
