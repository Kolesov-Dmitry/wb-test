package tools

import (
	"io"
	"testing"

	"dima.org/wb-test/internals/app/mock"
)

func TestWorkersPool(t *testing.T) {
	cases := []struct {
		workersAmount int
		inputStream   io.Reader
		expected      int
	}{
		{
			workersAmount: 3,
			inputStream:   mock.UrlsStreamFive,
			expected:      5,
		},
		{
			workersAmount: 5,
			inputStream:   mock.UrlsStreamTwo,
			expected:      2,
		},
	}

	for caseIdx, tc := range cases {
		pool := NewURLWorkersPool(tc.workersAmount, mock.NewWorker(), tc.inputStream)
		pool.Start()

		actual := pool.Result().(int)
		if actual != tc.expected {
			t.Fatalf("Test Case #%d failed: expected: %d, actual: %d\n", caseIdx+1, tc.expected, actual)
		}
	}
}
