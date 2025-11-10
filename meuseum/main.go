package main

import (
	"fmt"
	"net/http"
	"text/template"

	"meuseum.com/begin/data"
)

func handleTamplate(w http.ResponseWriter, r *http.Request) {
	output, err := template.ParseFiles("tamplate/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server error"))
		return
	}
	output.Execute(w, data.GetAllList()[1])
}

func main() {
	server := http.NewServeMux()
	// server2 := http.NewServeMux()
	server.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<h1 style="color:red">Hello world frvzxcvxczom go server</h1>`))
	})

	fs := http.FileServer(http.Dir("./public"))

	server.Handle("/", fs)
	server.HandleFunc("/tamplate", handleTamplate)
	// server2.HandleFunc("/hello2", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte(`<h1 style="color:red">Hello world from go server</h1>`))
	// })

	// go func() {
	// 	if err := http.ListenAndServe(":3030", server2); err != nil {
	// 		fmt.Println("Error while opening server")
	// 	}
	// }()
	if err := http.ListenAndServe(":3300", server); err != nil {
		fmt.Println("Error while opening server")
	}
}
