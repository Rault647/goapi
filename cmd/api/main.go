package main

import (
	"fmt"
	"net/http"

	"github.com/Rault647/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)        // Prints file and line number when printing to log
	var r *chi.Mux = chi.NewRouter() // Turns pointers to a Mux type (a struct to set up API)
	handlers.Handler(r)              // Function within the internal/handlers directory of the project

	fmt.Println("Starting GO API service...")

	fmt.Println(`
 ______     ______         ______     ______   __
/\	___\	 /\  __ \       /\  __ \   /\  == \ /\ \
\ \ \__ \  \ \ \/\ \      \ \  __ \  \ \  _-/ \ \ \
 \ \_____\  \ \_____\      \ \_\ \_\  \ \_\    \ \_\
  \/_____/   \/_____/       \/_/\/_/   \/_/     \/_/ `)

	err := http.ListenAndServe("localhost:8000", r) // Start the server with the http package
	if err != nil {
		log.Error(err)
	}
}
