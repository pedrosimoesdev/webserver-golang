package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Values struct {
	AppVersion string
	Row        string
}

func main() {
	log.SetFlags(log.Lshortfile)

	appVersion := "bad versions with bad developer loooool"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		data := Values{
			AppVersion: appVersion,
			Row:        string("test 123"),
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("http://127.0.0.1:3000")
	log.Fatalln(http.ListenAndServe("127.0.0.1:3000", nil))
}
