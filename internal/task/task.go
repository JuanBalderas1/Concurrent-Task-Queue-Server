package task

type Task struct {
	ID         int
	Type       string
	Payload    string
	Status     string
	Attempts   int
	MaxRetries int
	Error      string
}
