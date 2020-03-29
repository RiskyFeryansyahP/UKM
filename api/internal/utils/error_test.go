package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrapErrorJson(t *testing.T) {
	tests := []struct {
		message    error
		statuscode int
		want       struct {
			result Error
		}
	}{
		{
			message:    errors.New("failed register"),
			statuscode: 400,
			want: struct {
				result Error
			}{
				result: Error{
					StatusCode: 400,
					Message:    "failed register",
				},
			},
		},
	}

	for _, v := range tests {
		e := WrapErrorJson(v.message, v.statuscode)
		require.Equal(t, v.want.result.Message, e.Message)
		require.Equal(t, v.want.result.StatusCode, e.StatusCode)
	}
}
