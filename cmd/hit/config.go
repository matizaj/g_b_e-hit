package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
)

type parseFunc func(string) error
type positiveIntValue int
type config struct {
	url string
	n int
	c int 
	rps int
}

func asPositiveIntValue(p *int) *positiveIntValue {
	return (*positiveIntValue)(p)
}

func (n *positiveIntValue) String() string {
	return strconv.Itoa(int(*n))
}
func (n *positiveIntValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	if err != nil {
		return err
	}

	if v<= 0 {
		return errors.New("Only positive values")
	}

	*n=positiveIntValue(v)

	return nil
}

func parseArgs(c *config, args []string) error {
	// flagSet := map[string]parseFunc {
	// 	"url": stringVar(&c.url),
	// 	"n": intVar(&c.n),
	// 	"c": intVar(&c.c),
	// 	"rps": intVar(&c.rps),
	// }

	fs := flag.NewFlagSet("hit", flag.ContinueOnError)
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "usage %s [options] url\n", fs.Name())
		fs.PrintDefaults()
	}
	// fs.StringVar(&c.url, "url", "", "HTTP server `url` (required)")
	fs.Var(asPositiveIntValue(&c.c), "c", "Concurrency level")
	fs.Var(asPositiveIntValue(&c.rps), "r", "requests per second")
	fs.Var(asPositiveIntValue(&c.n), "n", "number of requests")

	if err:=fs.Parse(args); err != nil {
		return  err
	}

	c.url = fs.Arg(0)

	// for _, arg := range args {
	// 	name, val, _ := strings.Cut(arg, "=")
	// 	name = strings.TrimPrefix(name, "-")

	// 	setVar, ok := flagSet[name]
	// 	if !ok {
	// 		return fmt.Errorf("flag provided but not defined: -%s",name)
	// 	}

	// 	if err := setVar(val); err != nil {
	// 		return fmt.Errorf("invalid value %q for flag -%s: %w",val, name, err) 
	// 	}
	// }
	// return nil
	return fs.Parse(args)
}

func stringVar(p *string) parseFunc {
	return func(s string) error {
		*p=s
		return nil
	}
}

func intVar(v *int) parseFunc {
	return func(i string) error{
		var err error
		*v, err = strconv.Atoi(i)
		return err
	}
}