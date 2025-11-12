package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	"meuseum.com/begin/api"
	"meuseum.com/begin/data"
)

func handleTamplate(w http.ResponseWriter, r *http.Request) {
	output, err := template.ParseFiles("tamplate/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server error"))
		return
	}
	output.Execute(w, data.GetAllList())
}

func main() {

	go func() {

		router := gin.Default()

		router.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})
		router.Run()
	}()

	server := http.NewServeMux()
	// server2 := http.NewServeMux()
	server.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<h1 style="color:red">Hello world frvzxcvxczom go server</h1>`))
	})

	fs := http.FileServer(http.Dir("./public"))

	server.Handle("/", fs)
	server.HandleFunc("/api/exhibition", api.Get)
	server.HandleFunc("/api/add", api.Post)
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
