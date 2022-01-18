package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/gr8abbasi/go-microservice/handlers"
)

//var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	//env.Parse()

	// logger
	l := log.New(os.Stdout, "microservice-log: ", log.LstdFlags)

	// product handler
	ph := handlers.NewProducts(l)

	// New mux router
	sm := mux.NewRouter()

	// GET route
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	// POST route
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareValidateProduct)

	// PUT route
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareValidateProduct)

	// config http server
	s := &http.Server{
		Addr:         ":8081",
		Handler:      sm,
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
