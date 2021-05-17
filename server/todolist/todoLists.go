package todolist

import (
	// "fmt"
	"net/http"
	"strconv"

	"github.com/Lagbana/noted/data"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/Lagbana/noted/server/api"
)

type Request struct {
	Status 	  string `json:"status,omitempty"` 
	Title     string `json:"title"`
}

type Item struct {
	ID			int64 			`db:"id"           json:"id"`
	TodoListID	int64 			`db:"todo_list_id" json:"todo_list_id"`
	Title		string			`db:"title"        json:"title"`
	Todo		string			`db:"todo"         json:"todo"`
	Comment		string			`db:"comment"      json:"comment"`
	ListStatus	data.ListStatus `db:"list_status"  json:"list_status"`
	ItemStatus	data.TodoStatus `db:"todo_status"  json:"todo_status"`  	
} 

type Response []Item  

func (u *Request) Bind(r *http.Request) error {
	return nil
}
func (u *Request) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// NewTodoList creates a new todolist
func NewTodoList(w http.ResponseWriter, r *http.Request) {
	newTodoListRequest := &Request{} 
	if err := render.Bind(r, newTodoListRequest); err != nil {
		render.Render(w, r, api.ErrInvalidRequest(err))
	}

	// necesarry fields required to create a new todolist
	todoList := &data.TodoList{
		Status: data.ListStatusActive, 
		Title: newTodoListRequest.Title,
	}

	if err := data.DB.Save(todoList); err != nil {
		render.Render(w, r, api.ErrDatabase(err))
	}

	api.Render(w, r, newTodoListRequest)
}

func GetTodoList(w http.ResponseWriter, r *http.Request) {
	var todoListID int64
	var err error  

	listID := chi.URLParam(r, "listID")

	if todoListID, err = strconv.ParseInt(listID, 10, 64); err != nil {
	   render.Render(w, r, api.ErrInternalServerError(err)) 
	} 

	todoList, err := data.DB.TodoList.FindByID(todoListID) 
	if err != nil {
		render.Render(w, r, api.ErrDatabase(err))
	}  

	render.Render(w, r, todoList)          
}
