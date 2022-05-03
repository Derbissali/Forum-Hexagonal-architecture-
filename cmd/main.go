package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"tidy/config"
	"tidy/dbase"
)

func main() {
	db := dbase.CheckDB()

	router := config.Config(db)
	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	go func() {
		err := http.ListenAndServe(":8080", router)
		if err != nil {
			log.Printf("%v", err)
			// log.Fatal("ListenAndServe ERROR")
		}
	}()
	log.Println("port: 8080 is listening")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("port: 8080 is Sutting Down")
	if _, err := db.Exec(`DELETE FROM session`); err != nil {
		fmt.Println(err)

	}

	if err := db.Close(); err != nil {
		fmt.Println(err)
	}
}
