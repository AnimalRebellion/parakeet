package main

import (
	"fmt"
	"log"

	"github.com/AnimalRebellion/parakeet/proxy"
	"github.com/gin-gonic/gin"
)

func main() {
	var ns proxy.Server
	var err error
	err = ns.Connect()
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	r := gin.Default()
	err = proxy.TestApi(r.Group("/v1"), &ns)
	if err != nil {
		log.Fatal(err)
	}
	//db.ConnectDatabase()
	err = r.Run()
	if err != nil {
		log.Fatal("Failed to run server! Error = " + err.Error())
	}
}
