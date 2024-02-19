package server

import (
	"context"

	"helloadmin/pkg/log"
)

type Job struct {
	log *log.Logger
}

func NewJob(
	log *log.Logger,
) *Job {
	return &Job{
		log: log,
	}
}

func (j *Job) Start(ctx context.Context) error {
	// eg: kafka consumer
	return nil
}

func (j *Job) Stop(ctx context.Context) error {
	return nil
}
