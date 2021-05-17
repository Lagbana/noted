package todoitem

import (
	"net/http"
	"strconv"

	"github.com/Lagbana/noted/data"
	"github.com/Lagbana/noted/server/api"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type TodoItemRequest struct {
	Status 	   string `json:"status,omitempty"` 
	Todo       string `json:"todo,omitempty"`
	Comment    string `json:"comment,omitempty"`
}

func (u *TodoItemRequest) Bind(r *http.Request) error {
	return nil
}

func (u *TodoItemRequest) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// NewTodoItem creates a new todoitem
func NewTodoItem(w http.ResponseWriter, r *http.Request) {
  // fetch the url parameter `"userID"` from the request of a matching
  // routing pattern. An example routing pattern could be: /users/{userID}
	listID := chi.URLParam(r, "listID")
	var todoListID int64
	var err error 

	if todoListID, err = strconv.ParseInt(listID, 10, 64); err != nil {
	   render.Render(w, r, api.ErrInternalServerError(err)) 
	} 

   newTodoItemRequest := &TodoItemRequest{}
	if err = render.Bind(r, newTodoItemRequest); err != nil {
		render.Render(w, r, api.ErrInvalidRequest(err))
	}

	// necessary fields required to create a new todo item
	todoItem := &data.TodoItem{ 
		Status: data.TodoStatusActive, 
		Todo: newTodoItemRequest.Todo,  
		Comment: newTodoItemRequest.Comment,
		TodoListID: todoListID,
	}

	if err := data.DB.Save(todoItem); err != nil {
		render.Render(w, r, api.ErrDatabase(err))
	}

	api.Render(w, r, newTodoItemRequest)
}

// UpdateTodoItem updates a new todoitem
func UpdateTodoItem(w http.ResponseWriter, r *http.Request) {
	var todoItemID int64
	var todoListID int64
	var err error 
	itemID := chi.URLParam(r, "itemID")
	listID := chi.URLParam(r, "listID")

	if todoItemID, err = strconv.ParseInt(itemID, 10, 64); err != nil {
	   render.Render(w, r, api.ErrInternalServerError(err)) 
	} 

	if todoListID, err = strconv.ParseInt(listID, 10, 64); err != nil {
	   render.Render(w, r, api.ErrInternalServerError(err)) 
	} 

	updatedTodoItem := &TodoItemRequest{}
	if err = render.Bind(r, updatedTodoItem); err != nil {
		render.Render(w, r, api.ErrInvalidRequest(err))
	}

	// necessary fields required to update a todo item
	todoItem := &data.TodoItem{ 
		// Status: " completed", 
		Status: data.TodoStatus(updatedTodoItem.Status), 
		Todo: updatedTodoItem.Todo,  
		Comment: updatedTodoItem.Comment,
		TodoListID: todoListID,
	}

	if _, err = data.DB.TodoItem.UpdateByID(todoItemID, todoItem); err != nil {
		render.Render(w, r, api.ErrDatabase(err))
	}  

	api.Render(w, r, updatedTodoItem) 
}