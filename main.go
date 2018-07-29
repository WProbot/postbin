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
	var (
		port int
		key  string
	)

	flag.IntVar(&port, "port", 8080, "specifies the port number where to serve requests from")
	flag.StringVar(&key, "key", "", "if passed it will use to filter request. if not provided a random key will be used")

	flag.Parse()

	if len(key) == 0 {
		k, err := uuid.NewV4()
		if err != nil {
			log.Fatal(err)
		}
		key = k.String()
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

	if len(r.Header) != 0 {
		fmt.Fprintln(output, "Headers:")
		for k, v := range r.Header {
			fmt.Fprintf(output, "%s:%s\n", k, v[0])
		}
	}

	if len(r.URL.Query()) != 0 {
		fmt.Fprintln(output, "Query String Values:")
		for k, v := range r.URL.Query() {
			fmt.Fprintf(output, "%s:%s\n", k, v[0])
		}
	}

	if data, err := ioutil.ReadAll(r.Body); err == nil && len(data) != 0 {
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
