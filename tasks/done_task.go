package tasks

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ljubushkin/go_final_project/date"
)

func writeJSONResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	fmt.Fprint(w, message)
}

func DoneTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONResponse(w, http.StatusMethodNotAllowed, `{"error":"Invalid request method"}`)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		writeJSONResponse(w, http.StatusBadRequest, `{"error":"Task ID is required"}`)
		return
	}

	now := time.Now()

	var task Task
	err := DB.QueryRow("SELECT date, repeat FROM scheduler WHERE id = ?", id).Scan(&task.Date, &task.Repeat)
	if err != nil {
		writeJSONResponse(w, http.StatusNotFound, `{"error":"Task not found"}`)
		return
	}

	if task.Repeat == "" {
		if _, err := DB.Exec("DELETE FROM scheduler WHERE id = ?", id); err != nil {
			writeJSONResponse(w, http.StatusInternalServerError, `{"error":"Failed to delete task"}`)
			return
		}
		writeJSONResponse(w, http.StatusOK, `{}`)
		return
	}

	nextDate, err := date.NextDate(now, task.Date, task.Repeat)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, fmt.Sprintf(`{"error":"%s"}`, err.Error()))
		return
	}

	if _, err := DB.Exec("UPDATE scheduler SET date = ? WHERE id = ?", nextDate, id); err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, `{"error":"Failed to update task date"}`)
		return
	}

	writeJSONResponse(w, http.StatusOK, `{}`)
}
