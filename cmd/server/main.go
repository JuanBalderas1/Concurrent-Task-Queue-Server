package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"concurrent-task-queue-server/internal/api"
	"concurrent-task-queue-server/internal/member"
	"concurrent-task-queue-server/internal/queue"
	"concurrent-task-queue-server/internal/task"
	"concurrent-task-queue-server/internal/worker"
)

func main() {
	fmt.Println("Chimera Task Server starting...")

	taskQueue := queue.NewTaskQueue(10)
	taskStore := task.NewStore()
	memberStore := member.NewStore()

	startWorkers(
		3,
		taskQueue.Tasks,
		taskStore,
	)

	taskHandler := &api.TaskHandler{
		TaskQueue:   taskQueue.Tasks,
		TaskStore:   taskStore,
		MemberStore: memberStore,
	}

	memberHandler := &api.MemberHandler{
		Store: memberStore,
	}

	mux := http.NewServeMux()

	mux.HandleFunc(
		"/tasks",
		taskHandler.HandleTasks,
	)

	mux.HandleFunc(
		"/tasks/",
		taskHandler.GetTask,
	)

	mux.HandleFunc(
		"/members",
		memberHandler.HandleMembers,
	)

	mux.HandleFunc(
		"/members/",
		memberHandler.GetMember,
	)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	serverErrors := make(chan error, 1)

	go func() {
		fmt.Println(
			"Listening on http://localhost:8080",
		)

		serverErrors <- server.ListenAndServe()
	}()

	shutdownSignal := make(
		chan os.Signal,
		1,
	)

	signal.Notify(
		shutdownSignal,
		os.Interrupt,
		syscall.SIGTERM,
	)

	select {
	case signal := <-shutdownSignal:
		fmt.Printf(
			"\nShutdown signal received: %s\n",
			signal,
		)

	case err := <-serverErrors:
		if !errors.Is(
			err,
			http.ErrServerClosed,
		) {
			log.Fatalf(
				"Server error: %v\n",
				err,
			)
		}
	}

	shutdownServer(server)
}

func startWorkers(
	workerCount int,
	taskQueue chan task.Task,
	taskStore *task.Store,
) {
	for workerID := 1; workerID <= workerCount; workerID++ {
		worker.StartWorker(
			workerID,
			taskQueue,
			taskStore,
		)
	}
}

func shutdownServer(server *http.Server) {
	fmt.Println(
		"Chimera Task Server shutting down...",
	)

	shutdownContext, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	err := server.Shutdown(shutdownContext)
	if err != nil {
		log.Printf(
			"Graceful shutdown failed: %v\n",
			err,
		)

		err = server.Close()
		if err != nil {
			log.Printf(
				"Forced shutdown failed: %v\n",
				err,
			)
		}
	}

	fmt.Println(
		"Chimera Task Server stopped safely.",
	)
}
