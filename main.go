package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"

	"goji.io"
	"goji.io/pat"
)

func ping_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PONG")
}

func get_job_by_id_handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	jobid := pat.Param(ctx, "id")
	fmt.Fprintf(w, "Ive got job , %s!", jobid)
}

func main() {
	mux := goji.NewMux()

	mux.HandleFuncC(pat.Get("/jobs/:id"), get_job_by_id_handler)

	mux.HandleFunc(pat.Get("/ping"), ping_handler)

	http.ListenAndServe(":8080", mux)
}
