package data

import (
	"time"
	"net/http"

	"github.com/upper/db/v4"
)

// TodoList holds postgres data structure for todo_list
type TodoList struct {
	ID        int64      `db:"id,omitempty" json:"id"`
	Status    ListStatus `db:"status,omitempty" json:"status"`
	Title     string     `db:"title,omitempty" json:"title"`
	CreatedAt *time.Time `db:"created_at,omitempty" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at,omitempty" json:"updatedAt"`
}

type Item struct {
	ID			int64 			`db:"id"           json:"id"`
	TodoListID	int64 			`db:"todo_list_id" json:"todo_list_id"`
	Title		string			`db:"title"        json:"title"`
	Todo		string			`db:"todo"         json:"todo"`
	Comment		string			`db:"comment"      json:"comment"`
	ListStatus	ListStatus      `db:"list_status"  json:"list_status"`
	ItemStatus	TodoStatus 		`db:"todo_status"  json:"todo_status"`  	
} 

type TodoListItems []Item

// ListStatus is list status
type ListStatus string

func (u *TodoListItems) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}  

const (
	// ListStatusActive is todo_list status active
	ListStatusActive ListStatus = "active"
	// ListStatusCompleted is todo_list status completed
	ListStatusCompleted ListStatus = "completed"
	// ListStatusCancelled is todo_list status cancelled
	ListStatusCancelled ListStatus = "cancelled"
)

var _ = interface {
	db.Record
	db.BeforeUpdateHook
	db.BeforeCreateHook
}(&TodoList{})

// Store initializes a session interface that defines methods for database adapters.
func (tl *TodoList) Store(sess db.Session) db.Store {
	return TodoLists(sess)
}

// BeforeCreate retrives and sets GetTimeUTCPointer before creating a new todolist
func (tl *TodoList) BeforeCreate(sess db.Session) error {
	if err := tl.BeforeUpdate(sess); err != nil {
		return err
	}
	tl.UpdatedAt = nil
	tl.CreatedAt = GetTimeUTCPointer()
	return nil
}

// BeforeUpdate retrives and sets GetTimeUTCPointer before updating a todolist
func (tl *TodoList) BeforeUpdate(sess db.Session) error {
	tl.UpdatedAt = GetTimeUTCPointer()
	return nil
}

// TodoListStore holds thread todolist database collection
type TodoListStore struct {
	db.Collection
}

var _ = interface {
	db.Store
}(&TodoListStore{})

// TodoLists retrieves all todolists
func TodoLists(sess db.Session) *TodoListStore {
	return &TodoListStore{sess.Collection("todo_list")}
}

// FindByID retrieves a particular todolist by id
func (store TodoListStore) FindByID(ID int64) (*TodoListItems, error) {
	var todoList TodoListItems
	sess := store.Session()

	q := sess.SQL().Select(
		"ti.id", 
		"ti.todo_list_id", 
		"ti.comment", 
		"ti.todo", 
		"ti.status AS todo_status", 
		"tl.status AS list_status", 
		"tl.title",
	).
	From("todo_items AS ti", "todo_list AS tl").
	Where("tl.id = ? AND ti.todo_list_id = ?", ID, ID).
	GroupBy("ti.todo_list_id", "ti.id", "ti.comment", "ti.todo", "todo_status", "list_status", "tl.title").
	OrderBy("ti.id") 
	
	if err := q.All(&todoList); err != nil {
		return nil, err
	}

	return &todoList, nil
}

// FindOne retrieves todolist by certain conditions from TodoListStore
func (store TodoListStore) FindOne(cond ...interface{}) (*TodoList, error) {
	var todolist *TodoList
	if err := store.Find(cond...).One(&todolist); err != nil {
		return nil, err
	}
	return todolist, nil
}

// FindAll retrieves todolists by certain conditions from TodoListStore
func (store TodoListStore) FindAll(cond ...interface{}) ([]*TodoList, error) {
	var todolists []*TodoList
	if err := store.Find(cond...).All(&todolists); err != nil {
		return nil, err
	}
	return todolists, nil
}
