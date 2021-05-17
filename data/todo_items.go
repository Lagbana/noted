package data

import (
	"net/http"
	"time"
	"github.com/upper/db/v4"
)

// TodoItem holds postgres data structure for todo_item
type TodoItem struct {
	ID           int64      `db:"id,omitempty" json:"id"`
	TodoListID	 int64		`db:"todo_list_id,omitempty" json:"todo_list_id"`
	Status       TodoStatus `db:"status" json:"status"`
	Todo         string     `db:"todo,omitempty" json:"todo"`
	Comment      string     `db:"comment,omitempty" json:"comment"`
	CreatedAt    *time.Time `db:"created_at,omitempty" json:"createdAt"`
	UpdatedAt    *time.Time `db:"updated_at,omitempty" json:"updatedAt"`
}

// TodoStatus is todo status
type TodoStatus string

const (
	// TodoStatusActive is todo_item status active
	TodoStatusActive TodoStatus = "active"
	// TodoStatusCompleted is todo_item status completed
	TodoStatusCompleted TodoStatus = "completed"
	// TodoStatusCancelled is todo_item status cancelled
	TodoStatusCancelled TodoStatus = "cancelled"
)

var _ = interface {
	db.Record
	db.BeforeUpdateHook
	db.BeforeCreateHook
}(&TodoItem{})

func (u *TodoItem) Bind(r *http.Request) error {
	return nil
}

// Store initializes a session interface that defines methods for database adapters.
func (tl *TodoItem) Store(sess db.Session) db.Store {
	return TodoItems(sess)
}

// BeforeCreate retrives and sets GetTimeUTCPointer before creating a new todoitem
func (tl *TodoItem) BeforeCreate(sess db.Session) error {
	if err := tl.BeforeUpdate(sess); err != nil {
		return err
	}
	tl.UpdatedAt = nil
	tl.CreatedAt = GetTimeUTCPointer()
	return nil
}

// BeforeUpdate retrives and sets GetTimeUTCPointer before updating a todoitem
func (tl *TodoItem) BeforeUpdate(sess db.Session) error {
	tl.UpdatedAt = GetTimeUTCPointer()
	return nil
}

// TodoItemStore holds thread todoitem database collection
type TodoItemStore struct {
	db.Collection
}

var _ = interface {
	db.Store
}(&TodoItemStore{})

// TodoItems retrieves all todoitems
func TodoItems(sess db.Session) *TodoItemStore {
	return &TodoItemStore{sess.Collection("todo_items")}
}

// FindByID retrieves a particular todoitem by id
func (store TodoItemStore) FindByID(ID int64) (*TodoItem, error) {
	return store.FindOne(db.Cond{"id": ID})
}

// FindOne retrieves todoitem by certain conditions from TodoItemStore
func (store TodoItemStore) FindOne(cond ...interface{}) (*TodoItem, error) {
	var todoitem *TodoItem
	if err := store.Find(cond...).One(&todoitem); err != nil {
		return nil, err
	}
	return todoitem, nil
}

// UpdateByID retrieves and updates todoitem by id
func (store TodoItemStore) UpdateByID(ID int64, todoitem *TodoItem) (*TodoItem, error) {
	// var todoitem *TodoItem
	if err := store.Find(db.Cond{"id": ID}).Update(todoitem); err != nil {
		return nil, err
	}
	return todoitem, nil
}

// FindAll retrieves todoitems by certain conditions from TodoItemStore
func (store TodoItemStore) FindAll(cond ...interface{}) ([]*TodoItem, error) {
	var todoitems []*TodoItem
	if err := store.Find(cond...).All(&todoitems); err != nil {
		return nil, err
	}
	return todoitems, nil
}
