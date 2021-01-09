package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"urlshortner/urlshort"
)

func main() {
	mux := http.NewServeMux()

	pathToUrls := map[string]string{
		"/urlshort-godoc": "https://go.dev/",
		"/yaml-godoc":     "https://pkg.go.dev/gopkg.in/yaml.v3",
	}
	mapHandler := urlshort.MapHandler(pathToUrls, mux)

	content, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		panic(err)
	}
	yamlHandler, err := urlshort.YAMLHandler(content, mapHandler)
	if err != nil {
		panic(err)
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://vksharma.net", http.StatusFound)
	})

	fmt.Println("Server started at 3001 port")
	http.ListenAndServe(":3001", yamlHandler)
}
