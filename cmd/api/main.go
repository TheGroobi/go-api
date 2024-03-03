package main

import (
	"fmt"
	"net/http"

	"github.com/TheGroobi/go-api/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

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

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}

}
