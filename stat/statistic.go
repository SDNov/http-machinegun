package stat

import "sync"

type Statistic struct {
	muSuccessCount sync.Mutex
	SuccessCount   int64

	muErrorCount sync.Mutex
	ErrorCount int64
}

func (s *Statistic) IncrementSuccessCounter() {
	s.muSuccessCount.Lock()
	defer s.muSuccessCount.Unlock()
	s.SuccessCount++
}

func (s *Statistic) IncrementErrorCounter() {
	s.muErrorCount.Lock()
	defer s.muErrorCount.Unlock()
	s.ErrorCount++
}
