package _933_recent_calls

import "testing"

func TestRecentCounter(t *testing.T) {
	tests := []struct {
		name  string
		times []int
		want  []int
	}{
		{
			name:  "Example 1",
			times: []int{1, 100, 3001, 3002},
			want:  []int{1, 2, 3, 3},
		},
		{
			name:  "Single call",
			times: []int{1},
			want:  []int{1},
		},
		{
			name:  "Calls within range",
			times: []int{100, 200, 300, 400},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "Calls outside range",
			times: []int{1, 3002, 6003},
			want:  []int{1, 1, 1},
		},
		{
			name:  "Boundary values",
			times: []int{1, 3001, 3002, 6001, 6002},
			want:  []int{1, 2, 2, 3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := Constructor()

			for i, timestamp := range tt.times {
				got := rc.Ping(timestamp)

				if got != tt.want[i] {
					t.Errorf("Ping(%d) = %d, want %d", timestamp, got, tt.want[i])
				}
			}
		})
	}
}
