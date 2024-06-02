package main

import (
	"devtool/internal/composer"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	app, err := composer.NewApplication()
	if err != nil {
		log.Fatal(err)
	}

	go app.Start()
	log.Println("Press (Ctrl+C) to shutdown application")
	<-sigs

	log.Println("Shutdown application")
	app.Shutdown()
}
