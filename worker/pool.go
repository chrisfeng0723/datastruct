/**
 * @Author: dudu
 * @Description: 
 * @File:  pool.go
 * @Version: 1.0.0
 * @Date: 2020/1/9 10:39
 */
package worker

import (
	"errors"
	"log"
	"sync/atomic"
)

type Task struct {
	Handler func(v ...interface{})

	Params []interface{}
}

type Pool struct {
	capacity uint64

	runningWorkers uint64

	state  int64
	taskC  chan *Task
	closeC chan bool

	PanicHandler func(interface{})
}

var ErrInvalidPoolCap = errors.New("invalid pool cap")

const (
	RUNNING = 1
	STOPED  = 0
)

func NewPool(capacity uint64) (*Pool, error) {
	if capacity <= 0 {
		return nil, ErrInvalidPoolCap
	}

	return &Pool{
		capacity: capacity,
		state:    RUNNING,
		taskC:    make(chan *Task, capacity),
		closeC:   make(chan bool),
	}, nil

}

func (p *Pool) GetCap() uint64 {
	return atomic.LoadUint64(&p.capacity)
}

func (p *Pool) run() {
	p.incRunning()

	go func() {
		defer func() {
			p.decRunning()
			if r := recover(); r != nil {
				if p.PanicHandler != nil {
					p.PanicHandler(r)
				} else {
					log.Printf("Worker panic:%s\n", r)
				}
			}
		}()

		for {
			select {
			case task, ok := <-p.taskC:
				if !ok {
					return
				}
				task.Handler(task.Params...)
			case <-p.closeC:
				return
			}
		}
	}()

}

func (p *Pool) incRunning() {
	atomic.AddUint64(&p.runningWorkers, 1)
}

func (p *Pool) decRunning() {
	atomic.AddUint64(&p.runningWorkers, ^uint64(0))
}

func (p *Pool) GetRunningWorkers() uint64 {
	return atomic.LoadUint64(&p.runningWorkers)
}

var ErrPoolAlreadyClosed = errors.New("pool already closed")

func (p *Pool) Put(task *Task) error {
	if p.state == STOPED {
		return ErrPoolAlreadyClosed
	}

	if p.GetRunningWorkers() < p.GetCap() {
		p.run()
	}

	p.taskC <- task
	return nil
}

func (p *Pool) Close() {
	p.state = STOPED

	for len(p.taskC) > 0 {

	}
	p.closeC <- true
	close(p.taskC)
}
