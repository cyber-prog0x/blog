package main

import (
	"log"
	"fmt"
	"net/http"
	"blog/routers"
	"blog/pkg/setting"
)

func main() {
	log.Printf("Nunu's Blog Version 2.0 Start! \n")

	router := routers.InitRouter()

	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
		Handler: router,
		ReadTimeout: setting.ReadTimeOut,
		WriteTimeout: setting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}
