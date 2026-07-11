package main

import (
	"strings"
	"testing"
)

type testEnv struct {
	stdout strings.Builder
	stderr strings.Builder
}

func testRun(args ...string)(*testEnv, error) {
	var testEnv testEnv
	err := run(&env{
		args: append([]string{"hit"}, args...),
		stdout: &testEnv.stdout,
		stderr: &testEnv.stderr,
		dryRun: true,
	})
	return &testEnv, err
}

func TestRunValidInput(t *testing.T){
	t.Parallel()
	tenv, err := testRun("http://www.test1.go")
	if err != nil {
		t.Fatalf("got %q\nwant nil err", err)
	}
	if n:= tenv.stdout.Len(); n == 0 {
		t.Errorf("stdout = 0 bytes;, want > 0")
	}
	if n := tenv.stderr.Len(); n != 0 {
		t.Errorf("stderr > 0 bytes;, want = 0; stderr = %d, with %s", n, tenv.stderr.String())
	}
}

func TestRunInvalidInput(t *testing.T) {
	t.Parallel()

	tenv, err := testRun("-c=2", "-n=1", "invalid-url")
	if err == nil {
		t.Fatalf("got nil; want err")
	}
	if n := tenv.stderr.Len(); n == 0 {  
        t.Error("stderr = 0 bytes; want >0")
    }
}