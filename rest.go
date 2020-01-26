
/*
* Sample API with POST endpoint.
* POST data is converted to string
* Cours : https://www.devopssec.fr/
*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	// flagPort is the open port the application listens on
	flagPort = flag.String("port", "8080", "Port to listen on")
)

var result string

// PostHandler converts post request body to string
func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		result = string(body)
		fmt.Fprint(w, string(body))
		log.Printf(result)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	flag.Parse()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", PostHandler)

	log.Printf("listening on port %s", *flagPort)
	log.Fatal(http.ListenAndServe(":"+*flagPort, mux))
}
