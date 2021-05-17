package todoitem

import (
	"github.com/go-chi/chi"
)

func Routes() chi.Router { 
	r := chi.NewRouter()

	r.Route("/{listID}", func(r chi.Router) {
		r.Post("/item/new", NewTodoItem) 
		r.Put("/item/{itemID}", UpdateTodoItem) 
	})

	// r.Get("/", GetAllTodoLists)

	return r
}
