package main

import (
	"fmt"
	repository2 "github.com/alexgtn/esi2021-lab5/pkg/repository"
	"github.com/alexgtn/esi2021-lab5/pkg/service"
	http2 "github.com/alexgtn/esi2021-lab5/pkg/transport/http"
	ws "github.com/alexgtn/esi2021-lab5/pkg/transport/websocket"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	// SQL driver
	_ "github.com/lib/pq"
)

const (
	logLevel        = "debug"
	httpServicePort = 8080
	websocketServicePort = 8081
)

func main() {
	// begin setup
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}
	log.SetLevel(level)

	log.Info("Start server")

	// construct application
	bookRepo := repository2.NewBookRepostory()
	bookService := service.NewBookService(bookRepo)
	bookHTTPHandler := http2.NewBookHandler(bookService)
	bookWebsocketHandler := ws.NewBookHandler(bookService)

	studentRepoFoo := repository2.NewStudentFooRepostory()
	studentRepoBar := repository2.NewStudentBarRepostory()
	studentService := service.NewStudentService(studentRepoFoo, studentRepoBar)
	studentHTTPHandler := http2.NewStudentHandler(studentService)
	studentWebsocketHandler := http2.NewStudentHandler(studentService)

	httpRouter := mux.NewRouter()
	websocketRouter := mux.NewRouter()

	bookHTTPHandler.RegisterRoutes(httpRouter)
	bookWebsocketHandler.RegisterRoutes(websocketRouter)

	studentHTTPHandler.RegisterRoutes(httpRouter)
	studentWebsocketHandler.RegisterRoutes(websocketRouter)

	// setup http server
	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpServicePort),
		Handler: httpRouter,
	}

	// setup websocket server
	wsSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", websocketServicePort),
		Handler: websocketRouter,
	}

	go func() {
		err = wsSrv.ListenAndServe()
		if err != nil {
			log.Fatalf("Could not start server")
		}
	}()

	err = httpSrv.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start server")
	}

	log.Infof("Stoped server")
}
