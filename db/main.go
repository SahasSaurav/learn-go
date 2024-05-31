package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Todo struct {
	ID          uuid.UUID
	Description string
	IsCompleted bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func CreateTodoTable(ctx context.Context, db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS todo (
		id UUID PRIMARY KEY,
		description TEXT NOT NULL,
		isCompleted BOOLEAN NOT NULL, 
		createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP NOT NULL
	)`

	_, err := db.ExecContext(ctx, query)
	return err
}

func InsertTodo(ctx context.Context, db *sql.DB, args Todo) (Todo, error) {
	query := `INSERT INTO todo 
	(id, description, isCompleted, createdAt, updatedAt)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, description, isCompleted, createdAt, updatedAt`

	row := db.QueryRowContext(ctx, query, args.ID, args.Description, args.IsCompleted, args.CreatedAt, args.UpdatedAt)
	var todo Todo
	err := row.Scan(&todo.ID, &todo.Description, &todo.IsCompleted, &todo.CreatedAt, &todo.UpdatedAt)
	return todo, err
}

func GetAllTodo(ctx context.Context, db *sql.DB) ([]Todo, error) {
	query := `SELECT id, description, isCompleted, createdAt, updatedAt FROM "todo"`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Description, &todo.IsCompleted, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

func DeleteTodo(ctx context.Context, db *sql.DB, id uuid.UUID) error {
	query := `DELETE FROM "todo" where id=$1;`
	_, err := db.ExecContext(ctx, query, id)
	return err
}

func main() {
	const (
		host     = "127.0.0.1"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "go_postgres"
	)

	dbConUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbConUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("Database connected successfully")

	ctx := context.Background()

	err = CreateTodoTable(ctx, db)
	if err != nil {
		log.Fatalf("Failed to create todo table: %v", err)
	}

	todo, err := InsertTodo(ctx, db, Todo{
		ID:          uuid.New(),
		Description: "make tea",
		IsCompleted: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		log.Fatalf("Failed to insert todo: %v", err)
	}

	fmt.Printf("%+v\n", todo)

	todos, err := GetAllTodo(ctx, db)
	if err != nil {
		log.Fatalf("Failed to get all todo: %v", err)
	}
	fmt.Printf("%+v\n", todos)

}
