package handlers

var SelectAllTodosQuery = "SELECT * FROM \"public\".\"todo\";"
var SelectAllTasksQuery = "SELECT * FROM \"public\".\"task\";"

func InsertTodoQuery(id string, title string) string {
	return "INSERT INTO todo(id, title) VALUES ('" + id + "','" + title + "');"
}

func InsertTaskQuery(id string, id2 string, id3 string, id4 string, task1 string, task2 string, task3 string) string {
	return "INSERT INTO task(taskid, description, id) VALUES ('" + id2 + "','" + task1 + "','" + id + "'),('" + id3 + "','" + task2 + "','" + id + "'),('" + id4 + "','" + task3 + "','" + id + "');"
}

func DeleteTodo(id string) string {
	return "DELETE FROM todo WHERE id = '" + id + "';"
}

func DeleteTask(id string) string {
	return "DELETE FROM task WHERE id = '" + id + "';"
}

func UpdateTaskStatus(taskid string, val string) string {
	return "UPDATE task SET status = " + val + " WHERE taskid = '" + taskid + "';"
}

func UpdateTaskDescription(taskid string, description string) string {
	return "UPDATE task SET description = '" + description + "' WHERE taskid = '" + taskid + "';"
}

func SearchTodoQuery(title string) string {
	return "SELECT * FROM todo WHERE title LIKE '%" + title + "%';"
}