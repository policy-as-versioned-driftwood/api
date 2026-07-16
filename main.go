// api team's app (real-estate epic, ticket 08). The good citizen: current
// dependencies (see go.mod -- go-chi/chi, a small real router, latest),
// policy 2.2.0. The contrast with ledger's Log4Shell-era log4j is the
// point.
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("api\n"))
	})
	log.Println("api starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
