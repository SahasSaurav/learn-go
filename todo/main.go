package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	id          int
	description string
	isCompleted bool
	createdAt   time.Time
	updateAt    time.Time
}

func initializeTodoList() *[]Todo {
	var tasks []Todo
	return &tasks
}

func addTodo(tasks *[]Todo, description string, isCompleted bool) *[]Todo {
	id := len(*tasks) + 1
	newTask := Todo{
		id:          id,
		description: description,
		isCompleted: isCompleted,
		createdAt:   time.Now(),
		updateAt:    time.Now(),
	}
	*tasks = append(*tasks, newTask)
	return tasks
}

func deleteTodo(tasks *[]Todo, id int) (*[]Todo, error) {
	if len(*tasks) == 0 {
		return tasks, errors.New("There is no task present in Todo list")
	}

	if id < 1 {
		return tasks, fmt.Errorf("There is no task present with id %d in Todo list", id)
	}

	var taskIndex = -1

	for idx, todo := range *tasks {
		if todo.id == id {
			taskIndex = idx
			break
		}
	}

	if taskIndex == -1 {
		return tasks, fmt.Errorf("There is no task present with id %d in Todo list", id)
	}

	*tasks = append((*tasks)[:taskIndex], (*tasks)[taskIndex+1:]...)
	return tasks, nil
}

func markComplete(tasks *[]Todo, id int) (*[]Todo, error) {
	if len(*tasks) == 0 {
		return tasks, fmt.Errorf("There is no task present in Todo list")
	}

	if id < 1 {
		return tasks, fmt.Errorf("There is no task present with id %d in Todo list", id)
	}

	for idx := range *tasks {
		if (*tasks)[idx].id == id {
			(*tasks)[idx].isCompleted = true
			(*tasks)[idx].updateAt = time.Now()
			break
		}
	}

	return tasks, nil
}

func printTodo(tasks *[]Todo) {
	fmt.Println("All tasks in todo list")
	for _, task := range *tasks {
		fmt.Printf("ID: %d, Description: %s, isCompleted %v, createdAt %s\n",
			task.id, task.description, task.isCompleted,
			task.createdAt.Format("02 January 2006"))
	}
}

func main() {
	tasks := initializeTodoList()
	tasks = addTodo(tasks, "Make a coffee", false)
	tasks = addTodo(tasks, "Learn go lang", false)
	tasks = addTodo(tasks, "Learn php", false)
	tasks, _ = markComplete(tasks, 3)
	printTodo(tasks)
	tasks, _ = deleteTodo(tasks, 1)
	printTodo(tasks)

	t1 := initializeTodoList()
	t1 = addTodo(t1, "Make a tea", false)
	printTodo(t1)
}
