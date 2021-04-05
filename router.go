package main

import (
	"net/http"

	"github.com/iwollmann/gostudy/handlers"
)

func RegisterRoutes(){
	mc := handlers.NewMatrixHandler()

	http.Handle("/echo", ComposeReverse(http.HandlerFunc(mc.Echo), mc.ValidateMiddleware, mc.AcceptOnlyPostMiddleware))
	http.Handle("/invert", ComposeReverse(http.HandlerFunc(mc.Invert), mc.ValidateMiddleware, mc.AcceptOnlyPostMiddleware))
	http.Handle("/flatten", ComposeReverse(http.HandlerFunc(mc.Flatten), mc.ValidateMiddleware, mc.AcceptOnlyPostMiddleware))
	http.Handle("/sum", ComposeReverse(http.HandlerFunc(mc.Sum), mc.ValidateMiddleware, mc.AcceptOnlyPostMiddleware))
	http.Handle("/multiply", ComposeReverse(http.HandlerFunc(mc.Multiply), mc.ValidateMiddleware, mc.AcceptOnlyPostMiddleware))
}

type Middleware func(http.Handler) http.Handler

func ComposeReverse(h http.Handler, midlewares ...Middleware) http.Handler {
	for _, mw := range midlewares {
	  h = mw(h)
	}
	return h
}