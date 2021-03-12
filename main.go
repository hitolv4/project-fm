package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hitolv4/fm-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "fm-api ", log.LstdFlags)
	rh := handlers.NewRepuestos(l)

	sm := http.NewServeMux()
	sm.Handle("/repuestos", rh)
	sm.Handle("/repuestos/", rh)

	log.Println("Starting Server")

	port := os.Getenv("PORT")

	s := &http.Server{
		Addr:         ":" + port, // 9090
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
