package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"beleap.dev/cube/task"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ErrResponse struct {
	HTTPStatusCode int
	Message        string
}

func (a *Api) StartTaskHandler(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	te := task.TaskEvent{}
	err := d.Decode(&te)
	if err != nil {
		msg := fmt.Sprintf("Error unmashalling body: %v\n", err)
		log.Print(msg)
		w.WriteHeader(400)
		e := ErrResponse{
			HTTPStatusCode: 400,
			Message:        msg,
		}
		json.NewEncoder(w).Encode(e)
		return
	}

	a.Worker.AddTask(te.Task)
	log.Printf("Added task %v\n", te.Task.ID)
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(te.Task)
}

func (a *Api) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(a.Worker.GetTasks())
}

func (a *Api) StopTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "taskID")
	if taskID == "" {
		log.Print("No taskID passed in request.\n")
		w.WriteHeader(400)
		// NOTE: not exists in book but I think return is required here.
		return
	}

	tID, _ := uuid.Parse(taskID)
	taskToStop, ok := a.Worker.Db[tID]
	if !ok {
		log.Printf("No task with ID %v found", tID)
		w.WriteHeader(401)
		// NOTE: not exists in book but I think return is required here.
		return
	}

	taskCopy := *taskToStop
	taskCopy.State = task.Completed
	a.Worker.AddTask(taskCopy)

	log.Printf("Added task %v to stop container %v\n", taskToStop.ID, taskToStop.ContainerID)
	w.WriteHeader(204)
}
