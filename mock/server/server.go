package server

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

func NewServer(r io.Reader) (*Server, error) {
	log.SetOutput(os.Stderr)
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, err
	}
	return &Server{
		listener:  l,
		mockInput: r,
		aliveBuf:  make([]byte, 4096),
		connMap:   make(map[string]net.Conn),
	}, nil
}

type Server struct {
	connMap   map[string]net.Conn
	listener  net.Listener
	mockInput io.Reader
	aliveBuf  []byte
	sync.Mutex
}

func (s *Server) Addr() net.Addr {
	return s.listener.Addr()
}

func (s *Server) ListenAndServe() error {
	log.Printf("mock Killerqueen stat server started on %s", s.listener.Addr())
	go s.listen()
	return s.serve()
}

func (s *Server) serve() error {
	go s.SendKeepAlives()
	for {
		for _, conn := range s.connMap {
			s.SendData(conn)
		}
	}
}

func (s *Server) SendData(conn net.Conn) {
	var data []byte
	scanner := bufio.NewScanner(s.mockInput)
	var count int
	for scanner.Scan() {
		count++
		line := strings.TrimRight(scanner.Text(), "\n")
		data = append(data, []byte(line)...)
		if count%10 > 0 {
			continue
		}
		log.Printf("sending data to client %s: %s\n", conn.RemoteAddr(), string(data))
		n, err := conn.Write(data)
		if err != nil {
			log.Printf("failed Write data to %s: %s\n", conn.RemoteAddr(), err)
			return
		}
		if n != len(data) {
			log.Printf("short Write to %s: %d, < %d\n", conn.RemoteAddr(), n, len(data))
		}
		data = nil
		time.Sleep(randMS(100, 6000))
	}
	if err := scanner.Err(); err != nil {
		log.Printf("data scanner problems: %s\n", err)
		return
	}
	if len(data) > 0 {
		_, err := conn.Write(data)
		if err != nil {
			log.Printf("failed Write data to %s: %s\n", conn.RemoteAddr(), err)
			return
		}
	}
}

func randMS(low, high int) time.Duration {
	return time.Millisecond*time.Duration(low) + time.Millisecond*time.Duration(rand.Intn(high))
}

func (s *Server) listen() error {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Printf("failed Accept: %s", err)
			continue
		}
		s.addConn(conn)
	}
}

func (s *Server) addConn(conn net.Conn) {
	s.Lock()
	s.connMap[conn.RemoteAddr().String()] = conn
	s.Unlock()
}

func (s *Server) delConn(conn net.Conn) {
	s.Lock()
	delete(s.connMap, conn.RemoteAddr().String())
	s.Unlock()
}

func (s *Server) SendKeepAlives() {
	ticker := time.NewTicker(randMS(200, 600))
	for {
		<-ticker.C
		for _, conn := range s.connMap {
			s.SendKeepAlive(conn)
		}
	}
}

const keepAlive = "![k[im alive],v[]]!"

func (s *Server) SendKeepAlive(conn net.Conn) {
	s.Lock()
	defer s.Unlock()
	_, err := conn.Write([]byte("![k[alive],v[]]!"))
	if err != nil {
		log.Printf("failed Write keep alive: %s\n", err)
		conn.Close()
		s.delConn(conn)
		return
	}
	err = conn.SetReadDeadline(time.Now().Add(time.Millisecond * 200))
	if err != nil {
		log.Printf("failed SetReadDeadline: %s", err)
	}
	n, err := conn.Read(s.aliveBuf)
	if err != nil {
		log.Printf("failed Read keep alive: %s\n", err)
		conn.Close()
		s.delConn(conn)
	}
	err = conn.SetReadDeadline(time.Time{})
	if err != nil {
		log.Printf("failed SetReadDeadline: %s", err)
	}
	resp := s.aliveBuf[:n]
	if !bytes.Equal(resp, []byte(keepAlive)) {
		log.Print("did not receive expected keep alive response, closing connection\n")
		log.Printf("want %s got %s\n", keepAlive, string(resp))
		conn.Close()
		s.delConn(conn)
	}
}
