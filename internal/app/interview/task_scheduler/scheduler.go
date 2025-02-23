package task_scheduler

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"sync"
	"time"
)

type Task struct {
	Id     string
	Status TaskStatus
	Cancel context.CancelFunc
}

type TaskStatus int

const (
	TaskPending   = 0
	TaskRunning   = 1
	TaskCompleted = 2
	TaskCancelled = 3
)

type Scheduler interface {
	ScheduleTask(task func(), runAt time.Time) (string, error)
	CancelTask(taskId string) error
	GetTaskStatus(taskId string) (*Task, error)
}

type schedulerImpl struct {
	mu    sync.RWMutex
	tasks map[string]*Task
}

func (s *schedulerImpl) ScheduleTask(exec func(), runAt time.Time) (string, error) {
	s.mu.Lock()
	ctx, cancel := context.WithCancel(context.Background())

	task := &Task{
		Id:     uuid.New().String(),
		Status: TaskPending,
		Cancel: cancel,
	}

	s.tasks[task.Id] = task
	s.mu.Unlock()

	go func() {
		sleepingDuration := time.Until(runAt)

		if sleepingDuration > 0 {
			select {
			case <-ctx.Done():
				return
			case <-time.After(sleepingDuration):
			}
		}

		s.mu.Lock()
		task.Status = TaskRunning
		s.mu.Unlock()

		exec()

		s.mu.Lock()
		task.Status = TaskCompleted
		s.mu.Unlock()
	}()

	return task.Id, nil
}

func (s *schedulerImpl) CancelTask(taskId string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, ok := s.tasks[taskId]
	if !ok {
		return errors.New("task not found")
	}
	task.Status = TaskCancelled
	task.Cancel()
	return nil
}

func (s *schedulerImpl) GetTaskStatus(taskId string) (*Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	task, ok := s.tasks[taskId]
	if !ok {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func New() Scheduler {
	return &schedulerImpl{}
}
