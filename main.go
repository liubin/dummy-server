package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func readFile(file string) (string, error) {
	if b, err := ioutil.ReadFile(file); err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	file := fmt.Sprintf("responses%s.json", r.URL.Path)
	file = strings.Replace(file, "{", "", -1)
	file = strings.Replace(file, "}", "", -1)
	fmt.Println("request url: ", file)
	if s, e := readFile(file); e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, strings.Replace(e.Error(), `"`, `\"`, -1))))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(s))
	}
}

func main() {
	var port = flag.Int("port", 8081, "port to serve on")
	flag.Parse()
	fmt.Println("Listening on: ", *port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
