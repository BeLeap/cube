package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"beleap.dev/cube/task"
	"beleap.dev/cube/worker"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

func main() {
	host := os.Getenv("CUBE_HOST")
	port, _ := strconv.Atoi(os.Getenv("CUBE_PORT"))

	fmt.Println("Starting Cube worker")

	w := worker.Worker{
		Queue: *queue.New(),
		Db:    make(map[uuid.UUID]*task.Task),
	}
	api := worker.Api{
		Address: host,
		Port:    port,
		Worker:  &w,
	}

	go runTasks(&w)
	api.Start()
}

func runTasks(w *worker.Worker) {
	for {
		if w.Queue.Len() != 0 {
			result := w.RunTask()
			if result.Error != nil {
				log.Printf("Error running task: %v\n", result.Error)
			}
		} else {
			log.Print("No tasks to process currently.\n")
		}
		log.Println("Sleeping for 10 seconds.")
		time.Sleep(10 * time.Second)
	}
}
