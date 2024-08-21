package server

import (
	"Asynq/tasks"
	"github.com/hibiken/asynq"
	"log"
)

func StartServer() {
	server := asynq.NewServer(asynq.RedisClientOpt{Addr: "127.0.0.1:6379"},
		asynq.Config{
			Concurrency: 5,
			Queues: map[string]int{
				"critical": 3,
				"default":  2,
				"low":      1,
			},
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.SumTask, tasks.HandleSumTask)

	if err := server.Run(mux); err != nil {
		log.Fatalln("server not started")
	}
}
