package taskes

import "time"

type Job struct {
	JobFunc    func() error
	Duration   time.Duration
	CreataedAt time.Time
	LastRun    *time.Time
}

func NewTask(job func() error, duration time.Duration) *Job {
	return &Job{
		JobFunc:    job,
		Duration:   duration,
		CreataedAt: time.Now(),
	}
}

type Tasks struct {
	Jobs []*Job
}

func New() *Tasks {
	return &Tasks{}
}

func (t *Tasks) Add(job func() error, duration time.Duration) {
	t.Jobs = append(t.Jobs, NewTask(job, duration))
}

func (t *Tasks) Run() error {
	for _, job := range t.Jobs {
		j := job
		go func(job *Job) {
			ticker := time.NewTicker(job.Duration)
			for range ticker.C {
				now := time.Now()
				j.LastRun = &now
				err := job.JobFunc()
				if err != nil {
					return
				}
			}
		}(j)
	}
	return nil
}
