// ///
// Chi
// ///

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// TodoItem ...
type TodoItem struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var items []TodoItem

// ItemsResource ...
type ItemsResource struct{}

// List (or GetAll) ...
func (res ItemsResource) List(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, items)
}

// Create ...
func (res ItemsResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Todo item creation..."))
}

// Get ...
func (res ItemsResource) Get(w http.ResponseWriter, r *http.Request) {
	// some mocking with the {id} requested
	id := chi.URLParam(r, "id")
	sid, _ := strconv.Atoi(id)
	todoItem := TodoItem{ID: sid, Title: "Title", Body: "Body ..."}

	render.JSON(w, r, todoItem)
}

// Update (put method) ...
func (res ItemsResource) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Todo item update by id"))
}

// Delete ...
func (res ItemsResource) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Todo item delete by id"))
}

// Routes ...
func (res ItemsResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", res.List)
	r.Post("/", res.Create)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", res.Get)
		r.Put("/", res.Update)
		r.Delete("/", res.Delete)
	})

	return r
}

// UsersResource ...
type UsersResource struct{}

func main() {
	port := os.Getenv("PORT")

	items = append(items, TodoItem{ID: 1, Title: "Title 1", Body: "Body 1 ..."})
	items = append(items, TodoItem{ID: 2, Title: "Title 2", Body: "Body 2 ..."})
	items = append(items, TodoItem{ID: 3, Title: "Title 3", Body: "Body 3 ..."})

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/ ..."))
	})

	r.Mount("/todos", ItemsResource{}.Routes())

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
