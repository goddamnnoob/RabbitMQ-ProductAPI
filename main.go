package main

import (
	"context"
	"goddamnnoob/RabbitMQ-ProductAPI/app"
	"goddamnnoob/RabbitMQ-ProductAPI/services"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {

	//Set environmental variables for connecting to DB and Rabbit MQ
	log.Println(" Setting Environmental Variables ")
	err := godotenv.Load("env.bash")
	if err != nil {
		log.Println(err.Error())
		log.Fatalln("Error in loading ENV file")
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	ctx, cancel := context.WithCancel(context.Background())

	log.Println("Starting Web server")
	go app.StartApp(ctx, &wg)

	log.Println("Starting Consumer ")
	go services.ConsumerService(ctx, &wg)

	//Keyboard Ctrl+C interrupt
	ichan := make(chan os.Signal, 1)
	signal.Notify(ichan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-ichan
	log.Println("Interrupt Received in main")

	cancel()

	wg.Wait()

	log.Println("All the services shutdown")
}
