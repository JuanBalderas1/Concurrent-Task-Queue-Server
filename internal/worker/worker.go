package worker

import (
	"fmt"
	"time"

	"concurrent-task-queue-server/internal/task"
)

func StartWorker(id int, tasks chan task.Task, store *task.Store) {
	go func() {
		for t := range tasks {
			t.Status = "processing"
			store.Save(t)

			fmt.Printf("Worker %d started task %d: %s\n", id, t.ID, t.Payload)

			time.Sleep(2 * time.Second)

			success := shouldSucceed(t)

			if success {
				t.Status = "completed"
				t.Error = ""
				store.Save(t)
				fmt.Printf("Worker %d completed task %d\n", id, t.ID)

			} else {
				t.Attempts++
				t.Status = "failed"
				t.Error = "simulated processing failure"
				store.Save(t)

				fmt.Printf(
					"Worker %d failed task %d (Attempt %d)\n",
					id,
					t.ID,
					t.Attempts,
				)

				if t.Attempts < t.MaxRetries {
					fmt.Printf(
						"Retrying task %d...\n",
						t.ID,
					)

					t.Status = "retrying"
					store.Save(t)
					tasks <- t

				} else {
					fmt.Printf(
						"Task %d permanently failed.\n",
						t.ID,
					)
				}
			}
		}
	}()
}

func shouldSucceed(t task.Task) bool {
	if t.Type == "sms" && t.Attempts == 0 {
		return false
	}
	if t.Type == "backup" && t.Attempts == 0 {
		return false
	}
	return true
}
