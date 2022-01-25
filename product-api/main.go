package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gr8abbasi/go-microservice/product-api/data"
	"github.com/gr8abbasi/go-microservice/product-api/handlers"
)

//var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	//env.Parse()

	// logger
	l := log.New(os.Stdout, "microservice-log: ", log.LstdFlags)
	v := data.NewValidation()

	// product handler
	ph := handlers.NewProducts(l, v)

	// New mux router
	sm := mux.NewRouter()

	// GET route
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.ListAll)
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.ListProduct)

	// POST route
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", ph.Create)
	postRouter.Use(ph.MiddlewareValidateProduct)

	// PUT route
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", ph.Update)
	putRouter.Use(ph.MiddlewareValidateProduct)

	// DELETE route
	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", ph.Delete)

	// Documentation handler
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS requests
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	// config http server
	s := &http.Server{
		Addr:         ":8081",
		Handler:      ch(sm),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
