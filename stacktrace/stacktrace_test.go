package stacktrace

import (
	"context"
	"runtime"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

// TestNew tests the functionality of New.
func TestNew(t *testing.T) {
	// Allow the test to be run in parallel with other tests.
	t.Parallel()

	// Create two StackTrace objects with different depths.
	a := recursiveCall(5)
	b := recursiveCall(10)

	// Check if the difference in depths is as expected.
	if len(b.pcs)-len(a.pcs) != 5 {
		t.Fatalf(
			"Unexpected depths of stack traces: %d and %d",
			len(a.pcs), len(b.pcs))
	}

	// Iterate over all the program counters.
	for i := range a.pcs {
		switch {
		case i < 6:
			if a.pcs[i] != b.pcs[i] {
				t.Fatalf(
					"Stack frames inside recursiveCall must be the same: "+
						"a[%d] != b[%d]",
					i, i)
			}

			if i >= 2 && a.pcs[i] != a.pcs[i-1] {
				t.Fatalf(
					"Stack frames in recursiveCall must be repeated: "+
						"a[%d] != a[%d]",
					i, i-1)
			}
		case i == 6:
			if a.pcs[i] == b.pcs[i] {
				t.Fatalf(
					"Program counters at TestNew must be different, but: "+
						"a[%d] == b[%d]",
					i, i)
			}
		default:
			if a.pcs[i] != b.pcs[i+5] {
				t.Fatalf(
					"Stack frames deeper than TestNew must be the same: "+
						"a[%d] != b[%d]",
					i, i+5)
			}
		}
	}
}

// TestStackTrace_ToProto tests ToProto generates an expected frame.
func TestStackTrace_ToProto(t *testing.T) {
	t.Parallel()

	st := New(context.Background())
	s := st.ToProto()
	_, file, line, _ := runtime.Caller(0)

	f := s.GetFrames()[0]

	if actual, expected := f.GetLine(), int32(line-2); actual != expected {
		t.Fatalf("Unexpected line number: expected=%d, actual=%v", expected, f)
	}

	if actual, expected := f.GetFile(), file; actual != expected {
		t.Fatalf("Unexpected file path: expected=%s, actual=%v", expected, f)
	}

	if !strings.HasSuffix(f.GetFunction(), "/stacktrace.TestStackTrace_ToProto") {
		t.Fatalf("Unexpected function name: %v", f)
	}
}

// recursiveCall is a helper function that creates a new StackTrace object by
// calling the New() function recursively until the specified depth is reached.
func recursiveCall(depth int) StackTrace {
	if depth <= 0 {
		return New(context.Background())
	}

	return recursiveCall(depth - 1)
}

// BenchmarkNew_Baseline benchmarks the performance of the errors.WithStack
// function of pkg/errors package.
func BenchmarkNew_Baseline(b *testing.B) {
	b.ResetTimer()

	err := errors.New("test")

	for i := 0; i < b.N; i++ {
		_ = errors.WithStack(err)
	}
}

func BenchmarkNew(b *testing.B) {
	b.ResetTimer()

	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		New(ctx)
	}
}

func BenchmarkNew_ManyFrames(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		recursiveCall(100)
	}
}
