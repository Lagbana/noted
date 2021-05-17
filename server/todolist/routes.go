package todolist

import (
	"github.com/go-chi/chi"
)

func Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/list", func(r chi.Router) {
		r.Post("/new", NewTodoList)
		r.Get("/{listID}", GetTodoList)
	})


	return r
}
