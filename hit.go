package hit

import (
	"fmt"
	"iter"
	"net/http"
	"time"
)

// Result is performance metrics of a single [http.Request].
type Result struct {
	Status int
	Bytes int64
	Duration time.Duration
	Error error
}

type Results iter.Seq[Result]

func Send(_ *http.Client, _ *http.Request) Result {
	const roundtripTime = 1000 * time.Millisecond
	time.Sleep(roundtripTime)
	return Result{
		Status: http.StatusOK,
		Bytes: 10,
		Duration: roundtripTime,
	}
}

func SendN(n int, req *http.Request) (Results, error){
	if n<=0 {
		return nil, fmt.Errorf("n must be positive: got %d\n", n)
	}
	return func(yield func(Result) bool) {
		for range n {
			result := Send(http.DefaultClient, req)
			if !yield(result) {
				return
			}
		}
	}, nil
}