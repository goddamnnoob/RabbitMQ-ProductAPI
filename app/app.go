package app

import (
	"fmt"
	"goddamnnoob/Zocket-assignment/handlers"
	"log"
	"net/http"
)

func StartApp() {

	mux := http.NewServeMux()

	mux.HandleFunc("/AddProduct", handlers.AddProduct)
	mux.HandleFunc("/HelloWorld", handlers.HelloWorld)
	err := http.ListenAndServe(":8787", mux)
	if err != nil {
		log.Fatal("Server Start Failed")
	} else {
		fmt.Println("Server Started")
	}
}
