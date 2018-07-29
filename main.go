package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

func main() {
	var port int

	flag.IntVar(&port, "port", 8080, "specifies the port number where to serve requests from")

	flag.Parse()

	key, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("key: ", key)
	fmt.Println("port: ", port)
	printDashes(os.Stdout)

	router := mux.NewRouter()
	router.HandleFunc(fmt.Sprintf("/%s", key), logRequest)
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func logRequest(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	output := &bytes.Buffer{}

	fmt.Fprintln(output, "Received at:")
	fmt.Fprintln(output, now.String())

	fmt.Fprintln(output, "Remote Address:")
	fmt.Fprintln(output, r.RemoteAddr)

	fmt.Fprintln(output, "Method:")
	fmt.Fprintln(output, r.Method)

	fmt.Fprintln(output, "Headers:")
	for k, v := range r.Header {
		fmt.Fprintf(output, "%s:%s\n", k, v[0])
	}

	fmt.Fprintln(output, "Query String Values:")
	for k, v := range r.URL.Query() {
		fmt.Fprintf(output, "%s:%s\n", k, v[0])
	}

	if data, err := ioutil.ReadAll(r.Body); err == nil {
		fmt.Fprintln(output, "Body:")
		fmt.Fprintln(output, string(data))
	}

	w.WriteHeader(http.StatusOK)

	w.Write(output.Bytes())
	fmt.Println(output.String())

	printDashes(os.Stdout)
}

func printDashes(writer io.Writer) {
	fmt.Fprintln(writer, strings.Repeat("=", 40))
}
