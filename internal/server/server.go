package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

func Start() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      newHandler(), // Use custom handler via gorilla/mux
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		IdleTimeout:  time.Second * 120,
	}

	// graceful shutdown
	quit := make(chan os.Signal, 1) // buffered to a single value
	signal.Notify(quit, os.Interrupt)
	/*
		using a waitgroup because server.Shutdown() causes an error in the main thread
		and execution switches off of gracefulShutdown so it never completes logging.
	*/
	var wg sync.WaitGroup
	wg.Add(1)
	go gracefulShutdown(server, &wg, quit)

	err := server.ListenAndServe() // this is blocking
	wg.Wait()
	if err != http.ErrServerClosed {
		log.Fatal("Server stopped unexpectededly:", err)
	}
}

func gracefulShutdown(server *http.Server, wg *sync.WaitGroup, quit <-chan os.Signal) {
	defer wg.Done()
	<-quit // Receive the signal (this will block until I hit ^C)
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server shutdown failed:", err)
	}
	log.Println("Server gracefully stopped")
}
