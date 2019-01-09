///
// Chi
///

// package main

// import (
// 	"net/http"

// 	"github.com/go-chi/chi"
// 	"github.com/go-chi/chi/middleware"
// )

// // Create a new instance of the logger. You can have any number of instances.

// // ItemsResource ...
// type ItemsResource struct{}

// // List (or GetAll) ...
// func (res ItemsResource) List(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Todos item list of stuff..."))
// }

// // Create ...
// func (res ItemsResource) Create(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Todo item creation..."))
// }

// // Get ...
// func (res ItemsResource) Get(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Todo item get by id"))
// }

// // Update (put method) ...
// func (res ItemsResource) Update(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Todo item update by id"))
// }

// // Delete ...
// func (res ItemsResource) Delete(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Todo item delete by id"))
// }

// // Routes ...
// func (res ItemsResource) Routes() chi.Router {
// 	r := chi.NewRouter()

// 	r.Get("/", res.List)
// 	r.Post("/", res.Create)

// 	r.Route("/{id}", func(r chi.Router) {
// 		r.Get("/", res.Get)
// 		r.Put("/", res.Update)
// 		r.Delete("/", res.Delete)
// 	})

// 	return r
// }

// // UsersResource ...
// type UsersResource struct{}

// func main() {
// 	r := chi.NewRouter()

// 	r.Use(middleware.RequestID)
// 	r.Use(middleware.RealIP)
// 	r.Use(middleware.Logger)
// 	r.Use(middleware.Recoverer)

// 	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("/ ..."))
// 	})

// 	r.Mount("/todos", ItemsResource{}.Routes())

// 	http.ListenAndServe(":8080", r)
// }

// ///
// Gin
// ///

// package main

// import "github.com/gin-gonic/gin"

// func main() {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "pong!",
// 		})
// 	})
// 	r.Run() // listen and serve on 0.0.0.0:8080
// }

// ///////
// DEFAULT
// ///////

package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
