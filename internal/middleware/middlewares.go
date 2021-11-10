package middleware

import (
	"github.com/gorilla/mux"
)

// Middleware represents new middleware object
type Middleware struct {
	MiddleFuncs *[]mux.MiddlewareFunc
}

// New is a factory function for middleware
func New(authKey string) *Middleware {
	m := new(Middleware)
	aum := newAuth(authKey, noAuth...)
	middlewares = append(middlewares, aum.authCheck)
	m.MiddleFuncs = &middlewares
	return m
}

var middlewares = []mux.MiddlewareFunc{
	jsonMiddleware,
	loggingMiddleware,
}

var noAuth = []string{
	"/api/sign_in",
	"/api/sign_up",
	"/api/refresh",
}
