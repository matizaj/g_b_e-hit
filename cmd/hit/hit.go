package main

import (
	"fmt"
	"os"
)

const logo = `  #1
 __  __     __     ______
/\ \_\ \   /\ \   /\__  _\
\ \  __ \  \ \ \  \/_/\ \/
 \ \_\ \_\  \ \_\    \ \_\
  \/_/\/_/   \/_/     \/_/`  

// const usage = `  #1
// Usage:
//   -url
//        HTTP server URL (required)
//   -n
//        Number of requests
//   -c
//        Concurrency level
//   -rps
//        Requests per second` 

func main() {
//     fmt.Printf("%s\n%s", logo, usage)
//     var c config
//     if err := parseArgs(&c, os.Args[1:]); err!=nil {
//           fmt.Printf("%s\n%s", err, usage)
//           os.Exit(1)
//     }
//     fmt.Printf("%s\n\nSending %d requests to %q (concurrency: %d)\n",logo, c.n, c.url, c.c)

     c:= config {
          n: 120,
          c: 1,
     }
     if err := parseArgs(&c, os.Args[1:]); err != nil {
          os.Exit(1)
     }
     fmt.Printf("%s\n\nSending %d requests to %q (concurrency: %d)\n",logo, c.n, c.url, c.c)
}