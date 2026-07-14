package main


type Process struct {
	clock []int
	pid int
	N int
}

func (p *Process) Start(N int, PID int)  {
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
	return 0
}
