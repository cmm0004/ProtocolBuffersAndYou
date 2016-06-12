package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/cmm0004/ProtocolBuffersAndYou/proto"
	"github.com/golang/protobuf/proto"

	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
)

func ping_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PONG")
}

func get_job_by_id_handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	jobdid := pat.Param(ctx, "jobdid")
	responsetype := r.URL.Query().Get("response")

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

	if strings.ToLower(responsetype) == "json" {
		// now for comparison, heres the same process but for sending json
		// notice in the job.pb.go definitions, the structs are given json
		// mappings, likely because outside you consume would be still using json,
		// and these are set up for your convienevce.
		// lets see how they compare.

		jsonstart := time.Now()
		jsondata, err := json.Marshal(foundjob)
		jsonelasped := time.Since(jsonstart)
		fmt.Printf("\ntime to serialize to json: %v", jsonelasped)

		if err != nil {
			log.Fatal("Error marshalling to protobufs, %v", err)
		}

		fmt.Printf("\nsize of json: %v", len(jsondata))

		fmt.Fprintf(w, "%s", jsondata)

	} else {
		start := time.Now()
		data, err := proto.Marshal(foundjob)
		elasped := time.Since(start)
		fmt.Printf("\ntime to serialize to protobufs: %v", elasped)

		if err != nil {
			log.Fatal("Error marshalling to protobufs, %v", err)
		}

		fmt.Printf("\nsize of proto: %v", len(data))

		fmt.Fprintf(w, "%v", data)
	}
}

func main() {
	mux := goji.NewMux()

	mux.HandleFuncC(pat.Get("/jobs/:jobdid"), get_job_by_id_handler)

	mux.HandleFunc(pat.Get("/ping"), ping_handler)

	http.ListenAndServe(":8080", mux)
}
