package hit

import (
	"net/http"
	"time"
)

type Summary struct {
	Requests int
	Errors int
	Bytes int64
	RPS float64
	Duration time.Duration
	Fastest time.Duration
	Slowest time.Duration
	Success float64
}
func Summarize(results Results) Summary {
	var s Summary
	if results == nil {
		return s
	}
	started := time.Now()
	for r := range results {
		s.Requests++
		s.Bytes+=r.Bytes
		if r.Error != nil || r.Status != http.StatusOK {  
            s.Errors++
        }
        if s.Fastest == 0 {
            s.Fastest = r.Duration
        }
        if r.Duration < s.Fastest {
            s.Fastest = r.Duration
        }
        if r.Duration > s.Slowest {
            s.Slowest = r.Duration
        }
	}
	if s.Requests > 0 {
		s.Success = (float64(s.Success-float64(s.Errors)/float64(s.Requests))*100)
	}

	s.Duration = time.Since(started)
	s.RPS = float64(s.Requests)/s.Duration.Seconds()
	return s
}