package hit

import (
	"context"
	"fmt"
	"io"
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
		client:= &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: o.Concurrency,
			},
			CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: 30*time.Second,
		}
		o.Send = func(r *http.Request) Result {
			return Send(client, r) 
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


func Send(client *http.Client, req *http.Request) Result {
	started := time.Now()
	var (
		bytes int64
		code int
	)

	resp, err := client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		code = resp.StatusCode
		bytes, err = io.Copy(io.Discard, resp.Body)
	}

	return Result {
		Duration: time.Since(started),
		Bytes: bytes,
		Status: code,
		Error: err,
	}
}
