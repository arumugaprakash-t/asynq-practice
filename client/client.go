package client

import (
	"Asynq/tasks"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
	"time"
)

func StartClient() {
	//create a async client to connect to redis
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "127.0.0.1:6379"})
	defer client.Close()

	//create a task
	task, err := tasks.NewSumTask(int64(5), int64(5))
	if err != nil {
		log.Fatalf("not able to create task %v", err)
	}
	//push it to the queue
	taskInfo, err := client.Enqueue(task, nil)
	if err != nil {
		log.Fatalf("not able to enque task %v", err)
	}
	fmt.Println("task enqueued successfully")
	fmt.Println("Task info ", taskInfo)

	//enque task to perform after x time
	task, err = tasks.NewSumTask(int64(5), int64(999))
	if err != nil {
		log.Fatalf("not able to create task %v", err)
	}

	taskInfo, err = client.Enqueue(task, asynq.ProcessIn(time.Minute*2))
	if err != nil {
		log.Fatalf("not able to enque task %v", err)
	}
	fmt.Println("task enqueued successfully")
	fmt.Println("Task info ", taskInfo)
}
