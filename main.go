package main

import (
	"context"
	"fmt"
	"github.com/dottics/serve/auth"
	"github.com/dottics/serve/openapi"
	"github.com/dottics/serve/servemux"
	"log"
	"net/http"
)

func Home(c context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Println("param", servemux.GetParam(c, "uuid"))
	w.WriteHeader(204)
	_, _ = w.Write([]byte(`home`))
}

func openAPI(c context.Context, w http.ResponseWriter, r *http.Request) {
	openapi.Handler(w, r)
}

func main() {
	mux := servemux.NewMux()
	mux.Use(servemux.CORS(nil, nil))
	mux.Use(auth.Middleware)
	// whitelisted
	mux.GetWhitelisted("/docs*", openAPI)

	// for testing
	//	mux.PostWhitelisted("/queue")

	// protected
	mux.Get("/", Home)
	mux.Get("/{uuid}", Home)
	//mux.PostWhitelisted("/queue", handlers.ProactiveMessageQueue)

	log.Println("server listening on port: 8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
