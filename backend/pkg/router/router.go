package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"telvina/APG5_WebCalc/pkg/presenter"
)

type Router struct {
	prs *presenter.Presenter
	mux *chi.Mux
}

func New(prs *presenter.Presenter, mux *chi.Mux) *Router {
	return &Router{
		prs: prs,
		mux: mux,
	}
}

func (rtr *Router) Run() {
	rtr.mux.Use(middleware.Logger)
	rtr.mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:8080"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	rtr.mux.Route("/api", func(r chi.Router) {
		r.Get("/calculator", rtr.calculatorHandler)
		r.Get("/credit", rtr.creditHandler)
		r.Get("/plot", rtr.plotHandler)
		r.Get("/previous_expression", rtr.previousExpressionHandler)
		r.Put("/clear_history", rtr.clearHistoryHandler)
	})
}
