package main

import (
	"fmt"
	"github.com/candidature/onboardings/config"
	"github.com/candidature/onboardings/models"
	"log"
	"net/http"
)
const port string= ":8080"

func init() {
	config.SetConfig()

}
func main() {

	config.SetInProduction(false)


	config.SetUseCache(true)


	models.Init()

	fmt.Println("Listening on ", port)



	//log.Fatal(http.ListenAndServe(port,nil))

	srv:=&http.Server{
		Addr: port,
		Handler: approutes(),
	}

	err:= srv.ListenAndServe()
	if err!= nil {
		log.Fatal("Failed to start server ", err)
	}

}
