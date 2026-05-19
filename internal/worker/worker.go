package worker

import (
	"fmt"
	"sync"
	"time"

	"concurrent-task-queue-server/internal/task"
)

func StartWorker(id int, tasks chan task.Task, wg *sync.WaitGroup) {
	go func() {
		for t := range tasks {

			t.Status = "processing"

			fmt.Printf("Worker %d started task %d: %s\n", id, t.ID, t.Payload)

			time.Sleep(2 * time.Second)

			success := shouldSucceed(t)

			if success {
				t.Status = "completed"

				fmt.Printf("Worker %d completed task %d\n", id, t.ID)
				wg.Done()

			} else {
				t.Attempts++
				t.Status = "failed"
				t.Error = "simulated processing failure"

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

					tasks <- t

				} else {
					fmt.Printf(
						"Task %d permanently failed.\n",
						t.ID,
					)
					wg.Done()
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
