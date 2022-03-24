package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

const Version = "dev"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	log.Printf("app=soapbox@%v\n", Version)

	time.AfterFunc(5*time.Second, func() {
		for i := 0; i < 5; i++ {
			log.Println(`err="unable to connect to database 192.168.0.32:3306"`)
			time.Sleep(time.Second)
		}
		os.Exit(1)
	})

	http.ListenAndServe("localhost:8080", nil)
}
