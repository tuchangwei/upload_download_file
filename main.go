package main

import (
	"context"
	"files_handler/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	router := mux.NewRouter()
	l := log.New(os.Stdout, "", log.LstdFlags)
	f := handlers.NewFile(l)
	homeRouter := router.Methods(http.MethodGet).Subrouter()
	homeRouter.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Welcome!!!"))
	})

	fileEndPoint := fmt.Sprintf("/images/{%s:[0-9]+}/{%s}", handlers.FileIDKey, handlers.FileNameKey)
	l.Println("fileEndPoint:", fileEndPoint)
	fileHandlerPostRouter := router.Methods(http.MethodPost).Subrouter()
	fileHandlerPostRouter.HandleFunc(fileEndPoint, f.Upload)

	fileHandlerGetRouter := router.Methods(http.MethodGet).Subrouter()

	fileHandlerGetRouter.Handle(fileEndPoint, http.StripPrefix("/images/",http.FileServer(http.Dir(handlers.FileSystemRoot))))

	server := http.Server{
		Addr: ":7777",
		Handler: router,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Println("Unable to start server, error:", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	sig := <- c
	l.Println("Shutting down server with signal:", sig)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
