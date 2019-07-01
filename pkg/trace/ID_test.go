package trace

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {

	tests := []struct {
		scenario string
		function func(t *testing.T)
	}{
		{
			scenario: "test ID is set ands retrieved as expected from the context",
			function: testTraceID,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.scenario, func(t *testing.T) {
			t.Parallel()
			test.function(t)
		})
	}
}

func testTraceID(t *testing.T) {
	ctx := context.Background()
	assert.Equal(t, Value(ctx), nil)
	newCtx := WithValue(ctx, "a value")
	assert.NotEqual(t, ctx, newCtx)
	assert.Equal(t, Value(ctx), nil)
	assert.Equal(t, Value(newCtx), "a value")
}
