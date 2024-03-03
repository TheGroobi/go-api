package main

import (
	"fmt"
	"net/http"

	"../../internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var router *chi.Mux = chi.NewRouter()
	handlers.Handler()

	fmt.Println("Starting GO API service...")

	fmt.Println(`
$$$$$$\   $$$$$$\         $$$$$$\  $$$$$$$\$$$$$$\ 
$$  __$$\ $$  __$$\       $$  __$$\ $$  __$$\_$$  _|
$$ /  \__|$$ /  $$ |      $$ /  $$ |$$ |  $$ |$$ |  
$$ |$$$$\ $$ |  $$ |      $$$$$$$$ |$$$$$$$  |$$ |  
$$ |\_$$ |$$ |  $$ |      $$  __$$ |$$  ____/ $$ |  
$$ |  $$ |$$ |  $$ |      $$ |  $$ |$$ |      $$ |  
\$$$$$$  | $$$$$$  |      $$ |  $$ |$$ |    $$$$$$\ 
 \______/  \______/       \__|  \__|\__|    \______|`)

	err := http.ListenAndServe("localhost:8000", router)
	if err != nil {
		log.Error(err)
	}

}
