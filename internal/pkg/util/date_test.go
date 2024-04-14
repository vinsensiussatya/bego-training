package util

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestParseLocation(t *testing.T) {
	t.Parallel()

	type args struct {
		hourOffset int
	}
	tests := []struct {
		name string
		args args
		want *time.Location
	}{
		{
			name: "test 0",
			args: args{
				hourOffset: 0,
			},
			want: time.UTC,
		},
		{
			name: "test 7",
			args: args{
				hourOffset: 7,
			},
			want: time.FixedZone("UTC+7", 7*60*60),
		},
		{
			name: "test -7",
			args: args{
				hourOffset: -7,
			},
			want: time.FixedZone("UTC-7", -7*60*60),
		},
		{
			name: "test -24",
			args: args{
				hourOffset: -24,
			},
			want: time.FixedZone("UTC-24", -24*60*60),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := ParseLocation(tt.args.hourOffset)
			require.Equal(t, tt.want, got)
		})
	}
}
