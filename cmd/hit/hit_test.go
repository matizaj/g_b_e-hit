package main

import "strings"

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