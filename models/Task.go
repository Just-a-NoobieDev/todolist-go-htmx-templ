package models

type SingleTask struct {
	Taskid string `json:"taskid"`
	Description string `json:"description"`
	Status bool `json:"status"`
	Id string `json:"singleTodoId"`
}

type TaskList []*SingleTask

// Methods
func NewSingleTask(taskid string, description string, id string) *SingleTask {
	return &SingleTask{
		Taskid: taskid,
		Description: description,
		Status: false,
		Id: id,
	}
}

func TaskListPerTodoId(id string, taskList []*SingleTask) map[string]interface{} {
	list := map[string]interface{}{}
	slice := []*SingleTask{}
	for _, task := range taskList {
		if task.Id == id {
			slice = append(slice, task)
		}
	}
	list[id] = slice
	return list
}

// func (t *TaskList) ResetList() {
// 	taskList = TaskList{}
// }

// func (t *TaskList) indexOf(taskid string) int {
// 	for i, task := range *t {
// 		if task.Taskid == taskid {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func (t *TaskList) AddTask(taskid string, description string, id string) *SingleTask {
// 	task := NewSingleTask(taskid, description, id)
// 	*t = append(*t, task)
// 	return task
// }

// func (t *TaskList) RemoveTask(taskid string) {
// 	index := t.indexOf(taskid)
// 	*t = append((*t)[:index], (*t)[index+1:]...)
// }