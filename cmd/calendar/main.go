package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/tarusov/sndbx-srvc-24-01/internal/config"
	"github.com/tarusov/sndbx-srvc-24-01/internal/httpapi"
	"github.com/tarusov/sndbx-srvc-24-01/internal/service"
	"github.com/tarusov/sndbx-srvc-24-01/internal/storage"
)

func main() {

	cfg, err := config.Read("config.json")
	if err != nil {
		log.Fatal(err)
	}

	handler := httpapi.NewHandler(service.New(storage.NewInMemStorage()))

	mux := http.NewServeMux()
	mux.Handle("/create_event", http.HandlerFunc(handler.CreateEvent))
	mux.Handle("/update_event", http.HandlerFunc(handler.UpdateEvent))
	mux.Handle("/delete_event", http.HandlerFunc(handler.DeleteEvent))
	mux.Handle("/events_for_day", http.HandlerFunc(handler.EventsForDay))
	mux.Handle("/events_for_week", http.HandlerFunc(handler.EventsForWeek))
	mux.Handle("/events_for_month", http.HandlerFunc(handler.EventsForMonth))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.HTTP.Port),
		Handler: httpapi.NewLogMW(mux),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println("HTTP server stopped")
		}
	}()

	doneCh := make(chan struct{})

	// Make signal watcher
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		for range sigCh {
			log.Println("Received an interrupt...")
			_ = server.Shutdown(ctx)
			doneCh <- struct{}{}
		}
	}()

	<-doneCh
}
