package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cmm0004/ProtocolBuffersAndYou/proto"
	"github.com/golang/protobuf/proto"

	"golang.org/x/net/context"

	"goji.io"
	"goji.io/pat"
)

func ping_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PONG")
}

func get_job_by_id_handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	jobdid := pat.Param(ctx, "jobdid")

	//lets pretend you have a db of jobs ready to be searched for that id.
	//here, i will just initialize one with some values using the struct
	//  built for me after compiling the .proto file, since the structure
	//  of the job data i want to send is the same as the structure of the
	//  job i would get from my ~pretend db. In real life, you likely would
	//  need an intermediary step to arrage your data into what you wanted to send.

	foundjob := &protos.Data{
		Error: nil,
		Jobs: []*protos.Job{
			&protos.Job{
				Did:    jobdid,
				Title:  "Queen of JobGeo, Harbinger of Despair",
				Active: true,
			},
		},
	}

	data, err := proto.Marshal(foundjob)
	if err != nil {
		log.Fatal("Error marshalling to protobufs, %v", err)
	}

	fmt.Fprintf(w, "%v", data)
}

func main() {
	mux := goji.NewMux()

	mux.HandleFuncC(pat.Get("/jobs/:jobdid"), get_job_by_id_handler)

	mux.HandleFunc(pat.Get("/ping"), ping_handler)

	http.ListenAndServe(":8080", mux)
}
