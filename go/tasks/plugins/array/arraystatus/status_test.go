/*
 * Copyright (c) 2018 Lyft. All rights reserved.
 */

package arraystatus

import (
	"testing"

	types "github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/core"
	"github.com/stretchr/testify/assert"
)

func TestArraySummary_MergeFrom(t *testing.T) {
	t.Run("Update when not equal", func(t *testing.T) {
		expected := ArraySummary{
			types.PhaseRunning: 1,
		}

		other := ArraySummary{
			types.PhaseRunning: 1,
			types.PhaseQueued:  0,
		}

		actual := ArraySummary{
			types.PhaseRunning:          2,
			types.PhasePermanentFailure: 2,
			types.PhaseQueued:           10,
		}

		updated := actual.MergeFrom(other)
		assert.True(t, updated)

		assert.Equal(t, expected, actual)
	})

	t.Run("Delete when 0", func(t *testing.T) {
		expected := ArraySummary{
			types.PhaseRunning: 1,
		}

		other := ArraySummary{
			types.PhaseRunning: 1,
			types.PhaseQueued:  0,
		}

		actual := ArraySummary{}
		updated := actual.MergeFrom(other)
		assert.True(t, updated)

		assert.Equal(t, expected, actual)
	})

	t.Run("Delete when other nil", func(t *testing.T) {
		expected := ArraySummary{}

		actual := ArraySummary{
			types.PhaseRunning: 10,
		}

		updated := actual.MergeFrom(nil)
		assert.True(t, updated)
		assert.Equal(t, expected, actual)
	})

	t.Run("Not Updated when equal", func(t *testing.T) {
		expected := ArraySummary{
			types.PhaseRunning: 1,
			types.PhaseQueued:  10,
		}

		other := ArraySummary{
			types.PhaseRunning:          1,
			types.PhaseQueued:           10,
			types.PhaseRetryableFailure: 0,
		}

		actual := ArraySummary{
			types.PhaseRunning: 1,
			types.PhaseQueued:  10,
		}
		updated := actual.MergeFrom(other)
		assert.False(t, updated)

		assert.Equal(t, expected, actual)
	})
}

func TestArraySummary_Inc(t *testing.T) {
	original := ArraySummary{
		types.PhaseRunning:          2,
		types.PhasePermanentFailure: 2,
		types.PhaseQueued:           10,
	}

	original.Inc(types.PhaseRunning)
	original.Inc(types.PhaseRetryableFailure)

	validatedCount := 0
	for phase, count := range original {
		switch phase {
		case types.PhaseRunning:
			assert.Equal(t, int64(3), count)
			validatedCount++
		case types.PhasePermanentFailure:
			assert.Equal(t, int64(2), count)
			validatedCount++
		case types.PhaseQueued:
			assert.Equal(t, int64(10), count)
			validatedCount++
		case types.PhaseRetryableFailure:
			assert.Equal(t, int64(1), count)
			validatedCount++
		}
	}

	assert.Equal(t, 4, validatedCount)
}
