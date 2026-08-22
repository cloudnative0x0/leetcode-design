package _933_recent_calls

type RecentCounter struct {
	calls []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		calls: make([]int, 0),
	}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.calls = append(rc.calls, t)

	for rc.calls[0] < t-3000 {
		rc.calls = rc.calls[1:]
	}

	return len(rc.calls)
}
