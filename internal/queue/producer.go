package queue

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Producer struct {
	client *redis.Client
}

func NewProducer(client *redis.Client) *Producer {
	return &Producer{
		client: client,
	}
}

func (p *Producer) Enqueue(ctx context.Context, req CreateJobRequest) (CreateJobResponse, error) {
	jobID := uuid.NewString()

	job := Job{
		Id:         jobID,
		Type:       req.Type,
		Payload:    req.Payload,
		Status:     StatusPending,
		RetryCount: 0,
		MaxRetries: 3,
	}
	data, err := json.Marshal(job)
	if err != nil {
		return CreateJobResponse{}, err
	}
	err = p.client.LPush(ctx, "jobs:queue", data).Err()
	if err != nil {
		return CreateJobResponse{}, err
	}
	return CreateJobResponse{JobID: job.Id}, nil
}
