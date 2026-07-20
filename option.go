package hit

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type SendFunc func(*http.Request) Result

type Options struct {
	Concurrency int
	RPS         int
	Send        SendFunc
}

func Defaults() Options {
	return withDefaults(Options{})
}

func withDefaults(o Options) Options {
	if o.Concurrency ==0 {
		o.Concurrency =1
	}

	if o.Send == nil {
		o.Send = func(r *http.Request) Result {
			return Send(http.DefaultClient, r) 
		}
	}
	return o  
}

func SendN(ctx context.Context,  n int, req *http.Request, opts Options) (Results, error){
	opts = withDefaults(opts)
	ctx, cancel := context.WithCancel(ctx)
	results:= runPipeline(ctx, n, req, opts)
	if n<=0 {
		return nil, fmt.Errorf("n must be positive: got %d\n", n)
	}
	return func(yield func(Result) bool) {
		defer cancel()
		for result :=range results {
			if !yield(result) {
				return
			}
		}
	}, nil
}


func Send(_ *http.Client, _ *http.Request) Result {
	const roundtripTime = 1000 * time.Millisecond
	time.Sleep(roundtripTime)
	return Result{
		Status: http.StatusOK,
		Bytes: 10,
		Duration: roundtripTime,
	}
}
