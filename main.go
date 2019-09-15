package main

import (
	"net/http"
	"log"
	"./utils"
)

const port = "8000"

func main(){
	defer func() {
		log.Println("Server started at " + port + " ...")
		log.Fatal(http.ListenAndServe(":" + port, nil))
	}()

	http.HandleFunc("/signin", utils.Signin)
	http.HandleFunc("/home", utils.Home)
	http.HandleFunc("/refresh", utils.Refresh)
}
