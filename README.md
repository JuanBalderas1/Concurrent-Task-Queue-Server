# Concurrent Task Queue Server

A backend task-processing service written in Go that accepts jobs through a REST API, places them into a channel-backed queue, and distributes them across a pool of concurrent worker goroutines.

The server tracks each task’s progress in a thread-safe in-memory store, retries unsuccessful tasks, and allows clients to retrieve task results later using a unique ID.

## Motivation

I built this project to strengthen my understanding of backend systems and Go concurrency.

The main goal was to explore how an asynchronous job-processing service can accept work independently from the workers that perform it. Building the project gave me hands-on practice with:

* Goroutines and channels
* Concurrent worker pools
* REST API design
* JSON request and response handling
* Retry logic
* Synchronization with `sync.RWMutex`
* Thread-safe shared-state management
* Organizing a Go project into focused internal packages

This project also helped me better understand how larger systems process background jobs such as notifications, report generation, data imports, and other tasks that do not need to finish during the original HTTP request.

## Project Flow

```text
POST /tasks
    ↓
JSON decoding and validation
    ↓
Channel-backed task queue
    ↓
Concurrent worker pool
    ↓
Task processing
    ↓
Retry handling
    ↓
Task status updates
    ↓
Thread-safe in-memory store
    ↓
GET /tasks/{id}
```

## Current Features

* Accepts new tasks through a REST API
* Decodes task information from JSON request bodies
* Queues tasks using a Go channel
* Processes multiple tasks concurrently with worker goroutines
* Tracks task status and processing attempts
* Retries failed tasks up to a configurable maximum
* Stores task information in memory
* Protects shared task data with `sync.RWMutex`
* Retrieves individual tasks by ID
* Returns task information as JSON

## Core Components

### API Layer

The API handlers receive HTTP requests, decode incoming JSON, submit tasks to the queue, and return task information to clients.

### Task Queue

The queue uses a Go channel to transfer submitted tasks from the API layer to the worker pool.

### Worker Pool

Multiple worker goroutines listen for queued tasks and process them concurrently. This allows the server to handle several independent tasks without processing them one at a time.

### Retry Handling

When a task fails, the server tracks the number of attempts and retries the task until it succeeds or reaches its configured retry limit.

### Task Store

Task status is stored in memory and protected by a `sync.RWMutex`, allowing multiple goroutines to safely read and update shared task data.

## Project Structure

```text
Concurrent Task Queue Server/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── api/
│   │   └── handlers.go
│   ├── queue/
│   │   └── queue.go
│   ├── task/
│   │   ├── task.go
│   │   └── store.go
│   └── worker/
│       └── worker.go
├── go.mod
└── README.md
```

## Technologies Used

* Go
* Goroutines
* Channels
* `net/http`
* `sync.RWMutex`
* JSON-based REST APIs
* PowerShell for example client requests

## Quick Start

### Prerequisites

Install a recent version of Go and confirm that it is available:

```powershell
go version
```

### Clone the Repository

```powershell
git clone https://github.com/JuanBalderas1/Concurrent-Task-Queue-Server.git
cd Concurrent-Task-Queue-Server
```

Replace the repository URL above if the project uses a different GitHub repository name.

### Start the Server

From the project’s root directory, run:

```powershell
go run ./cmd/server
```

The server will begin listening for HTTP requests on its configured local address.

## Usage

### Create a Task

Submit a task with a `POST` request:

```powershell
Invoke-RestMethod `
    -Uri http://localhost:8080/tasks `
    -Method Post `
    -ContentType "application/json" `
    -Body '{"ID":10,"Type":"sms","Payload":"Send login code","MaxRetries":3}'
```

Example task fields:

* `ID`: A unique numeric identifier for the task
* `Type`: The category of work being performed
* `Payload`: The information needed to complete the task
* `MaxRetries`: The maximum number of retry attempts

### Check a Task’s Status

Retrieve the task by its ID:

```powershell
Invoke-RestMethod `
    -Uri http://localhost:8080/tasks/10 `
    -Method Get
```

### Example Response

```json
{
  "ID": 10,
  "Type": "sms",
  "Payload": "Send login code",
  "Status": "completed",
  "Attempts": 1,
  "MaxRetries": 3,
  "Error": ""
}
```

Depending on when the task is retrieved, its status may indicate that it is queued, processing, completed, or failed.

## Potential Improvements

Possible future upgrades include:

* Persistent database storage
* Configurable worker counts
* Structured request and task logging
* Improved graceful shutdown handling
* Task prioritization
* Docker support
* Authentication and authorization
* Request validation
* Automated tests
* Task cancellation
* Scheduled or delayed tasks
* A dashboard for monitoring queued, active, completed, and failed tasks

## Contributing

Contributions, suggestions, and feedback are welcome.

To contribute:

1. Fork the repository.
2. Create a new branch for your change.
3. Make and test your updates.
4. Commit your changes with a clear description.
5. Push the branch to your fork.
6. Open a pull request explaining the proposed improvement.

For substantial changes, please open an issue first so the idea and implementation approach can be discussed.

## Author

Created by [JuanBalderas1](https://github.com/JuanBalderas1).

Thank you for reading!
