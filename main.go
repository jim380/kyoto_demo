package main

import (
	"log"
	"net/http"
	"os"

	page_index "github.com/jim380/kyoto-demo/pages/index"
	page_test "github.com/jim380/kyoto-demo/pages/test"
	"github.com/kyoto-framework/kyoto/render"
)

func main() {
	// Init serve mux
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/", render.PageHandler(page_index.PageIndex))
	mux.HandleFunc("/test", render.PageHandler(page_test.PageTest))

	// Run
	if os.Getenv("PORT") == "" {
		log.Println("Listening on localhost:25025")
		http.ListenAndServe("localhost:25025", mux)
	} else {
		log.Println("Listening on 0.0.0.0:" + os.Getenv("PORT"))
		http.ListenAndServe(":"+os.Getenv("PORT"), mux)
	}
}
