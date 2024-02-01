package models

import "time"

type SingleTodo struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
}

type TodoList []*SingleTodo



// Methods
func NewSingleTodo(id string, title string) *SingleTodo {
	return &SingleTodo{
		Id: id,
		Title: title,
		CreatedAt: time.Now(),
	}
}

// func AddTodoList(id string, title string) []*SingleTodo {
// 	todo := NewSingleTodo (id, title)
// 	todoList = append(todoList, todo)
// 	return todoList
// }

// func (t *TodoList) ResetList() {
// 	todoList = TodoList{}
// }

// func (t *TodoList) indexOf(id string) int {
// 	for i, todo := range *t {
// 		if todo.Id == id {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func (t *TodoList) AddTodo(id string, title string) *SingleTodo {
// 	todo := NewSingleTodo (id, title)
// 	*t = append(*t, todo)
// 	return todo
// }

// func (t *TodoList) RemoveTodo(id string) {
// 	index := t.indexOf(id)
// 	*t = append((*t)[:index], (*t)[index+1:]...)
// }