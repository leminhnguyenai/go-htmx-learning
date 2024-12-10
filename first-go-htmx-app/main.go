package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			Message string
		}{
			Message: "Skibidi bap bap",
		}

		tmpl.Execute(w, data)
	})

	http.HandleFunc("/quote", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("To be the best you got to beat the best"))
	})

	log.Println("The server is on httpp://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
