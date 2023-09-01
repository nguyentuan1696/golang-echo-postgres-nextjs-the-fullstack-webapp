package utils

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"thichlab-backend-docs/infrastructure/logger"
	"time"
)

type run func(env *CronjobEnvironment) error

type process struct {
	id  string
	err error
	run run
	env *CronjobEnvironment
}

type CronjobEnvironment struct {
	CacheRedis *redis.Client
}

// Start the process
func (proc *process) Start(channel chan<- *process) (err error) {
	// make my goroutine signal its death, wether it's a panic or a return
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				proc.err = err
			} else {
				proc.err = fmt.Errorf("panic happened with %v", r)
			}
		} else {
			proc.err = err
		}
		channel <- proc
	}()

	return proc.run(proc.env)
}

// Processes collects sub processes
type Processes struct {
	list []*process
}

// Add a process
func (procs *Processes) Add(id string, fn run, environment *CronjobEnvironment) {
	procs.list = append(procs.list, &process{
		id:  id,
		run: fn,
		env: environment,
	})
}

// Run all processes
func (procs *Processes) Run() {
	// make a buffered channel with the space for workers
	processes := len(procs.list)
	channel := make(chan *process, processes)

	// go through all registered processes
	for _, proc := range procs.list {
		go func() {
			err := proc.Start(channel)
			if err != nil {

			}
		}()
	}

	// read the channel, it will block until something is written,
	// then it will restart the process
	for proc := range channel {
		logger.Error("Process [%s] stopped with err :%v", proc.id, proc.err)
		// reset err
		proc.err = nil

		// give it a little time
		time.Sleep(1 * time.Second)

		// a goroutine has ended, restart it
		go func() {
			err := proc.Start(channel)
			if err != nil {

			}
		}()
	}
}
