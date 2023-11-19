package p2p

import (
	"net"
	"sync"
)

type Peers struct {
	conn net.Conn
}

type ServerConfig struct {
	Version       string
	ListenAddr    string
	APIListenAddr string
}

type Server struct {
	ServerConfig
	mu      sync.RWMutex
	peers   map[*net.Addr]*Peers
	addPeer chan *Peer
}

func Newserver(cfg *ServerConfig) *Server {
	return &Server{
		ServerConfig: *cfg,
		peers:        make(map[*net.Addr]*Peers),
		addPeer:      make(chan *Peer),
	}
}

func (s *Server) Start() {
	go s.Loop()
}
func (s *Server) Loop() {
	for {
		select {
		case peer := <-addPeer:
			s.peers
		}
	}
}

//46:00
