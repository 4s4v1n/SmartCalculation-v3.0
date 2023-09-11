package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"telvina/APG5_WebCalc/pkg/configurator"
	"telvina/APG5_WebCalc/pkg/presenter"
	"telvina/APG5_WebCalc/pkg/router"
)

func main() {
	prs := presenter.New(configurator.New())
	mux := chi.NewRouter()
	rtr := router.New(prs, mux)

	rtr.Run()

	srv := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case _ = <-interrupt:
		prs.SaveExpressions()
		prs.ReleaseLogger()
		prs.ReleaseModel()
		os.Exit(0)
	}
}
