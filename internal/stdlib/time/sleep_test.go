package time

import (
	"testing"
	"time"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetSleepFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		milliseconds float64
	}{
		{
			name:         "sleep for 100 milliseconds",
			milliseconds: 100,
		},
		{
			name:         "sleep for 0 milliseconds",
			milliseconds: 0,
		},
		{
			name:         "sleep for -1 milliseconds (clamped to 0)",
			milliseconds: -1,
		},
	}

	sleepFunc := getSleepFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			start := time.Now()

			_, err := sleepFunc.Handler(
				nil,
				[]datavalue.Value{datavalue.Number(test.milliseconds)},
			)

			elapsed := time.Since(start)

			if err != nil {
				t.Fatalf("expected no error from handler, got: \"%s\"", err.Error())
			}

			expectedDuration := time.Duration(test.milliseconds) * time.Millisecond

			if test.milliseconds < 0 {
				expectedDuration = 0
			}

			if elapsed < expectedDuration {
				t.Fatalf(
					"expected sleep duration to be at least %d, got %d",
					expectedDuration,
					elapsed,
				)
			}

			if elapsed > expectedDuration+50*time.Millisecond {
				t.Fatalf(
					"expected sleep duration to be at most %d, got %d",
					expectedDuration+50*time.Millisecond,
					elapsed,
				)
			}
		})
	}
}
