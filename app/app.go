package app

import (
	"context"
	"goddamnnoob/RabbitMQ-ProductAPI/handlers"
	"log"
	"net/http"
	"sync"
)

func StartApp(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	mux := http.NewServeMux()
	mux.HandleFunc("/AddProduct", handlers.AddProduct)
	mux.HandleFunc("/HelloWorld", handlers.HelloWorld)

	server := &http.Server{Addr: "127.0.0.1:8111", Handler: mux}

	go func(server *http.Server) {
		err := server.ListenAndServe()
		if err != nil {
			log.Println("Server Start Failed")
		}
	}(server)

	<-ctx.Done()

	if server.Shutdown(context.Background()) != nil {
		log.Println(" WebServer shutdown failed!! ")
	}
}
