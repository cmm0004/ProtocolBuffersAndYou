package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/cmm0004/ProtocolBuffersAndYou/proto"
	"github.com/golang/protobuf/proto"
)

// this is a confusing file, I know. but how else can we have the full process without
// calling our protobuffed service as if you would in practice?
// in real life, you would have a wrapper that contained your protobuff file (and
// allowed operations on your endpoints), so you
// didnt need to have copies of in your clients and services, cause without that file, this client
// wouldnt know what to do with the binary it gets.
//
// run the main binary first, and in a separate teminal window, run this one (main-client)
func main() {

	// calling our running instance of main
	response, err := http.Get("http://localhost:8080/jobs/J312345")
	if err != nil {
		log.Fatal("error retrieving response: %v ", err)
	}

	//reading from response data
	responsedata, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("error reading response body: %v", err)
	}

	// deserializing that into the Data struct defined in protos/job.pb.go
	incomingjob := &protos.Data{}
	err = proto.Unmarshal(responsedata, incomingjob)
	if err != nil {
		log.Fatal("error deserializing job response: %v", err)
	}

	fmt.Printf("\n%v", incomingjob)

}
