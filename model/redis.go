package model

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type ListMQ struct {
	client *redis.Client
}

func NewListMQ(client *redis.Client) *ListMQ {
	return &ListMQ{client: client}
}

func (mq *ListMQ) SendTask(task Task) {
	mq.client.LPush(context.Background(), task.Key(), Task2Json(task))
	mq.client.LPush(context.Background(), "all", Task2Json(task))
}

func (mq *ListMQ) GetTask(topic string) (Task, error) {
	var t = Task{}
	cmd := mq.client.RPop(context.Background(), topic)

	if cmd.Err() != nil {
		return Task{}, cmd.Err()
	}
	err := cmd.Scan(&t)
	if err != nil {
		return Task{}, err
	}
	return t, nil
}
