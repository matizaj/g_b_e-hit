package hit

import (
	"slices"
	"testing"
	"time"
)

func TestSummarizeFastetsResult(t *testing.T) {
	results := []Result {
		{Duration: time.Second * 2},
		{Duration: time.Second * 5},
	}

	sum := Summarize(Results(slices.Values(results)))

	if sum.Fastest != 2*time.Second {
		t.Errorf("Fastest=%v; wants 2s\n", sum.Fastest)
	}
}