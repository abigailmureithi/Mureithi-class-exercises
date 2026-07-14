package main

type Process struct {
	clock []int
	pid   int
	N     int
}

func (p *Process) Start(N int, PID int) {
	p.N = N
	p.pid = PID
	p.clock = make([]int, N)

}

func (p *Process) Internal() {
	p.clock[p.pid] += 1

}

func (p *Process) Send() []int {
	p.clock[p.pid] += 1

	send := make([]int, p.N)
	copy(send, p.clock)
	return send
}

func (p *Process) Receive(ts []int) {
	p.clock[p.pid] += 1

	for i := 0; i < p.N; i++ {
		if ts[i] > p.clock[i] {
			p.clock[i] = ts[i]
		}
	}
}

func Compare(ts1 []int, ts2 []int) int {
	ts1LessThanTs2 := false
	ts2LessThanTs1 := false

	for i := 0; i < len(ts1); i++ {
		if ts1[i] < ts2[i] {
			ts1LessThanTs2 = true
		} else if ts1[i] > ts2[i] {
			ts2LessThanTs1 = true
		}
	}

	if ts1LessThanTs2 == true && !ts2LessThanTs1 {
		return -1
	}

	if ts2LessThanTs1 == true && !ts1LessThanTs2 {
		return 1
	}

	return 0
}
