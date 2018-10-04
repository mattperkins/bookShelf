package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// fmt.Fprintln(w, "Test build...")
	})
	fmt.Println(http.ListenAndServe(":8080", nil))
}
