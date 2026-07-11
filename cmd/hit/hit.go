package main

import (
	"fmt"
	"os"
     "io"
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
     return nil
}