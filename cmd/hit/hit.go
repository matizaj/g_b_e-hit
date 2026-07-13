package main

import (
	"fmt"
	"go_by_example/hello-go/hit"
	"io"
	"math"
	"net/http"
	"os"
	"time"
)

const logo = `
 __  __     __     ______
/\ \_\ \   /\ \   /\__  _\
\ \  __ \  \ \ \  \/_/\ \/
 \ \_\ \_\  \ \_\    \ \_\
  \/_/\/_/   \/_/     \/_/`  


type env struct {
     stdout  io.Writer
     stderr  io.Writer
     args    []string
     dryRun  bool
}

func main() {
     if err := run(&env{
          stdout: os.Stdout,
          stderr: os.Stderr,
          args: os.Args,
     }); err != nil {
          os.Exit(1)
     }
}

func run(env *env) error {
      c:= config {
          n: 120,
          c: 1,
     }
     if err := parseArgs(&c, env.args[1:], env.stderr); err != nil {
          return err
     }

     if err := runHit(&c, env.stdout); err != nil {
          fmt.Fprintf(env.stderr, "\nerror occured: %v\n", err)
          return err
     }
     fmt.Fprintf(env.stdout, "%s\n\nSending %d requests to %q (concurrency: %d), test mode: %v\n",logo, c.n, c.url, c.c, c.dry)
     return  nil
}

func runHit(c *config, stdout io.Writer) error {
     req, err := http.NewRequest(http.MethodGet, c.url, http.NoBody)
     if err != nil {
          return fmt.Errorf("creating a new request: %w", err)
     }

     results, err := hit.SendN(c.n, req, hit.Options{
          Concurrency: c.c,
          RPS: c.rps,
     })
     if err != nil {
          return fmt.Errorf("sending request: %w", err)
     }
     printSummary(hit.Summarize(results), stdout)
     return nil
}


func printSummary(sum hit.Summary, stdout io.Writer) {
    fmt.Fprintf(stdout, `  #1
Summary:
    Success:  %.0f%%  #2
    RPS:      %.1f  #3
    Requests: %d
    Errors:   %d
    Bytes:    %d
    Duration: %s
    Fastest:  %s
    Slowest:  %s
`,
        sum.Success,
        math.Round(sum.RPS),
        sum.Requests,
        sum.Errors,
        sum.Bytes,
        sum.Duration.Round(time.Millisecond),
        sum.Fastest.Round(time.Millisecond),
        sum.Slowest.Round(time.Millisecond),
    )
}