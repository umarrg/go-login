package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"github.com/umarrg/microservice/middlewares"
	"github.com/umarrg/microservice/myError"
	"github.com/umarrg/microservice/routes"
)

func main() {
	router := httprouter.New()
	call()
	router.POST("/login", middlewares.JHeader(routes.Login))

	fmt.Println("Server listening on port 3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}

func call() {
	response, err := http.Get("https://api.instantwebtools.net/v1/airlines")

	if err != nil {
		myError.HandleErr(err)

	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		myError.HandleErr(err)
	}
	fmt.Println(string(responseData))

}
