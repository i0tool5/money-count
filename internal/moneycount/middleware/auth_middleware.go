package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"

	"simpleAPI/internal/moneycount/models"
	"simpleAPI/pkg/apictx"
	"simpleAPI/pkg/apierrors"
)

type authMiddleware struct {
	secret string
	noAuth []string
}

func newAuth(secret string, noAuth ...string) *authMiddleware {
	return &authMiddleware{
		secret: secret,
		noAuth: noAuth,
	}
}

func (a *authMiddleware) authCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		for _, uri := range a.noAuth {
			if strings.Contains(uri, r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}
		}
		authHdr := r.Header.Get("Authorization")
		if authHdr == "" {
			apierrors.HandleHTTPErr(w, apierrors.ErrNoAuthHeader,
				http.StatusBadRequest)
			return
		}

		splited := strings.Split(authHdr, " ")
		if len(splited) != 2 {
			apierrors.HandleHTTPErr(w, apierrors.ErrInvalidAuth,
				http.StatusBadRequest)
			return
		}

		jwtToken := splited[1]
		// handle jwtToken
		tk := new(models.Token) // TODO: move token to another place?

		token, err := jwt.ParseWithClaims(jwtToken, tk,
			func(t *jwt.Token) (interface{}, error) {
				return []byte(a.secret), nil
			})
		if err != nil {
			apierrors.HandleHTTPErr(w, err, http.StatusForbidden)
			return
		}

		if !token.Valid {
			apierrors.HandleHTTPErr(w, apierrors.ErrInvalidLogin, http.StatusForbidden)
			return
		}

		var ct apictx.UserCtx
		ctx := context.WithValue(r.Context(), ct, tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
