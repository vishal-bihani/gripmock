package jsondecoder_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bavix/gripmock/pkg/jsondecoder"
)

type demo struct {
	ID int `json:"id"`
}

func TestUnmarshalSlice(t *testing.T) {
	inputs := [][]byte{
		[]byte(`{"id": 1}`),
		[]byte(`      {"id": 1}`),
		[]byte(`{"id": 1}      `),
		[]byte(`       [{"id": 1}]`),
		[]byte(`[{"id": 1}]`),
	}

	for _, input := range inputs {
		var results = make([]demo, 0)

		err := jsondecoder.UnmarshalSlice(input, &results)

		require.NoError(t, err)
		require.Equal(t, 1, len(results))
		require.Equal(t, 1, results[0].ID)
	}
}
