package handlers

import (
	"strconv"

	"github.com/Just-A-NoobieDev/go-htmx-templ-todo/models"
	"github.com/Just-A-NoobieDev/go-htmx-templ-todo/util"
	"github.com/Just-A-NoobieDev/go-htmx-templ-todo/views"

	// "github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Initializer
var todos = models.TodoList{}
var tasksLists = models.TaskList{}
var taskList = map[string]interface{}{}

func GetAllTaskPerTodoId() {
	for _, todo := range todos {
		tasks := models.TaskList{}
		for _, task := range tasksLists {
			if task.Id == todo.Id {
				tasks = append(tasks, task)
			}
		}
		taskList[todo.Id] = tasks
	}
}

func GetAllTodos(c echo.Context) error {
	todos = models.TodoList{}
	tasksLists = models.TaskList{}
	db := util.ConnectToDB()
	db.Query(SelectAllTodosQuery).Rows(&todos)
	db.Query(SelectAllTasksQuery).Rows(&tasksLists)
	defer db.Close()

	GetAllTaskPerTodoId()
	
	component := views.Index(todos, taskList)
	return component.Render(c.Request().Context() ,c.Response().Writer)
}

func GetAllTodosTotal(c echo.Context) error {
	component := views.Total(todos)
	return component.Render(c.Request().Context() ,c.Response().Writer)
}

func GetTodoById(c echo.Context) error {
	var id = c.Param("id")
	todo1 := models.SingleTodo{}
	for _, todo := range todos {
		if todo.Id == id {
			todo1 = *todo
		}
	}
	component := views.Modal(&todo1, taskList)
	return component.Render(c.Request().Context() ,c.Response().Writer)
}

func CreateTodo(c echo.Context) error {
	db := util.ConnectToDB()
	id := uuid.New()
	id2 := uuid.New()
	id3 := uuid.New()
	id4 := uuid.New()
	title := c.FormValue("title")
	task1 := c.FormValue("task1")
	task2 := c.FormValue("task2")
	task3 := c.FormValue("task3")

	stmt1 := InsertTodoQuery(id.String(), title)
	stmt2 := InsertTaskQuery(id.String(), id2.String(), id3.String(), id4.String(), task1, task2, task3)

	db.Query(stmt1).Run()
	db.Query(stmt2).Run()
	defer db.Close()

	todos = append(todos, models.NewSingleTodo (id.String(), title))
	tasksLists = append(tasksLists, models.NewSingleTask (id2.String(), task1, id.String()))
	tasksLists = append(tasksLists, models.NewSingleTask (id3.String(), task2, id.String()))
	tasksLists = append(tasksLists, models.NewSingleTask (id4.String(), task3, id.String()))

	GetAllTaskPerTodoId()
	
	component := views.RenderTodos(todos, taskList)
	return component.Render(c.Request().Context() ,c.Response().Writer)
}

func UpdateTodoStatus(c echo.Context) error {
	var taskId = c.QueryParam("taskId")
	db := util.ConnectToDB()


	var task1 = models.SingleTask{}

	for _, task := range tasksLists {
		if task.Taskid == taskId {
			task.Status = !task.Status

			stmt := UpdateTaskStatus(taskId, strconv.FormatBool(task.Status))
			defer db.Query(stmt).Run()
			defer db.Close()

			task1 = *task
		}
	}


	component := views.Task(&task1)
	return component.Render(c.Request().Context() ,c.Response().Writer)
}

func UpdateTodoDesc(c echo.Context) error {
	var taskId = c.FormValue("task-id")
	db := util.ConnectToDB()

	var task1 = models.SingleTask{}

	for _, task := range tasksLists {
		if task.Taskid == taskId {
			task.Description = c.FormValue("description")

			stmt := UpdateTaskDescription(taskId, c.FormValue("description"))
			defer db.Query(stmt).Run()
			defer db.Close()

			task1 = *task
		}
	}

	component := views.Task(&task1)
	return component.Render(c.Request().Context() ,c.Response().Writer)
}

func DeleteBtn(c echo.Context) error {
	var id = c.Param("id")
	component := views.DeleteBtn(id)
	return component.Render(c.Request().Context() ,c.Response().Writer)
}

func DeleteTodoById(c echo.Context) error {
	var id = c.Param("id")

	db := util.ConnectToDB()

	stmt := DeleteTodo(id)
	stmt2 := DeleteTask(id)

	db.Query(stmt).Run()
	db.Query(stmt2).Run()

	defer db.Close()

	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
		}
	}

	component := views.RenderTodos(todos, taskList)
	return component.Render(c.Request().Context() ,c.Response().Writer)
}

func RefreshSingleTodo(c echo.Context) error {
	var id = c.Param("id")
	todo1 := models.SingleTodo{}
	for _, todo := range todos {
		if todo.Id == id {
			todo1 = *todo
		}
	}
	
	component := views.RenderTodo(&todo1, taskList)
	return component.Render(c.Request().Context() ,c.Response().Writer)
}

func RefreshTodoList(c echo.Context) error {
	component := views.RenderTodos(todos, taskList)
	return component.Render(c.Request().Context() ,c.Response().Writer)
}

func EditTodo(c echo.Context) error {
	var taskId = c.QueryParam("taskId")

	var task1 = models.SingleTask{}

	for _, task := range tasksLists {
		if task.Taskid == taskId {
			task1 = *task
		}
	}

	component := views.EditForm(&task1)
	return component.Render(c.Request().Context() ,c.Response().Writer)
}

func SearchTodo(c echo.Context) error {
	var search = c.FormValue("todo")
	todoss := models.TodoList{}
	db := util.ConnectToDB()
	stmt := SearchTodoQuery(search)
	defer db.Query(stmt).Rows(&todoss)
	defer db.Close()

	GetAllTaskPerTodoId()

	return c.JSON(200, todoss)

	// component := views.RenderTodos(todos, taskList)
	// return component.Render(c.Request().Context() ,c.Response().Writer)

	// if search == "" {
	// 	todos = models.TodoList{}
	
	// 	db := util.ConnectToDB()
	// 	stmt := SearchTodoQuery(search)
	// 	defer db.Query(stmt).Rows(&todos)
	// 	defer db.Close()

	// 	GetAllTaskPerTodoId()
	// 	component := views.RenderTodos(todos, taskList)
	// 	return component.Render(c.Request().Context() ,c.Response().Writer)
	// } else {
	// 	component := views.RenderTodos(todos, taskList)
	// 	return component.Render(c.Request().Context() ,c.Response().Writer)
	// }
}