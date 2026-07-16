package main

import (
	"fmt"
	"log"
	"maps"
	"math/rand"
	"net"
	"net/http"
	"net/rpc"
	"sync"
	"time"
)

type PeerServer struct {
	Address string
	Client  *rpc.Client
}

type Args struct {
	GossipLive map[string]int
	Round      int
	Sender     string
}

type Server struct {
	live    map[string]int
	lock    sync.Mutex
	Round   int
	Address string
	peers   []PeerServer
}

func (t *Server) Heartbeat(args *Args, reply *int) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	fmt.Printf("recv hb from %v\n", args.Sender)

	if args.Round > t.Round {
		t.Round = args.Round
	}

	t.live[args.Sender] = t.Round

	for node, r := range args.GossipLive {
		if r > t.live[node] {
			t.live[node] = r
		}
	}

	return nil

}

func (t *Server) sendHeartbeat(to PeerServer) {
	t.lock.Lock()

	t.Round++

	clonedMap := maps.Clone(t.live)

	args := Args{
		GossipLive: clonedMap,
		Round:      t.Round,
		Sender:     t.Address,
	}
	t.lock.Unlock()

	var reply int
	err := to.Client.Call("Server.Heartbeat", args, &reply)
	if err != nil {
		log.Println("RPC error", err)
	}

}

func (t *Server) GenerateReport() {
	t.lock.Lock()
	defer t.lock.Unlock()
	log.Println("REPORT!")
	log.Println("ROUND", t.Round)
	log.Println(t.live)

	fmt.Println("---- LIVE SERVERS ----")
	for serverAddr, lastRound := range t.live {
		if serverAddr == t.Address {
			continue
		}
		if t.Round-lastRound <= 10 {
			fmt.Println(serverAddr)
		}
	}

}

func main() {

	server := new(Server)
	rpc.Register(server)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	go http.Serve(l, nil)

	my_address := "10.239.51.175:1234"
	server.Address = my_address
	server.Round = 0
	server.peers = make([]PeerServer, 0)
	server.live = make(map[string]int)
	peer_addresses := []string{"10.239.202.97:1234", "10.193.194.14:1234", "10.239.178.155:1234", "10.239.92.144:1234"}

	time.Sleep(10 * time.Second) // WAIT to start other servers

	for _, addr := range peer_addresses {
		if addr == my_address {
			continue
		}
		client, err := rpc.DialHTTP("tcp", addr)
		if err != nil {
			log.Println("dialing:", err)
		}
		server.peers = append(server.peers, PeerServer{addr, client})
	}

	/*
		TODO: call send heartbeats to a random server every second
			- NOTE: ensure that this code is non-blocking!
		TODO: call generate report every 5 seconds
	*/
	go func() {
		for {
			server.GenerateReport()
			time.Sleep(5 * time.Second)
		}
	}()

	for {
		server.sendHeartbeat(server.peers[rand.Intn(len(server.peers))])
		time.Sleep(1 * time.Second)
	}

}
