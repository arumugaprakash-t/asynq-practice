package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
)

const (
	SumTask = "sum:two"
)

type SumTwo struct {
	A int64
	B int64
}

func NewSumTask(a, b int64) (*asynq.Task, error) {
	marshalledPayload, err := json.Marshal(SumTwo{A: a, B: b})
	if err != nil {
		return nil, err
	}
	fmt.Println("Marshalled payload", marshalledPayload)

	return asynq.NewTask(SumTask, marshalledPayload, nil), nil
}

func HandleSumTask(ctx context.Context, task *asynq.Task) error {
	var sum SumTwo
	fmt.Println("Got task", task.Payload())
	err := json.Unmarshal(task.Payload(), &sum)
	if err != nil {
		return err
	}
	fmt.Printf("Sum of %d and %d is = %d\n", sum.A, sum.B, sum.A+sum.B)
	return nil
}
