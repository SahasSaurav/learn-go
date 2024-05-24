package Todo

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

type TodoList struct {
	todos []Todo
}

func initializeTodoList() TodoList {
	return TodoList{}
}

func (tl *TodoList) addTodo(description string, isCompleted bool) {
	id := len(tl.todos) + 1
	newTask := Todo{
		id:          id,
		description: description,
		isCompleted: isCompleted,
		createdAt:   time.Now(),
		updateAt:    time.Now(),
	}
	tl.todos = append(tl.todos, newTask)
}

func (tl *TodoList) deleteTodo(id int) error {
	if len(tl.todos) == 0 {
		return errors.New("There is no task present in Todo list")
	}

	if id < 1 {
		return fmt.Errorf("There is no task present with id %d in Todo list", id)
	}

	var taskIndex = -1

	for idx, todo := range tl.todos {
		if todo.id == id {
			taskIndex = idx
			break
		}
	}

	if taskIndex == -1 {
		return fmt.Errorf("There is no task present with id %d in Todo list", id)
	}

	tl.todos = append(tl.todos[:taskIndex], tl.todos[taskIndex+1:]...)
	return nil
}

func (tl *TodoList) markComplete(id int) error {
	if len(tl.todos) == 0 {
		return fmt.Errorf("There is no task present in Todo list")
	}

	if id < 1 {
		return fmt.Errorf("There is no task present with id %d in Todo list", id)
	}

	for idx := range tl.todos {
		if tl.todos[idx].id == id {
			tl.todos[idx].isCompleted = true
			tl.todos[idx].updateAt = time.Now()
			return nil
		}
	}

	return fmt.Errorf("Task with id %d not found", id)
}

func (tl *TodoList) printTodo() {
	fmt.Println("All tasks in todo list")
	for _, task := range tl.todos {
		fmt.Printf("ID: %d, Description: %s, isCompleted %v, createdAt %s\n",
			task.id, task.description, task.isCompleted,
			task.createdAt.Format("02 January 2006"))
	}
}

func Todo() {
	todoList := initializeTodoList()
	todoList.addTodo("Make a coffee", false)
	todoList.addTodo("Learn go lang", false)
	todoList.addTodo("Learn php", false)
	todoList.markComplete(3)
	todoList.printTodo()
	err := todoList.deleteTodo(1)
	if err != nil {
		fmt.Println("Error:", err)
	}
	todoList.printTodo()

	t1 := initializeTodoList()
	t1.addTodo("Make a tea", false)
	t1.printTodo()
}
