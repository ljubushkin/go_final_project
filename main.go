package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/ljubushkin/go_final_project/auth"
	"github.com/ljubushkin/go_final_project/database"
	"github.com/ljubushkin/go_final_project/date"
	"github.com/ljubushkin/go_final_project/tasks"
	_ "modernc.org/sqlite"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		tasks.AddTaskHandler(w, r)
	case http.MethodGet:
		tasks.GetTaskHandler(w, r)
	case http.MethodPut:
		tasks.EditTaskHandler(w, r)
	case http.MethodDelete:
		tasks.DeleteTaskHandler(w, r)
	default:
		http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	tasks.DB, err = sql.Open("sqlite", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer tasks.DB.Close()

	_, err = os.Stat(dbFile)
	if err != nil {
		database.CreateDatabase(tasks.DB)
	} else {
		log.Println("Database already exists")
	}

	http.HandleFunc("/api/signin", auth.SigninHandler)
	http.Handle("/api/nextdate", http.HandlerFunc(date.ApiNextDate))
	http.Handle("/api/task", auth.Auth(http.HandlerFunc(TaskHandler)))
	http.Handle("/api/tasks", auth.Auth(http.HandlerFunc(tasks.GetTasksHandler)))
	http.Handle("/api/task/done", auth.Auth(http.HandlerFunc(tasks.DoneTaskHandler)))

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./web"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "7540"
	}

	log.Printf("Server is starting on port %s...\n", port)

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
