-- OVERVIEW --
This is a backend systems project that simulates an asynchronous job processing service. It accepts tasks through a REST API, 
places them into a channel-backed queue, distributes them across multiple worker goroutines, retries failed job, and stores
task status in memory so tasks can be checked later by ID. 

This project was built to practice backend engineering concepts so I could better understand concurrency, worker pools,
REST API design, synchronization, and safe shared-state management. 


-- PROJECT FLOW --

POST /tasks
    ↓
JSON decoding
    ↓
Task queue (channel)
    ↓
Concurrent worker pool
    ↓
Retry handling
    ↓
Task status updates
    ↓
Thread-safe in-memory store
    ↓
GET /tasks/{id}


-- PROJECT STRUCTURE --

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

-- CURRENT FEATURES --


-- CORE COMPONENTS --


-- EXAMPLE USAGE --
1) Start the server:

  go run ./cmd/server

2) Create a task:

  Invoke-RestMethod -Uri http://localhost:8080/tasks -Method Post -ContentType "application/json" -Body '{"ID":10,"Type":"sms","Payload":"Send login code","MaxRetries":3}'

3) Check task status:

  Invoke-RestMethod -Uri http://localhost:8080/tasks/10 -Method Get

Example response:

  {
  "ID": 10,
  "Type": "sms",
  "Payload": "Send login code",
  "Status": "completed",
  "Attempts": 1,
  "MaxRetries": 3,
  "Error": ""
  }


-- TECHNOLOGIES USED --
Go, goroutines, channels, net/http, sync.RWMutex, and JSON APIs.

-- IMPROVEMENTS THAT COULD BE MADE --
Possible upgrades that could be done include persistent database storage, configurable worker counts, request logging, 
graceful shutdown improvements, task prioritization, Docker support, authentication, and a dashboard for monitoring queued 
and completed tasks.

-- AUTHOR --
JuanBalderas1

THANK YOU FOR READING! 
