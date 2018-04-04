package scheduler

import "imooc.com/learngo/crawler/engine"

type SimpleScheduler struct {
	workChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkChan(c chan engine.Request) {
	s.workChan = c
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	//s.workChan <- request
	go func() { s.workChan <- request }()
}
