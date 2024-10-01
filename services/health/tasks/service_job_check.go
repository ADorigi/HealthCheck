package tasks

import (
	"context"

	"github.com/adorigi/healthcheck/pkg/healthcheck"
	"github.com/adorigi/workerpool"
	"go.uber.org/zap"
)

type ServiceJobCheckTask struct {
	logger *zap.Logger
	workerpool.TaskProperties
	jobName   string
	namespace string
}

func NewServiceJobCheckTask(logger *zap.Logger, taskProperties workerpool.TaskProperties, jobName, namespace string) *ServiceJobCheckTask {
	return &ServiceJobCheckTask{
		logger:         logger,
		TaskProperties: taskProperties,
		jobName:        jobName,
		namespace:      namespace,
	}
}

func (s ServiceJobCheckTask) Properties() workerpool.TaskProperties {
	return s.TaskProperties
}

func (s ServiceJobCheckTask) Run(_ context.Context) error {

	s.logger.Info("Processing Job", zap.String("JobID", s.TaskProperties.ID.String()))
	healthcheck.GetJobStatus(s.namespace, s.jobName)
	return nil
}
