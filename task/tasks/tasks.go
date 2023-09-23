package tasks

import (
	"encoding/json"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

const (
	filePath = "./tasks.json"
)

func (t *TaskList) AddTask(task Task) {
	t.Tasks = append(t.Tasks, task)
}

func (t *TaskList) RemoveTask(task Task) {
	for i, v := range t.Tasks {
		if v.ID == task.ID {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
		}
	}
}

func (t *TaskList) CompleteTask(task Task) {
	for i, v := range t.Tasks {
		if v.ID == task.ID {
			t.Tasks[i].Completed = true
		}
	}
}

func (t *TaskList) UncompleteTask(task Task) {
	for i, v := range t.Tasks {
		if v.ID == task.ID {
			t.Tasks[i].Completed = false
		}
	}
}

func (t *TaskList) GetTask(id int) Task {
	for _, v := range t.Tasks {
		if v.ID == id {
			return v
		}
	}
	return Task{}
}

func (t *TaskList) GetTasks() []Task {
	return t.Tasks
}

func (t *TaskList) GetCompletedTasks() []Task {
	var completedTasks []Task
	for _, v := range t.Tasks {
		if v.Completed {
			completedTasks = append(completedTasks, v)
		}
	}
	return completedTasks
}

func (t *TaskList) GetUncompletedTasks() []Task {
	var uncompletedTasks []Task
	for _, v := range t.Tasks {
		if !v.Completed {
			uncompletedTasks = append(uncompletedTasks, v)
		}
	}
	return uncompletedTasks
}

func (t *TaskList) GetNextID() int {
	var maxID int
	for _, v := range t.Tasks {
		if v.ID > maxID {
			maxID = v.ID
		}
	}
	return maxID + 1
}

func CreateTaskList() TaskList {
	return TaskList{}
}

func LoadFile() (TaskList, error) {
	f, err := os.Open("tasks.json")
	if err != nil {
		return TaskList{}, err
	}
	defer f.Close()
	var taskList TaskList
	err = json.NewDecoder(f).Decode(&taskList)
	if err != nil {
		return TaskList{}, err
	}
	return taskList, nil
}

func FileExists() bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	return false
}

func CreateFile() error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	tasks := TaskList{
		Tasks: []Task{},
	}
	err = json.NewEncoder(f).Encode(tasks)
	if err != nil {
		return err
	}
	return nil
}

func SaveToFile(taskList TaskList) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(taskList)
	if err != nil {
		return err
	}
	return nil
}
